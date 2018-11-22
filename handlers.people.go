package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func showSantaPick(c *gin.Context)  {

	person:= getSecretSantaPick()

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"person.html",
		// Pass the data that the page uses
		gin.H{
			"title":   "Home Page",
			"Name": person,
		},
	)

}
