// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

// Package superspades implements a variant of Spades where a player from each team
// must go nil every round.
//
// There are two ways to use this library. You can either manage several
// games at once, or you can manage a single game at a time.
//
// To run a simple game:
//     DefaultGame = &Game{}
//     ShuffleAndDeal()
//
//     DefaultGame.CurrentPlayer = 1 // must set CurrentPlayer before bidding
//
//     Bid(4)
//     Bid(4) // other 2 players automatically will bid nil
//
//     DefaultGame.CurrentPlayer = 1 // must set CurrentPlayer before playing
//
//     // play first playable card in each player's hand for whole game
//     for i := 0; i < 52; i++ { Play(0) }
//
//     s1 := Score{}.Add(DefaultGame.Books[0], DefaultGame.Bids[0], DefaultGame.Books[2], DefaultGame.Bids[2])
//     s2 := Score{}.Add(DefaultGame.Books[1], DefaultGame.Bids[1], DefaultGame.Books[3], DefaultGame.Bids[3])
//
//     fmt.Printf("score: %d to %d\n", s1.ToInt(), s2.ToInt())
package superspades