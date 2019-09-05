// Copyright 2019 The OpenSDS Authors.
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

// +build integration

package integration

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//Function to run the Ginkgo Test
func TestFileShareIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	//var UID string
	var _ = BeforeSuite(func() {
		fmt.Println("Before Suite Execution")

	})
	AfterSuite(func() {
		By("After Suite Execution....!")
	})

	RunSpecs(t, "File Share Integration Test Suite")
}