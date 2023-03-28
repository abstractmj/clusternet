/*
Copyright 2023 The Clusternet Authors.

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

package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/clusternet/clusternet/cmd/clusternet-controller-manager/app"
	"github.com/clusternet/clusternet/pkg/utils"
	"k8s.io/component-base/cli"
	_ "k8s.io/component-base/logs/json/register" // for JSON log format registration
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	command := app.NewControllerManagerCommand(utils.GracefulStopWithContext())
	code := cli.Run(command)
	os.Exit(code)
}