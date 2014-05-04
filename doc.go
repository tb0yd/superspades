// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

// Package superspades implements the card game Super Spades, also known
// as "Suicide Spades," a variant of Spades where a player from each team
// must go nil every round.
//
// There are two ways to use this library. You can either manage several
// games at once, or you can manage a single game at a time. Additionally,
// you can either run the game via method calls or interactively with 
// channels.
package superspades