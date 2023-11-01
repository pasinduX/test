package api

import (
  "TestApp/utils"
  "github.com/gofiber/fiber/v2"

  "TestApp/dao"
  )

// @Summary      GET Supplier input: Supplier
// @Description  GET Supplier API
// @Tags         GET Supplier - Supplier
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Supplier "Status OK"
// @Success      202  {array}   dto.Model_Supplier "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindSupplier [GET]

    func FindSupplierApi(c *fiber.Ctx) error {


objectId := c.Query("objectId")
    returnValue, err := dao.DB_FindSupplierbyObjectId(objectId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}