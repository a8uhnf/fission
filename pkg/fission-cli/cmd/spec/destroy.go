/*
Copyright 2019 The Fission Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package spec

import (
	"github.com/fission/fission/pkg/controller/client"
	"github.com/fission/fission/pkg/fission-cli/cliwrapper/cli"
	"github.com/fission/fission/pkg/fission-cli/cmd"
	"github.com/fission/fission/pkg/fission-cli/util"
)

type DestroySubCommand struct {
	client *client.Client
}

// Destroy destroys everything in the spec.
func Destroy(flags cli.Input) error {
	opts := &DestroySubCommand{
		client: cmd.GetServer(flags),
	}
	return opts.do(flags)
}

func (opts *DestroySubCommand) do(flags cli.Input) error {
	return opts.run(flags)
}

func (opts *DestroySubCommand) run(flags cli.Input) error {
	// get specdir
	specDir := cmd.GetSpecDir(flags)

	// read everything
	fr, err := ReadSpecs(specDir)
	util.CheckErr(err, "read specs")

	// set desired state to nothing, but keep the UID so "apply" can find it
	emptyFr := FissionResources{}
	emptyFr.DeploymentConfig = fr.DeploymentConfig

	// "apply" the empty state
	_, _, err = applyResources(opts.client, specDir, &emptyFr, true)
	util.CheckErr(err, "delete resources")

	return nil
}
