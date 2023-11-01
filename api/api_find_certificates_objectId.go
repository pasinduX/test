package api

import (
  "TestApp/utils"
  "github.com/gofiber/fiber/v2"

  "TestApp/dao"
  )

// @Summary      GET Certificates input: Certificates
// @Description  GET Certificates API
// @Tags         GET Certificates - Certificates
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Certificates "Status OK"
// @Success      202  {array}   dto.Model_Certificates "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindCertificates [GET]

    func FindCertificatesApi(c *fiber.Ctx) error {


objectId := c.Query("objectId")
    returnValue, err := dao.DB_FindCertificatesbyObjectId(objectId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}