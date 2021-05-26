package main

import (
	"database/sql"
	"fmt"
	. "gofiber/start/models"
)

func CreateCarroIfNotExists(conn *sql.DB) error {
	var q string
	q = `CREATE TABLE IF NOT EXISTS Carro(
carro_uuid varchar(37) PRIMARY KEY,
marca varchar(20) NOT NULL,
nome varchar(20) NOT NULL,
cor varchar(20) NOT NULL,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP) ENGINE=INNODB;`

	_, err := conn.Query(q)
	return err
}

func InsertCarro(c Carro, conn *sql.DB) error {
	_, err := conn.Query("INSERT INTO `godb`.`carro` VALUES (UUID(), ?, ?, ?, null);", c.Marca, c.Nome, c.Cor)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func GetCarros(c Carro, conn *sql.DB) []Carro {

	if c.Marca == "" {
		c.Marca = "%"
	}
	if c.Nome == "" {
		c.Nome = "%"
	}
	if c.Cor == "" {
		c.Cor = "%"
	}

	var vetorCarros []Carro

	results, err := conn.Query("SELECT * FROM `godb`.`carro` WHERE marca LIKE ? AND nome LIKE ? AND cor LIKE ?  ORDER BY marca ASC, nome ASC;", c.Marca, c.Nome, c.Cor)

	if err != nil {
		fmt.Println(err.Error())
	}

	//fmt.Println("Resultados: ", results)

	var descartar string
	var tempCar Carro //eu crio um carro temporário só pra transformar ROW em UM objeto carro, e já apenso ele ao meu vetorCarros que será retornado
	for i := 0; results.Next(); i++ {
		results.Scan(&descartar, &tempCar.Marca, &tempCar.Nome, &tempCar.Cor, &descartar)
		vetorCarros = append(vetorCarros, tempCar)
	}

	return vetorCarros
}

func UpdateCarro(c Carro, conn *sql.DB) {
	fmt.Println("Falta implementar UPDATE nas queries")
}

func DeleteCarro(c Carro, conn *sql.DB) {
	results, err := conn.Query("DELETE FROM `godb`.`carro` WHERE (`marca` = ? AND `nome` = ? AND `cor` = ?);", c.Marca, c.Nome, c.Cor)
	if err != nil {
		fmt.Println(results)
	}
}
