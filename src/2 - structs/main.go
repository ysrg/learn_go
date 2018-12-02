package main

import (
  "encoding/json"
  "github.com/gorilla/mux"
  "log"
  "math/rand"
  "net/http"
  "strconv"
)

// Book struct (Model)
type Book struct {
  ID     string  `json:"id"`
  ISBN   string  `json:"isbn"`
  TITLE  string  `json:"title"`
  AUTHOR *author `json:"author"`
}

type author struct {
  Firstname string `json:"firstname"`
  Lastname  string `json:"lastname"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application-json")
  json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application-json")
  params := mux.Vars(r) // Get params

  // Loops through books and finds id
  for _, i := range books {
    if i.ID == params["id"] {
      json.NewEncoder(w).Encode(i)
      return
    }
  }
  json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application-json")
  var book Book
  _ = json.NewDecoder(r.Body).Decode(&book)
  book.ID = strconv.Itoa(rand.Intn(100000))
  books = append(books, book)
  json.NewEncoder(w).Encode(book)

}

func updateBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application-json")
  params := mux.Vars(r)
  for i, item := range books {
    if item.ID == params["id"] {
      books = append(books[:i], books[i+1:]...)
      var book Book
      _ = json.NewDecoder(r.Body).Decode(&book)
      book.ID = params["id"]
      books = append(books, book)
      json.NewEncoder(w).Encode(book)
      return
    }
  }
  json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application-json")
  params := mux.Vars(r)
  for i, item := range books {
    if item.ID == params["id"] {
      books = append(books[:i], books[i+1:]...)
    }
    break
  }
  json.NewEncoder(w).Encode(books)
}

func main() {
  // Router init
  r := mux.NewRouter()

  books = append(books, Book{ID: "1", ISBN: "32442", TITLE: "Be good", AUTHOR: &author{Firstname: "John", Lastname: "Doe"}})
  books = append(books, Book{ID: "2", ISBN: "12442", TITLE: "Be bad", AUTHOR: &author{Firstname: "Eric", Lastname: "Cartman"}})

  // Route handlers
  r.HandleFunc("/api/books", getBooks).Methods("GET")
  r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
  r.HandleFunc("/api/books", createBook).Methods("POST")
  r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
  r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
  log.Fatal(http.ListenAndServe(":8000", r))
}
