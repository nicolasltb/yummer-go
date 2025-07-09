package database

import (
	"database/sql"
	"fmt"
	"time"
)

func SeedData(db *sql.DB) {
	// Seed Clientes
	_, err := db.Exec(`
		INSERT INTO clientes (nome, email, telefone) VALUES
		('João Silva', 'joao.silva@example.com', '11987654321'),
		('Maria Souza', 'maria.souza@example.com', '21998765432'),
		('Carlos Pereira', 'carlos.pereira@example.com', '31976543210')
	`)
	if err != nil {
		fmt.Printf("Error seeding clientes: %v\n", err)
	} else {
		fmt.Println("Clientes seeded successfully.")
	}

	// Seed Restaurantes
	_, err = db.Exec(`
		INSERT INTO restaurantes (nome, endereco, tipo_cozinha, horario_funcionamento) VALUES
		('Restaurante A', 'Rua A, 123', 'Italiana', '10:00-23:00'),
		('Restaurante B', 'Avenida B, 456', 'Japonesa', '11:00-22:00'),
		('Restaurante C', 'Praça C, 789', 'Brasileira', '09:00-21:00')
	`)
	if err != nil {
		fmt.Printf("Error seeding restaurantes: %v\n", err)
	} else {
		fmt.Println("Restaurantes seeded successfully.")
	}

	// Seed Mesas
	_, err = db.Exec(`
		INSERT INTO mesas (restaurante_id, numero, capacidade, disponivel) VALUES
		(1, 1, 4, TRUE),
		(1, 2, 2, TRUE),
		(2, 1, 6, TRUE),
		(3, 1, 8, TRUE)
	`)
	if err != nil {
		fmt.Printf("Error seeding mesas: %v\n", err)
	} else {
		fmt.Println("Mesas seeded successfully.")
	}

	// Seed Reservas
	// Ensure DataHora is in a format compatible with SQLite (YYYY-MM-DD HH:MM:SS)
	now := time.Now()
	futureTime1 := now.Add(24 * time.Hour).Format("2006-01-02 15:04:05")
	futureTime2 := now.Add(48 * time.Hour).Format("2006-01-02 15:04:05")

	_, err = db.Exec(fmt.Sprintf(`
		INSERT INTO reservas (cliente_id, mesa_id, data_hora, numero_pessoas) VALUES
		(1, 1, '%s', 2),
		(2, 3, '%s', 4)
	`, futureTime1, futureTime2))
	if err != nil {
		fmt.Printf("Error seeding reservas: %v\n", err)
	} else {
		fmt.Println("Reservas seeded successfully.")
	}
}
