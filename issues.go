package main

import (
    "net/http"
    "fmt"
    "log"
    "os"
    "encoding/json"
)

type Issue struct {
    Url string
    Title string
    Events_Url string
}

func getIssues() []Issue {
    org := os.Getenv("ORG")
    repo := os.Getenv("REPO")
    token := os.Getenv("TOKEN")
    url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", org, repo)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatalln(err)
    }

    req.Header.Set("Authorization", fmt.Sprintf("token %s", token))
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatalln(err)
    }

    defer resp.Body.Close()

    var issues []Issue
    if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
        log.Fatalln(err)
    }

    return issues
}
