package util

import "slices"

func SortDesc(s []int) {
	slices.SortFunc(s, func(a, b int) int {
		var res int

		if a < b {
			res = 1
		} else if a > b {
			res = -1
		}

		return res

	})
}
