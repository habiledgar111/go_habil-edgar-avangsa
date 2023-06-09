package controller

import (
	"code_competence/config"
	"code_competence/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllBarangfromKategori(c echo.Context) error {
	var kategori model.Kategori
	id, _ := strconv.Atoi(c.Param("id"))

	err := config.DB.Model(&model.Kategori{}).Preload("Barang").Find(&kategori, id).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cant get data",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":             "success get data",
		"kateogri and barang": kategori,
	})
}
