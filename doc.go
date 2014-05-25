// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

// Package superspades implements a variant of Spades where a player from each team
// must go nil every round.
//
// There are two ways to use this library. You can either manage several
// games at once, or you can manage a single game at a time.
//
// To run a simple game:
//  package main
//
//  import (
//      ss "github.com/tb0yd/superspades"
//      "fmt"
//  )
//
//  func main() {
//      ss.Reset()
//      ss.ShuffleAndDeal()
//
//      ss.SetCurrentPlayer(1) // must set CurrentPlayer before bidding
//
//      ss.Bid(4)
//      ss.Bid(4) // other 2 players automatically will bid nil
//
//      ss.SetCurrentPlayer(1) // must set CurrentPlayer before playing
//
//      // play first playable card in each player's hand for whole game
//      for i := 0; i < 52; i++ { ss.Play(0) }
//
//      s1, s2 := ss.Scores()
//
//      fmt.Printf("score: %d to %d\n", s1.ToInt(), s2.ToInt())
//  }
package superspades