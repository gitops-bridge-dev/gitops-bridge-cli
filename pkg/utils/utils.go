/*
Copyright © 2023 GitOps Bridge Project https://github.com/gitops-bridge-dev

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

package utils

import "errors"

// VerifyAddOnRepo verifies the add on repo
func VerifyAddOnRepo(s string) error {
	// Switch statement to check the add on repo
	switch s {
	case "gitops-bridge-argocd-control-plane-template":
		return nil
	case "":
		return errors.New("please provide an addon repo")
	default:
		// return an error if the add on repo is not valid
		return errors.New("invalid addon repo")
	}
}
