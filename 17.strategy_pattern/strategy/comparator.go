package strategy

type Comparator interface {
	Compare(h1, h2 Human) int
}

type AgeComparator struct{}

func (c *AgeComparator) Compare(h1, h2 Human) int {
	if h1.age > h2.age {
		return 1
	} else if h1.age == h2.age {
		return 0
	} else {
		return -1
	}
}

type HeightComparator struct{}

func (c *HeightComparator) Compare(h1, h2 Human) int {
	if h1.height > h2.height {
		return 1
	} else if h1.height == h2.height {
		return 0
	} else {
		return -1
	}
}

type WeightComparator struct{}

func (c *WeightComparator) Compare(h1, h2 Human) int {
	if h1.weight > h2.weight {
		return 1
	} else if h1.weight == h2.weight {
		return 0
	} else {
		return -1
	}
}
