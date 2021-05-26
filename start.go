package main

import (
	"database/sql"
	"fmt"
	. "gofiber/start/models" //importo com ponto antes: se eu não fizer isso, preciso sempre usar models.Carro... fazendo isso, posso usar somente Carro como se fosse nome de classe

	_ "github.com/go-sql-driver/mysql" //importando com underline antes: importa esse pacote somente pra inicialização dele. NO caso de bancos de dados, precisamos desse pacote só pra iniciar uma conexão
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	//iniciar conexão ao BD, guarda no heap, e me retorna um ponteiro pra ela (ou seja, ela é global enquanto esse servidor rodar)
	ConexaoDatabase, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/godb") //formato: "usuario:senha@tcp(127.0.0.1:porta)/nomeDoSCHEMA"
	if err != nil {
		panic(err.Error()) //se der erro, cancela tudo com PANIC passando a string Error
	}
	CreateCarroIfNotExists(ConexaoDatabase)
	defer ConexaoDatabase.Close() //já põe na pilha de execução a ordem de fechar a conexão quando essa função main acabar

	// Inicia o Template Engine HTML padrão do GOe
	engine := html.New("./templates", ".html")

	//aqui inicia o Framework Fiber
	app := fiber.New(fiber.Config{Views: engine})

	//rotas
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	app.Get("/carros", func(c *fiber.Ctx) error {
		return c.Render("carros", fiber.Map{
			"Frase": "Mais vale uma melancia na mão do que duas no caminhão",
		})
	})

	app.Post("/carros", func(c *fiber.Ctx) error {
		carro := new(Carro) //esse new cria uma instância no heap e retorna o endereço dela, portanto carro é ponteiro pra models.Carro
		c.BodyParser(carro) //pega os valores no body enviado via POST e coloca no objeto que eu defini
		fmt.Println(carro.ToString())

		InsertCarro(*carro, ConexaoDatabase)

		return c.Render("carros", fiber.Map{
			"Frase": "Você cadastrou o seguinte carro:",
			"Marca": carro.Marca,
			"Nome":  carro.Nome,
			"Cor":   carro.Cor,
		})
	})

	app.Get("/consulta", func(c *fiber.Ctx) error {
		return c.Render("consulta", fiber.Map{
			"Lista": nil,
		})
	})

	app.Post("/consulta", func(c *fiber.Ctx) error {
		carro := new(Carro)
		c.BodyParser(carro)
		listaCarro := GetCarros(*carro, ConexaoDatabase)
		//fmt.Println(listaCarro)
		return c.Render("consulta", fiber.Map{
			"Lista": listaCarro,
		})
	})

	app.Delete("/consulta", func(c *fiber.Ctx) error {
		carro := new(Carro)
		c.BodyParser(carro)
		fmt.Println(carro)
		DeleteCarro(*carro, ConexaoDatabase)
		return c.SendStatus(fiber.StatusOK)

	})

	app.Static("/public", "./public")

	app.Listen("127.0.0.1:3000")

}
