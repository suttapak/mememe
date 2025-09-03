package utils

import "os"

func IsHaveControllerModuleFile() bool {
	// check if controller module file exists
	// if exists return true
	// else
	// return false
	workingDir, err := os.Getwd()
	if err != nil {
		return false
	}
	_, err = os.Stat(workingDir + "/internal/controller/module.go")
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
	workingDir, err := os.Getwd()
	if err != nil {
		return false
	}
	_, err = os.Stat(workingDir + "/internal/service/module.go")
	if err == nil {
		return true
	}

	return false
}

func IsHaveRepositoryModuleFile() bool {
	// check if controller module file exists
	// if exists return true
	// else
	// return false
	workingDir, err := os.Getwd()
	if err != nil {
		return false
	}
	_, err = os.Stat(workingDir + "/internal/repository/module.go")
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
	workingDir, err := os.Getwd()
	if err != nil {
		return false
	}
	_, err = os.Stat(workingDir + "/internal/route/module.go")
	if err == nil {
		return true
	}

	return false
}
