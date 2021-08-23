package main

func main() {
    server := NewServer(":3000", Logger())
    server.Handle("/", HandleRoot)
    server.Handle("/api", HandleHome, CheckAuth())
    server.Listen()
}
