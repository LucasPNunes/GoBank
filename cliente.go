package main

import "time"

type Cliente struct {
	Nome      string      `json:"nome"`
	Sobrenome string      `json:"sobrenome"`
	CPF       string      `json:"cpf"`
	User      User        `json:"user"`
	Conta     Conta       `json:"conta"`
	Historico []Historico `json:"historico"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Historico struct {
	Tipo  string    `json:"data"`
	Valor float64   `json:"valor"`
	Date  time.Time `json:"date"`
}
