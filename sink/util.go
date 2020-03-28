package sink

// Merge create a single map by combining all the given maps
func Merge(ms ...map[string][]float64) map[string][]float64 {
	if ms == nil || len(ms) == 0 {
		return map[string][]float64{}
	}
	if len(ms) == 1 {
		return ms[0]
	}
	res := map[string][]float64{}
	for _, m := range ms {
		for k, v := range m {
			res[k] = append(res[k], v...)
		}
	}
	return res
}
