package main

//historico??
import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	clientes := getclientes()
	option := welcome()
	if option == "login" {
		usuario := login(&clientes)
		menu(usuario, &clientes)
	} else {
		clientes, user := register(&clientes)
		menu(user, &clientes)
	}
	saveChanges(&clientes)
}

func menu(usuario *Cliente, clientes *[]Cliente) {
	rodando := true
	for rodando {
		fmt.Printf("\n----------------------------------------------------")
		fmt.Printf("\n\t\t\tGoBank")
		fmt.Printf("\n----------------------------------------------------")
		fmt.Println("\n\t 1 - Informações do Cliente")
		fmt.Println("\n\t 2 - Informações da Conta")
		fmt.Println("\n\t 3 - Transferir")
		fmt.Println("\n\t 4 - Depositar")
		fmt.Println("\n\t 5 - Sacar")
		fmt.Println("\n\t 6 - Histórico de ações")
		fmt.Println("\n\t 7 - Sair")

		var opcao int
		fmt.Print("\n\nEscolha uma opção: ")
		fmt.Scan(&opcao)
		switch opcao {
		case 1:
			fmt.Printf("\nNome: %s %s", usuario.Nome, usuario.Sobrenome)
			fmt.Printf("\nCPF: %s", usuario.CPF)
			fmt.Printf("\nContas: %s", usuario.Conta.Numero)
			break
		case 2:
			fmt.Printf("\nNúmero da Conta: %s", usuario.Conta.Numero)
			fmt.Printf("\nSaldo: %.2f", usuario.Conta.Saldo)

			break
		case 3:
			fmt.Println("Digite o numero da Conta para qual quer trasferir:")
			var num string
			fmt.Scan(&num)
			Conta, err := findContaByNumber(clientes, num)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(Conta)
			fmt.Println("Digite o valor:")
			var valor float64
			fmt.Scan(&valor)
			erro := usuario.Conta.transferir(valor, Conta)
			if erro != nil {
				fmt.Println(err)
				break
			}
			criarHistorico("Transferência", valor, usuario)
			break
		case 4:
			fmt.Println("Digite o valor:")
			var valor float64
			fmt.Scan(&valor)
			err := usuario.Conta.depositar(valor)
			if err != nil {
				fmt.Println(err)
				break
			}
			criarHistorico("Deposito", valor, usuario)
			break

		case 5:
			fmt.Println("Digite o valor:")
			var valor float64
			fmt.Scan(&valor)
			err := usuario.Conta.sacar(valor)
			if err != nil {
				fmt.Println(err)
				break
			}
			criarHistorico("Saque", valor, usuario)
			break
		case 6:
			for i := 0; i < len(usuario.Historico); i++ {
				fmt.Printf("\n%s no valor %0.2f no dia %s.", usuario.Historico[i].Tipo, usuario.Historico[i].Valor, usuario.Historico[i].Date)
			}
			break
		case 7:
			rodando = false
			fmt.Println("\n\n\t\t\tObrigado por utilizar nosso sistema!")
			break
		default:
			fmt.Println("\nOpção inválida!")
			break
		}
	}
}

func getclientes() []Cliente {
	clientesFile, err := os.Open("clientes.json")
	if err != nil {
		panic(err)
	}

	fileByte, err := io.ReadAll(clientesFile)
	if err != nil {
		panic(err)
	}
	defer clientesFile.Close()

	var clientes []Cliente
	json.Unmarshal(fileByte, &clientes)
	return clientes
}

func saveChanges(clientes *[]Cliente) {
	file, err := os.OpenFile("clientes.json", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(clientes, "", "  ")
	if err != nil {
		panic(err)
	}
	file.Write(jsonData)
}

func findContaByNumber(clientes *[]Cliente, num string) (*Cliente, error) {
	for i := range *clientes {
		if (*clientes)[i].Conta.Numero == num {
			return &(*clientes)[i], nil
		}
	}
	return nil, fmt.Errorf("\nConta não encontrada.")
}

func login(clientes *[]Cliente) *Cliente {
	authenticaded := false
	var Cliente *Cliente
	var Username, Password string
	for authenticaded == false {
		fmt.Print("\nDigite seu usuário (CPF): ")
		fmt.Scan(&Username)
		fmt.Print("\nDigite sua senha: ")
		fmt.Scan(&Password)

		for i := range *clientes {
			if (*clientes)[i].User.Username == Username && (*clientes)[i].User.Password == Password {
				Cliente = &(*clientes)[i]
				authenticaded = true
			}
		}
		if Cliente == nil {
			fmt.Println("\nUsuário ou senha inválidos!")
		}
	}
	return Cliente
}

func welcome() string {
	fmt.Println("\n----------------------------------------------------")
	fmt.Printf("\t\t\tGoBank")
	fmt.Println("\n----------------------------------------------------")
	fmt.Println("\n\t\tBem-vindo ao GoBank!")
	fmt.Println("\nEste é um sistema de gerenciamento de Contas bancárias em Go.")
	fmt.Println("\nPara utilizar, faça seu login ou se registre.")
	fmt.Println("\n----------------------------------------------------")

	rodando := true
	for rodando {
		fmt.Print("\n\nDeseja fazer login ou registrar? (1 - Login, 2 - Registrar, 3 - Sair): ")
		var opcao int
		fmt.Scan(&opcao)
		switch opcao {
		case 1:
			return "login"
		case 2:
			return "register"
		case 3:
			fmt.Println("\n\n\tObrigado por utilizar nosso sistema!")
			os.Exit(0)
		default:
			fmt.Println("\nOpção inválida!")
		}
	}
	return ""
}

func register(clientes *[]Cliente) ([]Cliente, *Cliente) {
	var Nome, Sobrenome, cpf, senha, confirmarSenha string
	fmt.Println("\nDigite seu Nome: ")
	fmt.Scan(&Nome)
	fmt.Println("\nDigite seu Sobrenome: ")
	fmt.Scan(&Sobrenome)
	fmt.Println("\nDigite seu CPF (somente números): ")
	fmt.Scan(&cpf)
	fmt.Println("\nDigite sua senha: ")
	fmt.Scan(&senha)
	fmt.Println("\nComfirme sua senha: ")
	fmt.Scan(&confirmarSenha)

	if senha != confirmarSenha {
		fmt.Println("\nSenhas não conferem!")
		os.Exit(1)
	}

	if checkRegister(*clientes, cpf) == false {
		fmt.Println("\nCPF já cadastrado!")
		os.Exit(1)
	}

	c := Cliente{
		Nome:      Nome,
		Sobrenome: Sobrenome,
		CPF:       cpf,
		User: User{
			Username: cpf,
			Password: senha,
		},
		Conta: Conta{
			Numero: cpf,
			Saldo:  0,
		},
		Historico: []Historico{},
	}
	*clientes = append(*clientes, c)
	i := len(*clientes) - 1
	return *clientes, &(*clientes)[i]
}

func checkRegister(clientes []Cliente, cpf string) bool {
	for _, c := range clientes {
		if c.CPF == cpf {
			return false
		}
	}
	return true
}

func criarHistorico(tipo string, valor float64, c *Cliente) {
	h := Historico{
		Tipo:  tipo,
		Valor: valor,
		Date:  time.Now(),
	}
	c.Historico = append(c.Historico, h)
}
