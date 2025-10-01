package ui

import (
	"fmt"
	"sort"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/juparave/gotodo/internal/model"
)

var (
	titleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))
	doneStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	todoStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("10"))
)

// formatDuration formats a duration as "Xd Yh ago" for human readability
func formatDuration(d time.Duration) string {
	days := int(d.Hours() / 24)
	hours := int(d.Hours()) % 24
	minutes := int(d.Minutes()) % 60

	if days > 0 {
		if hours > 0 {
			return fmt.Sprintf("%dd %dh ago", days, hours)
		}
		return fmt.Sprintf("%dd ago", days)
	}
	if hours > 0 {
		if minutes > 0 {
			return fmt.Sprintf("%dh %dm ago", hours, minutes)
		}
		return fmt.Sprintf("%dh ago", hours)
	}
	if minutes > 0 {
		return fmt.Sprintf("%dm ago", minutes)
	}
	return "just now"
}

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
		if long {
			fmt.Printf(" %2d. %s  (%s)\n", i+1, todoStyle.Render(t.Text), formatDuration(time.Since(t.CreatedAt)))
		} else {
			fmt.Printf(" %2d. %s\n", i+1, todoStyle.Render(t.Text))
		}
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

// RenderHelp prints a short, styled usage hint when the user runs the CLI
// without arguments.
func RenderHelp() {
	// small helper styles
	desc := lipgloss.NewStyle().Foreground(lipgloss.Color("7"))
	cmd := lipgloss.NewStyle().Foreground(lipgloss.Color("12")).Bold(true)
	word := lipgloss.NewStyle().Foreground(lipgloss.Color("10"))

	fmt.Println(titleStyle.Render("gotodo — simple todo CLI"))
	fmt.Println(desc.Render("A tiny filesystem-aware todo manager that stores lists as JSON files."))
	fmt.Println()
	fmt.Println(cmd.Render("Usage:") + " gotodo <command> [args]")
	fmt.Println()
	fmt.Println(cmd.Render("Commands:"))
	fmt.Println(word.Render("  init") + "    Initialize a .gotodo.json in the current directory")
	fmt.Println(word.Render("  add <task>") + "  Add a todo (multi-word allowed)")
	fmt.Println(word.Render("  list") + "   Show todos (flags: --done-limit, --long, --file)")
	fmt.Println(word.Render("  done <id|n>") + "  Mark a todo done (n = open index, flags: --file)")
	fmt.Println(word.Render("  rm <id|n>") + "    Remove a todo (n = open index, flags: --force, --file)")
	fmt.Println()
	fmt.Println(cmd.Render("Examples:"))
	fmt.Println("  " + todoStyle.Render("gotodo add \"Write README\"") + "  — add a todo")
	fmt.Println("  " + todoStyle.Render("gotodo list --long") + "  — show todos with timestamps")
	fmt.Println()
	fmt.Println(desc.Render("Note: gotodo uses a .gotodo.json file at the Git repository root if inside a repo, otherwise in the current directory."))
}
