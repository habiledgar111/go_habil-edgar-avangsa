package controller

import (
	"code_competence/config"
	"code_competence/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllBarang(c echo.Context) error {
	var barang []model.Barang

	err := config.DB.Model(&model.Barang{}).Preload("Kategori").Find(&barang).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get data",
		"barang":  barang,
	})
}

func GetBarangID(c echo.Context) error {
	var barang model.Barang
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "cant get id barang",
			"error":   err,
		})
	}

	err = config.DB.Model(&model.Barang{}).Find(&barang, id).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "error when get data",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success when get data",
		"barang":  barang,
	})
}

func CreateBarang(c echo.Context) error {
	var barang model.Barang
	json_map := make(map[string]interface{})

	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Massage": "json cant empty",
		})
	}

	jumlah, _ := strconv.Atoi((fmt.Sprintf("%v", json_map["jumlah"])))
	harga, _ := strconv.ParseFloat(fmt.Sprintf("%v", json_map["harga"]), 3)

	barang = model.Barang{
		Nama:      fmt.Sprintf("%v", json_map["nama"]),
		Deskripsi: fmt.Sprintf("%v", json_map["deskripsi"]),
		Jumlah:    jumlah,
		Harga:     harga,
		Kategori: model.Kategori{
			Nama: fmt.Sprintf("%v", json_map["katergori"]),
		},
	}

	result := config.DB.Create(&barang)
	if result.RowsAffected < 1 {
		return c.JSON(http.StatusInternalServerError, "cant save data")
	}

	return c.JSON(http.StatusOK, "success save data")
}

func DeleteBarang(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	result := config.DB.Delete(&model.Barang{}, id)

	if result.RowsAffected < 1 {
		return c.JSON(http.StatusInternalServerError, "cant delete data")
	}

	return c.JSON(http.StatusAccepted, "success delete barang")
}
