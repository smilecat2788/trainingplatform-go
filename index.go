package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := sql.Open("sqlite3", "D:/Code/go/mydatabase.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	createTableSQL :=
		`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		password TEXT
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		fmt.Println(err)
	}

	router := gin.Default()
	router.Static("/static", "./static")
	// Serve static files (CSS, JS, images, etc.)
	router.LoadHTMLGlob("index.html")
	//router.LoadHTMLFiles("./static/2.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
		//c.File("./static/2.html")
	})

	router.POST("/login", func(c *gin.Context) {
		// Get the input value from the form
		vefi := false
		username := c.PostForm("user")
		password := c.PostForm("psw")
		rows, err := db.Query("SELECT id, username, password FROM users")
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()
		var id int
		var user, psw string
		for rows.Next() {
			rows.Scan(&id, &user, &psw)
			if user == username && psw == password {
				vefi = true
			}
		}
		if vefi {
			c.File("./static/2.html")
			//c.HTML(http.StatusOK, "index.html", nil)
		} else {
			c.HTML(500, "index.html", nil)
		}

	})

	router.Run(":8080")
}
