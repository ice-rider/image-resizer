package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)


func main() {
	cfg := parseFlags()

	if cfg.InputPath == "" && !cfg.TUI {
		flag.Usage()
		fmt.Println("\nNo input provided. Use -tui for interactive mode or specify -input.")
		os.Exit(1)
	}

	if cfg.InputPath != "" && !cfg.TUI {
		proc := NewProcessor(cfg)
		logCh := make(chan string)
		go func() {
			for msg := range logCh {
				fmt.Println(msg)
			}
		}()
		proc.Process(logCh)
		fmt.Printf("Finished. Success: %d, Errors: %d\n", proc.Success, len(proc.Errors))
		if len(proc.Errors) > 0 {
			os.Exit(1)
		}
		return
	}

	p := tea.NewProgram(NewModel(cfg), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running TUI: %v\n", err)
		os.Exit(1)
	}
}