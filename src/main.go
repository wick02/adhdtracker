package main

import (
	"database/sql"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
    "strings"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type Response struct {
	Data interface{} `json:"data,omitempty"`
	Err  string      `json:"error,omitempty"`
}

//go:embed index.html
var content embed.FS

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db, err := sql.Open("sqlite3", "./adhdtracker.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, status TEXT NOT NULL)")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		indexContent, err := fs.ReadFile(content, "index.html") // Change variable name from 'content' to 'indexContent'
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write(indexContent)
	})

    http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            getTasksHandler(db, w, r)
        case http.MethodPost:
            addTaskHandler(db, w, r)
        default:
            http.NotFound(w, r)
        }
    })
    
    http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodPut:
            updateTaskStatusHandler(db, w, r)
        case http.MethodDelete:
            deleteTaskHandler(db, w, r)
        default:
            http.NotFound(w, r)
        }
    })
    
    
    

	fmt.Printf("Listening on :%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}


func getTasksHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT id, name, status FROM tasks")
    if err != nil {
        writeErrorResponse(w, http.StatusInternalServerError, "Error retrieving tasks from database")
        return
    }
    defer rows.Close()

    tasks := make([]Task, 0) // Initialize tasks slice with a non-nil value
    for rows.Next() {
        var task Task
        err := rows.Scan(&task.ID, &task.Name, &task.Status)
        if err != nil {
            writeErrorResponse(w, http.StatusInternalServerError, "Error reading task from database")
            return
        }
        tasks = append(tasks, task)
    }

    writeJSONResponse(w, http.StatusOK, Response{Data: tasks})
}


func addTaskHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	result, err := db.Exec("INSERT INTO tasks (name, status) VALUES (?, ?)", task.Name, "Not Done")
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Error adding task to database")
		return
	}

	task.ID, _ = result.LastInsertId()
	task.Status = "Not Done"

	writeJSONResponse(w, http.StatusCreated, Response{Data: task})
}

func updateTaskStatusHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
    var task Task
    err := json.NewDecoder(r.Body).Decode(&task)
    if err != nil {
        writeErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    defer r.Body.Close()

    id := strings.TrimPrefix(r.URL.Path, "/tasks/")
    if id == "" {
        writeErrorResponse(w, http.StatusBadRequest, "Invalid task ID")
        return
    }

    _, err = db.Exec("UPDATE tasks SET status = ? WHERE id = ?", task.Status, id)
    if err != nil {
        writeErrorResponse(w, http.StatusInternalServerError, "Error updating task status in database")
        return
    }

    writeJSONResponse(w, http.StatusOK, Response{Data: "Task status updated successfully"})
}

func deleteTaskHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
    id := strings.TrimPrefix(r.URL.Path, "/tasks/")
    if id == "" {
        writeErrorResponse(w, http.StatusBadRequest, "Invalid task ID")
        return
    }

    _, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
    if err != nil {
        writeErrorResponse(w, http.StatusInternalServerError, "Error deleting task from database")
        return
    }

    writeJSONResponse(w, http.StatusOK, Response{Data: "Task deleted successfully"})
}



func writeJSONResponse(w http.ResponseWriter,statusCode int, response Response) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(response)
    }
    
    func writeErrorResponse(w http.ResponseWriter, statusCode int, message string) {
    writeJSONResponse(w, statusCode, Response{Err: message})
    }
