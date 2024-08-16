package main

import (
	"database/sql"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

type Note struct {
	ID   int
	Name string
	Text string
}

type NoteRepository interface {
	Create(note Note) error
	GetAll() ([]Note, error)
	Delete(id int) error
	Update(note Note) error
}

type SQLiteNoteRepository struct {
	db *sql.DB
}

func NewSQLiteNoteRepository(databasePath string) (*SQLiteNoteRepository, error) {
	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		return nil, err
	}
	return &SQLiteNoteRepository{db: db}, nil
}

func (r *SQLiteNoteRepository) Create(note Note) error {
	query := `INSERT INTO note (name, text) VALUES (?, ?)`
	_, err := r.db.Exec(query, note.Name, note.Text)
	return err
}

func (r *SQLiteNoteRepository) GetAll() ([]Note, error) {
	rows, err := r.db.Query("SELECT id, name, text FROM note")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []Note
	for rows.Next() {
		var note Note
		if err := rows.Scan(&note.ID, &note.Name, &note.Text); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	return notes, nil
}

func (r *SQLiteNoteRepository) Delete(id int) error {
	query := `DELETE FROM note WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *SQLiteNoteRepository) Update(note Note) error {
	query := `UPDATE note SET name = ?, text = ? WHERE id = ?`
	_, err := r.db.Exec(query, note.Name, note.Text, note.ID)
	return err
}

func (r *SQLiteNoteRepository) Close() {
	if err := r.db.Close(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("Записная книжка")

	repo, err := NewSQLiteNoteRepository("/Users/maximbabichev/Library/DBeaverData/workspace6/.metadata/sample-database-sqlite-1/Chinook.db")
	if err != nil {
		log.Fatal(err)
	}
	defer repo.Close()

	createTable(repo.db)

	notes, err := repo.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	var selectedNoteID int = -1

	noteList := widget.NewList(
		func() int {
			return len(notes)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(notes[i].Name)
		},
	)

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Название")
	textEntry := widget.NewMultiLineEntry()
	textEntry.SetPlaceHolder("Текст")

	noteList.OnSelected = func(id widget.ListItemID) {
		if id >= 0 && id < len(notes) {
			selectedNoteID = int(id)
			nameEntry.SetText(notes[id].Name)
			textEntry.SetText(notes[id].Text)
		}
	}

	addButton := widget.NewButton("Добавить", func() {
		note := Note{Name: nameEntry.Text, Text: textEntry.Text}
		err := repo.Create(note)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		nameEntry.SetText("")
		textEntry.SetText("")
		notes, _ = repo.GetAll()
		noteList.Refresh()
	})

	saveButton := widget.NewButton("Сохранить изменения", func() {
		if selectedNoteID >= 0 && selectedNoteID < len(notes) {
			note := Note{ID: notes[selectedNoteID].ID, Name: nameEntry.Text, Text: textEntry.Text}
			err := repo.Update(note)
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			notes, _ = repo.GetAll()
			noteList.Refresh()
		}
	})

	deleteButton := widget.NewButton("Удалить", func() {
		if selectedNoteID >= 0 && selectedNoteID < len(notes) {
			err := repo.Delete(notes[selectedNoteID].ID)
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			notes, _ = repo.GetAll()
			noteList.Refresh()
			selectedNoteID = -1 // сбросить выбор после удаления
			nameEntry.SetText("")
			textEntry.SetText("")
		}
	})

	scrollableList := container.NewScroll(noteList)
	scrollableList.SetMinSize(fyne.NewSize(400, 300)) // Устанавливаем минимальный размер для области просмотра списка

	w.SetContent(container.NewVBox(
		scrollableList,
		nameEntry,
		textEntry,
		container.NewHBox(addButton, saveButton, deleteButton),
	))

	w.Resize(fyne.NewSize(600, 800)) // Увеличиваем размер окна
	w.ShowAndRun()
}

func createTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS note (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		text TEXT NOT NULL
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
