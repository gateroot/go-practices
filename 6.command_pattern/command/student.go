package command

type Student struct{}

func (s *Student) Experiment() {
	addSalt := &AddSaltCommand{}
	addWater := &AddWaterCommand{}
	makeSaltWater := &MakeSaltWaterCommand{}

	addSalt.SetBeaker(NewBeaker(100, 0))
	addWater.SetBeaker(NewBeaker(0, 10))
	makeSaltWater.SetBeaker(NewBeaker(90, 10))

	addSalt.Execute()
	addWater.Execute()
	makeSaltWater.Execute()
}
