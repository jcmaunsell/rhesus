package main

import (
	"fmt"
	"github.com/jcmaunsell/rhesus/repl"
	"os"
	"os/user"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		panic(fmt.Errorf("could not get username: %w", err))
	}
	fmt.Printf("Hi %s! Welcome to Rhesus.\n", usr.Username)
	fmt.Println("You can enter Monkey commands here.")
	repl.New(os.Stderr).Start(os.Stdin, os.Stdout)
}
