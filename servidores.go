package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
    inicio := time.Now()

    servidores := []string{
        "https://platzi.com",
        "https://google.com",
        "https://facebook.com",
        "https://instagram.com",
    }

    for _, servidor := range servidores {
        revisarServidor(servidor)
    }

    tiempo := time.Since(inicio)
    fmt.Printf("Tiempo de ejecucion %s\n", tiempo)
}

func revisarServidor(servidor string) {
    _, err := http.Get(servidor)
    if err != nil {
        fmt.Println(servidor, "no está disponible")
    } else {
        fmt.Println(servidor, "está funcionando normalmente")
    }
}
