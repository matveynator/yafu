package main

import (
  "bufio"
  "fmt"
  "math/big"
  "os"
  "runtime"
  "strings"
  "sync"
  "time"
  "yafu/pkg/config"
)

func isMaybePrime(p *big.Int) bool {
  return p.ProbablyPrime(1) // один тест Миллера-Рабина
}

func isPrime(p *big.Int) bool {
  return p.ProbablyPrime(50) // 50 тестов Миллера-Рабина для уверенности
}

func worker(id int, jobs <-chan *big.Int, results chan<- string, wg *sync.WaitGroup) {
  for {
    select {
    case candidate, ok := <-jobs:
      if !ok {
        return
      }
      startTime := time.Now()
      maybePrime := isMaybePrime(candidate)
      strictPrime := isPrime(candidate)
      duration := time.Since(startTime)
      results <- fmt.Sprintf("Worker %d: %s: isMaybePrime: %v, isPrime: %v, Duration: %d ns", id, candidate.String(), maybePrime, strictPrime, duration.Nanoseconds())
      wg.Done()
    default:
      time.Sleep(50 * time.Millisecond) // небольшая задержка, чтобы не загружать CPU
    }
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
  results := make(chan string, len(candidates))

  var wg sync.WaitGroup
  for i := 0; i < numWorkers; i++ {
    go worker(i, jobs, results, &wg)
  }

  for _, candidate := range candidates {
    wg.Add(1)
    jobs <- candidate
  }
  close(jobs)

  go func() {
    wg.Wait()
    close(results)
  }()

  for r := range results {
    fmt.Println(r)
  }
}

