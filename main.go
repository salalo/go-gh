package main

import "fmt"

func main() {
    fmt.Println("Fetching todays issues you've been working on...")

    issues := getIssues()

    for _, issue := range issues {
        events := getIssueEvents(issue.Events_Url)

        for _, event := range events {
        fmt.Println(event)
        }

    }
}
