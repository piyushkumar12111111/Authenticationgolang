package handlers

import (
	"encoding/json"
	"net/http"
	

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/piyushkumar/authenticationmayursir/models"
)

var bookList = make(map[string]models.BookList )


func Addbook(w http.ResponseWriter, r *http.Request) {
	var book models.BookList
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}


	//! if book already exists 

	for _, value := range bookList {
		if value.BookName == book.BookName && value.Author == book.Author{
			http.Error(w, "Book already exists", http.StatusBadRequest)
			return
		}
	}
	

	//! Generate a new UUID for the report ID
	uuid := uuid.New().String()
	book.BookId = uuid
	//book.UpdatedAt = time.Now()

	book.BookName = book.BookName
	book.Author = book.Author
	


	mu.Lock()
	bookList[book.BookId] = book
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}


//! get book by id
func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	mu.Lock()
	book, ok := bookList[id]
	mu.Unlock()
	if !ok {
		http.NotFound(w, r)

		json.NewEncoder(w).Encode("Book not found 🔒")
		return
	}

	json.NewEncoder(w).Encode(book)
}


//! delete book by id 

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	mu.Lock()
	delete(bookList, id)
	mu.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

//! update book by id

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	var updatedBook models.BookList
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	book, exists := bookList[id]
	if !exists {
		mu.Unlock()
		http.NotFound(w, r)
		return
	}

	// Maintain the ID of the original report
	updatedBook.BookId = book.BookId
	bookList[id] = updatedBook
	//updatedBook.UpdatedAt = time.Now()
	mu.Unlock()

	json.NewEncoder(w).Encode(updatedBook)
}

//! get all books 


func GetAllBooks(w http.ResponseWriter, r *http.Request) {

	mu.Lock()
	defer mu.Unlock()

	json.NewEncoder(w).Encode(bookList)
}


