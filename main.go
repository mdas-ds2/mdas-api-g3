package main

import (
	command "github.com/mdas-ds2/mdas-api-g3/pokemon-types/infrastructure/command"
)

func main() {
	getTypesByNameCommand := command.GetTypesByNameCommand{}
	getTypesByNameCommand.Run()
}
