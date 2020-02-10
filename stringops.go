package main

// StringsSlicesDifference give you diff between 2 strings slices
// Using slice1 as source and slice2 as target
func stringsSlicesDifference(slice1 []string, slice2 []string) []string {
	var diff []string
	for _, s1 := range slice1 {
		found := false
		for _, s2 := range slice2 {
			if s1 == s2 {
				found = true
				break
			}
		}
		// String not found. We add it to return slice
		if !found {
			diff = append(diff, s1)
		}
	}

	return diff
}

//unique returns slice of string of just unique values
func unique(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
