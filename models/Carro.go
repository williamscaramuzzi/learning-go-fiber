package models

type Carro struct {
	Nome  string `json:"nome" form:"nome"`
	Marca string `json:"marca" form:"marca"`
	Cor   string `json:"cor" form:"cor"`
}
type CarroPage struct {
	Nome  string
	Marca string
}

func (c Carro) ToString() (retorno string) {
	retorno = c.Marca + ` ` + c.Nome + ` ` + c.Cor
	return
}
func NewCarro() {

}
