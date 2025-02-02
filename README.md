# Task Tracker (JSON Edition - Go) 📋

[![Go Version](https://img.shields.io/badge/Go-1.20%2B-blue)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

Uma ferramenta CLI minimalista para gerenciar tarefas com armazenamento em JSON. Desenvolvida em Go seguindo a **Arquitetura Hexagonal**.

---

## 📋 Tabela de Conteúdos
- [Descrição](#-descrição)
- [Instalação](#-instalação)
- [Uso](#-uso)
- [Funcionalidades](#-funcionalidades)
- [Exemplos](#-exemplos)
- [Tecnologias](#-tecnologias)
- [Testes](#-testes)
- [Contribuição](#-contribuição)
- [Licença](#-licença)
- [Agradecimentos](#-agradecimentos)
- [Contato](#-contato)

---

## 🚀 Descrição
O **Task Tracker CLI** permite que você gerencie tarefas diárias diretamente do terminal.  
Recursos principais:
- Armazena tarefas em um arquivo JSON (`tasks.json`)
- Suporta status: `todo`, `in-progress`, `done`
- Operações CRUD (criar, ler, atualizar, excluir)
- Fácil integração com automações e scripts

---

## 🔧 Instalação

1. **Clone o repositório**:
   ```bash
   git clone https://github.com/denner-s/task-tracker.git
   cd task-tracker
   ```

2. **Instale as dependências**:
   ```bash
   go mod tidy
   ```

3. **Compile o projeto**:
   ```bash
   go build -o task-cli main.go
   ```

---

## 🖥️ Uso

### Comandos Disponíveis:
| Comando                | Descrição                              | Exemplo                          |
|------------------------|---------------------------------------|----------------------------------|
| `add "descrição"`      | Adiciona uma nova tarefa              | `./task-cli add "Ler livro"`     |
| `update <ID> "descrição"` | Atualiza a descrição da tarefa    | `./task-cli update 1 "Ler 2 capítulos"` |
| `delete <ID>`          | Exclui uma tarefa                     | `./task-cli delete 1`            |
| `mark-in-progress <ID>`| Marca tarefa como "em andamento"      | `./task-cli mark-in-progress 1`  |
| `mark-done <ID>`       | Marca tarefa como "concluída"         | `./task-cli mark-done 1`         |
| `list`                 | Lista todas as tarefas                | `./task-cli list`                |
| `list <status>`        | Filtra por status (`todo`, `in-progress`, `done`) | `./task-cli list done` |

---

## ⚙️ Funcionalidades
- ✅ **CRUD Completo**: Adicione, atualize, exclua e liste tarefas
- ✅ **Status Dinâmico**: Transição entre `todo` → `in-progress` → `done`
- ✅ **Armazenamento Persistente**: Dados salvos automaticamente em `tasks.json`
- ✅ **IDs Automáticos**: Geração incremental de IDs únicos
- ✅ **Timestamps**: Registro de criação e última atualização
- 🛠️ **Tratamento de Erros**: Mensagens claras para operações inválidas

---

## 📊 Exemplos

### Adicionar Tarefa
```bash
./task-cli add "Estudar Go"
# Saída: Tarefa adicionada com sucesso (ID: 1)
```

### Listar Tarefas Pendentes
```bash
./task-cli list todo
# Saída:
# ID: 1, Descrição: Estudar Go, Status: todo, Criado em: 2023-10-01T10:00:00Z, Atualizado em: 2023-10-01T10:00:00Z
```

### Marcar como Concluída
```bash
./task-cli mark-done 1
# Saída: Tarefa marcada como concluída (ID: 1)
```

---

## 🛠️ Tecnologias
- **Linguagem**: [Go 1.20+](https://golang.org/)
- **Armazenamento**: JSON (via `encoding/json`)
- **Arquitetura**: Hexagonal (Core/Ports/Adapters)
- **Bibliotecas**: `os`, `time`, `strconv`

---

## 🧪 Testes
Execute os testes unitários com:
```bash
go test -v ./...
```
*(Nota: Testes detalhados serão implementados em futuras versões)*

---

## 🤝 Contribuição
1. Faça um **fork** do projeto.
2. Crie uma branch: `git checkout -b feat/nova-feature`.
3. Commit suas mudanças: `git commit -m 'feat: Minha nova feature'`.
4. Push para a branch: `git push origin feat/nova-feature`.
5. Abra um **Pull Request**.

**Reportar Bugs**: [Issues](https://github.com/seu-usuario/task-tracker/issues)

---

## 📜 Licença
Distribuído sob a licença MIT. Veja [LICENSE](LICENSE) para detalhes.

---

## 🙌 Agradecimentos
- Roadmap.sh (https://roadmap.sh/projects/task-tracker)

---

## 📞 Contato
**Denner Sousa**  
- Email: denner.sousa@outlook.com
- GitHub: [@denner-s](https://github.com/denner-s)  


