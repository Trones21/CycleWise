package main

import (
    "fmt"
    "os"
    "cyclewise/templates"
)

func main() {
    args := os.Args[1:]
    if len(args) == 0 {
        fmt.Println("Usage: cyclewise [command] [options]")
        return
    }

    command := args[0]
    switch command {
    case "plan":
        templates.GeneratePlan()
    case "retro":
        templates.GenerateRetro()
    default:
        fmt.Println("Unknown command:", command)
    }
}
