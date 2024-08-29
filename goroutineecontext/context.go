package goroutineecontext

/*
O pacote `context` em Go é uma ferramenta poderosa para gerenciar prazos, cancelamentos e a passagem de valores entre goroutines de forma segura e eficiente.
Ele é amplamente utilizado para controlar a vida útil de operações assíncronas, como chamadas a APIs, consultas a bancos de dados ou qualquer outra operação que
 precise ser cancelada ou tenha um tempo limite.

### Como o `context` Funciona

O `context` é essencialmente uma forma de passar informações sobre a operação em andamento para as goroutines. Ele permite que você:

1. **Propague cancelamentos**: Se uma operação precisa ser cancelada, todas as goroutines que dependem desse contexto serão informadas.
2. **Defina prazos (deadlines)**: Você pode estabelecer um prazo máximo para a execução de uma operação.
3. **Compartilhe valores**: O `context` também permite o armazenamento de valores específicos para uma operação.

### Criando Contextos

O pacote `context` oferece várias formas de criar contextos, sendo os principais:

- **`context.Background()`**: Cria um contexto vazio, normalmente usado como ponto de partida para outros contextos.
- **`context.TODO()`**: Similar ao `Background()`, mas indica que o código precisa ser revisado, pois o contexto específico ainda não foi decidido.

### Cancelamento com `context.WithCancel`

Vamos criar um exemplo onde usamos `context.WithCancel` para cancelar uma operação.

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func tarefa(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Tarefa %d cancelada\n", id)
			return
		default:
			fmt.Printf("Tarefa %d em andamento\n", id)
			time.Sleep(1 * time.Second) // Simula trabalho
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	for i := 1; i <= 3; i++ {
		go tarefa(ctx, i)
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Cancelando as tarefas...")
	cancel()

	time.Sleep(1 * time.Second) // Espera as goroutines terminarem
	fmt.Println("Programa encerrado")
}
```

### Explicação

- **`context.WithCancel`**: Criamos um contexto com uma função de cancelamento. Quando `cancel()` é chamado, o canal `ctx.Done()` é fechado, sinalizando
 a todas as goroutines que elas devem encerrar.
- **Goroutines**: As goroutines escutam o canal `ctx.Done()` para saber quando devem parar a execução.

### Prazos com `context.WithTimeout`

Agora, vejamos como usar `context.WithTimeout` para definir um prazo máximo para uma operação.

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func tarefa(ctx context.Context, id int) {
	select {
	case <-time.After(2 * time.Second): // Simula uma tarefa que leva 2 segundos
		fmt.Printf("Tarefa %d concluída\n", id)
	case <-ctx.Done():
		fmt.Printf("Tarefa %d cancelada devido ao timeout\n", id)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	for i := 1; i <= 3; i++ {
		go tarefa(ctx, i)
	}

	time.Sleep(3 * time.Second) // Espera o tempo suficiente para ver os resultados
	fmt.Println("Programa encerrado")
}
```

### Explicação

- **`context.WithTimeout`**: Cria um contexto que será automaticamente cancelado após o prazo especificado (1 segundo, neste caso).
- **Cancelamento Automático**: As goroutines que ainda não completaram após 1 segundo são canceladas automaticamente.

### Passando Valores com `context.WithValue`

O `context` também permite armazenar e compartilhar valores entre diferentes partes de seu programa.

```go
package main

import (
	"context"
	"fmt"
)

func tarefa(ctx context.Context, id int) {
	valor := ctx.Value("chave").(string)
	fmt.Printf("Tarefa %d recebeu valor: %s\n", id, valor)
}

func main() {
	ctx := context.WithValue(context.Background(), "chave", "dados importantes")

	for i := 1; i <= 3; i++ {
		go tarefa(ctx, i)
	}

	// Espera para garantir que as goroutines terminem
	fmt.Scanln()
}
```

### Explicação

- **`context.WithValue`**: Cria um contexto que carrega um par chave/valor, permitindo que as goroutines acessem esse valor durante sua execução.
- **Segurança**: Usar `context.WithValue` para passar dados entre funções e goroutines evita a necessidade de variáveis globais ou outros métodos menos seguros.

### Conclusão

O `context` é uma ferramenta essencial em Go para gerenciar a execução de operações assíncronas, especialmente quando você precisa coordenar goroutines com cancelamentos,
prazos ou valores compartilhados.
Ele melhora a segurança e a eficiência do seu código, tornando-o mais robusto e fácil de manter.
*/