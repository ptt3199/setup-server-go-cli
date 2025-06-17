// Install accessories

package cmd

import (
	"fmt"
	"strconv"
	"strings"
)

func installAccessories() {
	fmt.Println("Choose accessories to install:")
	accessories := []string{
		"tmux",
		"git",
		"curl",
		"htop",
		"ufw",
		"zsh",
	}

	for i, accessory := range accessories {
		fmt.Printf("%d. %s\n", i+1, accessory)
	}

	fmt.Print("Enter the number of the accessory to install (e.g. 1, 2, 3, etc.) or 'all' to install all accessories: ")
	var choice string
	fmt.Scanln(&choice)
	choice = strings.TrimSpace(choice)

	if choice == "all" {
		for _, accessory := range accessories {
			runCommand(fmt.Sprintf("sudo apt-get install -y %s", accessory))
		}
		return
	}

	choiceInt, err := strconv.Atoi(choice)
	if err != nil {
		fmt.Println("Invalid choice")
		return
	}

	accessory := accessories[choiceInt-1]
	runCommand(fmt.Sprintf("sudo apt-get install -y %s", accessory))

	fmt.Println("Accessories installed successfully")
}
