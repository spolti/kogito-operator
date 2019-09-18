// Copyright 2019 Red Hat, Inc. and/or its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import "fmt"

// checkProjectLocally will verify if the project/namespace exists in the CLI context.
// Won't fecth the cluster to verify if the project/namespace exists. This is a local validation only.
func checkProjectLocally(project string) (localProject string, err error) {
	if len(project) == 0 {
		if len(config.Namespace) == 0 {
			return "", fmt.Errorf("Couldn't find any Project in the current context. Use 'kogito use-project NAME' to set the Kogito Project where the service will be deployed or pass '--project NAME' flag to this one")
		}
		return config.Namespace, nil
	}
	return project, nil
}