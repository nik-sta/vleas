package util

import "github.com/rhea-0b1/vleas/model"

func RemoveDuplicates(elements []model.Dependency) []model.Dependency { // change string to int here if required
	// Use map to record duplicates as we find them.
	encountered := map[model.Dependency]bool{} // change string to int here if required
	var result []model.Dependency              // change string to int here if required

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}
