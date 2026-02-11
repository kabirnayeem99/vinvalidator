package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	reader := bufio.NewReader(os.Stdin)
	printBanner()

	var mockCancel context.CancelFunc

	for {
		input, err := promptAndRead(reader)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			continue
		}

		switch input {
		case "":
			continue

		case "EXIT", "QUIT":
			fmt.Println("Exiting VIN Validator")
			return

		case "MOCK":
			if mockCancel != nil {
				fmt.Println("MOCK mode already running. Type STOP to stop it.")
				continue
			}

			mockCtx, cancel := context.WithCancel(ctx)
			mockCancel = cancel

			fmt.Println("MOCK mode started. Type STOP to stop it.")
			go runMockMode(mockCtx, 50*time.Millisecond)

		case "STOP":
			if mockCancel == nil {
				fmt.Println("MOCK mode is not running.")
				continue
			}
			mockCancel()
			mockCancel = nil
			fmt.Println("MOCK mode stopped.")

		default:
			validateAndPrint(VIN(input))
		}
	}
}

func printBanner() {
	fmt.Println("VIN Validator")
	fmt.Println("Type a VIN and press Enter.")
	fmt.Println("Commands: MOCK (start), STOP (stop), EXIT (quit).")
	fmt.Println("Press Ctrl+C to exit.")
	fmt.Println("----------------------")
}

func promptAndRead(r *bufio.Reader) (string, error) {
	fmt.Print("Enter VIN: ")
	line, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	return normalizeInput(line), nil
}

func normalizeInput(s string) string {
	return strings.TrimSpace(strings.ToUpper(s))
}

func validateAndPrint(v VIN) {
	if err := ValidateVIN(v); err != nil {
		fmt.Printf("❌ %s (%v)\n", v, err)
		return
	}
	fmt.Printf("✅ %s\n", v)
}

func runMockMode(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			v := VIN(GenerateMockVIN())
			validateAndPrint(v)
		}
	}
}
