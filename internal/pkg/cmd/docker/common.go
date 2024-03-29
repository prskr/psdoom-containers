// Copyright © 2019 Peter Kurfer peter.kurfer@googlemail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	_ "github.com/docker/docker/reference"
)

func getDockerClient() (*client.Client, error) {
	return client.NewClientWithOpts()
}

func getContainers(client *client.Client) ([]types.Container, error) {
	ctx := context.Background()
	return client.ContainerList(ctx, types.ContainerListOptions{
		All: true,
	})
}



func firstOrEmpty(sa []string) string {
	if len(sa) < 1 {
		return ""
	}
	return sa[0]
}
