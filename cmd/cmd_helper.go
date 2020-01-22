package cmd

import (
	"fmt"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func promptForPassword() string {
	fmt.Println("Please enter your password:")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	return string(bytePassword)
}
