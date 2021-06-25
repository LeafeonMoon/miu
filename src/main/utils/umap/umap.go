package umap

func ContainsKey(vals map[string]interface{}, key string) bool {
	for k := range vals {
		if k == key {
			return true
		}
	}
	return false
}
