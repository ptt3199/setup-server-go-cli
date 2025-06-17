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

	// Handle multiple numbers separated by spaces
	choices := strings.Fields(choice)
	var selectedAccessories []string

	for _, choiceStr := range choices {
		choiceInt, err := strconv.Atoi(choiceStr)
		if err != nil || choiceInt < 1 || choiceInt > len(accessories) {
			fmt.Printf("Invalid choice: %s\n", choiceStr)
			continue
		}
		selectedAccessories = append(selectedAccessories, accessories[choiceInt-1])
	}

	if len(selectedAccessories) == 0 {
		fmt.Println("No valid accessories selected")
		return
	}

	// Install all selected accessories
	for _, accessory := range selectedAccessories {
		fmt.Printf("Installing %s...\n", accessory)
		runCommand(fmt.Sprintf("sudo apt-get install -y %s", accessory))
	}

	fmt.Println("Accessories installed successfully")
}
