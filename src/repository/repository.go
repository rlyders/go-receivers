package repository

import (
	"github.com/rlyders/go-receivers/src/model"
)

func CreateCatAlso(id int) *model.Cat {
	cat := new(model.Cat)
	cat.setId(id) // <<<=== COMPILE ERROR: cat.setId undefined (type *model.Cat has no field or method setId)compilerMissingFieldOrMethod
	return cat
}
