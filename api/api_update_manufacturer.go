package api

import (
  "TestApp/utils"
  "github.com/gofiber/fiber/v2"

  "TestApp/dto"
    "github.com/go-playground/validator/v10"
  "TestApp/dao"
  )

// @Summary      PUT Manufacturer input: Manufacturer
// @Description  PUT Manufacturer API
// @Tags         PUT Manufacturer - Manufacturer
// @Accept       json
// @Produce      json
// @Param        data body dto.Model_Manufacturer false "string collection" 
// @Success      200  {array}   dto.Model_Manufacturer "Status OK"
// @Success      202  {array}   dto.Model_Manufacturer "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /UpdateManufacturer [PUT]

    func UpdateManufacturerApi(c *fiber.Ctx) error {



    inputObj := dto.Manufacturer{}
    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    

    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err := dao.DB_UpdateManufacturer(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}