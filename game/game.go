package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/ramfox/fullmoon/store"
	"golang.org/x/crypto/ssh/terminal"
)

func Play(s *store.State, r *bufio.Reader, w *bufio.Writer) {
	for {
		WriteGreen(w, s.Guessed()+"\n")
		WriteWhite(w, "Guess a letter or the phrase: ")
		st, _ := r.ReadString('\n')
		str := strings.Replace(st, "\n", "", -1)

		if str == "" {
			return
		}

		if str == "exit" {
			os.Exit(1)
		}

		if len(str) > 1 {
			if str == s.Reveal() {
				WriteWin(w, s.Reveal())
				os.Exit(1)
			}
			WriteRed(w, fmt.Sprintf("Wrong! The magic word is not '%s'. Guess again.", str))
			continue
		}

		res := s.Guess(str)
		if res != "" {
			WriteRed(w, res)
		}

		if s.Guessed() == s.Reveal() {
			WriteWin(w, s.Reveal())
			os.Exit(1)
		}
	}
}

func Setup(r *bufio.Reader, w *bufio.Writer) (*store.State, error) {
	w.WriteString("Enter magic word: ")
	w.Flush()
	mw, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return nil, fmt.Errorf("error reading magic word: %s", err)
	}
	return store.NewState(string(mw))
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
	return WriteGreen(w, fmt.Sprintf("\nYou have correctly guessed '%s' as the magic word!\nCongratulations!\n", word))
}

func Clear(w *bufio.Writer) {
	w.WriteString(("\033[H\033[2J"))
	w.Flush()
}
