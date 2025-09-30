package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/juparave/gotodo/internal/model"
)

var (
	titleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))
	doneStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	todoStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("10"))
)

func RenderList(items []*model.Todo) {
	fmt.Println(titleStyle.Render("Todos"))

	// separate open and done
	open := []*model.Todo{}
	done := []*model.Todo{}
	for _, t := range items {
		if t.Done {
			done = append(done, t)
		} else {
			open = append(open, t)
		}
	}

	fmt.Println("\nOpen:")
	if len(open) == 0 {
		fmt.Println("  (no open todos)")
	}
	for i, t := range open {
		fmt.Printf(" %2d. %s\n", i+1, todoStyle.Render(t.Text))
	}

	fmt.Println("\nDone:")
	if len(done) == 0 {
		fmt.Println("  (no done todos)")
	}
	for i, t := range done {
		fmt.Printf(" %2d. %s\n", i+1, doneStyle.Render(t.Text))
	}
}
