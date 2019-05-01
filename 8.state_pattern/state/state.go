package state

type StatePatternYumichan struct {
	state State
}

func (yumi *StatePatternYumichan) MorningGreet() string {
	return yumi.state.MorningGreet()
}

func (yumi *StatePatternYumichan) GetProtectionForCold() string {
	return yumi.state.GetProtectionForCold()
}

func (yumi *StatePatternYumichan) ChangeState(state State) {
	yumi.state = state
}

type State interface {
	MorningGreet() string
	GetProtectionForCold() string
}

type BadMoodState struct{}

func (s *BadMoodState) MorningGreet() string {
	return "おお"
}

func (s *BadMoodState) GetProtectionForCold() string {
	return "ももひき"
}

type OrdinaryState struct{}

func (*OrdinaryState) MorningGreet() string {
	return "おっす！"
}

func (*OrdinaryState) GetProtectionForCold() string {
	return "走る"
}
