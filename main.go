package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
		Use: "hellow",
		Short: "hellow CLI",
		Long: "A Command Line Application that says hellow",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to hellow!\n")
		},
}

var hellowCmd = &cobra.Command{
	Use: "sayhellow",
	Short: "Say hello world",
	Long: "Greet the user with hellow world!",
	Run: func(cmd *cobra.Command, args []string) {
		var username string
		fmt.Println("What is your name?")
		fmt.Scanf("%s", &username)
		fmt.Printf("\n> Hi %s, Hello World!\n", username)
	},
}

func init()  {
		rootCmd.AddCommand(hellowCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
			fmt.Println(err)
			os.Exit(1)
	}
}

func main(){
	Execute()
}
