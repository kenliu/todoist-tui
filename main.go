package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kenliu/todoist-tui/tui"
	"os"
)

func main() {
	p := tea.NewProgram(tui.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
