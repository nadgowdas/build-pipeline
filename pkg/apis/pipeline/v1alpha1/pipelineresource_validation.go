/*
Copyright 2018 The Knative Authors

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

package v1alpha1

import (
	"strings"

	"github.com/knative/pkg/apis"
	"k8s.io/apimachinery/pkg/api/equality"
)

func (r *PipelineResource) Validate() *apis.FieldError {
	if err := validateObjectMetadata(r.GetObjectMeta()); err != nil {
		return err.ViaField("metadata")
	}

	return r.Spec.Validate()
}

func (rs *PipelineResourceSpec) Validate() *apis.FieldError {
	if equality.Semantic.DeepEqual(rs, &PipelineResourceSpec{}) {
		return apis.ErrMissingField(apis.CurrentField)
	}
	if rs.Type == PipelineResourceTypeCluster {
		var usernameFound, cadataFound bool
		for _, param := range rs.Params {
			switch {
			case strings.EqualFold(param.Name, "URL"):
				if err := validateURL(param.Value, "URL"); err != nil {
					return err
				}
			case strings.EqualFold(param.Name, "Username"):
				usernameFound = true
			case strings.EqualFold(param.Name, "CAData"):
				cadataFound = true
			}
		}

		for _, secret := range rs.SecretParams {
			switch {
			case strings.EqualFold(secret.FieldName, "Username"):
				usernameFound = true
			case strings.EqualFold(secret.FieldName, "CAData"):
				cadataFound = true
			}
		}

		if !usernameFound {
			return apis.ErrMissingField("username param")
		}
		if !cadataFound {
			return apis.ErrMissingField("CAData param")
		}
	}
	if rs.Type == PipelineResourceTypeStorage {
		foundTypeParam := false
		var location string
		for _, param := range rs.Params {
			switch {
			case strings.EqualFold(param.Name, "type"):
				if !allowedStorageType(param.Value) {
					return apis.ErrInvalidValue(param.Value, "spec.params.type")
				}
				foundTypeParam = true
			case strings.EqualFold(param.Name, "Location"):
				location = param.Value
			}
		}

		if !foundTypeParam {
			return apis.ErrMissingField("spec.params.type")
		}
		if location == "" {
			return apis.ErrMissingField("spec.params.location")
		}
	}

	for _, allowedType := range AllResourceTypes {
		if allowedType == rs.Type {
			return nil
		}
	}

	return apis.ErrInvalidValue("spec.type", string(rs.Type))
}

func allowedStorageType(gotType string) bool {
	return string(PipelineResourceTypeGCS) == gotType
}
