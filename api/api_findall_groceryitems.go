package api

import (
  "TestApp/utils"
  "github.com/gofiber/fiber/v2"

  "TestApp/dao"
  )

// @Summary      GET GroceryItems input: GroceryItems
// @Description  GET GroceryItems API
// @Tags         GET GroceryItems - GroceryItems
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_GroceryItems "Status OK"
// @Success      202  {array}   dto.Model_GroceryItems "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindallGroceryItems [GET]

    func FindallGroceryItemsApi(c *fiber.Ctx) error {


returnValue, err := dao.DB_FindallGroceryItems()
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}