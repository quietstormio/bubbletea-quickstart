package main

import (
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Variables that hold styles for the counter and title
var (
	counterStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFA500"))
	titleStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFFFFF"))
)

// Define the model
type model struct {
	counter int
}

// Give the model an initial value
func initialModel() model {
	return model{
		counter: 0,
	}
}

// Init command to run when the program starts
func (m model) Init() tea.Cmd {
	return nil
}

// Function to clear the screen when the user quits the program
func clearScreen() {
	cmd := exec.Command("clear") // Use "cls" for Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			clearScreen()
			return m, tea.Quit
		case "up", "k":
			m.counter++
		case "down", "j":
			m.counter--
		}
	}
	return m, nil
}

func (m model) View() string {
	title := titleStyle.Render("Quiet Storm Counter")
	counter := counterStyle.Render(fmt.Sprintf("Count: %d", m.counter))
	return fmt.Sprintf("%s\n\n%s", title, counter)
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Something went wrong: %s", err)
	}
}
