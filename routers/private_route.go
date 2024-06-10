package routers

import (
	"mini-core/middleware"
	"mini-core/middleware/go-utils/database"
	out "mini-core/modules/approve/routes"
	accounts "mini-core/modules/batch_upload/routes/accounts"
	clients "mini-core/modules/batch_upload/routes/clients"
	user "mini-core/modules/create_account/routes"
	rout "mini-core/modules/search/routes"
	routes "mini-core/modules/select/routes"
	route "mini-core/modules/update/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupPrivateRoutes(app *fiber.App) {

	app.Use(middleware.HeaderResponse())
	database.Data = database.Database{}
	database.ConnectDB()

	// CORS middleware with a more permissive configuration
	corsMiddleware := cors.New(cors.Config{
		AllowMethods: "GET,POST", // Allow GET and POST methods
		AllowOrigins: "*",        // Allow all origins
	})

	ewalletweb := app.Group("/E-Wallet/Web")
	v1Endpoint := ewalletweb.Group("/API")

	// Apply general CORS middleware to the whole group
	v1Endpoint.Use(corsMiddleware)

	v1Endpoint.Get("/download-client-template", clients.DownloadClientTemplate)

	// Use CORS middleware specifically for the /upload-client and /upload-account endpoints
	v1Endpoint.Post("/upload-client", cors.New(cors.Config{
		AllowMethods: "POST",
		AllowOrigins: "*",
	}), clients.UploadClient)

	v1Endpoint.Get("/get_client", routes.GetClient)
	v1Endpoint.Post("/update-client", route.UpdateClient)
	v1Endpoint.Post("/search-client", rout.SearchClient)

	v1Endpoint.Get("/download-accounts-template", accounts.DownloadAccountsTemplate)

	// Use CORS middleware specifically for the /upload-account endpoint
	v1Endpoint.Post("/upload-account", cors.New(cors.Config{
		AllowMethods: "POST",
		AllowOrigins: "*",
	}), accounts.UploadAccount)

	v1Endpoint.Get("/get_account", routes.GetAccount)
	v1Endpoint.Post("/update-account", route.UpdateAccount)
	v1Endpoint.Post("/search-account", rout.SearchAccount)

	v1Endpoint.Post("/approve-client", out.ApproveClients)
	v1Endpoint.Post("/disapprove-client", out.DisapproveClients)
	v1Endpoint.Post("/view-client-account", rout.ViewClientAccount)

	v1Endpoint.Post("/create-account", user.RegisterNewUser)
	v1Endpoint.Post("/login-user", user.LoginUser)
	v1Endpoint.Post("/logout-user", middleware.JWTMiddleware(), user.Logout)

	v1Endpoint.Get("/get-registered-users", user.GetUsers)
	v1Endpoint.Post("/delete-clients", route.DeleteClients)
}
