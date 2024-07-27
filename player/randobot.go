package player

import (
	"math/rand"
	"time"
)

type RandoBot struct {
	rand *rand.Rand
	seed int64
}

func NewRandoBot(opts ...func(*RandoBot)) *RandoBot {
	defaultSeed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(defaultSeed))
	rb := &RandoBot{r, defaultSeed}

	for _, opt := range opts {
		opt(rb)
	}

	return rb
}

func WithSeed(seed int64) func(*RandoBot) {
	return func(rb *RandoBot) {
		rb.rand.Seed(seed)
		rb.seed = seed
	}
}

func (r RandoBot) GetMove(validMoves [][2]string) (string, string) {
	randomIndex := r.rand.Intn(len(validMoves))
	pick := validMoves[randomIndex]
	return pick[0], pick[1]
}

func (r RandoBot) State() map[string]any {
	return map[string]any{
		"name": "RandoBot",
		"seed": r.seed,
	}
}
