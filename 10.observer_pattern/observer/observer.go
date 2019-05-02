package observer

import "fmt"

type Observer interface {
	Update(subject *Subject)
}

type Teacher struct{}

func (t *Teacher) Update(subject *Subject) {
	fmt.Println("status: ", subject.distance)
}
