package has_node_of_type

func any(values []bool) bool {
	for _, value := range values {
		if value {
			return true
		}
	}

	return false
}
