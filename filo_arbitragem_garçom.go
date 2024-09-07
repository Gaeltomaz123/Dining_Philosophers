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
}

// Função para filósofo pensar
func (p *Philosopher) think() {
    fmt.Printf("Filósofo %d está pensando.\n", p.id)
    time.Sleep(time.Millisecond * 500)
}

// Função para filósofo comer
func (p *Philosopher) eat() {
    fmt.Printf("Filósofo %d está comendo.\n", p.id)
    time.Sleep(time.Millisecond * 500)
    p.eatCount++
}

// Função principal de execução do filósofo
func (p *Philosopher) dine(wg *sync.WaitGroup, forks []sync.Mutex, waiter *sync.Mutex) {
    defer wg.Done()

    for i := 0; i < 3; i++ {
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

    // Criar filósofos
    philosophers := make([]Philosopher, numPhilosophers)
    for i := 0; i < numPhilosophers; i++ {
        philosophers[i] = Philosopher{
            id:        i + 1,
            leftFork:  i,
            rightFork: (i + 1) % numPhilosophers,
        }
    }

    // Iniciar filósofos
    for i := 0; i < numPhilosophers; i++ {
        wg.Add(1)
        go philosophers[i].dine(&wg, forks, waiter)
    }

    wg.Wait()

    fmt.Println("Todos os filósofos terminaram de comer.")
}
