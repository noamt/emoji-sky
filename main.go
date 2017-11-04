package main

import(
  "fmt"
  "time"
  "math/rand"
  "strings"
)

const sun = "\U00002600"
const moon = "\U0001F315"
const cloud = "\U00002601"

//Bald Eagle, Dove, Duck, Bird
var midSkyByDay = [4]string {"\U0001F985", "\U0001F54A", "\U0001F986", "\U0001F426"}

//Bee, Butterfly
var lowSkyByDay = [2]string {"\U0001F41D", "\U0001F98B"}

func main() {
  fmt.Printf("%s%s%s%s\n\n%s%s", sunOrMoon(), clouds(), midSky(), midSky(), lowSky(), lowSky())
}

func sunOrMoon() (string) {
  t := time.Now().UTC()
  h := t.Hour()
  if (h >= 5 && h < 19) {
    return fmt.Sprintf("    %s    \n", sun)
  }
  return fmt.Sprintf("    %s    \n", moon)
}

func clouds() (string) {
  r := newlySeededRandom()
  c := r.Intn(9)
  cT := []string {" ", " ", " ", " ", " ", " ", " ", " ", " ", "\n"}

  for i := 0; i < c; i++ {
		cT[r.Intn(9)] = cloud
	}

  return strings.Join(cT, "")
}

func midSky() (string) {
  r := newlySeededRandom()
  m := r.Intn(5)
  mT := []string {" ", " ", " ", " ", " ", " ", " ", " ", " ", "\n"}

  for i := 0; i < m; i++ {
		mT[r.Intn(9)] = midSkyByDay[r.Intn(4)]
	}

  return strings.Join(mT, "")
}

func lowSky() (string) {
  r := newlySeededRandom()
  l := r.Intn(5)
  lT := []string {" ", " ", " ", " ", " ", " ", " ", " ", " ", "\n"}

  for i := 0; i < l; i++ {
		lT[r.Intn(9)] = lowSkyByDay[r.Intn(2)]
	}

  return strings.Join(lT, "")
}

func newlySeededRandom() (*rand.Rand) {
  s := rand.NewSource(time.Now().UnixNano())
  return rand.New(s)
}
