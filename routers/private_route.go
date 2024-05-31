package routers

import (
	"mini-core/middleware"
	"mini-core/middleware/go-utils/database"
	accounts "mini-core/modules/batch_upload/routes/accounts"
	clients "mini-core/modules/batch_upload/routes/clients"
	rout "mini-core/modules/search/routes"
	routes "mini-core/modules/select/routes"
	route "mini-core/modules/update/routes"

	"github.com/gofiber/fiber/v2"
)

func SetupPrivateRoutes(app *fiber.App) {

	app.Use(middleware.HeaderResponse())
	database.Data = database.Database{}
	database.ConnectDB()

	ewalletweb := app.Group("/E-Wallet/Web")
	v1Endpoint := ewalletweb.Group("/API")

	v1Endpoint.Get("/download-client-template", clients.DownloadClientTemplate)
	v1Endpoint.Post("/upload-client", clients.UploadClient)
	v1Endpoint.Get("/get_client", routes.GetClient)
	v1Endpoint.Post("/update-client", route.UpdateClient)
	v1Endpoint.Post("/search-client", rout.SearchClient)

	v1Endpoint.Get("/download-accounts-template", accounts.DownloadAccountsTemplate)
	v1Endpoint.Post("/upload-account", accounts.UploadAccount)
	v1Endpoint.Get("/get_account", routes.GetAccount)
	v1Endpoint.Post("/update-account", route.UpdateAccount)
	v1Endpoint.Post("/search-account", rout.SearchAccount)

}
