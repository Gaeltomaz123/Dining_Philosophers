package main

import (
    "fmt"
    "sync"
    "time"
)

// Estrutura Filósofo
type Philosopher struct {
    id          int
    leftFork    int
    rightFork   int
    eatCount    int
    offset      int
}

// Função para filósofo pensar
func (p *Philosopher) think() {
    fmt.Printf("%*sT%d\n", p.offset, "", p.eatCount+1)
    time.Sleep(time.Duration(2) * time.Second)
}

// Função para filósofo comer
func (p *Philosopher) eat() {
    fmt.Printf("%*sE%d\n", p.offset, "", p.eatCount+1)
    time.Sleep(time.Duration(3) * time.Second)
    p.eatCount++
}

// Função principal de execução do filósofo
func (p *Philosopher) dine(wg *sync.WaitGroup, forks []sync.Mutex, waiter *sync.Mutex) {
    defer wg.Done()

    for i := 0; i < 5; i++ { // Alterar para 5 execuções
        p.think()

        // O filósofo pede permissão ao garçom para pegar os garfos
        waiter.Lock()

        // Pega os garfos adjacentes
        forks[p.leftFork].Lock()
        forks[p.rightFork].Lock()

        // Come
        p.eat()

        // Solta os garfos
        forks[p.rightFork].Unlock()
        forks[p.leftFork].Unlock()

        // Informa ao garçom que terminou
        waiter.Unlock()
    }
}

func main() {
    var wg sync.WaitGroup

    // Definir número de filósofos e garfos
    numPhilosophers := 5
    forks := make([]sync.Mutex, numPhilosophers)
    waiter := &sync.Mutex{} // Garçom

    // Definir deslocamento para formatação
    offsetStep := 10

    // Criar filósofos
    philosophers := make([]Philosopher, numPhilosophers)
    for i := 0; i < numPhilosophers; i++ {
        philosophers[i] = Philosopher{
            id:        i + 1,
            leftFork:  i,
            rightFork: (i + 1) % numPhilosophers,
            offset:    i * offsetStep,
        }
    }

    // Exibir cabeçalho formatado
    fmt.Printf("[P1]%*s[P2]%*s[P3]%*s[P4]%*s[P5]\n", offsetStep-4, "", offsetStep-4, "", offsetStep-4, "", offsetStep-4, "")

    start := time.Now()
    // Iniciar filósofos
    for i := 0; i < numPhilosophers; i++ {
        wg.Add(1)
        go philosophers[i].dine(&wg, forks, waiter)
    }

    wg.Wait()

    elapsed := time.Since(start)

    fmt.Printf("\nWaiter Dinner took %s\n\n", elapsed)
}
