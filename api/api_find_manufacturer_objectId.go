package api

import (
  "TestApp/utils"
  "github.com/gofiber/fiber/v2"

  "TestApp/dao"
  )

// @Summary      GET Manufacturer input: Manufacturer
// @Description  GET Manufacturer API
// @Tags         GET Manufacturer - Manufacturer
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Manufacturer "Status OK"
// @Success      202  {array}   dto.Model_Manufacturer "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindManufacturer [GET]

    func FindManufacturerApi(c *fiber.Ctx) error {


objectId := c.Query("objectId")
    returnValue, err := dao.DB_FindManufacturerbyObjectId(objectId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}