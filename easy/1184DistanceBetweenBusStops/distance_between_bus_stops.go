package distancebetweenbusstops

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	i, j := 0, 0
	if start > destination {
		start, destination = destination, start
	}

	positive := 0
	i, j = start, destination
	for i < j {
		positive += distance[i]
		i++
	}

	reverse := 0
	i, j = start, destination
	for j%len(distance) != i {
		if j >= len(distance) {
			j = j - len(distance)
		}
		reverse += distance[j]
		j++
	}
	if positive < reverse {
		return positive
	}
	return reverse
}
