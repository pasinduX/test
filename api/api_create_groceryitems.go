package api

import (
  "TestApp/utils"
  "github.com/gofiber/fiber/v2"

  "TestApp/functions"
    "TestApp/dto"
    "github.com/go-playground/validator/v10"
  "TestApp/dao"
  )

// @Summary      POST GroceryItems input: GroceryItems
// @Description  POST GroceryItems API
// @Tags         POST GroceryItems - GroceryItems
// @Accept       json
// @Produce      json
// @Param        data body dto.Model_GroceryItems false "string collection" 
// @Success      200  {array}   dto.Model_GroceryItems "Status OK"
// @Success      202  {array}   dto.Model_GroceryItems "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /CreateGroceryItems [POST]

    func CreateGroceryItemsApi(c *fiber.Ctx) error {



    inputObj := dto.GroceryItems{}
    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    
if err := functions.UniqueCheck(inputObj, "GroceryItemss", []string{ "groceryitemsId"}) ; err!=nil{
          return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
        }
    
    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err := dao.DB_CreateGroceryItems(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}