// os módulos podem ter o path do github, pois o mesmo pode ser baixado por outras pessoas
module packages_and_modules

go 1.23.3

// O **`go mod`** é o sistema de gerenciamento de dependências introduzido no Go a partir da versão 1.11.
// Ele utiliza módulos para organizar e controlar as dependências de um projeto, substituindo o modelo antigo baseado em **`GOPATH`**.
// Um módulo é definido por um arquivo **`go.mod`**, que fica localizado na raiz do projeto e descreve informações importantes sobre o módulo e suas dependências.

// ### Principais Funcionalidades do `go mod`:
// 1. **Inicializar um Módulo**:
//    O comando `go mod init <module_name>` cria o arquivo **`go.mod`** e inicia um módulo no projeto.

// 2. **Gerenciamento de Dependências**:
//    - Quando um pacote externo é utilizado, ele é automaticamente adicionado ao **`go.mod`**.
//    - O comando `go get <package>` permite adicionar, atualizar ou remover dependências manualmente.

// 3. **Controle de Versões**:
//    O **`go.mod`** especifica as versões exatas das dependências usadas, garantindo reprodutibilidade.

// 4. **Resolução Automática**:
//    Dependências são baixadas e armazenadas no cache local, otimizando a utilização.

// 5. **Verificação de Consistência**:
//    O arquivo **`go.sum`** é gerado automaticamente e armazena hashes das versões das dependências, garantindo que não sejam alteradas.

// ### Comandos Comuns:
// - **`go mod init <module_name>`**: Inicializa um novo módulo.
// - **`go get <package>`**: Adiciona ou atualiza dependências.
// - **`go mod tidy`**: Remove dependências não utilizadas e adiciona as necessárias.
// - **`go list -m all`**: Lista todas as dependências do módulo.
// - **`go mod graph`**: Mostra o grafo de dependências do projeto.

// ### Benefícios:
// - Trabalha fora do **`GOPATH`**, permitindo projetos em qualquer local no sistema de arquivos.
// - Garante que todos os desenvolvedores utilizem as mesmas versões de dependências.
// - Oferece um processo de build mais previsível e confiável.

// Em resumo, o **`go mod`** trouxe mais organização, flexibilidade e controle para o desenvolvimento de aplicações Go, tornando a gestão de dependências moderna e eficiente.

require github.com/google/uuid v1.6.0


// go.sum
// Verificação de Integridade:

// Ele armazena checksums (somas de verificação) das dependências do seu projeto para garantir que as versões exatas das bibliotecas usadas sejam consistentes e confiáveis.
// Essas somas são verificadas contra um repositório público chamado Go checksum database para prevenir manipulações ou alterações maliciosas.
// Garantia de Reprodutibilidade:

// Com o go.sum, qualquer pessoa que clone seu repositório e execute go mod tidy ou go build obterá exatamente as mesmas dependências, assegurando que o código funcione da mesma forma em diferentes ambientes.
// Complemento ao go.mod:

// O arquivo go.mod especifica as dependências e suas versões.
// O go.sum serve como um complemento, validando que essas dependências não foram alteradas desde a última vez que foram baixadas.