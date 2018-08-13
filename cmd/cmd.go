package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ramfox/fullmoon/game"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fullmoon",
	Short: "Guess the magic word before the next full moon or you will turn into a blood thirsty monster!",
	Long: `   A silly game created by @ramfox to strenghten her go skills
    and learn how to use libp2p. Basically hangman.`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		writer := bufio.NewWriter(os.Stdout)

		gameState := game.Setup(reader, writer)
		writer.WriteString(("\033[H\033[2J"))
		writer.Flush()
		game.Play(gameState, reader, writer)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
