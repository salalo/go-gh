package main

import (
    "fmt"
    "os"
)

type Issue struct {
    Url string
    Title string
    Events_Url string
}

type Event struct {
    Actor Actor
    Event string
    Created_At string
}

type Actor struct {
    login string
}

func getIssues() []Issue {
    org := os.Getenv("ORG")
    repo := os.Getenv("REPO")
    url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", org, repo)

    var issues []Issue
    getWithHeaders(url, &issues)

    return issues
}

func getIssueEvents(eventUrl string) []Event {
    var events []Event
    getWithHeaders(eventUrl, &events)

    return events
}
