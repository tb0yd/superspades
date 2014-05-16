package superspades

//import "fmt"

// Game to be used in programs that only keep track of a single game.
var DefaultGame *Game = &Game{}

// Make a bid in the default game. It is assumed that the one bidding
// is the current player. If the bid is out of range (1-14),
// an error is thrown. 14 is the nil bid.
func Bid(amt uint8) error {
    game, err := DefaultGame.Bid(amt)
    DefaultGame = &game
    return err
}

// Deal cards in default game to all 4 players evenly. User is responsible
// for seeding random number generator.
func Deal() {
    game := DefaultGame.Deal()
    DefaultGame = &game
    return
}

// Returns a slice of Cards representing the cards in player n's hand in the
// default game.
func Hand(n uint8) []Card {
    return DefaultGame.Hand(n)
}

// Returns true if there are no more moves left in the default game.
func IsOver() bool {
    return DefaultGame.IsOver()
}

// Play the card in the current player's PlayableCards(), indexed by cardNo.
// Acts on the default game. If the current player doesn't have
// that many cards playable, an error is thrown.
func Play(cardNo int) error {
    game, err := DefaultGame.Play(cardNo)
    DefaultGame = &game
    return err
}

// Returns a slice of cards in the current player's hand for the default game. Only
// cards which the player is allowed to play will appear.
func PlayableCards() []Card {
    return DefaultGame.PlayableCards()
}

// Set current player on default game
func SetCurrentPlayer(n int) {
    DefaultGame.CurrentPlayer = uint8(n)
    return
}

// Deal cards in default game to all 4 players evenly. Random number
// generator is automatically seeded.
func ShuffleAndDeal() {
    game := DefaultGame.ShuffleAndDeal()
    DefaultGame = &game
    return
}

// Returns all the cards on the table for the default game, indexed by player.
func Table() [4]Card {
    return DefaultGame.Table()
}

