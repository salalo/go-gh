package main

import "fmt"

func main() {
    fmt.Println("Fetching todays issues you've been working on...")

    for _, issue := range getIssues() {
        c := make(chan []Event)
        go getIssueEvents(issue.Events_Url, c)
        events := <- c
        close(c)

        for _, event := range events {
            // One positive event is enough to list issue as touched.
            if issueTouched(&event) && len(issue.Events) == 0 {
                // Aappend for new features
                issue.Events = append(issue.Events, event)
            }
        }

        if len(issue.Events) > 0 {
            fmt.Println(issue.Title)
        }
    }
}
