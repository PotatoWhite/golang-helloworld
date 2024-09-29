// cmd/helloworld/main.go
package main

import "helloworld/pkg/greeter"

func main() {
	// Greeter 패키지를 사용하여 인사 메시지 출력
	greeter := greeter.NewGreeter("Potato")
	greeter.Greet()
}
