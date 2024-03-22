package command

func areArgumentsProvided(command *Command, arguments map[string]*string) []Argument {
	requiredArguments := (*command).Arguments()
	if requiredArguments == nil {
		return nil
	}

	var missingArguments []Argument

	for i := range requiredArguments {
		argument := requiredArguments[i]

		_, ok := arguments[argument.Key]
		if !ok {
			missingArguments = append(missingArguments, argument)
		}
	}

	return missingArguments
}
