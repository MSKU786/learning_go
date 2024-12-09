package controllers

import (
	"encoding/json"
	"fmt"
	"go-bookstore/pkg/models"
	"go-bookstore/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter,  r *http.Request) {
		newBooks := models.GetAllBooks();
		res,_  := json.Marshal(newBooks);
		w.Header().Set("Content-Type", "pkglication/josn")
		w.WriteHeader(http.StatusOK);
		w.Write(res);
}


func GetBookById(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r);
		bookId := vars["bookId"]
		ID, err := strconv.ParseInt(bookId,0,0)

		if (err != nil) {
			fmt.Println("Error while parsing")
		}
		bookDetails,_ := models.GetBookById(ID)
		res, _ := json.Marshal(bookDetails)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res);
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
		CreateBook := &models.Book{}
		utils.ParseBody(r, CreateBook) 
		b := CreateBook.CreateBook()
		res, _ := json.Marshal(b);
		w.WriteHeader(http.StatusOK)
		w.Write(res);
}