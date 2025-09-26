package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/kubermatic/go-interviews/patterns/pkg/utils"
)

// CLI demonstrates command-line argument parsing and subcommands
// This is commonly asked about in Go interviews

type CLI struct {
	verbose bool
	timeout time.Duration
}

func main() {
	cli := &CLI{}

	// Define global flags
	flag.BoolVar(&cli.verbose, "verbose", false, "Enable verbose output")
	flag.DurationVar(&cli.timeout, "timeout", 30*time.Second, "Timeout for operations")

	// Parse flags
	flag.Parse()

	// Get subcommand
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	subcommand := os.Args[1]

	// Remove the subcommand from os.Args so flag.Parse() works for subcommands
	os.Args = os.Args[1:]

	ctx, cancel := context.WithTimeout(context.Background(), cli.timeout)
	defer cancel()

	switch subcommand {
	case "validate":
		cli.runValidate(ctx)
	case "retry":
		cli.runRetry(ctx)
	case "server":
		cli.runServer(ctx)
	default:
		fmt.Printf("Unknown subcommand: %s\n", subcommand)
		printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) runValidate(ctx context.Context) {
	var input string
	flag.StringVar(&input, "input", "", "Input to validate")
	flag.Parse()

	if input == "" {
		fmt.Println("Error: input is required")
		os.Exit(1)
	}

	if cli.verbose {
		fmt.Printf("Validating input: %s\n", input)
	}

	// Use the validation utility
	validator := &utils.StringValidator{
		Field:    "input",
		Value:    input,
		MinLen:   1,
		MaxLen:   100,
		Required: true,
	}

	if err := validator.Validate(); err != nil {
		fmt.Printf("Validation failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Validation passed!")
}

func (cli *CLI) runRetry(ctx context.Context) {
	var maxAttempts int
	var baseDelay time.Duration

	flag.IntVar(&maxAttempts, "max-attempts", 3, "Maximum number of retry attempts")
	flag.DurationVar(&baseDelay, "base-delay", 100*time.Millisecond, "Base delay between retries")
	flag.Parse()

	if cli.verbose {
		fmt.Printf("Running retry with %d max attempts and %v base delay\n", maxAttempts, baseDelay)
	}

	// Create retry config
	config := &utils.RetryConfig{
		MaxAttempts: maxAttempts,
		BaseDelay:   baseDelay,
		MaxDelay:    5 * time.Second,
		Multiplier:  2.0,
		Jitter:      true,
	}

	// Simulate a function that fails a few times then succeeds
	attempt := 0
	fn := func() error {
		attempt++
		if cli.verbose {
			fmt.Printf("Attempt %d\n", attempt)
		}

		if attempt < 3 {
			return fmt.Errorf("simulated failure on attempt %d", attempt)
		}

		fmt.Println("Success!")
		return nil
	}

	if err := utils.Retry(ctx, config, fn); err != nil {
		fmt.Printf("Retry failed: %v\n", err)
		os.Exit(1)
	}
}

func (cli *CLI) runServer(ctx context.Context) {
	var host string
	var port int

	flag.StringVar(&host, "host", "localhost", "Server host")
	flag.IntVar(&port, "port", 8080, "Server port")
	flag.Parse()

	if cli.verbose {
		fmt.Printf("Starting server on %s:%d\n", host, port)
	}

	// Simulate server startup
	fmt.Printf("Server would start on %s:%d\n", host, port)
	fmt.Println("Press Ctrl+C to stop")

	// Wait for context cancellation
	<-ctx.Done()
	fmt.Println("Server stopped")
}

func printUsage() {
	fmt.Println("Usage: cli <subcommand> [flags]")
	fmt.Println()
	fmt.Println("Subcommands:")
	fmt.Println("  validate    Validate input")
	fmt.Println("  retry       Demonstrate retry logic")
	fmt.Println("  server      Start a server")
	fmt.Println()
	fmt.Println("Global flags:")
	fmt.Println("  -verbose    Enable verbose output")
	fmt.Println("  -timeout    Timeout for operations (default: 30s)")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  cli validate -input 'hello world'")
	fmt.Println("  cli retry -max-attempts 5 -base-delay 200ms")
	fmt.Println("  cli server -host 0.0.0.0 -port 9090")
}
