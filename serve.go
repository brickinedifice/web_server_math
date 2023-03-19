package main

import (
        //"fmt"
        "log"
        "net/http"
)

/*func handler(w http.ResponseWriter, r *http.Request) {

        http.ServeFile(w, r, "./")

        return
}*/

func main() {
    log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("./"))))
}
