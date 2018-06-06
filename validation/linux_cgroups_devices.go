package main

import (
	"github.com/opencontainers/runtime-tools/cgroups"
	"github.com/opencontainers/runtime-tools/validation/util"
)

func main() {
	var major1, minor1, major2, minor2, major3, minor3 int64 = 10, 229, 8, 20, 10, 200
	g, err := util.GetDefaultGenerator()
	if err != nil {
		util.Fatal(err)
	}
	g.SetLinuxCgroupsPath(cgroups.AbsCgroupPath)
	g.AddLinuxResourcesDevice(true, "c", &major1, &minor1, "rwm")
	g.AddLinuxResourcesDevice(true, "b", &major2, &minor2, "rw")
	g.AddLinuxResourcesDevice(true, "b", &major3, &minor3, "r")
	err = util.RuntimeOutsideValidate(g, util.ValidateLinuxResourcesDevices)
	if err != nil {
		util.Fatal(err)
	}
}