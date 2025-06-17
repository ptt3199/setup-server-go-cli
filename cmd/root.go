package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func runCommand(cmd string) error {
	fmt.Printf("Running: %s\n", cmd)
	out, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		log.Fatalf("❌ Error running: %s\n%s", cmd, out)
	}
	fmt.Printf("✅ Done: %s\n", cmd)
	return nil
}

var rootCmd = &cobra.Command{
	Use:   "setup-server",
	Short: "A CLI tool to setup your server with essential tools",
	Long: `Setup Server CLI allows you to install and configure essential tools
for your server either all at once or by selecting individual components.`,
	Run: func(cmd *cobra.Command, args []string) {
		showInteractiveMenu()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func showInteractiveMenu() {
	fmt.Println("🚀 Welcome to Server Setup CLI!")
	fmt.Println("Choose an option:")
	fmt.Println("1. Install accessories")
	fmt.Println("2. Install Docker")
	fmt.Println("3. Exit")
	fmt.Print("Enter your choice (1-3): ")

	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		installAccessories()
	case "2":
		installDocker()
	case "3":
		fmt.Println("👋 Goodbye!")
		return
	default:
		fmt.Println("❌ Invalid choice. Please try again.")
		showInteractiveMenu()
	}
}
