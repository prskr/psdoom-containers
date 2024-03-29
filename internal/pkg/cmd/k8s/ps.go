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

package k8s

import (
	"fmt"
	"github.com/baez90/psdoom-containers/internal/pkg/api/k8s"
	"github.com/baez90/psdoom-containers/internal/pkg/hashing"
	"github.com/spf13/cobra"
	v1meta "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"path/filepath"

	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

// psCmd represents the ps command
var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := k8s.GetKubeClient()

		if err != nil {
			return
		}

		pods, err := client.CoreV1().Pods("").List(v1.ListOptions{})
		if err != nil {
			return
		}

		for _, pod := range pods.Items {
			podNameHash, err := hashing.MapStringToInt(string(pod.UID))
			if err != nil || pod.Status.Phase != v1meta.PodRunning {
				continue
			}
			// format <user> <pid> <processname> <is_daemon=[1|0]>
			fmt.Printf("%s %d %s 1\n", pod.Namespace, podNameHash, pod.Name)
		}
	},
}

func init() {
	k8sCmd.AddCommand(psCmd)
	if home := k8s.HomeDir(); home != "" {
		k8s.KubeConfigPathFlag = psCmd.Flags().String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		k8s.KubeConfigPathFlag = psCmd.Flags().String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
}
