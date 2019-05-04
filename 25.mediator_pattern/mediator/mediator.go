package mediator

import "fmt"

type Mediator interface {
	AddColleague(colleague Colleague)
	Consultation(colleague1 Colleague, to, msg string)
}

type Colleague interface {
	GetName() string
	Twitter(to, msg string)
	SetMediator(mediator Mediator)
}

type Saito struct {
	colleagueMap map[string]Colleague
}

func NewSaito() *Saito {
	return &Saito{make(map[string]Colleague)}
}

func (saito *Saito) AddColleague(colleague Colleague) {
	saito.colleagueMap[colleague.GetName()] = colleague
}

func (saito *Saito) Consultation(colleague1 Colleague, to, msg string) {
	colleague2 := saito.colleagueMap[to]
	fmt.Printf("receive message from %s to %s: %s", colleague1.GetName(), colleague2.GetName(), msg)
}

type Tanaka struct {
	mediator Mediator
}

func (tanaka *Tanaka) GetName() string {
	return "Tanaka"
}

func (tanaka *Tanaka) SetMediator(mediator Mediator) {
	tanaka.mediator = mediator
}

func (tanaka *Tanaka) Twitter(to, msg string) {
	tanaka.mediator.Consultation(tanaka, to, msg)
}

type Sato struct {
	mediator Mediator
}

func (sato *Sato) GetName() string {
	return "Sato"
}

func (sato *Sato) SetMediator(mediator Mediator) {
	sato.mediator = mediator
}

func (sato *Sato) Twitter(to, msg string) {
	sato.mediator.Consultation(sato, to, msg)
}
