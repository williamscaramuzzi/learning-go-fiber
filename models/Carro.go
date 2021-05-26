package models

type Carro struct {
	Marca string `json:"marca" form:"marca"`
	Nome  string `json:"nome" form:"nome"`
	Cor   string `json:"cor" form:"cor"`
}

func (c Carro) ToString() (retorno string) {
	retorno = c.Marca + ` ` + c.Nome + ` ` + c.Cor
	return
}
func NewCarro() {

}
