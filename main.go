package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sig
		fmt.Println("\nExiting VIN Validator")
		os.Exit(0)
	}()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("VIN Validator")
	fmt.Println("Press Ctrl+C to exit.")
	fmt.Println("----------------------")

	for {
		fmt.Print("Enter VIN: ")

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Printf("Error reading input: %v\n", err.Error())
			continue
		}

		v := VIN(strings.TrimSpace(strings.ToUpper(input)))

		if err := ValidateVIN(v); err != nil {
			fmt.Println("Invalid VIN:", err.Error())
			continue
		}

		fmt.Println("✅ Valid VIN")

	}

}
