package controllers

import (
	"fmt"
	"net/http"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateBook")
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetBook")
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateBookByID")
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateBook")
}

func DeleteBookByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteBookByID")
}
