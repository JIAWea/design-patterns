package strategy

const (
	DOG = iota + 1
	CAT
)

// NewAnimalFactory 生成具体策略的工厂
type NewAnimalFactory struct {
}

// GetAnimal 将具体的策略实现封装起来，交由工厂来实现
func (a NewAnimalFactory) GetAnimal(animalType int) AnimalFactory {
	var animal AnimalFactory
	switch animalType {
	case DOG:
		d := &Dog{}
		d.Name = "allen"
		animal = d
	case CAT:
		animal = &Cat{}
	}
	return animal
}
