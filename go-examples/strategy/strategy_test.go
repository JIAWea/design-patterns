package strategy

import (
	"testing"
)

func TestNewAnimalFactory_GetAnimal(t *testing.T) {
	type args struct {
		animalType int
	}
	tests := []struct {
		name string
		args args
	}{
		{"name", args{CAT}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewAnimalFactory{}
			got := a.GetAnimal(tt.args.animalType)
			got.Speak()
		})
	}
}
