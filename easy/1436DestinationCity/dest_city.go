package destcity

func destCity(paths [][]string) string {
	if len(paths) == 1 {
		return paths[0][1]
	}

	pathMap := make(map[string]string)
	for _, path := range paths {
		pathMap[path[0]] = path[1]
	}

	path := paths[0][0]
	for {
		_, ok := pathMap[path]
		if ok {
			path = pathMap[path]
		} else {
			return path
		}
	}
}
