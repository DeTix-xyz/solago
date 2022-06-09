package utils

import (
	"reflect"

	"github.com/deezdegens/solago"
)

func IndexOf[T solago.Account](elements []T, targets ...T) solago.AccountIndexes {
	indices := []uint8{}

	for _, target := range targets {
		for index, element := range elements {
			if reflect.DeepEqual(element, target) {
				indices = append(indices, uint8(index))
			}
		}
	}

	return indices
}
