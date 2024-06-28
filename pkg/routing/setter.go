package routing

import (
	"fmt"
	"phonebook/internal/modules/phone/routes"
	"phonebook/internal/modules/user"
	"phonebook/pkg/config"
)

func Set() {
	var configuration = config.Get()
	user.Routes(router)
	routes.Routes(router)
	router.Logger.Fatal(router.Start(fmt.Sprintf("%s:%s", configuration.Server.Host, configuration.Server.Port)))
}
