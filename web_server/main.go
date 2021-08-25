package main

func main() {
	server := NewServer(":3000", Logger())
	server.Handle("/", HTTPMethodGET, HandleRoot)
	server.Handle("/api", HTTPMethodGET, HandleHome, CheckAuth())
	server.Listen()
}
