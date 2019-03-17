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
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
)

// killCmd represents the kill command
var killCmd = &cobra.Command{
	Use:   "kill",
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
			mappedName, err := hashing.MapStringToInt(string(pod.UID))
			if err != nil {
				return
			}

			if strconv.Itoa(int(mappedName)) == args[0] {
				err := client.CoreV1().Pods(pod.Namespace).Delete(pod.Name, &v1.DeleteOptions{})
				if err != nil {
					fmt.Println(err)
				}
				return
			}
		}
	},
}

func init() {
	k8sCmd.AddCommand(killCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// killCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// killCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
