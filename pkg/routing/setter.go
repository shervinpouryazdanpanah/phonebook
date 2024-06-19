package routing

import (
	"fmt"
	"phonebook/internal/modules/phone"
	"phonebook/internal/modules/user"
	config "phonebook/pkg/config"
)

func Set() {
	var configuration = config.Get()
	user.Routes(router)
	phone.Routes(router)
	router.Logger.Fatal(router.Start(fmt.Sprintf("%s:%s", configuration.Server.Host, configuration.Server.Port)))
}
