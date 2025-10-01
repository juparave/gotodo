# Roadmap

This document outlines planned features and improvements for future releases of `gotodo`.

## Planned Features

### High Priority

- [ ] **TUI Mode**: Interactive terminal interface using Bubbletea
  - `gotodo tui` - Launch interactive mode
  - Visual todo list with keyboard navigation
  - In-place editing of todo text
  - Drag-and-drop reordering
  - Interactive priority setting
  - Due date picker
  - Tag management interface
  - Bulk selection and operations (mark done, delete, tag)
  - Search and filtering
  - Keyboard shortcuts for power users

- [ ] **Edit Command**: Allow editing existing todo text
  - `gotodo edit <id|index> "new description"`
  - Interactive editing mode

- [ ] **Reorder Command**: Change the order of todos
  - `gotodo reorder <from-index> <to-index>`
  - Drag-and-drop style reordering

- [ ] **Priority Levels**: Add priority system to todos
  - Low, Medium, High priority levels
  - Sort by priority in list view
  - Priority indicators in display

### Medium Priority

- [ ] **Due Dates**: Add deadline support
  - `gotodo add "task" --due 2025-10-15`
  - Show overdue items prominently
  - Sort by due date

- [ ] **Categories/Tags**: Organize todos by categories
  - `gotodo add "task" --tag work`
  - Filter by tags: `gotodo list --tag work`
  - Multiple tags per todo

- [ ] **Bulk Operations**: Handle multiple todos at once
  - `gotodo done 1 2 3` (mark multiple as done)
  - `gotodo rm --all-done` (remove all completed todos)

### Low Priority

- [ ] **Subtasks**: Break down todos into smaller tasks
  - Nested todo structure
  - Progress tracking for parent tasks

- [ ] **Recurring Todos**: Automatically create repeating tasks
  - Daily, weekly, monthly recurrence
  - `gotodo add "daily standup" --repeat daily`

- [ ] **Search and Filter**: Advanced filtering options
  - `gotodo search "keyword"`
  - Filter by date ranges, completion status

- [ ] **Export/Import**: Data portability
  - Export to JSON, CSV, Markdown
  - Import from other todo apps

- [ ] **Web Interface**: Optional web UI
  - Serve local web interface
  - REST API for integrations

- [ ] **Sync Features**: Cloud synchronization
  - GitHub integration
  - Google Drive sync (original concept)
  - Generic cloud storage

### Technical Improvements

- [ ] **Configuration File**: User preferences
  - Default list location
  - Color themes
  - Default behaviors

- [ ] **Shell Completion**: Auto-completion for commands
  - Bash, Zsh, Fish completion scripts
  - Dynamic completion for todo IDs

- [ ] **Testing**: Expand test coverage
  - Integration tests for CLI commands
  - Cross-platform testing
  - Performance benchmarks

- [ ] **CI/CD**: Automated releases
  - GitHub Actions workflows
  - Automated cross-platform builds
  - Release automation with GoReleaser

### Version Planning

- **v0.1.0**: Edit, reorder, priority levels
- **v0.2.0**: Due dates, categories, bulk operations
- **v0.3.0**: Subtasks, recurring todos
- **v1.0.0**: Stable release with core features complete

## Contributing

Want to help implement these features? See our [Contributing Guide](CONTRIBUTING.md) and check out the [Issues](https://github.com/juparave/gotodo/issues) for specific tasks.

## Feedback

Have ideas for new features? [Open an issue](https://github.com/juparave/gotodo/issues/new) or start a discussion!