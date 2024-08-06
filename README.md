# GoBank

GoBank é um pequeno sistema de gerenciamento de contas bancárias simples desenvolvido em Go (Golang). Este projeto foi criado com o objetivo de aprender e explorar as funcionalidades e sintaxe da linguagem Go, e manipulação de arquivos JSON para armazenamento de dados.

## Funcionalidades

- **Cadastro de Usuário:** Permite que novos usuários se registrem no sistema.
- **Login:** Usuários existentes podem fazer login utilizando seu nome de usuário e senha.
- **Gerenciamento de Conta:**
  - Visualização de informações do cliente.
  - Visualização de informações da conta.
  - Transferências entre contas.
  - Depósitos em conta.
  - Saques de conta.
- **Histórico de Transações:** Armazena e exibe o histórico de ações realizadas na conta do usuário.
- **Armazenamento Persistente:** Utiliza um arquivo JSON para salvar e carregar dados de clientes.

## Como Executar

1. **Baixe o Go em https://go.dev/doc/install**
2. **Clone o repositório:**
   ```bash
   git clone https://github.com/seu-usuario/gobank.git
   cd gobank
   go run main.go conta.go cliente.go
