package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// State is
type State uint32

const (
	// locked state initialized to 0
	locked State = iota

	// unlocked state = 1
	unlocked

	coinCommand = "coin"
	pushCommand = "push"
)

func main() {
	state := locked
	// prompt user to enter the command to the turnstile
	reader := bufio.NewReader(os.Stdin)

	prompt(locked)

	for {
		cmd, err := reader.ReadString('\n')

		if err != nil {
			log.Fatalln(err)
			return
		}

		state = step(state, strings.TrimSpace(cmd))
	}
}

func step(state State, cmd string) State {
	if cmd != coinCommand && cmd != pushCommand {
		fmt.Println("Unknown command, please try again.")
		return state
	}

	switch state {
	case locked:
		if cmd == coinCommand {
			fmt.Println("Unlocked, ready to pass through.")

			state = unlocked
		} else {
			fmt.Println("Access denied! Please put a coin first.")
		}
	case unlocked:
		//  if the state is unlocked, push command is allowed
		if cmd == coinCommand {
			fmt.Println("The gate is already open. Don't waste your coins.")
		} else {
			fmt.Println("You can pass. The gate is unlocked.")
			state = locked
		}
	}

	return state
}

func prompt(state State) {
	m := map[State]string{
		locked:   "Locked",
		unlocked: "Unlocked",
	}

	fmt.Printf("The current state is [%s], please input the command [coin | push]\n", m[state])
}
