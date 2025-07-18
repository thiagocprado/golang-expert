module github.com/thiagocprado/golang-expert/packaging/mod

go 1.24.3

// o go.mod é um arquivo de configuração do Go Modules
// ele define o módulo, a versão do Go e as dependências do projeto

// go mod tidy remove as dependências não utilizadas e adiciona as necessárias

// o go.sum é um arquivo que contém o hash das dependências
// ele é gerado automaticamente pelo Go Modules
// garante que as dependências sejam baixadas corretamente
// é importante manter o go.sum atualizado 

// dependencias indiretas são aquelas que não são usadas diretamente 
// no projeto, mas são necessárias para o funcionamento de outras dependências
// elas são adicionadas automaticamente pelo Go Modules

require github.com/google/uuid v1.6.0
