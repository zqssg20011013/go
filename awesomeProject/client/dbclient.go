package main

import (
	_ "google.golang.org/grpc"
)

func main() {
	result := like("koo")
	print(result)
}
