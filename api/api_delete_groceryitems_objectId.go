package api

import (
  "TestApp/utils"
  "github.com/gofiber/fiber/v2"

  "TestApp/dao"
  )

// @Summary      DELETE GroceryItems input: GroceryItems
// @Description  DELETE GroceryItems API
// @Tags         DELETE GroceryItems - GroceryItems
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_GroceryItems "Status OK"
// @Success      202  {array}   dto.Model_GroceryItems "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /DeleteGroceryItems [DELETE]

    func DeleteGroceryItemsApi(c *fiber.Ctx) error {


objectId := c.Query("objectId")
    err := dao.DB_DeleteGroceryItems(objectId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}