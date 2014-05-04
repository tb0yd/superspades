package superspades

// Struct which represents a game state. The zero value is a game
// that has not been started yet. CurrentPlayer is the player's seat
// number, either 1, 2, 3, 4, or 0 which represents that the game
// is finished (or that a starting player hasn't been assigned). 
// Bids and Books are both arrays of 4 uint8 values, where the index
// is the player's seat number minus 1.
type Game struct {
    CurrentPlayer uint8
    Bids [4]uint8
    Books [4]uint8

    // cardPositions is a list of positions for each of the 52 cards.
    // Card order is defined in card.go.
    cardPositions [52]uint8
}

const (
    cardInDeck = iota
    cardInHand1
    cardInHand2
    cardInHand3
    cardInHand4
    cardOnTable
    cardInBook
)

const (
    GameNoBids = iota
    GameNotDealt
    GameNoCurrentPlayer
    GameInPlay
    GameFinished
)

func NewGame() Game {
    return Game{}
}

func (g Game) State() (int, string) {
    if g.cardPositions == [52]uint8{} {
        return GameNotDealt, MsgNotDealt
    }

    if g.Bids == [4]uint8{} {
        return GameNoBids, MsgNoBids
    }

    if g.Books[0] + g.Books[1] + g.Books[2] + g.Books[3] >= 13 {
        return GameFinished, MsgFinished
    }

    if g.CurrentPlayer == 0 {
        panic(PanicMsgNoCurrentPlayer) // developer error -> panic
    }

    return GameInPlay, MsgInPlay
}
