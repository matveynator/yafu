package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"runtime"
	"strings"
	"sync"
  "yafu/pkg/config"
)

// isMaybePrime проверяет, возможно ли, что число простое.
func isMaybePrime(p *big.Int) bool {
	return true // Здесь должна быть реализация
}

// isPrime проверяет, простое ли число.
func isPrime(p *big.Int) bool {
	return p.ProbablyPrime(0)
}

func worker(id int, jobs <-chan *big.Int, wg *sync.WaitGroup) {
	for candidate := range jobs {
		// Здесь логика обработки кандидата, например проверка на простоту.
		fmt.Printf("Worker %d: %s\n", id, candidate.String())
		wg.Done()
	}
}

func main() {

  config := Config.ParseFlags()

	file, err := os.Open(config.FILE_PATH)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var candidates []*big.Int
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		n, ok := new(big.Int).SetString(line, 10)
		if ok {
			candidates = append(candidates, n)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	numWorkers := runtime.NumCPU()
	jobs := make(chan *big.Int, len(candidates))

	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		go worker(i, jobs, &wg)
	}

	for _, candidate := range candidates {
		wg.Add(1)
		jobs <- candidate
	}

	close(jobs)
	wg.Wait()
}

