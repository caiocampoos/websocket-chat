package main

import (
        "flag"
        "log"
        "net/http"
)

var addr = flag.String("addr", "8080", "http service address")

func servHome(w http.ResponseWriter, r *http.Request) {
        log.Println(r.URL)
        if r.URL.Path != "/" {
                http.Error(w, "Not Found", htt.StatusNotFound)
                return
        }
        if r.Method != http.MethodGet {
                http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
                return
        }
        http.ServeFile(w, r, "home.html")
}

