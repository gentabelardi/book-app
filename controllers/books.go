package controllers

import (
	"net/http"
	"github.com/gentabelardi/book-app/models"
	"github.com/gentabelardi/book-app/helper"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"encoding/json"
	"strconv"
)

var ResponseError = helper.ResponseError
var ResponseJson = helper.ResponseJson

func GetAllBook(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	
	if err := models.DB.Find(&books).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
	}

	ResponseJson(w, http.StatusOK, books)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
	}
	var book models.Book

	if err := models.DB.First(&book, id).Error; err != nil {
		switch err {
			case gorm.ErrRecordNotFound:
				ResponseError(w, http.StatusNotFound, "Book not found")
				return
			default:
				ResponseError(w, http.StatusInternalServerError, err.Error())
				return
		}
	}
	ResponseJson(w, http.StatusOK, book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	decode := json.NewDecoder(r.Body)
	if err := decode.Decode(&book); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if err := models.DB.Create(&book).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJson(w, http.StatusCreated, book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
	}
	var book models.Book

	if err := models.DB.Delete(&book, id).Error; err != nil {
		switch err {
			case gorm.ErrRecordNotFound:
				ResponseError(w, http.StatusNotFound, "Book not found")
				return
			default:
				ResponseError(w, http.StatusInternalServerError, err.Error())
				return
		}
	}
	ResponseJson(w, http.StatusOK, "Successfully deleted")

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
	}

	var book models.Book

	decode := json.NewDecoder(r.Body)
	if err := decode.Decode(&book); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if models.DB.Where("id = ?", id).Updates(&book).RowsAffected == 0 {
		ResponseError(w, http.StatusBadRequest, "Book not found")
		return
	}

	book.ID = int(id)

	ResponseJson(w, http.StatusCreated, book)
}