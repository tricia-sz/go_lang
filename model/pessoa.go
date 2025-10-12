package model

import (
	"time"
)

type Pessoa struct {
	Nome             string
	Endereco         Endereco
	DataDeNascimento time.Time
	Idade            int
}

func (pessoa1 *Pessoa) CalculaIdade() {
	anoDeNascimento := pessoa1.DataDeNascimento.Year()
	anoAtual := time.Now().Year()
	pessoa1.Idade = anoAtual - anoDeNascimento
}
