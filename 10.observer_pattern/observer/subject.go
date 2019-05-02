package observer

type Subject struct {
	observers []Observer
	distance  int
}

func (s *Subject) AddObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) NotifyObservers() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

type Student struct {
	Subject
}

func (s *Student) Run(distance int) {
	s.distance += distance
	s.NotifyObservers()
}
