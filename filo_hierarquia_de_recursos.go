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
func (p *Philosopher) dine(wg *sync.WaitGroup, forks []sync.Mutex) {
    defer wg.Done()

    for i := 0; i < 5; i++ { // Alterar para 5 execuções
        p.think()

        // Filósofo pega os garfos na ordem da hierarquia de recursos
        if p.leftFork < p.rightFork {
            forks[p.leftFork].Lock()
            forks[p.rightFork].Lock()
        } else {
            forks[p.rightFork].Lock()
            forks[p.leftFork].Lock()
        }

        p.eat()

        // Solta os garfos na ordem inversa
        if p.leftFork < p.rightFork {
            forks[p.rightFork].Unlock()
            forks[p.leftFork].Unlock()
        } else {
            forks[p.leftFork].Unlock()
            forks[p.rightFork].Unlock()
        }
    }
}

func main() {
    var wg sync.WaitGroup

    // Definir número de filósofos e garfos
    numPhilosophers := 5
    forks := make([]sync.Mutex, numPhilosophers)

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
        go philosophers[i].dine(&wg, forks)
    }

    wg.Wait()

    elapsed := time.Since(start)

    fmt.Printf("\nHierarchy Dinner took %s\n\n", elapsed)
}
