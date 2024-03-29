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

package daemon

import (
	"context"
	"github.com/baez90/psdoom-containers/internal/pkg/api/k8s/generated"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"strconv"
)

var killDaemonCmd = &cobra.Command{
	Use:   "kill",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		i, err := strconv.Atoi(args[0])

		con, err := grpc.Dial("127.0.0.1:1357", grpc.WithInsecure())
		if err != nil {
			logrus.Error("failed connect to k8s daemon", err)
			return
		}
		client := k8sApi.NewK8SDaemonClient(con)
		_, err = client.KillPod(context.Background(), &k8sApi.PodDeletion{
			PodId: int32(i),
		})

		if err != nil {
			logrus.Error("Error while sending kill command", err)
		}
	},
}

func init() {
	k8sDaemonCmd.AddCommand(killDaemonCmd)
}
