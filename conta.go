package main

import (
	"errors"
	"fmt"
)

type Conta struct {
	Numero string  `json:"numero"`
	Saldo  float64 `json:"Saldo"`
}

func (conta *Conta) sacar(valor float64) error {
	if valor <= 0 {
		return errors.New("\nO valor deve ser positivo maior que 0")
	}
	if conta.Saldo < valor {
		return errors.New("\nSaldo insuficiente")
	}
	conta.Saldo -= valor
	fmt.Printf("\nSaque realizado com sucesso")
	return nil
}

func (conta *Conta) depositar(valor float64) error {
	if valor <= 0 {
		return errors.New("\nO valor deve ser positivo maior que 0")
	}
	conta.Saldo += valor
	fmt.Printf("\nDeposito realizado com sucesso")
	return nil
}

func (remetente *Conta) transferir(valor float64, beneficiario *Cliente) error {
	if valor <= 0 {
		return errors.New("\nO valor deve ser positivo maior que 0")
	}
	if remetente.Saldo < valor {
		return errors.New("\nSaldo insuficiente")
	}
	(*remetente).Saldo -= valor
	(*beneficiario).Conta.Saldo += valor
	return nil
}
