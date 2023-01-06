package lightsOn

func turnOn(bulbs []int) int {
	counter := 0
	state := 0
	for i := range bulbs {
		if bulbs[i]^state == 0 {
			counter++
			state = state ^ 1
		}
	}
	return counter
}
