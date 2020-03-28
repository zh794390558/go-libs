package main

import (
	"fmt"
	"os"
	"bufio"
	"path/filepath"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
  file, err := os.Create(path)
  if err != nil {
    return err
  }
  defer file.Close()

  w := bufio.NewWriter(file)
  for _, line := range lines {
    fmt.Fprintln(w, line)
  }

  return nil
}

func main() {
	fmt.Println("Base: ", filepath.Base("/tmp/main.go"))
	fmt.Println("Ext: ", filepath.Ext("/tmp/main.go"))
	fmt.Println(filepath.Abs("/tmp/main.go"))
	fmt.Println(os.Hostname())

	filename := "tmp.go"
	file := filepath.Join("./uploaded/", filename)
	fmt.Println(file)

	lines, _ := readLines("./main.go")
	for _, line := range lines{
		fmt.Println(line)
  }
  

  tmp := "/tmp/ip-2-1584862655.168315.wav"
  dir, file := filepath.Split(tmp)
  fmt.Println(dir, file)
	return
}