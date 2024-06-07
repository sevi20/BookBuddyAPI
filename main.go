package main

import (
	"BOOKBUDDYAPI/database"
	"BOOKBUDDYAPI/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response map[string]any

func main() {
	database.Init()

	app := gin.Default()

	app.GET("/books", func(context *gin.Context) {
		fmt.Println("Books")

		result, err := models.GetAllBooks()
		if err != nil {
			context.JSON(400, Response{
				"message": "Cannot serve your request",
			})
			return
		}

		context.JSON(200, Response{
			"message": "All books in the database",
			"books":   result,
		})
	})

	app.GET("/book/:id", func(context *gin.Context) {
		idParam := context.Param("id")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			context.JSON(400, Response{
				"message": "Invalid book ID",
			})
			return
		}

		result, err := models.GetBookByID(int(id))
		if err != nil {
			context.JSON(400, Response{
				"message": "Cannot serve your request",
			})
			return
		}

		context.JSON(200, Response{
			"message": "Book with this ID is in the database",
			"book":    result,
		})
	})

	app.POST("/books", func(context *gin.Context) {
		var bookObject models.Book
		err := context.ShouldBindJSON(&bookObject)
		if err != nil {
			context.JSON(400, Response{
				"message": "Invalid object",
			})
			return
		}

		err = bookObject.Save()
		if err != nil {
			context.JSON(400, Response{
				"message": "Cannot insert book object",
			})
			return
		}

		context.JSON(200, Response{
			"message": "Book created successfully",
			"object":  bookObject,
		})
	})

	app.DELETE("/book/:id", func(context *gin.Context) {
		idParam := context.Param("id")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			context.JSON(400, Response{
				"message": "Invalid book ID",
			})
			return
		}

		err = models.DeleteBookByID(int(id))
		if err != nil {
			context.JSON(400, Response{
				"message": "Cannot delete book",
			})
			return
		}

		context.JSON(200, Response{
			"message": "Book deleted successfully!",
		})
	})

	app.PUT("/book/:id", func(context *gin.Context) {
		idParam := context.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			context.JSON(400, Response{
				"message": "Invalid book ID",
			})
			return
		}

		var bookUpdate models.Book
		err = context.ShouldBindJSON(&bookUpdate)
		if err != nil {
			context.JSON(400, Response{
				"message": "Invalid JSON body",
			})
			return
		}

		err = models.UpdateBookByID(id, &bookUpdate)
		if err != nil {
			context.JSON(400, Response{
				"message": "Failed to update book",
			})
			return
		}

		context.JSON(200, Response{
			"message": "Book updated successfully",
		})
	})

	err := app.Run(":8087")
	if err != nil {
		fmt.Println("SERVER exception")
		fmt.Println(err)
	}
}
