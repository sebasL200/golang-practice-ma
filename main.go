package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"

    _ "modernc.org/sqlite"

    "golang-practice-ma/internal/service"
    "golang-practice-ma/internal/store"
    "golang-practice-ma/internal/transport"
)

func main() {
    db, err := sql.Open("sqlite", "./libros.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }

    q := `CREATE TABLE IF NOT EXISTS libros (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        autor TEXT NOT NULL
    );`

    if _, err := db.Exec(q); err != nil {
        log.Fatal(err.Error())
    }

    libroStore := store.New(db)
    
    myLogger := &RunLogger{} 
    
    libroService := service.New(libroStore, myLogger)
    
    libroHandler := transport.New(libroService)

   
    http.HandleFunc("/books", libroHandler.HandleBooks)      
    http.HandleFunc("/books/", libroHandler.HandleBookByID)  

    fmt.Println("Servidor corriendo en http://localhost:8080")
    fmt.Println("API Endpoints Disponibles:")
    fmt.Println("- GET /books        : Obtiene todos los libros")
    fmt.Println("- POST /books       : Crea un nuevo libro")
    fmt.Println("- GET /books/{id}   : Obtiene un libro por ID")
    fmt.Println("- PUT /books/{id}   : Actualiza un libro por ID")
    fmt.Println("- DELETE /books/{id}: Elimina un libro por ID")

    log.Fatal(http.ListenAndServe(":8080", nil))
}

type RunLogger struct{}

func (l *RunLogger) Log(msg, errStr string) {
    if errStr != "" {
        log.Printf("[ERROR] %s: %s", msg, errStr)
    } else {
        log.Printf("[INFO] %s", msg)
    }
}