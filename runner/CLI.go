package runner

import (
	"config-forge/configs"
	"config-forge/configurator"
	"config-forge/utils"
	"fmt"
	"os"
)

func CLI() {
	sanityCheckReturn := utils.CheckArguments(os.Args)
	if sanityCheckReturn == configs.SANITY_CHECK_PASSED {
		informations := map[string]string{
			"name":         os.Args[2],
			"serverName":   os.Args[3],
			"documentRoot": os.Args[4],
		}
		configu := configurator.ConfigBuilder(os.Args[1], informations)

		er := utils.EditFile(configu)

		if er != nil {
			fmt.Println(er)
		}
		return
	}
	fmt.Println(sanityCheckReturn)
}
