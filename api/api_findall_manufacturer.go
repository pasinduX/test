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
// @Router      /FindallManufacturer [GET]

    func FindallManufacturerApi(c *fiber.Ctx) error {


returnValue, err := dao.DB_FindallManufacturer()
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}