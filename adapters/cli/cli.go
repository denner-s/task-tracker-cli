package cli

import (
	"fmt"
	"github.com/denner-s/task-tracker-cli/core"
	"os"
	"strconv"
	"time"
)

type CLI struct {
	service *core.TaskService
}

func NewCLI(service *core.TaskService) *CLI {
	return &CLI{service: service}
}

func (c *CLI) HandleCommand(args []string) {
	if len(args) == 0 {
		fmt.Println("Comando inválido. Use add, update, delete, mark-in-progress, mark-done ou list")
		os.Exit(1)
	}
	switch args[0] {
	case "add":
		c.handleAdd(args[1:])
	case "update":
		c.handleUpdate(args[1:])
	case "delete":
		c.handleDelete(args[1:])
	case "mark-in-progress":
		c.handleMarkInProgress(args[1:])
	case "mark-done":
		c.handleMarkDone(args[1:])
	case "list":
		c.handleList(args[1:])
	default:
		fmt.Println("Comando inválido")
		os.Exit(1)
	}
}

func (c *CLI) handleAdd(args []string) {
	if len(args) < 1 {
		fmt.Println("Uso: add <descrição>")
		os.Exit(1)
	}
	task, err := c.service.AddTask(args[0])
	if err != nil {
		fmt.Printf("Erro ao adicionar tarefa: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Tarefa adicionada com sucesso (ID: %d)\n", task.ID)
}

func (c *CLI) handleUpdate(args []string) {
	if len(args) < 2 {
		fmt.Println("Uso: update <ID> <nova descrição>")
		os.Exit(1)
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("ID inválido: %v\n", err)
		os.Exit(1)
	}
	task, err := c.service.UpdateTask(id, args[1])
	if err != nil {
		fmt.Printf("Erro ao atualizar tarefa: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Tarefa atualizada com sucesso (ID: %d)\n", task.ID)
}

func (c *CLI) handleDelete(args []string) {
	if len(args) < 1 {
		fmt.Println("Uso: delete <ID>")
		os.Exit(1)
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("ID inválido: %v\n", err)
		os.Exit(1)
	}
	if err := c.service.DeleteTask(id); err != nil {
		fmt.Printf("Erro ao excluir tarefa: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Tarefa excluída com sucesso (ID: %d)\n", id)
}

func (c *CLI) handleMarkInProgress(args []string) {
	c.handleMarkStatus(args, "em andamento", core.StatusInProgress)
}

func (c *CLI) handleMarkDone(args []string) {
	c.handleMarkStatus(args, "concluída", core.StatusDone)
}

func (c *CLI) handleMarkStatus(args []string, status string, targetStatus core.TaskStatus) {
	if len(args) < 1 {
		fmt.Printf("Uso: mark-%s <ID>\n", status)
		os.Exit(1)
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("ID inválido: %v\n", err)
		os.Exit(1)
	}
	task, err := c.service.UpdateStatus(id, targetStatus)
	if err != nil {
		fmt.Printf("Erro ao marcar tarefa como %s: %v\n", status, err)
		os.Exit(1)
	}
	fmt.Printf("Tarefa marcada como %s (ID: %d)\n", status, task.ID)
}

func (c *CLI) handleList(args []string) {
	var filter string
	if len(args) > 0 {
		filter = args[0]
	}
	tasks, err := c.service.ListTasks(filter)
	if err != nil {
		fmt.Printf("Erro ao listar tarefas: %v\n", err)
		os.Exit(1)
	}
	if len(tasks) == 0 {
		fmt.Println("Nenhuma tarefa encontrada")
		return
	}
	for _, task := range tasks {
		fmt.Printf("ID: %d, Descrição: %s, Status: %s, Criado em: %s, Atualizado em: %s\n",
			task.ID, task.Description, task.Status, task.CreatedAt.Format(time.RFC3339), task.UpdatedAt.Format(time.RFC3339))
	}
}
