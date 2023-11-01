package api

import (
  "TestApp/utils"
  "github.com/gofiber/fiber/v2"

  "TestApp/dao"
  )

// @Summary      DELETE Supplier input: Supplier
// @Description  DELETE Supplier API
// @Tags         DELETE Supplier - Supplier
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Supplier "Status OK"
// @Success      202  {array}   dto.Model_Supplier "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /DeleteSupplier [DELETE]

    func DeleteSupplierApi(c *fiber.Ctx) error {


objectId := c.Query("objectId")
    err := dao.DB_DeleteSupplier(objectId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}