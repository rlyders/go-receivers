package model

type ICat interface {
	setA(int)
}

type Cat struct {
	id int
}

func (c *Cat) setId(id int) {
	c.id = id
}

func CreateCat(id int) *Cat {
	cat := new(Cat)
	cat.setId(id) // <<<=== OK: this works just fine
	return cat
}
