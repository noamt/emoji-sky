package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

const space = "\U00003000"

const sun = "☀️"
const moon = "🌕"
const cloud = "☁️"

var midSkyByDay = []string{"🦅", "🦆", "🕊", "🐦"}
var midSkyByNight = []string{"🦉"}

var lowSkyByDay = []string{"🐝", "🦋"}
var lowSkyByNight = []string{"🦇"}

func main() {
	printTheSky()
}

func printTheSky() {
	r := newlySeededRandom()
	t := time.Now().UTC()
	h := t.Hour()

	status := fmt.Sprintf("%s\n%s\n%s\n%s%s\n%s\n%s\n%s", sunOrMoon(h), clouds(r),
		midSky(r, h), midSky(r, h), space, space, lowSky(r, h), lowSky(r, h))
	if os.Getenv("DEVELOPMENT") == "TRUE" {
		log.Println("Printing status to stdout")
		println(status)
		return
	}

	log.Println("Posting status to Twitter")
	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	_, _, err := client.Statuses.Update(status, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func sunOrMoon(hour int) string {
	r := row()

	if isDay(hour) {
		r[4] = sun
	} else {
		r[4] = moon
	}
	return joinRow(r)
}

func clouds(r *rand.Rand) string {
	cT := row()
	c := r.Intn(len(cT))

	for i := 0; i < c; i++ {
		cT[r.Intn(len(cT))] = cloud
	}

	return joinRow(cT)
}

func midSky(r *rand.Rand, hour int) string {
	if isDay(hour) {
		return sky(midSkyByDay, r)
	}
	return sky(midSkyByNight, r)
}

func lowSky(r *rand.Rand, hour int) string {
	if isDay(hour) {
		return sky(lowSkyByDay, r)
	}
	return sky(lowSkyByNight, r)
}

func sky(animals []string, rand *rand.Rand) string {
	n := rand.Intn(3)
	r := row()
	lR := len(r)
	lA := len(animals)
	for i := 0; i < n; i++ {
		r[rand.Intn(lR)] = animals[rand.Intn(lA)]
	}

	return joinRow(r)
}

func newlySeededRandom() *rand.Rand {
	s := rand.NewSource(time.Now().UnixNano())
	return rand.New(s)
}

func row() []string {
	return []string{space, space, space, space, space, space, space, space, space}
}

func joinRow(r []string) string {
	return strings.Join(r, "")
}

func isDay(h int) bool {
	return h >= 5 && h < 19
}
