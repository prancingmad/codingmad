package handlers

import (
	"codingmad/internal/db"
	"codingmad/internal/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func GetNotes(w http.ResponseWriter, r *http.Request) {
	// Optional filtering by language via query parameter
	language := r.URL.Query().Get("language")

	var rows *sql.Rows
	var err error
	if language != "" {
		rows, err = db.DB.Query("SELECT id, title, content, language FROM notes WHERE language = ? ORDER BY id DESC", language)
	} else {
		rows, err = db.DB.Query("SELECT id, title, content, language FROM notes ORDER BY id DESC")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		var n models.Note
		if err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.Language); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		notes = append(notes, n)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var n models.Note
	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Ensure language is not empty
	if n.Language == "" {
		n.Language = "General"
	}

	result, err := db.DB.Exec(
		"INSERT INTO notes (title, content, language) VALUES (?, ?, ?)",
		n.Title, n.Content, n.Language,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	n.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(n)
}

func GetNoteByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/notes/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	var n models.Note
	err = db.DB.QueryRow("SELECT id, title, content FROM notes WHERE id = ?", id).
		Scan(&n.ID, &n.Title, &n.Content)

	if err == sql.ErrNoRows {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(n)
}

func RegisterNoteRoutes(mux *http.ServeMux) {
	// /notes GET/POST
	mux.HandleFunc("/notes", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Route hit:", r.Method, "/notes")
		switch r.Method {
		case http.MethodGet:
			GetNotes(w, r)
		case http.MethodPost:
			CreateNote(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// /notes/{id} GET
	mux.HandleFunc("/notes/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Route hit:", r.Method, r.URL.Path)
		GetNoteByID(w, r)
	})
}
