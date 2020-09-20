package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/siddontang/go/log"
)

type RequestAPI struct {
	database *sql.DB
}

// Requesting will create a request, if the book exists in the books table.
func (a *RequestAPI) Requesting(ctx *gin.Context) {
	email := ctx.PostForm("email")
	if email == "" {
		ctx.String(http.StatusBadRequest, "must include an email")
		return
	}

	title := ctx.PostForm("title")
	if title == "" {
		ctx.String(http.StatusBadRequest, "must include an title")
		return
	}

	availability, err := a.fetchBook(title)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.String(http.StatusNotFound, fmt.Sprintf("book %s does not exist", title))
			return
		} else {
			log.Error(err.Error())
			ctx.Status(http.StatusInternalServerError)
			return
		}
	}

	id, err := a.insertRequest(email, title)
	if err != nil {
		log.Error(err.Error())
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":        id, // ideally db autoinc id would not be directly exposed
		"available": availability,
		"title":     title,
		"timestamp": "123",
	})
}

// Fetching will retrieve a request, if it exists.
func (a *RequestAPI) Fetching(ctx *gin.Context) {
	// TODO
	ctx.Status(http.StatusNotImplemented)
	return
}

// Fetching will retrieve all request.
func (a *RequestAPI) FetchingAll(ctx *gin.Context) {
	// TODO
	ctx.Status(http.StatusNotImplemented)
	return
}

// Deleting will delete a request, if it exists.
func (a *RequestAPI) Deleting(ctx *gin.Context) {
	// TODO
	ctx.Status(http.StatusNotImplemented)
	return
}

func (a *RequestAPI) insertRequest(email, title string) (int64, error) {
	stmt, err := a.database.Prepare("INSERT INTO requests(email, title) values(?,?)")
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(email, title)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (a *RequestAPI) fetchBook(title string) (avail bool, err error) {
	stmt, err := a.database.Prepare("SELECT available FROM books WHERE title = ?")
	if err != nil {
		return false, err
	}

	row := stmt.QueryRow(title)
	err = row.Scan(&avail)
	if err != nil {
		return false, err
	}

	return
}

// TODO
func (a *RequestAPI) fetchRequest(id int) (avail bool, err error) {
	return
}
