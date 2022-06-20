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
                http.Error(w, "Not Found", http.StatusNotFould)
                return
        }
        if r.Method != http.MethodGet {
                http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
                return
        }
        http.ServeFile(w, r, "home.html")
}


func main() {
        flag.Parse()
        hub := newHub()
        go hub.run()
        http.HandleFunc("/", servHome)
        http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
                serveWs(hub, w, r)
        })
        err := http.ListenAndServe(*addr, nil)
        if err != nil {
                log.Fatal("ListenAndServe: ", err)
        }
}
