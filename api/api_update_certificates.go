package api

import (
  "TestApp/utils"
  "github.com/gofiber/fiber/v2"

  "TestApp/dto"
    "github.com/go-playground/validator/v10"
  "TestApp/dao"
  )

// @Summary      PUT Certificates input: Certificates
// @Description  PUT Certificates API
// @Tags         PUT Certificates - Certificates
// @Accept       json
// @Produce      json
// @Param        data body dto.Model_Certificates false "string collection" 
// @Success      200  {array}   dto.Model_Certificates "Status OK"
// @Success      202  {array}   dto.Model_Certificates "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /UpdateCertificates [PUT]

    func UpdateCertificatesApi(c *fiber.Ctx) error {



    inputObj := dto.Certificates{}
    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    

    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err := dao.DB_UpdateCertificates(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}