package superspades

import (
    "testing"
    "math/rand"
    "fmt"
)

func displayGame(g Game) string {
    return fmt.Sprintf("Current player: %d\n", g.CurrentPlayer) +
           fmt.Sprintf("Bids:           %x\n", g.Bids) +
           fmt.Sprintf("Books:          %x\n", g.Books) +
           fmt.Sprintf("Hand:           %s\n", g.Hand(g.CurrentPlayer-1))+
           fmt.Sprintf("Playable:       %s\n", g.PlayableCards())+
           fmt.Sprintf("Table:          %s\n", g.Table())
}

func TestUsage(t *testing.T) {
    rand.Seed(123)

    s1 := Score{}
    s2 := Score{}

    g := Game{}.Deal()

    //bidding round
    g.CurrentPlayer = 1

    g, err := g.Bid(5)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Bid(4)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    //playing round
    g.CurrentPlayer = 1

    g, err = g.Play(8)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }


    g, err = g.Play(0)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(0)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(0)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(2)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(0)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(0)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(4)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(7)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(6)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(3)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(0)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(7)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(2)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(6)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(7)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(1)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(0)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(1)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(0)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(3)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(0)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(0)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(1)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(5)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(0)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(4)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(0)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(0)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(0)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(0)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(4)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(3)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(4)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(1)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(3)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(0)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(2)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(0)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(0)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(2)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(0)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(2)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(0)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(0)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(1)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(0)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(0)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(0)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(0)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(0)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Play(0)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    s1 = s1.Add(g.Books[0], g.Bids[0], g.Books[2], g.Bids[2])
    s2 = s2.Add(g.Books[1], g.Bids[1], g.Books[3], g.Bids[3])

    if s1.ToInt() != -54 || s2.ToInt() != 140 {
        t.Errorf("expected -54 to 140, was %d to %d\n", s1.ToInt(), s2.ToInt())
    }
}
