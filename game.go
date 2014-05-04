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
    cardOnTableForPlayer1
    cardOnTableForPlayer2
    cardOnTableForPlayer3
    cardOnTableForPlayer4
    cardInBook
)

func NewGame() Game {
    return Game{}
}
