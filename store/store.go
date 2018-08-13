package store

import (
	"fmt"
	"strings"
)

type State struct {
	magic_word string
	letters    Letters
	rounds     int
	guessed    []string
}

func NewState(mw string) *State {
	return &State{strings.ToLower(mw), NewLetters(), 0, NewGuessed(mw)}
}

func NewGuessed(mw string) []string {
	guessed := make([]string, len(mw))

	// create list of possible indices for the spaces letter
	indices := []int{}
	temp_mw := mw

	for {
		index := strings.Index(temp_mw, " ")

		if index == -1 {
			break
		}

		indices = append(indices, index)

		if index+1 >= len(temp_mw) {
			break
		}

		temp_mw = strings.Replace(temp_mw, " ", "*", 1)
	}

	for _, num := range indices {
		guessed[num] = " "
	}

	return guessed
}

type Letters map[string]bool

func NewLetters() Letters {
	return Letters{}
}

func (l Letters) IsUsed(letter string) bool {
	return l[letter]
}

func (l Letters) MarkUsed(letter string) {
	l[letter] = true
}

func (s *State) Guess(letter string) string {
	letter = strings.ToLower(letter)

	// check if the letter has been guessed already
	if s.letters.IsUsed(letter) {
		return fmt.Sprintf("You have guessed the letter '%s' already.", letter)
	}

	// record that the letter has been guessed
	// and increment the number of rounds
	s.letters.MarkUsed(letter)
	s.rounds++

	// create list of possible indices of the guessed letter
	indices := []int{}
	temp_mw := s.magic_word

	for {
		index := strings.Index(temp_mw, letter)

		if index == -1 {
			break
		}

		indices = append(indices, index)

		if index+1 >= len(temp_mw) {
			break
		}

		temp_mw = strings.Replace(temp_mw, letter, "*", 1)
	}

	if len(indices) == 0 {
		return fmt.Sprintf("Wrong! Guess again.")
	}

	// add the letters to the proper places in the guessed slice
	for _, num := range indices {
		s.guessed[num] = letter
	}

	return ""
}

func (s *State) Reveal() string {
	return s.magic_word
}

func (s *State) Guessed() string {
	g := ""
	for _, letter := range s.guessed {
		if letter == "" {
			g += "_"
			continue
		}
		g += letter
	}
	return g
}