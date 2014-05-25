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

// Reset default game to new game.
func Reset() {
    DefaultGame = &Game{}
    return
}

// Set current player on default game
func SetCurrentPlayer(n int) {
    DefaultGame.CurrentPlayer = uint8(n)
    return
}

// Return the scores of the default game for teams 1 and 2, respectively. Returns
// zeroed scores if the game isn't over yet.
func Scores() (s1 Score, s2 Score) {
    if !DefaultGame.IsOver() {
        return
    }

    s1 = Score{}.Add(DefaultGame.Books[0], DefaultGame.Bids[0], DefaultGame.Books[2], DefaultGame.Bids[2])
    s2 = Score{}.Add(DefaultGame.Books[1], DefaultGame.Bids[1], DefaultGame.Books[3], DefaultGame.Bids[3])

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

