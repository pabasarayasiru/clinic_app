package controllers

import "strconv"

func toUint(s string) uint {
	id, _ := strconv.Atoi(s)
	return uint(id)
}
