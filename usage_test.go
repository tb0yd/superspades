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
    var err error

    playMove := func(g Game,  m int) Game {
        g, err = g.Play(m)
        if err != nil {
            t.Errorf("%s\n", err.Error())
        }
        return g
    }

    playAllMoves := func(g Game, moves []int) Game {
        for _, move := range moves {
            g = playMove(g, move)
        }
        return g
    }

    s1 := Score{}
    s2 := Score{}

    rand.Seed(123)
    g := Game{}.Deal()

    //bidding round
    g.CurrentPlayer = 1

    g, err = g.Bid(5)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Bid(4)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    //playing round
    g.CurrentPlayer = 1

    g = playAllMoves(g, []int{
        8,0,0,0,
        2,0,0,4,
        7,6,3,0,
        7,2,6,7,
        1,0,1,0,
        3,0,0,1,
        5,0,4,0,
        0,0,0,4,
        3,4,1,3,
        0,2,0,0,
        2,0,2,0,
        0,1,0,0,
        0,0,0,0,
    })

    s1 = s1.Add(g.Books[0], g.Bids[0], g.Books[2], g.Bids[2])
    s2 = s2.Add(g.Books[1], g.Bids[1], g.Books[3], g.Bids[3])

    if s1.ToInt() != -54 || s2.ToInt() != 140 {
        t.Errorf("expected -54 to 140, was %d to %d\n", s1.ToInt(), s2.ToInt())
    }

    rand.Seed(124)
    g = Game{}.Deal()
    g.CurrentPlayer = 2

    g, err = g.Bid(14)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Bid(4)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    g, err = g.Bid(6)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }
    g.CurrentPlayer = 2

    g = playAllMoves(g, []int{
        4,4,4,0,
        0,3,3,0,
        3,3,0,3,
        6,0,2,5,
        5,0,1,4,
        1,3,2,2,
        6,3,0,2,
        4,2,2,5,
        4,1,4,0,
        3,2,1,1,
        2,0,2,2,
        1,1,1,0,
        0,0,0,0,
    })

    s1 = s1.Add(g.Books[0], g.Bids[0], g.Books[2], g.Bids[2])
    s2 = s2.Add(g.Books[1], g.Bids[1], g.Books[3], g.Bids[3])


    if s1.ToInt() != -117 || s2.ToInt() != 300 {
        t.Errorf("expected -117 to 300, was %d to %d\n", s1.ToInt(), s2.ToInt())
    }
}

func TestDefaultGameUsage(t *testing.T) {
    rand.Seed(124)
    Deal()
    SetCurrentPlayer(2)
    var err error

    s1 := Score{}
    s2 := Score{}

    playMove := func( m int) {
        err = Play(m)
        if err != nil {
            t.Errorf("%s\n", err.Error())
        }
    }

    playAllMoves := func(moves []int) {
        for _, move := range moves {
            playMove(move)
        }
    }

    err = Bid(14)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    err = Bid(4)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }

    err = Bid(6)
    if err != nil {
        t.Errorf("%s\n", err.Error())
    }
    SetCurrentPlayer(2)

    playAllMoves([]int{
        4,4,4,0,
        0,3,3,0,
        3,3,0,3,
        6,0,2,5,
        5,0,1,4,
        1,3,2,2,
        6,3,0,2,
        4,2,2,5,
        4,1,4,0,
        3,2,1,1,
        2,0,2,2,
        1,1,1,0,
        0,0,0,0,
    })

    s1 = s1.Add(DefaultGame.Books[0], DefaultGame.Bids[0], DefaultGame.Books[2], DefaultGame.Bids[2])
    s2 = s2.Add(DefaultGame.Books[1], DefaultGame.Bids[1], DefaultGame.Books[3], DefaultGame.Bids[3])

    if s1.ToInt() != -63 || s2.ToInt() != 160 {
        t.Errorf("expected -63 to 160, was %d to %d\n", s1.ToInt(), s2.ToInt())
    }
}

func TestSimplestUsage(t *testing.T) {
    DefaultGame = &Game{}
    ShuffleAndDeal()

    // must set CurrentPlayer before bidding
    DefaultGame.CurrentPlayer = 1

    Bid(4)
    Bid(4) // other 2 players automatically will bid nil

    // must set CurrentPlayer before playing
    DefaultGame.CurrentPlayer = 1

    // play first playable card in hand for whole game
    for i := 0; i < 52; i++ {
        Play(0)
    }

    // get scores
    Score{}.Add(DefaultGame.Books[0], DefaultGame.Bids[0], DefaultGame.Books[2], DefaultGame.Bids[2])
    Score{}.Add(DefaultGame.Books[1], DefaultGame.Bids[1], DefaultGame.Books[3], DefaultGame.Bids[3])
}
