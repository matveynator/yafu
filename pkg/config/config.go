package Config

import (
  "os"
  "fmt"
  "flag"

)
var CompileVersion string

type Settings struct {
  APP_NAME, VERSION, FILE_PATH string
}

func isFlagPassed(name string) bool {
  found := false
  flag.Visit(func(f *flag.Flag) {
    if f.Name == name {
      found = true
    }
  })
  return found
}


func ParseFlags() (config Settings)  {
  config.APP_NAME = "yafu"
  flagVersion := flag.Bool("version", false, "Output version information")
  flag.StringVar(&config.FILE_PATH, "file", "candidates.txt", "Provide path to file with canadidates.")


  //process all flags
  flag.Parse()

  if *flagVersion  {
    if config.VERSION != "" {
      fmt.Println("Version:", config.VERSION)
    }
    os.Exit(0)
  }

  if config.FILE_PATH == "" {
    fmt.Println("Укажите файл с кандидатами")
    os.Exit(0)
  }

  // Startup banner START:
  fmt.Println("Starting %s ", config.APP_NAME)
  if config.VERSION != "" {
    fmt.Println("Version %s ", config.VERSION)
  }

  return
}
