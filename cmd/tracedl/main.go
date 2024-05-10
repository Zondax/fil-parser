package main

import (
	"fmt"
	"strings"

	"github.com/zondax/fil-parser/types"
	"github.com/zondax/golem/pkg/cli"
	"golang.org/x/mod/semver"
)

func main() {
	appSettings := cli.AppSettings{}

	cli := cli.New[Config](appSettings)
	defer cli.Close()

	cli.GetRoot().AddCommand(GetStartCommand(cli))

	cli.Run()
}

func processNodeVersion(fullVersion string) (*types.NodeInfo, error) {
	var majorMinor string
	splitVersion := strings.Split(fullVersion, "+")
	if len(splitVersion) < 2 {
		return nil, fmt.Errorf("could not get node version, invalid version format detected: %s", fullVersion)
	} else {
		majorMinor = fmt.Sprintf("v%s", splitVersion[0])
		majorMinor = semver.MajorMinor(majorMinor)
		if majorMinor == "" {
			return nil, fmt.Errorf("could not get node version, invalid version format detected: %s", fullVersion)
		}
	}

	return &types.NodeInfo{
		NodeFullVersion:       fullVersion,
		NodeMajorMinorVersion: majorMinor,
	}, nil
}
