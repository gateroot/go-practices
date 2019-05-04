package builder

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder}
}

func (d *Director) Construct() {
	d.builder.AddSolvent(100)
	d.builder.AddSolute(40)
	d.builder.AbandonSolution(70)
	d.builder.AddSolvent(100)
	d.builder.AddSolute(15)
}

type Builder interface {
	AddSolute(soluteAmount float64)
	AddSolvent(solventAmount float64)
	AbandonSolution(solutionAmount float64)
	GetResult() interface{}
}

type SaltWaterBuilder struct {
	saltWater SaltWater
}

func (builder *SaltWaterBuilder) AddSolute(soluteAmount float64) {
	builder.saltWater.Salt += soluteAmount
}

func (builder *SaltWaterBuilder) AddSolvent(solventAmount float64) {
	builder.saltWater.Water += solventAmount
}

func (builder *SaltWaterBuilder) AbandonSolution(solutionAmount float64) {
	sw := builder.saltWater
	saltDelta := solutionAmount * (sw.Salt / (sw.Water + sw.Salt))
	waterDelta := solutionAmount * (sw.Salt / (sw.Water + sw.Salt))
	sw.Salt -= saltDelta
	sw.Water -= waterDelta
}

func (builder SaltWaterBuilder) GetResult() interface{} {
	return builder.saltWater
}

func NewSaltWaterBuilder() Builder {
	return &SaltWaterBuilder{SaltWater{0, 0}}
}

type SaltWater struct {
	Water float64
	Salt  float64
}
