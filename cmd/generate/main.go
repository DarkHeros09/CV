package main

import (
	"cv-website/builder"
	"fmt"
)

func main() {
	builder.RunStaticGenerator()
	fmt.Println("dist/ generated successfully.")
}
