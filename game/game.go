package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ramfox/fullmoon/store"
)

func Play(s *store.State, r *bufio.Reader, w *bufio.Writer) {
	for {
		WriteGreen(w, s.Guessed()+"\n")
		WriteWhite(w, "Guess a letter: ")
		st, _ := r.ReadString('\n')
		str := strings.Replace(st, "\n", "", -1)

		if str == "" {
			return
		}

		if str == "exit" {
			os.Exit(1)
		}

		if len(str) > 1 {
			WriteRed(w, "Only one letter can be guessed at a time")
			continue
		}

		res := s.Guess(str)
		if res != "" {
			WriteRed(w, res)
		}

		if s.Guessed() == s.Reveal() {
			WriteGreen(w, fmt.Sprintf("\nYou have correctly guessed '%s' as the magic word!\nCongratulations!\n", s.Reveal()))
			os.Exit(1)
		}
	}
}

func Setup(r *bufio.Reader, w *bufio.Writer) *store.State {
	w.WriteString("Enter magic word: ")
	w.Flush()
	mw, _ := r.ReadString('\n')
	return store.NewState(strings.Replace(mw, "\n", "", -1))
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
