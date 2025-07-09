
package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"yummer-go/database"
	"yummer-go/models"
)

func CreateCliente(c *gin.Context) {
	var cliente models.Cliente
	if err := c.ShouldBindJSON(&cliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query, err := database.DB.Prepare("INSERT INTO clientes (nome, email, telefone) VALUES (?, ?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	result, err := query.Exec(cliente.Nome, cliente.Email, cliente.Telefone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	cliente.ID = int(id)

	c.JSON(http.StatusCreated, cliente)
}

func GetClientes(c *gin.Context) {
	var clientes []models.Cliente
	rows, err := database.DB.Query("SELECT id, nome, email, telefone FROM clientes")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var cliente models.Cliente
		if err := rows.Scan(&cliente.ID, &cliente.Nome, &cliente.Email, &cliente.Telefone); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		clientes = append(clientes, cliente)
	}

	c.JSON(http.StatusOK, clientes)
}

func GetCliente(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var cliente models.Cliente
	err := database.DB.QueryRow("SELECT id, nome, email, telefone FROM clientes WHERE id = ?", id).Scan(&cliente.ID, &cliente.Nome, &cliente.Email, &cliente.Telefone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cliente not found"})
		return
	}

	c.JSON(http.StatusOK, cliente)
}

func UpdateCliente(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var cliente models.Cliente
	if err := c.ShouldBindJSON(&cliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query, err := database.DB.Prepare("UPDATE clientes SET nome = ?, email = ?, telefone = ? WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = query.Exec(cliente.Nome, cliente.Email, cliente.Telefone, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cliente.ID = id
	c.JSON(http.StatusOK, cliente)
}

func DeleteCliente(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	query, err := database.DB.Prepare("DELETE FROM clientes WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = query.Exec(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cliente deleted successfully"})
}
