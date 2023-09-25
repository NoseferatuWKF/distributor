package main

import (
    "fmt"
    "log"
    "bufio"
    "os"
)

func main() {
    var input []string

    defer output(&input)

    ingest(&input)
}

func die(s string) {
    log.Fatalln("\x1b[;31;1m" + s + "\x1b[;0m")
}

func ingest(input *[]string) {
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Err() != nil {
        die("Malformed Input")
    }

    for {
        scanner.Scan()
        line := scanner.Text()

        if len(line) == 0 {
            break
        }

        *input = append(*input, line)
    }

    // this should not break UX, ie; the browser should be opened up immdiately
    // while running tasks in the background
    go processInput(input)
}

func processInput(input *[]string) {
    validate(input)
    distribute(input)
}

func validate(input *[]string) {
    // check if the url redirects as it might be a sign of a actor
    // check if arguments are valid
    // check if delivery destinations are valid
}

func distribute(input *[]string) {
    // targets for delivery usually dbs
    // no need to support raw files as there are already utils for that
    targets := os.Args[1:]

    for _, t := range targets{
        go fmt.Println(t)
    }
}

// pipe to next process stdin
func output(input *[]string) {
    for _, v := range *input {
        fmt.Println(v)
    }
}
