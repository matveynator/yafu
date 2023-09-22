package main

import (
  "bufio"             // Для чтения из файла
  "fmt"               // Для форматированного ввода/вывода
  "math/big"          // Для работы с большими целыми числами
  "os"                // Для работы с операционной системой, например, с файлами
  "runtime"           // Для получения информации о рантайме
  "strings"           // Для работы со строками
  "sync"              // Для синхронизации горутин
  "time"              // Для измерения времени
  "yafu/pkg/config"   // Импортирование конфигурации (специфично для вашего проекта)
)

// Функция для проверки, является ли число возможно простым
func isMaybePrime(p *big.Int) bool {
  return p.ProbablyPrime(1) // Один тест Миллера-Рабина
}

// Функция для строгой проверки, является ли число простым
func isPrime(p *big.Int) bool {
  return p.ProbablyPrime(50) // 50 тестов Миллера-Рабина для уверенности
}

// Рабочая горутина, которая проверяет числа на простоту
func worker(id int, jobs <-chan *big.Int, results chan<- string, wg *sync.WaitGroup) {
  for {
    select {
    case candidate, ok := <-jobs:
      if !ok {
        return // Выход из горутины, если канал jobs закрыт
      }
      startTime := time.Now() // Засекаем время начала
      maybePrime := isMaybePrime(candidate)
      strictPrime := isPrime(candidate)
      duration := time.Since(startTime) // Вычисляем продолжительность выполнения
      results <- fmt.Sprintf("Worker %d: %s: isMaybePrime: %v, isPrime: %v, Duration: %d ns", id, candidate.String(), maybePrime, strictPrime, duration.Nanoseconds())
      wg.Done() // Уменьшаем счетчик ожидающих горутин
    default:
      time.Sleep(50 * time.Millisecond) // Небольшая задержка, чтобы не загружать CPU
    }
  }
}

func main() {
  config := Config.ParseFlags() // Парсим флаги командной строки

  // Открываем файл
  file, err := os.Open(config.FILE_PATH)
  if err != nil {
    fmt.Println("Ошибка при открытии файла:", err)
    return
  }
  defer file.Close() // Закрываем файл при завершении main()

  scanner := bufio.NewScanner(file)
  var candidates []*big.Int
  // Считываем каждую строку из файла
  for scanner.Scan() {
    line := strings.TrimSpace(scanner.Text())
    n, ok := new(big.Int).SetString(line, 10)
    if ok {
      candidates = append(candidates, n) // Добавляем кандидата в список
    }
  }
  if err := scanner.Err(); err != nil {
    fmt.Println("Ошибка при чтении файла:", err)
    return
  }

  numWorkers := runtime.NumCPU() // Получаем количество ядер CPU
  jobs := make(chan *big.Int, len(candidates))
  results := make(chan string, len(candidates))

  var wg sync.WaitGroup
  // Запускаем рабочие горутины
  for i := 0; i < numWorkers; i++ {
    go worker(i, jobs, results, &wg)
  }

  // Помещаем кандидатов в канал jobs
  for _, candidate := range candidates {
    wg.Add(1) // Увеличиваем счетчик ожидающих горутин
    jobs <- candidate
  }
  close(jobs) // Закрываем канал jobs

  // Закрываем канал results, когда все горутины завершатся
  go func() {
    wg.Wait()
    close(results)
  }()

  // Выводим результаты
  for r := range results {
    fmt.Println(r)
  }
}

