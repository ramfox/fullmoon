package state

import (
	"fmt"
	"strings"
	"unicode"
)

type State struct {
	magic_word string
	letters    Letters
	rounds     int
	guessed    []string
	phase      MoonPhase
}

func NewState(mw string) (*State, error) {
	if strings.Contains(mw, "_") {
		return nil, fmt.Errorf("magic word cannot contain an underscore.")
	}

	mw = strings.Trim(mw, " \n\t")
	return &State{strings.ToLower(mw), NewLetters(), 0, NewGuessed(mw), WaningGibbous}, nil
}

func NewGuessed(mw string) []string {

	guessed := make([]string, len(mw))

	for i, char := range mw {
		if unicode.IsLetter(char) {
			guessed[i] = "_"
			continue
		}
		guessed[i] = string(char)
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

func (s *State) GuessWord(guess string) (bool, string) {
	if guess == s.Reveal() {
		return true, ""
	}
	return false, fmt.Sprintf("Wrong! The magic word is not '%s'. Guess again.", guess)
}

func (s *State) GuessLetter(letter string) (bool, string) {
	if len(letter) > 1 {
		letter = string(letter[0])
	}
	// check if the letter has been guessed already
	if s.letters.IsUsed(letter) {
		return true, fmt.Sprintf("You have guessed the letter '%s' already.", letter)
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
		return false, fmt.Sprintf("Wrong! Guess again.")
	}

	// add the letters to the proper places in the guessed slice
	for _, num := range indices {
		s.guessed[num] = letter
	}

	return true, ""
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

func (s *State) WrongGuess() {
	s.phase++
}

func (s *State) GameOver() bool {
	return s.phase == FullMoon
}

func (s *State) Phase() string {
	return s.phase.String()
}
