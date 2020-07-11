package weight

import (
	"math/rand"
	"sort"
	"time"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// Weight is a record
type Weight struct {
	N string
	W int
}

type Weights struct {
	Items []Weight
}

//type ByWeights Weights

func (w Weights) Len() int {
	return len(w.Items)
}

func (w Weights) Swap(i, j int) {
	w.Items[i], w.Items[j] = w.Items[j], w.Items[i]
}

func (w Weights) Less(i, j int) bool {
	return w.Items[i].W < w.Items[j].W
}

func (ws *Weights) Random() Weight {
	rSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	return ws.Items[rSource.Intn(len(ws.Items))]
}

func (ws *Weights) WeightedRandom() Weight {
	rSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	sort.Sort(Weights(*ws))

	// as we iterate through the sorted (by weight) list of weights,
	//	keep track of the running sum, and create a simple list
	sum := 0
	total := make([]int, len((*ws).Items))
	for i, w := range (*ws).Items {
		sum += w.W
		total[i] = sum
	}
	r := rSource.Intn(sum)

	// the magic is here, it's like a primitive way to do bucketing:
	//		https://golang.org/pkg/sort/#SearchInts
	i := sort.SearchInts(total, r)
	return ws.Items[i]
}
