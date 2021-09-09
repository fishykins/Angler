package internal

func checkWord(word string) {

	canAppend := true

	if isPropperNounCandidate(word) {
		for _, pn := range potentialNouns {
			if pn == word {
				canAppend = false
				break
			}
		}

		if canAppend {
			potentialNouns = append(potentialNouns, word)
		}
	}
}
