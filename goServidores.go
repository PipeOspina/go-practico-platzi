package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
    inicio := time.Now()
    canal := make(chan string)

    servidores := []string{
        "https://platzi.com",
        "https://google.com",
        "https://facebook.com",
        "https://instagram.com",
    }

    for _, servidor := range servidores {
        go revisarServidor(servidor, canal)
    }

    for i := 0; i < len(servidores); i++ {
        fmt.Println(<-canal)
    }

    tiempo := time.Since(inicio)
    fmt.Printf("Tiempo de ejecucion %s\n", tiempo)
}

func revisarServidor(servidor string, canal chan string) {
    _, err := http.Get(servidor)
    if err != nil {
        canal <- servidor + " no está disponible"
    } else {
        canal <- servidor + " está funcionando normalmente"
    }
}
