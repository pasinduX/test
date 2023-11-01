package api

import (
  "TestApp/utils"
  "github.com/gofiber/fiber/v2"

  "TestApp/dao"
  )

// @Summary      DELETE Manufacturer input: Manufacturer
// @Description  DELETE Manufacturer API
// @Tags         DELETE Manufacturer - Manufacturer
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Manufacturer "Status OK"
// @Success      202  {array}   dto.Model_Manufacturer "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /DeleteManufacturer [DELETE]

    func DeleteManufacturerApi(c *fiber.Ctx) error {


objectId := c.Query("objectId")
    err := dao.DB_DeleteManufacturer(objectId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}