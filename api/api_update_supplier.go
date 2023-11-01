package api

import (
  "TestApp/utils"
  "github.com/gofiber/fiber/v2"

  "TestApp/dto"
    "github.com/go-playground/validator/v10"
  "TestApp/dao"
  )

// @Summary      PUT Supplier input: Supplier
// @Description  PUT Supplier API
// @Tags         PUT Supplier - Supplier
// @Accept       json
// @Produce      json
// @Param        data body dto.Model_Supplier false "string collection" 
// @Success      200  {array}   dto.Model_Supplier "Status OK"
// @Success      202  {array}   dto.Model_Supplier "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /UpdateSupplier [PUT]

    func UpdateSupplierApi(c *fiber.Ctx) error {



    inputObj := dto.Supplier{}
    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    

    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err := dao.DB_UpdateSupplier(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}