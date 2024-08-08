package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/manifoldco/promptui"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")
	if scanner.Scan(){
		input := scanner.Text()
		fmt.Printf("Greetings, %s\nWelcome to this app!\n", input)
	}
	fmt.Println()
	displayPrompts()
	// getUser()
	// getCurrentDir()
	// listFiles()
}

func getUser(){
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Your user name is %s\n", currentUser.Username)
}


func getCurrentDir(){
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return 
	}
	fmt.Printf("Current directory is %s\n", currentDir)
}

func listFiles(){
	dir := "."

	files, err := os.ReadDir(dir)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Listing files in current directory: ")
	for _, file := range files{
		fmt.Println(file.Name())
	}
}

func displayPrompts(){
	prompts := []string{
		"Find out your local user name",
		"Display your current working directory",
		"Display all the files within your directory",
	}

	promptFuncs := map[string]func(){
		prompts[0]: getUser,
		prompts[1]: getCurrentDir,
		prompts[2]: listFiles,
	}

	prompt := promptui.Select{
		Label: "--- Choose from the below options ---",
		Items: prompts,
	}

	index, result, err := prompt.Run()
	if err != nil{
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	fmt.Printf("You have selected prompt %d: %s\n", index+1, result)

	if function, found := promptFuncs[result]; found{
		function()
	} else{
		fmt.Println("No function found for the selected prompt")
	}
}