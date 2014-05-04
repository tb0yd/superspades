package superspades

import (
    "testing"
)

func TestBid(t *testing.T) {
    g := NewGame()
    g.CurrentPlayer = 1
    g, err := g.Bid(4)

    if err != nil {
        t.Errorf("error: %s\n", err.Error())
    }

    if g.Bids != [4]uint8{4,0,14,0} {
        t.Errorf("Bids weren't 4,0,14,0 -- were %d,%d,%d,%d\n", g.Bids[0], g.Bids[1], g.Bids[2], g.Bids[3])
    }

    if g.CurrentPlayer != 2 {
        t.Errorf("Current player wasn't 2, was %d\n", g.CurrentPlayer)
    }

    g, err = g.Bid(14)

    if err != nil {
        t.Errorf("error: %s\n", err.Error())
    }

    if g.Bids != [4]uint8{4,14,14,0} {
        t.Errorf("Bids weren't 4,14,14,0 -- were %d,%d,%d,%d\n", g.Bids[0], g.Bids[1], g.Bids[2], g.Bids[3])
    }

    if g.CurrentPlayer != 4 {
        t.Errorf("Current player wasn't 4, was %d\n", g.CurrentPlayer)
    }

    g, err = g.Bid(7)

    if err != nil {
        t.Errorf("error: %s\n", err.Error())
    }

    if g.Bids != [4]uint8{4,14,14,7} {
        t.Errorf("Bids weren't 4,14,14,0 -- were %d,%d,%d,%d\n", g.Bids[0], g.Bids[1], g.Bids[2], g.Bids[3])
    }

    if g.CurrentPlayer != 0 {
        t.Errorf("Current player wasn't 0, was %d\n", g.CurrentPlayer)
    }
}
