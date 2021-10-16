package utils

import "strconv"

func ParseUint(arr []string) ([]uint, error) {
	var numberUint []uint
	for _, s := range arr {
		id, err := strconv.ParseUint(s, 10, 32)

		if err != nil {
			return nil, err
		}

		numberUint = append(numberUint, uint(id))
	}

	return numberUint, nil
}
