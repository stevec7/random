package publisher

import (
	"context"
)

type data struct {
	filename string
	heat int
}

type Publisher interface {
	Publish(context.Context, []data) error
}


func main() {
	_ = context.Background()
}
