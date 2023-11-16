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

package gobgh

// import the go-github library
import (
	"context"

	"github.com/google/go-github/v56/github"
)

// ForkRepo forks the repo
func ForkRepo(token, org, repo string) (string, error) {
	ctx := context.Background()
	client := github.NewClient(nil).WithAuthToken(token)

	// Fork the repo
	// TODO: Add support for forking into a specific organization/user and if it should be private or not
	fork, res, err := client.Repositories.CreateFork(ctx, org, repo, nil)

	// Status code 202 is technically a success
	if err != nil && res.StatusCode != 202 {
		return "", err
	}

	// Check if the fork was successfully created
	_, _, err = client.Repositories.Get(ctx, *fork.GetOwner().Login, repo)
	if err != nil {
		return "", err
	}

	return fork.GetHTMLURL(), nil
}
