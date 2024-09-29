package greeter

import "fmt"

// Greeter 구조체 정의
type Greeter struct {
	Name string
}

// NewGreeter는 새로운 Greeter 객체를 생성하는 함수
func NewGreeter(name string) *Greeter {
	return &Greeter{Name: name}
}

// Greet는 Greeter 구조체의 메서드로, 인사를 출력함
func (g *Greeter) Greet() {
	fmt.Printf("Hello, %s! Welcome to the Go project.\n", g.Name)
}
