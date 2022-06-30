package main

import (
    "fmt"
    "os"
)

type Issue struct {
    Url string
    Title string
    Events_Url string
    Events []Event
}

type Event struct {
    Event string
    Created_At string
    Actor struct {
        Login string
    }
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

func issueTouched(event *Event) bool {
    // these are actual gh api event action names
    ghEvents := []string { 
        "added_to_project",
        "closed",
        "reopened",
        "commented",
        "commited",
        "mentioned",
        "labeled",
        "moved_columns_in_project",
    }

    if event.Actor.Login == "salalo" && contains(ghEvents, event.Event) {
        return true
    } else {
        return false
    }
}

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}
