package tasks

func splitCommandString(command string) []string {

	if len(command) == 0 {
		return []string{}
	}

	components := make([]string, 0)
	startIndex := 0
	currentIndex := 0

	parenStack := make([]byte, 0)

	for currentIndex < len(command) {
		// fmt.Println(parenStack, currentIndex, startIndex)
		if command[currentIndex] == ' ' && len(parenStack) == 0 {
			component := command[startIndex:currentIndex]
			startIndex = currentIndex + 1

			if component != "" {
				components = append(components, cleanComponent(component))
			}
		} else if command[currentIndex] == '\'' {
			if len(parenStack) == 0 {
				parenStack = append(parenStack, '\'')
			} else {
				lastElement := parenStack[len(parenStack)-1]

				if lastElement == command[currentIndex] {
					parenStack = parenStack[:len(parenStack)-1]
				} else {
					parenStack = append(parenStack, command[currentIndex])
				}
			}
		} else if command[currentIndex] == '"' {
			if len(parenStack) == 0 {
				parenStack = append(parenStack, '"')
			} else {
				lastElement := parenStack[len(parenStack)-1]

				if lastElement == command[currentIndex] {

					parenStack = parenStack[:len(parenStack)-1]

				} else {
					parenStack = append(parenStack, command[currentIndex])
				}
			}
		}
		currentIndex += 1
	}

	lastComponent := command[startIndex:currentIndex]

	components = append(components, lastComponent)

	return components

}
