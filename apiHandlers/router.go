package apiHandlers

import (
	"TestApp/api"

	"encoding/gob"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Router(app *fiber.App) {
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	store := session.New()
	//encryptcookie.New(encryptcookie.Config{
	//	Key: "secret-thirty-2-character-string",
	//})
	gob.Register(map[string]interface{}{})

	api := app.Group("/TestApp/api")
	check := app.Group("/")
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	DefaultMappings(check, store)
	RouteMappings(api,store)
}

func RouteMappings(cg fiber.Router, store *session.Store) {
cg.Post("/CreateGroceryItems", api.CreateGroceryItemsApi)
cg.Put("/UpdateGroceryItems", api.UpdateGroceryItemsApi)
cg.Delete("/DeleteGroceryItems", api.DeleteGroceryItemsApi)
cg.Get("/FindGroceryItems", api.FindGroceryItemsApi)
cg.Get("/FindallGroceryItems", api.FindallGroceryItemsApi)
cg.Post("/CreateShop", api.CreateShopApi)
cg.Put("/UpdateShop", api.UpdateShopApi)
cg.Delete("/DeleteShop", api.DeleteShopApi)
cg.Get("/FindShop", api.FindShopApi)
cg.Get("/FindallShop", api.FindallShopApi)
cg.Post("/CreateSupplier", api.CreateSupplierApi)
cg.Put("/UpdateSupplier", api.UpdateSupplierApi)
cg.Delete("/DeleteSupplier", api.DeleteSupplierApi)
cg.Get("/FindSupplier", api.FindSupplierApi)
cg.Get("/FindallSupplier", api.FindallSupplierApi)
cg.Post("/CreateManufacturer", api.CreateManufacturerApi)
cg.Put("/UpdateManufacturer", api.UpdateManufacturerApi)
cg.Delete("/DeleteManufacturer", api.DeleteManufacturerApi)
cg.Get("/FindManufacturer", api.FindManufacturerApi)
cg.Get("/FindallManufacturer", api.FindallManufacturerApi)
cg.Post("/CreateCertificates", api.CreateCertificatesApi)
cg.Put("/UpdateCertificates", api.UpdateCertificatesApi)
cg.Delete("/DeleteCertificates", api.DeleteCertificatesApi)
cg.Get("/FindCertificates", api.FindCertificatesApi)
cg.Get("/FindallCertificates", api.FindallCertificatesApi)


}

func DefaultMappings(cg fiber.Router, store *session.Store) {
	cg.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"message": "service is up and running", "version": "1.0"})
	})
}