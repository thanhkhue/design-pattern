package strategy

import "fmt"

type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator Operator
}

func (o *Operation) Operate(leftValue, rightValue int) int {
	return o.Operator.Apply(leftValue, rightValue)
}

type Addition struct{}

func (a *Addition) Apply(leftValue, rightValue int) int {
	return leftValue + rightValue
}

type Multiple struct{}

func (m *Multiple) Apply(leftValue, rightValue int) int {
	return leftValue * rightValue
}

func init() {
	add := Operation{&Addition{}}
	fmt.Println(add.Operate(3, 5))

	multi := Operation{&Multiple{}}
	fmt.Println(multi.Operate(3, 5))
}
