package jumping_on_the_clouds

func jumpingOnClouds(c []int32) int32 {
	var jump int32
	i := 0
	for {
		jump++
		if i+1 == len(c)-1 || i+2 == len(c)-1 {
			println(i, "jump to goal")
			break
		}
		if c[i+2] == 1 {
			println(i, "jump 1")
			i += 1
		} else {
			println(i, "jump 2")
			i += 2
		}
	}
	return jump
}
