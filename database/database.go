package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./yummer.db")
	if err != nil {
		log.Fatal(err)
	}

	createTables()
}

func createTables() {
	clienteTable := `CREATE TABLE IF NOT EXISTS clientes (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"nome" TEXT,
		"email" TEXT,
		"telefone" TEXT
	);`

	restauranteTable := `CREATE TABLE IF NOT EXISTS restaurantes (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"nome" TEXT,
		"endereco" TEXT,
		"tipo_cozinha" TEXT,
		"horario_funcionamento" TEXT
	);`

	mesaTable := `CREATE TABLE IF NOT EXISTS mesas (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"restaurante_id" INTEGER,
		"numero" INTEGER,
		"capacidade" INTEGER,
		"disponivel" BOOLEAN,
		FOREIGN KEY(restaurante_id) REFERENCES restaurantes(id)
	);`

	reservaTable := `CREATE TABLE IF NOT EXISTS reservas (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"cliente_id" INTEGER,
		"mesa_id" INTEGER,
		"data_hora" DATETIME,
		"numero_pessoas" INTEGER,
		FOREIGN KEY(cliente_id) REFERENCES clientes(id),
		FOREIGN KEY(mesa_id) REFERENCES mesas(id)
	);`

	query, err := DB.Prepare(clienteTable)
	if err != nil {
		log.Fatal(err)
	}
	query.Exec()

	query, err = DB.Prepare(restauranteTable)
	if err != nil {
		log.Fatal(err)
	}
	query.Exec()

	query, err = DB.Prepare(mesaTable)
	if err != nil {
		log.Fatal(err)
	}
	query.Exec()

	query, err = DB.Prepare(reservaTable)
	if err != nil {
		log.Fatal(err)
	}
	query.Exec()
}
