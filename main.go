package main

import (
	repositories "clean-go/repositories"
	"fmt"
)

func main() {
	// Importação dos Repositórios. (TX)

	// Criar construtor do repositorio enviando a eles os objetos a ser aceitos.

	// EX: TransactionRepository necessita da entidade Transaction para executar suas funções.

	get := repositories.Repository.GetTransactions

	fmt.Println(get)

	// --------------------------------------------------------------------------

	// Importação dos Casos de Uso. (GET TX)

	// Importação dos Controlladores. (GET TX)

	// Associação dos Repositório nos Casos de Uso.

	// Associação dos Casos de Uso nos Controlladores.

	// Definição das Rotas.
}
