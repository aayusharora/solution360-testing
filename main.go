package main

import (
	"fmt"
	"github.com/tsenart/vegeta/lib"
	"github.com/urfave/cli"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

// Panic if there is an error
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var (
		users int
	)

	// The Go random number generator source is deterministic, so we need to seed
	// it to avoid getting the same output each time
	rand.Seed(time.Now().UTC().UnixNano())

	// Configure our command line app
	app := cli.NewApp()
	app.Name = "Pokemon User Data Generator"
	app.Usage = "generate a stream of test data for vegeta. Type 'pokemon help' for details"

	// Add -users flag, which defaults to 5
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "users",
			Value:       5,
			Usage:       "Number of users to simulate",
			Destination: &users,
		},
	}

	// Our app's main action
	app.Action = func(c *cli.Context) error {

		// Combine verb and URL to a target for Vegeta
		verb := c.Args().Get(0)
		url := c.Args().Get(1)
		target := fmt.Sprintf("%s %s", verb, url)

		if len(target) > 1 {

			for i := 1; i < users; i++ {

				// Generate request data

				// Generate a map of the request body that we'll convert to JSON


				// Convert the map to JSON
				payload := `{"anonymousId":"\"aayush\"","channel":"Chrome 70",
							"context":{"ip":"103.46.201.226","userAgent":"Mozilla/5.0 (Macintosh Intel Mac OS X 10_13_6)
AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36"},"traits":{"name":"jacal","email":"",
	"userId":"1234","address":{"city":"Delhi","state":"National Capital Territory of Delhi","country":"India"}},"messageID":"",
	"receivedAt":null,"sentAt":null,"timestamp":"%s","pageInfo":{"title":"MC41NTQ4NDY4NjYwODQzNjA4","url":"http://localhost:4000/game"},"event":{"name":"","additional":""},"version":"1.0","apiKey":"MC41NTQ4NDY4NjYwODQzNjA4","destinationBQ":null,"extra":{"d1":"Hero","d2":"Jackal","d3":"h","d4":"","d5":"","d6":"","d7":"","d8":"","d9":"","d10":"","d11":"","d12":"","d13":"","d14":"","d15":"","d16":"","d17":"","d18":"","d19":"","d20":"","d21":"","d22":"Zero","d23":"","d24":"","d25":"","d26":"","d27":"","d28":"","d29":"","d30":"","d31":"","d32":"","d33":"","d34":"","d35":"","d36":"","d37":"","d38":"","d39":"","d40":"","d41":"","d42":"","d43":"","d44":"","d45":"","d46":"","d47":"","d48":"","d49":"","d50":"","d51":"","d52":"","d53":"","d54":"","d55":"","d56":"","d57":"","d58":"","d59":"","d60":"","d61":"","d62":"","d63":"","d64":"","d65":"","d66":"","d67":"","d68":"","d69":"","d70":"","d71":"","d72":"","d73":"","d74":"","d75":"","d76":"","d77":"","d78":"","d79":"","d80":"","d81":"","d82":"","d83":"","d84":"","d85":"","d86":"","d87":"","d88":"","d89":"","d90":"","d91":"","d92":"","d93":"","d94":"","d95":"","d96":"","d97":"","d98":"","d99":"","d100":""}}`
				// Create a tmp directory to write our JSON files

				body := fmt.Sprintf(payload, time.Now())

				err := os.MkdirAll("tmp", 0755)
				filename := fmt.Sprintf("tmp/%s.json", time.Now())
				err = ioutil.WriteFile(filename, []byte(body), 0644)
				check(err)

				check(err)
				fmt.Println(body)
				rate := vegeta.Rate{Freq: 10000, Per: time.Second}
				duration := time.Second/10000
				targeter := vegeta.NewStaticTargeter(vegeta.Target{
					Method: "POST",
					URL:    "https://go-test-dot-solution360-event.appspot.com/pubsub/publish?token=token69",
					Body: []byte(body),
				})
				attacker := vegeta.NewAttacker()

				var metrics vegeta.Metrics
				for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
					metrics.Add(res)
				}
				metrics.Close()

				fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)

				// Print the attack target
				fmt.Println(target)


				// Print '@' followed by the absolute path to our JSON file, followed by
				// two newlines, which is the delimiter Vegeta uses
			}
		} else {
			// Return an error if we're missing the required command line arguments
			return cli.NewExitError("You must specify the target in format 'VERB url'", 1)
		}
		return nil
	}

	app.Run(os.Args)
}