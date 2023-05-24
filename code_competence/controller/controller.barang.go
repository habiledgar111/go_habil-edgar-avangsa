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

	err = config.DB.Model(&model.Barang{}).Preload("Kategori").Find(&barang, id).Error
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
	kategori, _ := strconv.Atoi((fmt.Sprintf("%v", json_map["kategori"])))

	barang = model.Barang{
		Nama:       fmt.Sprintf("%v", json_map["nama"]),
		Deskripsi:  fmt.Sprintf("%v", json_map["deskripsi"]),
		Jumlah:     jumlah,
		Harga:      harga,
		KategoriID: uint(kategori),
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

func UpdateBarang(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var barang model.Barang
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Massage": "json cant empty",
		})
	}
	config.DB.Where("id = ?", id).First(&barang)

	if barang.ID == 0 {
		return c.JSON(http.StatusBadRequest, "cant find data")
	}

	if json_map["nama"] != nil {
		barang.Nama = fmt.Sprintf("%v", json_map["nama"])
	}

	if json_map["deskripsi"] != nil {
		barang.Deskripsi = fmt.Sprintf("%v", json_map["deskripsi"])
	}

	if json_map["jumlah"] != nil {
		jumlah, _ := strconv.Atoi((fmt.Sprintf("%v", json_map["jumlah"])))
		barang.Jumlah = jumlah
	}

	if json_map["harga"] != nil {
		harga, _ := strconv.ParseFloat(fmt.Sprintf("%v", json_map["harga"]), 3)
		barang.Harga = harga
	}

	if json_map["kategori"] != nil {
		kategori, _ := strconv.Atoi(fmt.Sprintf("%v", json_map["kategori"]))
		barang.KategoriID = uint(kategori)
	}

	result := config.DB.Where("id = ?", id).Updates(&barang)

	if result.RowsAffected < 1 {
		return c.JSON(http.StatusInternalServerError, "cant update data")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update data",
		"update":  barang,
	})
}

func GetBarangfromName(c echo.Context) error {
	nama := c.QueryParam("nama")
	namabarang := "%" + nama + "%"
	var barang []model.Barang

	err := config.DB.Where("Nama LIKE ?", namabarang).Find(&barang).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, "cant get data")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "success get data",
		"barangs": barang,
	})
}
