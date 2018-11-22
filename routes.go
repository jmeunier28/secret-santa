package main

func initializeRoutes()  {
	// Handle the index route
	router.GET("/", showIndexPage)

	router.GET("/people/view/santa", showSantaPick)
}


