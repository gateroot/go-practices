package prototype

type Cloneable interface {
	CreateClone() Cloneable
}

type Paper struct {
	Name string
}

func (p Paper) CreateClone() Cloneable {
	paper := Paper{Name: p.Name}
	return paper
}

type PrototypeKeeper struct {
	prototypeMap map[string]Cloneable
}

func NewPrototypeKeeper(prototypeMap map[string]Cloneable) PrototypeKeeper {
	return PrototypeKeeper{prototypeMap}
}

func (keeper PrototypeKeeper) AddCloneable(key string, cloneable Cloneable) {
	keeper.prototypeMap[key] = cloneable
}

func (keeper PrototypeKeeper) GetClone(key string) Cloneable {
	prototype := keeper.prototypeMap[key]
	return prototype.CreateClone()
}
