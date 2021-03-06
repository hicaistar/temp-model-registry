/*
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

// +nirvana:api=descriptors:"Descriptor"

package apis

import (
	"github.com/caicloud/nirvana/definition"

	v1alpha1 "github.com/kleveross/klever-model-registry/pkg/registry/apis/v1alpha1/descriptors"
)

var (
	AllDescriptor []definition.Descriptor
)

func init() {
	AllDescriptor = append(
		AllDescriptor,
		Descriptor(),
		v1alpha1.HarborAPIDescriptor(),
		v1alpha1.HarborCDescriptor(),
		v1alpha1.HarborServiceDescriptor(),
		v1alpha1.HarborV2Descriptor(),
	)
}

// Descriptor contain klever model registry descriptors
func Descriptor() definition.Descriptor {
	return definition.Descriptor{
		Path:        "/api",
		Description: "It contains all api in v1alpha1",
		Consumes:    []string{definition.MIMEAll},
		Produces:    []string{definition.MIMEJSON},
		Children: []definition.Descriptor{
			v1alpha1.Descriptor(),
		},
	}
}
