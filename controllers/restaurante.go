
package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"yummer-go/database"
	"yummer-go/models"
)

func CreateRestaurante(c *gin.Context) {
	var restaurante models.Restaurante
	if err := c.ShouldBindJSON(&restaurante); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query, err := database.DB.Prepare("INSERT INTO restaurantes (nome, endereco, tipo_cozinha, horario_funcionamento) VALUES (?, ?, ?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	result, err := query.Exec(restaurante.Nome, restaurante.Endereco, restaurante.TipoCozinha, restaurante.HorarioFuncionamento)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	restaurante.ID = int(id)

	c.JSON(http.StatusCreated, restaurante)
}

func GetRestaurantes(c *gin.Context) {
	var restaurantes []models.Restaurante
	rows, err := database.DB.Query("SELECT id, nome, endereco, tipo_cozinha, horario_funcionamento FROM restaurantes")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var restaurante models.Restaurante
		if err := rows.Scan(&restaurante.ID, &restaurante.Nome, &restaurante.Endereco, &restaurante.TipoCozinha, &restaurante.HorarioFuncionamento); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		restaurantes = append(restaurantes, restaurante)
	}

	c.JSON(http.StatusOK, restaurantes)
}

func GetRestaurante(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var restaurante models.Restaurante
	err := database.DB.QueryRow("SELECT id, nome, endereco, tipo_cozinha, horario_funcionamento FROM restaurantes WHERE id = ?", id).Scan(&restaurante.ID, &restaurante.Nome, &restaurante.Endereco, &restaurante.TipoCozinha, &restaurante.HorarioFuncionamento)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Restaurante not found"})
		return
	}

	c.JSON(http.StatusOK, restaurante)
}

func UpdateRestaurante(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var restaurante models.Restaurante
	if err := c.ShouldBindJSON(&restaurante); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query, err := database.DB.Prepare("UPDATE restaurantes SET nome = ?, endereco = ?, tipo_cozinha = ?, horario_funcionamento = ? WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = query.Exec(restaurante.Nome, restaurante.Endereco, restaurante.TipoCozinha, restaurante.HorarioFuncionamento, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	restaurante.ID = id
	c.JSON(http.StatusOK, restaurante)
}

func DeleteRestaurante(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	query, err := database.DB.Prepare("DELETE FROM restaurantes WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = query.Exec(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Restaurante deleted successfully"})
}
