module github.com/thiagocprado/golang-expert/packaging/replace/system

go 1.24.3

// o go mod edit --replace é usado para substituir um módulo por outro localmente
// isso é útil para desenvolvimento local, testes ou quando você precisa usar
// uma versão específica de um módulo
// o comando go mod edit --replace <module>@<version> => <local-path> substitui o
// módulo especificado por um caminho local
// isso não afeta o repositório remoto, apenas a sua cópia local do módulo

replace github.com/thiagocprado/golang-expert/packaging/replace/math => ../math

require github.com/thiagocprado/golang-expert/packaging/replace/math v0.0.0-00010101000000-000000000000
