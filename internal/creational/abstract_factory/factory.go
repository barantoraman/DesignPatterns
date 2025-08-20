package abstractfactory

import "errors"

type FurnitureFactory interface {
	MakeChair() Chair
	MakeSofa() Sofa
}

func GetFurnitureFactory(furnitureType string) (FurnitureFactory, error) {
	switch furnitureType {
	case "modern":
		return &modern{}, nil
	case "victorian":
		return &victorian{}, nil
	default:
		return nil, errors.New("invalid furniture type name")
	}
}
