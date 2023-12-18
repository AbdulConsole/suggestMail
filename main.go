package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username) // Trim the newline character

	suggestions := suggestEmails(username)

	fmt.Println("Suggested email addresses:")
	for _, email := range suggestions {
		fmt.Println(email)
	}
}

func suggestEmails(username string) []string {
	domains := []string{"gmail.com", "yahoo.com", "outlook.com", "hotmail.com", "example.com"}
	var suggestions []string

	for _, domain := range domains {
		suggestions = append(suggestions, fmt.Sprintf("%s@%s", username, domain))
	}

	return suggestions
}
