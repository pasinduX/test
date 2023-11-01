package api

import (
  "TestApp/utils"
  "github.com/gofiber/fiber/v2"

  "TestApp/dto"
    "github.com/go-playground/validator/v10"
  "TestApp/dao"
  )

// @Summary      PUT Shop input: Shop
// @Description  PUT Shop API
// @Tags         PUT Shop - Shop
// @Accept       json
// @Produce      json
// @Param        data body dto.Model_Shop false "string collection" 
// @Success      200  {array}   dto.Model_Shop "Status OK"
// @Success      202  {array}   dto.Model_Shop "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /UpdateShop [PUT]

    func UpdateShopApi(c *fiber.Ctx) error {



    inputObj := dto.Shop{}
    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    

    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err := dao.DB_UpdateShop(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}