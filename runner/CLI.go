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
			"name":         os.Args[3],
			"serverName":   os.Args[4],
			"documentRoot": os.Args[5],
		}
		configu := configurator.ConfigBuilder(os.Args[2], informations)

		er := utils.EditFile(configu)

		if er != nil {
			fmt.Println(er)
		}
		return
	}
	fmt.Println(sanityCheckReturn)
}
