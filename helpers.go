package main

import (
    "net/http"
    "fmt"
    "log"
    "os"
    "encoding/json"
)
func getWithHeaders[T comparable](url string, data T) {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatalln(err)
    }

    token := os.Getenv("TOKEN")
    req.Header.Set("Authorization", fmt.Sprintf("token %s", token))
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatalln(err)
    }

    defer resp.Body.Close()

    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        log.Fatalln(err)
    }
}
