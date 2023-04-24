/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "my test cli",
	Short: "this is just a test cli",
	Long:  `this is just a test clik long long long `,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		amir, err := cmd.Flags().GetBool("amir")
		marat, err := cmd.Flags().GetBool("marat")
		manas, err := cmd.Flags().GetBool("manas")
		uzair, err := cmd.Flags().GetBool("uzair")
		dmitriy, err := cmd.Flags().GetBool("dmitriy")
		florian, err := cmd.Flags().GetBool("florian")
		wina, err := cmd.Flags().GetBool("wina")
		wade, err := cmd.Flags().GetBool("wade")
		matias, err := cmd.Flags().GetBool("matias")
		maria, err := cmd.Flags().GetBool("maria")
		smith, err := cmd.Flags().GetBool("smith")
		andre, err := cmd.Flags().GetBool("andre")
		milad, err := cmd.Flags().GetBool("milad")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if amir {
			fmt.Println("react native is my only love 😑")
			return
		}
		if marat {
			fmt.Println("ping pong 😏")
			return
		}
		if manas {
			fmt.Println("karim 🐣")
			return
		}
		if uzair {
			fmt.Println("aba waghoom zo")
			return
		}
		if dmitriy {
			fmt.Println("don't hesitate 🤠")
			return
		}
		if florian {
			fmt.Println("Flow SharedFlow StateFlow 🤟")
			return
		}
		if wade {
			fmt.Println("Hallowwwww, shalalalaaaaaaa 👀")
			return
		}
		if wina {
			fmt.Println("berlin ): ")
			return
		}
		if matias {
			fmt.Println("beach ")
			return
		}
		if maria {
			fmt.Println("mariaaaaaaaaaaaaaa ")
			return
		}
		if smith {
			fmt.Println("baby ... ")
			return
		}
		if andre {
			fmt.Println("404 ... ")
			return
		}
		if milad {
			fmt.Println("come onnnnnn ... ")
			return
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.myapp.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "a", false, "nothing")
	rootCmd.Flags().BoolP("amir", "b", false, "nothing")
	rootCmd.Flags().BoolP("marat", "c", false, "nothing")
	rootCmd.Flags().BoolP("manas", "d", false, "nothing")
	rootCmd.Flags().BoolP("uzair", "e", false, "nothing")
	rootCmd.Flags().BoolP("dmitriy", "f", false, "nothing")
	rootCmd.Flags().BoolP("florian", "g", false, "nothing")
	rootCmd.Flags().BoolP("wade", "z", false, "nothing")
	rootCmd.Flags().BoolP("wina", "i", false, "nothing")
	rootCmd.Flags().BoolP("matias", "j", false, "nothing")
	rootCmd.Flags().BoolP("maria", "k", false, "nothing")
	rootCmd.Flags().BoolP("smith", "l", false, "nothing")
	rootCmd.Flags().BoolP("andre", "m", false, "nothing")
	rootCmd.Flags().BoolP("milad", "n", false, "nothing")
}
