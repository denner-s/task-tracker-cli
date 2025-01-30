package main

import (
	"github.com/denner-s/task-tracker-cli/adapters/cli"
	"github.com/denner-s/task-tracker-cli/adapters/storage"
	"github.com/denner-s/task-tracker-cli/core"
	"os"
)

func main() {
	repo := storage.NewFileTaskRepository(`tasks.json`)
	service := core.NewTaskService(repo)
	cli_ := cli.NewCLI(service)
	cli_.HandleCommand(os.Args[1:])
}
