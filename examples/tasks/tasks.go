package main

import (
	"bufio"
	"fmt"
	"net/http"
	"sync"

	"github.com/catcher3/cegla"
	"github.com/catcher3/cegla/atr"
	"github.com/catcher3/cegla/atr/css/tw"
	"github.com/catcher3/cegla/atr/htmx"
	"github.com/catcher3/cegla/tags"
	"github.com/catcher3/cegla/ui"
	"github.com/gin-gonic/gin"
)

// --- 1. State / Database Simulation ---

type Task struct {
	ID    int
	Title string
}

var (
	mu     sync.Mutex
	tasks  []Task
	nextID = 1
)

// --- 2. Gin Helper for Cegla ---

// RenderHTML bridges Gin and Cegla by writing directly to the HTTP response stream.
func RenderHTML(c *gin.Context, status int, node cegla.Node) {
	c.Status(status)
	c.Header("Content-Type", "text/html; charset=utf-8")

	w := bufio.NewWriter(c.Writer)

	// Ensure DOCTYPE is written if rendering full HTML (optional, depends on your HTML tag implementation)
	if node.Name() == "html" {
		w.WriteString("<!DOCTYPE html>\n")
	}

	if err := node.Render(c.Request.Context(), w); err != nil {
		fmt.Printf("cegla render error: %v\n", err)
	}

	// Crucial: flush the buffer to the client!
	w.Flush()
}

// --- 3. UI Components ---

// Layout is the main HTML wrapper with Tailwind CDN and HTMX included.
// Layout is the main HTML wrapper with Tailwind CDN and HTMX included.
func Layout(title string, content cegla.FlowContent) cegla.HTML { // <-- Изменили cegla.Node на cegla.FlowContent
	return cegla.HTML{
		atr.Lang("en"),
		cegla.Head{
			tags.Title{cegla.Text(title)},
			tags.Script{atr.Src("https://unpkg.com/htmx.org@1.9.10")},
			tags.Script{atr.Src("https://cdn.tailwindcss.com")}, // Using CDN for demo
		},
		cegla.Body{
			tw.Class("bg-gray-100 text-gray-900 min-h-screen p-8"),
			ui.AvatarGroup{
				Class: "mt-10",
				Children: []cegla.FlowContent{
					ui.Avatar{
						Source:         "assets/img/user.jpeg",
						ContainerClass: "w-12 rounded-full",
					},
					ui.Avatar{
						Source:           "assets/img/user.jpeg",
						ContainerClass:   "w-12 rounded-full",
						Placeholder:      "+5",
						PlaceholderClass: "bg-neutral text-neutral-content",
					},
				},
			},
			content,
		},
	}
}

// TaskForm is the input form to create a new task.
func TaskForm() tags.Form {
	return tags.Form{
		tw.Class("flex gap-2 mb-6"),

		// HTMX attributes
		htmx.Post("/tasks"),
		htmx.Target("#task-list"),
		htmx.Swap("beforeend"), // Append new item to the end of the list

		// When HTMX request finishes, clear the input
		atr.Custom("hx-on::after-request", "this.reset()"),

		tags.Input{
			atr.TypeAttr("text"),
			atr.NameAttr("title"),
			atr.Placeholder("What needs to be done?"),
			atr.Custom("required", "true"),
			tw.Class("flex-1 p-2 border border-gray-300 rounded shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"),
		},
		tags.Button{
			atr.TypeAttr("submit"),
			tw.Class("bg-blue-600 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded shadow-sm transition"),
			cegla.Text("Add Task"),
		},
	}
}

// TaskItem represents a single task row.
func TaskItem(t Task) tags.Div {
	return tags.Div{
		atr.ID(fmt.Sprintf("task-%d", t.ID)),
		tw.Class("flex justify-between items-center bg-white p-4 mb-2 rounded shadow-sm"),

		tags.Span{
			tw.Class("text-lg"),
			cegla.Text(t.Title),
		},
		tags.Button{
			// HTMX attributes for deletion
			htmx.Delete(fmt.Sprintf("/tasks/%d", t.ID)),
			htmx.Target(fmt.Sprintf("#task-%d", t.ID)),
			htmx.Swap("outerHTML"), // Replace the whole item with nothing (delete it from DOM)

			tw.Class("text-red-500 hover:text-red-700 font-medium"),
			cegla.Text("Delete"),
		},
	}
}

// TaskList wraps all tasks in a container.
func TaskList(items []Task) tags.Div {
	list := tags.Div{
		atr.ID("task-list"),
		tw.Class("flex flex-col"),
	}

	for _, t := range items {
		list = append(list, TaskItem(t))
	}

	return list
}

// MainApp wraps the form and the list.
func MainApp(items []Task) tags.Div {
	return tags.Div{
		tw.Class("max-w-xl mx-auto mt-10 p-6 bg-white rounded-lg shadow-xl"),

		tags.H1{
			tw.Class("text-3xl font-extrabold text-center mb-8 text-gray-800"),
			cegla.Text("Cegla HTMX Todo"),
		},
		TaskForm(),
		TaskList(items),
	}
}

// --- 4. HTTP Handlers & Router ---

func main() {
	r := gin.Default()

	// Initial data
	mu.Lock()
	tasks = append(tasks, Task{ID: nextID, Title: "Learn Cegla"})
	nextID++
	mu.Unlock()

	// 1. Указываем, что папка "assets" доступна по URL-пути "/assets"
	// Все файлы внутри папки проекта "assets" будут доступны по адресу /assets/...
	r.Static("/assets", "./assets")

	// 1. Full page load
	r.GET("/", func(c *gin.Context) {
		mu.Lock()
		currentTasks := make([]Task, len(tasks))
		copy(currentTasks, tasks)
		mu.Unlock()

		page := Layout("Task Manager", MainApp(currentTasks))
		RenderHTML(c, http.StatusOK, page)
	})

	// 2. Add a new task via HTMX
	r.POST("/tasks", func(c *gin.Context) {
		title := c.PostForm("title")
		if title == "" {
			c.Status(http.StatusBadRequest)
			return
		}

		mu.Lock()
		newTask := Task{ID: nextID, Title: title}
		nextID++
		tasks = append(tasks, newTask)
		mu.Unlock()

		// HTMX expects only the HTML snippet for the new item
		RenderHTML(c, http.StatusOK, TaskItem(newTask))
	})

	// 3. Delete a task via HTMX
	r.DELETE("/tasks/:id", func(c *gin.Context) {
		idStr := c.Param("id")

		mu.Lock()
		defer mu.Unlock()

		for i, t := range tasks {
			if fmt.Sprintf("%d", t.ID) == idStr {
				// Remove from slice
				tasks = append(tasks[:i], tasks[i+1:]...)
				break
			}
		}

		// HTMX swap outerHTML with empty response effectively deletes the element
		c.Status(http.StatusOK)
	})

	fmt.Println("🚀 Server running on http://localhost:8080")
	r.Run("127.0.0.1:8080")
}
