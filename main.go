package main

import "fmt"

func main() {
    fmt.Println("Fetching todays issues you've been working on...")

    
    issues := getIssues()

    for _, issue := range issues {
        fmt.Println(issue)
    }
}
