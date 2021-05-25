package main

import (
	"fmt"
	"github.com/gofiber/template/html"
	"gofiber/start/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize standard Go html template engine
	engine := html.New("./templates", ".html")

	//aqui inicia o Framework Fiber
	app := fiber.New(fiber.Config{Views: engine})

	//rotas
	app.Get("/", func(c *fiber.Ctx) error {
		if err := c.SendFile("./templates/index.html"); err != nil {
			c.Next()
			return err
		}
		return nil
	})

	app.Get("/carros", func(c *fiber.Ctx) error {
		return c.Render("carros", fiber.Map{
			"Frase": "Mais vale uma melancia na mão do que duas no caminhão",
		})
	})

	app.Post("/carros", func(c *fiber.Ctx) error {
		carro := new(models.Carro) //esse new cria uma instância no heap e retorna o endereço dela, portanto carro é ponteiro pra models.Carro
		c.BodyParser(carro)        //pega os valores no body enviado via POST e coloca no objeto que eu defini
		fmt.Println(carro.ToString())
		return c.Render("carros", fiber.Map{
			"Frase": "Você cadastrou o seguinte carro:",
			"Marca": carro.Marca,
			"Nome":  carro.Nome,
			"Cor":   carro.Cor,
		})
	})

	app.Static("/public", "./public")

	app.Listen("127.0.0.1:3000")

}
