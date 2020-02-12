package app

import (
	ping_controllers "github.com/mky-cyber/bookstore_users-api/controllers/ping"
	users_controllers "github.com/mky-cyber/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping_controllers.Ping)

	router.GET("/users/:user_id", users_controllers.GetUser)
	// router.GET("/users/search", controllers.SearchUser)
	router.POST("/users", users_controllers.CreateUser)

}
