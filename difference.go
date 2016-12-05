package set

// Create a set as the difference of the sets `a` and `b`. See
// https://en.wikipedia.org/wiki/Complement_%28set_theory%29
//
// If the set a is countable the resulting set is also countable and if a is
// not countable the resulting set is not also.
func Difference(a Set, b Set) (Set, error) {

	// a is countable
	if a.Countable() {
		newSet := elementSet{make(map[interface{}]bool)}

		elements, err := a.List()
		if err != nil {
			return newSet, err
		}

		for element := range elements {
			if inB, err := b.Contains(element); err != nil {
				return newSet, err
			} else if !inB { // not in b, so in difference
				newSet.elements[element] = true
			}
		}
		return newSet, nil
	}

	// If a is not countable
	newContains := func(x interface{}) (bool, error) {
		if inA, err := a.Contains(x); err != nil {
			return false, err
		} else if !inA {
			return false, nil
		}

		// If here, x is in a
		if inB, err := b.Contains(x); err != nil {
			return false, err
		} else if inB {
			return false, nil
		}

		// if here, x not in b
		return true, nil

	}
	return functionSet{newContains}, nil

}
