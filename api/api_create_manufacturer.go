package api

import (
  "TestApp/utils"
  "github.com/gofiber/fiber/v2"

  "TestApp/functions"
    "TestApp/dto"
    "github.com/go-playground/validator/v10"
  "TestApp/dao"
  )

// @Summary      POST Manufacturer input: Manufacturer
// @Description  POST Manufacturer API
// @Tags         POST Manufacturer - Manufacturer
// @Accept       json
// @Produce      json
// @Param        data body dto.Model_Manufacturer false "string collection" 
// @Success      200  {array}   dto.Model_Manufacturer "Status OK"
// @Success      202  {array}   dto.Model_Manufacturer "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /CreateManufacturer [POST]

    func CreateManufacturerApi(c *fiber.Ctx) error {



    inputObj := dto.Manufacturer{}
    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    
if err := functions.UniqueCheck(inputObj, "Manufacturers", []string{ "manufacturerId"}) ; err!=nil{
          return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
        }
    
    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err := dao.DB_CreateManufacturer(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}