package api

import (
  "TestApp/utils"
  "github.com/gofiber/fiber/v2"

  "TestApp/dao"
  )

// @Summary      DELETE Certificates input: Certificates
// @Description  DELETE Certificates API
// @Tags         DELETE Certificates - Certificates
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Certificates "Status OK"
// @Success      202  {array}   dto.Model_Certificates "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /DeleteCertificates [DELETE]

    func DeleteCertificatesApi(c *fiber.Ctx) error {


objectId := c.Query("objectId")
    err := dao.DB_DeleteCertificates(objectId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}