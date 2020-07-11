package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	"github.com/stevec7/random/weightedexercise/pkg/weight"
)

// i knew how to map integers to chars, but lets just copy/paste
//	https://www.calhoun.io/creating-random-strings-in-go/
const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		//b[i] = charset[seededRand.Intn(len(charset))]
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func randString(length int) string {
	return stringWithCharset(length, charset)
}

type choice struct {
	value   int
	counter int
}

func main() {
	// make 10 weight structs and append them to weights
	ws := weight.Weights{}

	for i := 0; i < 100; i++ {
		name := randString(3)
		wght := rand.Intn(16384)
		w := weight.Weight{N: name, W: wght}
		ws.Items = append(ws.Items, w)
	}

	chosen := map[string]choice{}
	for i := 0; i < 16384; i++ {
		c := ws.WeightedRandom()
		if _, ok := chosen[c.N]; !ok {
			chosen[c.N] = choice{
				counter: 1,
				value:   c.W,
			}
		} else {
			tmp := chosen[c.N]
			tmp.counter++
			chosen[c.N] = tmp
		}
	}
	keys := []string{}
	for k := range chosen {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return chosen[keys[i]].counter < chosen[keys[j]].counter
	})
	for _, k := range keys {
		fmt.Printf("k: %s, weight: %d, chosen: %d\n", k, chosen[k].value, chosen[k].counter)
	}
}
