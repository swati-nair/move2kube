/*
Copyright IBM Corporation 2020

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

package optimize

import (
	irtypes "github.com/konveyor/move2kube/internal/types"
)

// replicaOptimizer sets the minimum number of replicas
type replicaOptimizer struct {
}

const (
	minReplicas = 2
)

func (ep replicaOptimizer) optimize(ir irtypes.IR) (irtypes.IR, error) {
	for k, scObj := range ir.Services {
		if scObj.Replicas < minReplicas {
			scObj.Replicas = minReplicas
		}
		ir.Services[k] = scObj
	}

	return ir, nil
}
