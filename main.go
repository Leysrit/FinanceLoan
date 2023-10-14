package main

import (
	"Finance/handler"
	"Finance/impl/module"
	"Finance/utility"
	"log"
)

func main() {
	db := utility.ConnectDB()

	if err := utility.MigrationDB(db); err != nil {
		log.Panicln("Error when do migration:", err)
	}

	dataModule := module.NewDataModuleImpl(db)
	serviceModule := module.NewServiceModuleImpl(dataModule)

	handler.StartHandler(serviceModule)
}
