package memento

type Originator struct {
	State string
}

func (e *Originator) CreateMemento() *Memento {
	return &Memento{State: e.State}
}

func (e *Originator) RestoreMemento(m *Memento) {
	e.State = m.GetSavedState()
}

func (e *Originator) SetState(State string) {
	e.State = State
}

func (e *Originator) GetState() string {
	return e.State
}
