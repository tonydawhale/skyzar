package utils

func ObjectKeys(obj interface{}) []string {
	keys := make([]string, len(obj.(map[string]string)))
	i := 0
	for k := range obj.(map[string]string) {
		keys[i] = k
		i++
	}
	return keys
}