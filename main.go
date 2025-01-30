package main

import (
	"github.com/denner-s/task-tracker-cli/adapters/cli"
	"github.com/denner-s/task-tracker-cli/adapters/storage"
	"github.com/denner-s/task-tracker-cli/core"
	"os"
)

func main() {
	repo := storage.NewFileRepository("tasks.json")
	service := core.NewTaskService(repo)
	cliApp := cli.NewCLI(service)
	cliApp.HandleCommand(os.Args[1:])
}
