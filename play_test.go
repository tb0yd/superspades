package superspades

import (
    "testing"
    "github.com/jteeuwen/deck"
)

func TestPlayChecksState(t *testing.T) {
    g := NewGame()

    if _, err := g.Play(card{deck.NewCard(deck.Spades, 0)}); err == nil {
        t.Errorf("allowed play without cards dealt")
    }
}
