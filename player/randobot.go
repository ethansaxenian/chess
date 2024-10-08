package player

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ethansaxenian/chess/move"
	"github.com/ethansaxenian/chess/piece"
)

type RandoBot struct {
	rand      *rand.Rand
	seed      int64
	moveDelay time.Duration
}

func NewRandoBot(opts ...func(*RandoBot)) *RandoBot {
	defaultSeed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(defaultSeed))
	rb := &RandoBot{r, defaultSeed, 0}

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

func WithMoveDelay(delay time.Duration) func(*RandoBot) {
	return func(rb *RandoBot) {
		rb.moveDelay = delay
	}
}

func (r RandoBot) GetMove(validMoves []move.Move) move.Move {
	time.Sleep(r.moveDelay)
	randomIndex := r.rand.Intn(len(validMoves))
	pick := validMoves[randomIndex]
	return pick
}

func (r RandoBot) ChoosePromotionPiece(square string) piece.Piece {
	randomIndex := r.rand.Intn(len(piece.PossiblePromotions))
	pick := piece.PossiblePromotions[randomIndex]
	return pick
}

func (r RandoBot) String() string {
	return fmt.Sprintf("RandoBot%d", r.seed)
}

func (r RandoBot) IsBot() bool {
	return true
}
