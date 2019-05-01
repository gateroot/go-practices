package command

import "fmt"

type Command interface {
	Execute()
	SetBeaker(beaker Beaker)
}

type AddSaltCommand struct {
	beaker *Beaker
}

func (cmd *AddSaltCommand) SetBeaker(beaker *Beaker) {
	cmd.beaker = beaker
}

func (cmd *AddSaltCommand) Execute() {
	fmt.Println("[Add Salt Experiment]")
	for {
		if !cmd.beaker.IsMelted() {
			break
		}
		cmd.beaker.AddSalt(1)
		cmd.beaker.Mix()
	}
	cmd.beaker.Note()
}

type AddWaterCommand struct {
	beaker *Beaker
}

func (cmd *AddWaterCommand) SetBeaker(beaker *Beaker) {
	cmd.beaker = beaker
}

func (cmd *AddWaterCommand) Execute() {
	fmt.Println("[Add Water Experiment]")
	for {
		if cmd.beaker.isMelted {
			break
		}
		cmd.beaker.AddWater(10)
		cmd.beaker.Mix()
	}
	cmd.beaker.Note()
}

type MakeSaltWaterCommand struct {
	beaker *Beaker
}

func (cmd *MakeSaltWaterCommand) SetBeaker(beaker *Beaker) {
	cmd.beaker = beaker
}

func (cmd *MakeSaltWaterCommand) Execute() {
	fmt.Println("[Make Salt Water Experiment]")
	for {
		if !cmd.beaker.isMelted {
			break
		}
		cmd.beaker.AddSalt(1)
		cmd.beaker.Mix()
	}
	cmd.beaker.Note()
}
