package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/ramfox/fullmoon/state"
	"golang.org/x/crypto/ssh/terminal"
)

func Play(s *state.State, r *bufio.Reader, w *bufio.Writer) {
	for {
		WriteWhite(w, "\n\n")
		WriteGreen(w, fmt.Sprintf("%s\n\n%s\n", s.Phase(), s.Guessed()))
		WriteWhite(w, "Guess: ")
		st, _ := r.ReadString('\n')
		str := strings.Replace(st, "\n", "", -1)

		if len(str) < 1 {
			continue
		}

		if len(str) == 1 {
			goodGuess, res := s.GuessLetter(str)
			if !goodGuess {
				s.WrongGuess()
				WriteRed(w, res)

				if s.GameOver() {
					WriteLoss(w, s)
					return
				}
			}
		}

		if len(str) > 1 {
			goodGuess, res := s.GuessWord(str)
			if !goodGuess {
				s.WrongGuess()
				WriteRed(w, res)

				if s.GameOver() {
					WriteLoss(w, s)
					return
				}
			}
			WriteWin(w, s.Reveal())
			return
		}

		if s.Guessed() == s.Reveal() {
			WriteWin(w, s.Reveal())
			return
		}
	}
}

func Setup(r *bufio.Reader, w *bufio.Writer) (*state.State, error) {
	w.WriteString("Enter magic word: ")
	w.Flush()
	mw, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return nil, fmt.Errorf("error reading magic word: %s", err)
	}
	Clear(w)
	return state.NewState(string(mw))
}

func WriteRed(w *bufio.Writer, s string) error {
	w.WriteString(fmt.Sprintf("\x1b[31m%s\x1b[0m\n", s))
	return w.Flush()
}

func WriteGreen(w *bufio.Writer, s string) error {
	w.WriteString(fmt.Sprintf("\x1b[32m%s\x1b[0m", s))
	return w.Flush()
}

func WriteWhite(w *bufio.Writer, s string) error {
	w.WriteString(s)
	return w.Flush()
}

func WriteWin(w *bufio.Writer, word string) error {
	return WriteGreen(w, fmt.Sprintf("\nYou have correctly guessed '%s' as the magic word!\n\nCongratulations, the ritual to keep you sane has worked. You can rest easy. The monster is asleep.\n", word))
}

func Clear(w *bufio.Writer) {
	w.WriteString("\033[H\033[2J")
	w.Flush()
}

func WriteLoss(w *bufio.Writer, s *state.State) error {
	Clear(w)
	WriteGreen(w, fmt.Sprintf("%s\n\n%s\n", s.Phase(), s.Guessed()))
	return WriteRed(w, "\nOh no! the full moon has come. You feel the monster inside of you claw its way to the surface. Your bones break and reform, your skin bubbles, claws burst from your fingers. You see red.\n\nAWWOOOOOOOooooooo!\n\n")
}

func PlayAgain(r *bufio.Reader, w *bufio.Writer) {
	WriteWhite(w, "Press enter to play again or ctrl-c to quit.")
	s, _ := r.ReadByte()
	if string(s) != "\n" {
		os.Exit(1)
	}
	Clear(w)
}
