package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // Open the file for writing
    file, err := os.Create("output.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    // Create a new bufio.Writer
    writer := bufio.NewWriter(file)

    // Write lines to the file
    lines := []string{"Hello, world!", "This is a test.", "Buffered writing with bufio."}
    for _, line := range lines {
        _, err := writer.WriteString(line + "\n")
        if err != nil {
            fmt.Println("Error writing to file:", err)
            return
        }
    }

    // Make example to flush the buffer
    writer.Flush()
}
