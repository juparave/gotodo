package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/juparave/gotodo/internal/model"
	"github.com/juparave/gotodo/internal/store"
	"github.com/juparave/gotodo/internal/ui"
)

func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	listDoneLimit := listCmd.Int("done-limit", 3, "number of done items to show")
	listLong := listCmd.Bool("long", false, "show timestamps for done items")
	doneCmd := flag.NewFlagSet("done", flag.ExitOnError)
	initCmd := flag.NewFlagSet("init", flag.ExitOnError)

	if len(os.Args) < 2 {
		ui.RenderHelp()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "init":
		initCmd.Parse(os.Args[2:])
		cwd, _ := os.Getwd()
		path := filepath.Join(cwd, ".gotodo.json")
		s := store.NewJSONFileStore(path)
		if err := s.Init(); err != nil {
			fmt.Fprintln(os.Stderr, "init error:", err)
			os.Exit(1)
		}
		fmt.Println("Initialized:", path)

	case "add":
		addCmd.Parse(os.Args[2:])
		if addCmd.NArg() < 1 {
			fmt.Fprintln(os.Stderr, "usage: gotodo add \"task text\"")
			os.Exit(2)
		}
		// join all args as the task text so multi-word tasks work
		text := ""
		for i := 0; i < addCmd.NArg(); i++ {
			if i > 0 {
				text += " "
			}
			text += addCmd.Arg(i)
		}
		cwd, _ := os.Getwd()
		path := filepath.Join(cwd, ".gotodo.json")
		s := store.NewJSONFileStore(path)
		if err := s.Load(); err != nil {
			// try init
			if err := s.Init(); err != nil {
				fmt.Fprintln(os.Stderr, "store init error:", err)
				os.Exit(1)
			}
		}
		t := model.NewTodo(text)
		s.Add(t)
		if err := s.Save(); err != nil {
			fmt.Fprintln(os.Stderr, "save error:", err)
			os.Exit(1)
		}
		fmt.Println("Added:", t.ID)

	case "list":
		listCmd.Parse(os.Args[2:])
		cwd, _ := os.Getwd()
		path := filepath.Join(cwd, ".gotodo.json")
		s := store.NewJSONFileStore(path)
		if err := s.Load(); err != nil {
			fmt.Fprintln(os.Stderr, "no todo file found at", path)
			os.Exit(1)
		}
		items := s.All()
		ui.RenderList(items, *listDoneLimit, *listLong)

	case "done":
		doneCmd.Parse(os.Args[2:])
		if doneCmd.NArg() < 1 {
			fmt.Fprintln(os.Stderr, "usage: gotodo done <id|n>")
			os.Exit(2)
		}
		target := doneCmd.Arg(0)
		cwd, _ := os.Getwd()
		path := filepath.Join(cwd, ".gotodo.json")
		s := store.NewJSONFileStore(path)
		if err := s.Load(); err != nil {
			fmt.Fprintln(os.Stderr, "no todo file found at", path)
			os.Exit(1)
		}
		// try parse as index relative to the open list
		if idx, err := strconv.Atoi(target); err == nil {
			// build open list mapping to global indexes
			var openIdxes []int
			for i, t := range s.All() {
				if !t.Done {
					openIdxes = append(openIdxes, i)
				}
			}
			if idx < 1 || idx > len(openIdxes) {
				fmt.Fprintln(os.Stderr, "index out of range for open todos")
				os.Exit(2)
			}
			global := openIdxes[idx-1]
			if err := s.MarkDoneByIndex(global); err != nil {
				fmt.Fprintln(os.Stderr, "mark done error:", err)
				os.Exit(1)
			}
			fmt.Println("Marked done (open index):", idx)
		} else {
			if err := s.MarkDoneByID(target); err != nil {
				fmt.Fprintln(os.Stderr, "mark done error:", err)
				os.Exit(1)
			}
			fmt.Println("Marked done (id):", target)
		}

	default:
		fmt.Fprintln(os.Stderr, "unknown command", os.Args[1])
		os.Exit(2)
	}
}
