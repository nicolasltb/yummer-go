
package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"yummer-go/database"
	"yummer-go/models"
)

func CreateMesa(c *gin.Context) {
	var mesa models.Mesa
	if err := c.ShouldBindJSON(&mesa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query, err := database.DB.Prepare("INSERT INTO mesas (restaurante_id, numero, capacidade, disponivel) VALUES (?, ?, ?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	result, err := query.Exec(mesa.RestauranteID, mesa.Numero, mesa.Capacidade, mesa.Disponivel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	mesa.ID = int(id)

	c.JSON(http.StatusCreated, mesa)
}

func GetMesas(c *gin.Context) {
	var mesas []models.Mesa
	rows, err := database.DB.Query("SELECT id, restaurante_id, numero, capacidade, disponivel FROM mesas")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var mesa models.Mesa
		if err := rows.Scan(&mesa.ID, &mesa.RestauranteID, &mesa.Numero, &mesa.Capacidade, &mesa.Disponivel); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		mesas = append(mesas, mesa)
	}

	c.JSON(http.StatusOK, mesas)
}

func GetMesa(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var mesa models.Mesa
	err := database.DB.QueryRow("SELECT id, restaurante_id, numero, capacidade, disponivel FROM mesas WHERE id = ?", id).Scan(&mesa.ID, &mesa.RestauranteID, &mesa.Numero, &mesa.Capacidade, &mesa.Disponivel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Mesa not found"})
		return
	}

	c.JSON(http.StatusOK, mesa)
}

func UpdateMesa(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var mesa models.Mesa
	if err := c.ShouldBindJSON(&mesa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query, err := database.DB.Prepare("UPDATE mesas SET restaurante_id = ?, numero = ?, capacidade = ?, disponivel = ? WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = query.Exec(mesa.RestauranteID, mesa.Numero, mesa.Capacidade, mesa.Disponivel, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	mesa.ID = id
	c.JSON(http.StatusOK, mesa)
}

func DeleteMesa(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	query, err := database.DB.Prepare("DELETE FROM mesas WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = query.Exec(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Mesa deleted successfully"})
}
