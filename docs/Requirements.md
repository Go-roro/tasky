# 📝 `tasky`: CLI Task Manager — Requirements Document

## 📌 Overview

`tasky` is a command-line tool that helps users manage tasks efficiently. It supports creating, viewing, completing, deleting, and exporting tasks. All data is stored locally in a file (`tasks.json`), and users interact with the tool via commands and flags.

## 🎯 Goals

- Provide a fast and intuitive CLI interface for managing tasks.
- Ensure tasks are saved and loaded automatically between runs.
- Offer basic prioritization, deadline, and filtering capabilities.
- Be easily extensible and testable.

## 📂 Data Model

Go struct definition for a task:
```go
type Task struct {  
    ID        int          // Auto-incremented unique identifier  
    Title     string       // Short description  
    Done      bool         // Completion status  
    Due       *time.Time   // Optional deadline  
    Priority  int          // 1 = High, 2 = Medium, 3 = Low  
    CreatedAt time.Time    // Timestamp when task was created  
}
```

## 🔧 Core Features

### 1. `add` - Add a new task

**Syntax**  
tasky add --title "Buy groceries" --priority 2 --due "2025-07-10 18:00"

**Behavior**
- Title is required.
- Due date is optional (`YYYY-MM-DD HH:MM` format).
- Priority defaults to 2 (Medium) if not specified.
- Automatically assigns an incremental ID.

---

### 2. `list` - View all tasks

**Syntax**  
tasky list

**Optional Flags**
- `--all`: Include completed tasks
- `--sort priority|due|created`: Sort by given field
- `--filter "due:today"`: Filter to tasks due today
- `--json`: Output as JSON

```
**Example Output**  
[ ] 1. Buy groceries       🟡 Medium  Due: 2025-07-10 18:00  
[✓] 2. Fix laptop          🔴 High    Due: 2025-07-09 09:00  
[ ] 3. Call Alice          🟢 Low     No deadline
```
---

### 3. `done` - Mark a task as complete

**Syntax**  
tasky done --id 1

---

### 4. `delete` - Delete a task

**Syntax**  
tasky delete --id 2

---

### 5. `export` - Export tasks as JSON

**Syntax**  
tasky export --out tasks_export.json

---

### 6. `help` - Show help screen

Shows available commands and examples.

---

## 💡 Advanced Features (Optional/Stretch Goals)

### ⏰ Reminder for Due Soon
- `tasky list --filter "due:soon"` will show tasks due within 24 hours.

### 📦 Tagging Support
- Add tags like `#work`, `#home`
- Support `--tag` flag when adding and filtering

### 🔄 Recurring Tasks
- Option to repeat daily/weekly/monthly

---

## 💾 File Storage

- All tasks are stored in a JSON file in the working directory: `tasks.json`
- Structure is a JSON array of `Task` objects
- The program loads this file on start and saves on every change

---

## ✅ Test Cases & Sample Flows

### Add Task
tasky add --title "Prepare report" --due "2025-07-09 17:00" --priority 1

### Mark as Done
tasky done --id 1

### List Tasks
tasky list

### Export to JSON
tasky export --out mytasks.json

---

## ✨ Bonus UX Ideas
- Emoji indicators for priority (🔴 High, 🟡 Medium, 🟢 Low)
- Human-friendly due messages (e.g., “in 3h”, “due tomorrow”)
- Show `✓` for completed tasks
