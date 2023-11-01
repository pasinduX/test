package api

import (
  "TestApp/utils"
  "github.com/gofiber/fiber/v2"

  "TestApp/functions"
    "TestApp/dto"
    "github.com/go-playground/validator/v10"
  "TestApp/dao"
  )

// @Summary      POST Supplier input: Supplier
// @Description  POST Supplier API
// @Tags         POST Supplier - Supplier
// @Accept       json
// @Produce      json
// @Param        data body dto.Model_Supplier false "string collection" 
// @Success      200  {array}   dto.Model_Supplier "Status OK"
// @Success      202  {array}   dto.Model_Supplier "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /CreateSupplier [POST]

    func CreateSupplierApi(c *fiber.Ctx) error {



    inputObj := dto.Supplier{}
    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    
if err := functions.UniqueCheck(inputObj, "Suppliers", []string{ "supplierId", "phone", "email"}) ; err!=nil{
          return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
        }
    
    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err := dao.DB_CreateSupplier(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}