package utils

import "os"

func IsHaveControllerModuleFile() bool {
	// check if controller module file exists
	// if exists return true
	// else
	// return false
	workginDir, err := os.Getwd()
	if err != nil {
		return false
	}
	_, err = os.Stat(workginDir + "/internal/controller/module.go")
	if err == nil {
		return true
	}

	return false
}

func IsHaveServiceModuleFile() bool {
	// check if controller module file exists
	// if exists return true
	// else
	// return false
	workginDir, err := os.Getwd()
	if err != nil {
		return false
	}
	_, err = os.Stat(workginDir + "/internal/service/module.go")
	if err == nil {
		return true
	}

	return false
}

func IsHaveRepositoryeModuleFile() bool {
	// check if controller module file exists
	// if exists return true
	// else
	// return false
	workginDir, err := os.Getwd()
	if err != nil {
		return false
	}
	_, err = os.Stat(workginDir + "/internal/repository/module.go")
	if err == nil {
		return true
	}

	return false
}

func IsHaveRouteModuleFile() bool {
	// check if controller module file exists
	// if exists return true
	// else
	// return false
	workginDir, err := os.Getwd()
	if err != nil {
		return false
	}
	_, err = os.Stat(workginDir + "/internal/route/module.go")
	if err == nil {
		return true
	}

	return false
}
