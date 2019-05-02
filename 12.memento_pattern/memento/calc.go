package memento

type Calc struct {
	temp int
}

func (calc *Calc) CreateMemento() *Memento {
	return &Memento{calc.temp}
}

func (calc *Calc) SetMemento(memento *Memento) {
	calc.temp = memento.result
}

func (calc *Calc) Plus(plus int) {
	calc.temp += plus
}

func (calc *Calc) GetTemp() int {
	return calc.temp
}

type Memento struct {
	result int
}
