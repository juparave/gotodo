package ui

import (
	"fmt"
	"sort"

	"github.com/charmbracelet/lipgloss"
	"github.com/juparave/gotodo/internal/model"
)

var (
	titleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))
	doneStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	todoStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("10"))
)

func RenderList(items []*model.Todo, doneLimit int, long bool) {
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
		return
	}
	// show only the last doneLimit done todos (most recent first)
	if doneLimit <= 0 {
		doneLimit = 3
	}
	sort.Slice(done, func(i, j int) bool {
		return done[i].DoneAt.After(done[j].DoneAt)
	})
	limit := doneLimit
	if len(done) < limit {
		limit = len(done)
	}
	for i := 0; i < limit; i++ {
		t := done[i]
		if long {
			fmt.Printf(" %2d. %s  (%s)\n", i+1, doneStyle.Render(t.Text), t.DoneAt.Local().Format("2006-01-02 15:04:05"))
		} else {
			fmt.Printf(" %2d. %s\n", i+1, doneStyle.Render(t.Text))
		}
	}
}
