package api

import (
  "TestApp/utils"
  "github.com/gofiber/fiber/v2"

  "TestApp/dao"
  )

// @Summary      GET Shop input: Shop
// @Description  GET Shop API
// @Tags         GET Shop - Shop
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Shop "Status OK"
// @Success      202  {array}   dto.Model_Shop "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindallShop [GET]

    func FindallShopApi(c *fiber.Ctx) error {


returnValue, err := dao.DB_FindallShop()
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}