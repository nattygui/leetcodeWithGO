package minoperations

func minOperations(logs []string) int {
	floders := []string{}
	for _, log := range logs {
		switch log {
		case "./":
			continue
		case "../":
			floders = floders[:len(floders)-1]
		default:
			floders = append(floders, log)
		}
	}
	return len(floders)
}
