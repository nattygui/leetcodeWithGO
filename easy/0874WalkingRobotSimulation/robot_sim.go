package robotsim

type stand struct {
	x      int
	y      int
	direct int
}

func robotSim(commands []int, obstacles [][]int) int {
	site := &stand{
		x:      0,
		y:      0,
		direct: 0,
	}
	max := 0
	for _, command := range commands {
		switch {
		case command == -2:
			site.direct = (site.direct - 90 + 360) % 360
		case command == -1:
			site.direct = (site.direct + 90) % 360
		default:
			site.move(command, obstacles)
			if max < site.x*site.x+site.y*site.y {
				max = site.x*site.x + site.y*site.y
			}
		}
	}
	return max
}

func (s *stand) move(distance int, obstacles [][]int) {
	switch s.direct {
	case 0:
		min := 30001
		for _, obstacle := range obstacles {
			if obstacle[0] != s.x {
				continue
			}
			if s.y < obstacle[1] && (s.y+distance) >= obstacle[1] {
				if min > obstacle[1]-1 {
					min = obstacle[1] - 1
				}
			}
		}
		if min != 30001 {
			s.y = min
		} else {
			s.y += distance
		}
	case 90:
		min := 30001
		for _, obstacle := range obstacles {
			if obstacle[1] != s.y {
				continue
			}
			if s.x < obstacle[0] && (s.x+distance) >= obstacle[0] {
				if min > obstacle[0]-1 {
					min = obstacle[0] - 1
				}
			}
		}
		if min != 30001 {
			s.x = min
		} else {
			s.x += distance
		}
	case 180:
		min := -30001
		for _, obstacle := range obstacles {
			if obstacle[0] != s.x {
				continue
			}
			if s.y > obstacle[1] && (s.y-distance) <= obstacle[1] {
				if min < obstacle[1]+1 {
					min = obstacle[1] + 1
				}
			}
		}
		if min != -30001 {
			s.y = min
		} else {
			s.y -= distance
		}
	case 270:
		min := -30001
		for _, obstacle := range obstacles {
			if obstacle[1] != s.y {
				continue
			}
			if s.x > obstacle[0] && (s.x-distance) <= obstacle[0] {
				if min > obstacle[0]+1 {
					min = obstacle[0] + 1
				}
			}
		}
		if min != -30001 {
			s.x = min
		} else {
			s.x -= distance
		}
	}
}
