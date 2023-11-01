package api

import (
  "TestApp/utils"
  "github.com/gofiber/fiber/v2"

  "TestApp/dto"
    "github.com/go-playground/validator/v10"
  "TestApp/dao"
  )

// @Summary      PUT GroceryItems input: GroceryItems
// @Description  PUT GroceryItems API
// @Tags         PUT GroceryItems - GroceryItems
// @Accept       json
// @Produce      json
// @Param        data body dto.Model_GroceryItems false "string collection" 
// @Success      200  {array}   dto.Model_GroceryItems "Status OK"
// @Success      202  {array}   dto.Model_GroceryItems "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /UpdateGroceryItems [PUT]

    func UpdateGroceryItemsApi(c *fiber.Ctx) error {



    inputObj := dto.GroceryItems{}
    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    

    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err := dao.DB_UpdateGroceryItems(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}