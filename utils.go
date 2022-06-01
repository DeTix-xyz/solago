package solana

import "reflect"

func indexOf[T Account](elements []T, targets ...T) []uint8 {
	indices := []uint8{}

	for index, element := range elements {
		for _, target := range targets {
			if reflect.DeepEqual(element, target) {
				indices = append(indices, uint8(index))
			}
		}
	}

	return indices
}
