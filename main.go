package main

import (
    "net/http"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "os"
)

var (
    MyFile  *log.Logger
)

func Init(
    traceHandle io.Writer,
    infoHandle io.Writer,
    warningHandle io.Writer,
    errorHandle io.Writer) {

    file, _ := os.OpenFile("log/production.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    
    MyFile = log.New(file,
        "PREFIX: ",
        log.Ldate|log.Ltime|log.Lshortfile)    
}

func handler(w http.ResponseWriter, r *http.Request) { 
    MyFile.Println("Inside handler")
    fmt.Println("Inside handler")
    fmt.Fprintf(w, "Hello ShepHertz")
}

func main() {
    Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

    MyFile.Println("Here is log....")
    
    http.HandleFunc("/", handler)
    fmt.Println("Listening on port 3000...")
    if err := http.ListenAndServe("localhost:3000", nil); err != nil {
        fmt.Println("Error:", err)
    }
}
