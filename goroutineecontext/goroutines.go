package goroutineecontext

/*
 Concorrência:
 é como fazer malabarismo com várias tarefas ao mesmo tempo, mas não necessariamente ao mesmo tempo (você troca entre elas rapidamente).

 Paralelismo
 é executar várias tarefas ao mesmo tempo, usando recursos (como múltiplos processadores) para fazer tudo simultaneamente.

 Go Routines (Green Threads):
 - 2~4 kbytes para alocar.
 - Custo menor que uma thread convencional.

 Condição de corrida:
 Se as tarefas forem executadas ao mesmo tempo, ocorrência e paralelismo são equivalentes.


 Em Go, você geralmente começa com concorrência, e o Go runtime decide se pode executar suas goroutines em paralelo,
 dependendo dos recursos disponíveis (como o número de núcleos de CPU).
*/

/*
Exemplo sem WaitGroup
Neste exemplo, vamos lançar várias goroutines que simulam a execução de tarefas, mas sem nenhuma sincronização para garantir que todas as goroutines sejam concluídas antes do programa terminar.

go
Copy code
package main

import (
	"fmt"
	"time"
)

func tarefa(id int) {
	fmt.Printf("Tarefa %d iniciada\n", id)
	time.Sleep(2 * time.Second) // Simula uma tarefa que leva 2 segundos
	fmt.Printf("Tarefa %d concluída\n", id)
}

func main() {
	for i := 1; i <= 5; i++ {
		go tarefa(i)
	}

	// Espera artificialmente (não é uma boa prática)
	time.Sleep(3 * time.Second)
	fmt.Println("Programa encerrado")
}
Explicação
Goroutines: Usamos go tarefa(i) dentro de um loop para iniciar várias goroutines.
Problema: Sem o uso de WaitGroup, o programa principal pode terminar antes que todas as goroutines terminem. Aqui, usamos time.Sleep(3 * time.Second) para tentar esperar as goroutines, mas isso é ineficaz e pouco confiável.
Exemplo com WaitGroup
Agora vamos ver como usar WaitGroup para garantir que todas as goroutines terminem antes de o programa principal encerrar.

go
Copy code
package main

import (
	"fmt"
	"sync"
	"time"
)

func tarefa(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Marca a goroutine como concluída quando a função terminar
	fmt.Printf("Tarefa %d iniciada\n", id)
	time.Sleep(2 * time.Second) // Simula uma tarefa que leva 2 segundos
	fmt.Printf("Tarefa %d concluída\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Incrementa o contador do WaitGroup
		go tarefa(i, &wg)
	}

	wg.Wait() // Aguarda até que todas as goroutines tenham chamado wg.Done()
	fmt.Println("Programa encerrado")
}
Explicação
WaitGroup: Utilizamos sync.WaitGroup para controlar quando todas as goroutines terminam.
wg.Add(1): Incrementa o contador do WaitGroup antes de iniciar cada goroutine, indicando que há uma nova goroutine em execução.
wg.Done(): Cada goroutine chama wg.Done() (usando defer para garantir que será chamado mesmo se a goroutine falhar), o que decrementa o contador do WaitGroup.
wg.Wait(): O programa principal espera aqui até que o contador do WaitGroup volte a zero, garantindo que todas as goroutines tenham concluído antes de encerrar o programa.
Conclusão
O uso de WaitGroup permite que você coordene as goroutines de maneira segura e eficiente. Sem WaitGroup, seu programa pode encerrar antes que as goroutines tenham concluído suas tarefas, o que pode levar a resultados inesperados. Com WaitGroup, você garante que todas as goroutines sejam executadas completamente antes que o programa termine.
*/