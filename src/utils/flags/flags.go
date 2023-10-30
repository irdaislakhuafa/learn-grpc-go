package flags

import (
	"errors"
	"flag"
	"fmt"

	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/files"
)

// default env is "local"
func ParseFlags(configPath, configFile string) (*string, error) {
	env := ""
	flag.StringVar(&env, "env", "local", fmt.Sprintf("mode from dir name in %v/", configPath))
	flag.Parse()

	if dir := fmt.Sprintf("%v/%v", configPath, env); !files.IsFileExist(dir) {
		return nil, errors.New(fmt.Sprintf("env '%v' not found in dir '%v'", env, configPath))
	} else {
		if cfgFile := fmt.Sprintf("%v/%v", dir, configFile); !files.IsFileExist(cfgFile) {
			return nil, errors.New(fmt.Sprintf("cfg file '%v' not found in '%v'", configFile, dir))
		}
	}
	return &env, nil
}
