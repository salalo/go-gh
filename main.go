package main

import "fmt"

func main() {
    fmt.Println("Fetching todays issues you've been working on...")
    issues := getIssues()

    for _, issue := range issues {
        for _, event := range getIssueEvents(issue.Events_Url) {
            // One positive event is enough to list issue as touched.
            if issueTouched(&event) && len(issue.Events) == 0 {
                issue.Events = append(issue.Events, event)
            }
        }

        if len(issue.Events) > 0 {
            fmt.Println(issue.Title)
        }
    }
}
