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
	for i, t := range items {
		idx := i + 1
		if t.Done {
			fmt.Printf("%2d. %s\n", idx, doneStyle.Render(t.Text))
		} else {
			fmt.Printf("%2d. %s\n", idx, todoStyle.Render(t.Text))
		}
	}
}
