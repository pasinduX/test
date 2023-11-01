package api

import (
  "TestApp/utils"
  "github.com/gofiber/fiber/v2"

  "TestApp/dao"
  )

// @Summary      DELETE Shop input: Shop
// @Description  DELETE Shop API
// @Tags         DELETE Shop - Shop
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Shop "Status OK"
// @Success      202  {array}   dto.Model_Shop "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /DeleteShop [DELETE]

    func DeleteShopApi(c *fiber.Ctx) error {


objectId := c.Query("objectId")
    err := dao.DB_DeleteShop(objectId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}