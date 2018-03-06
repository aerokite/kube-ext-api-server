
/*
Copyright 2017 The Kubernetes Authors.

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



package stack_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/client-go/rest"
	"github.com/kubernetes-incubator/apiserver-builder/pkg/test"

	"github.com/aerokite/kube-ext-api-server/pkg/apis"
	"github.com/aerokite/kube-ext-api-server/pkg/client/clientset_generated/clientset"
	"github.com/aerokite/kube-ext-api-server/pkg/openapi"
	"github.com/aerokite/kube-ext-api-server/pkg/controller/sharedinformers"
	"github.com/aerokite/kube-ext-api-server/pkg/controller/stack"
)

var testenv *test.TestEnvironment
var config *rest.Config
var cs *clientset.Clientset
var shutdown chan struct{}
var controller *stack.StackController
var si *sharedinformers.SharedInformers

func TestStack(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "Stack Suite", []Reporter{test.NewlineReporter{}})
}

var _ = BeforeSuite(func() {
	testenv = test.NewTestEnvironment()
	config = testenv.Start(apis.GetAllApiBuilders(), openapi.GetOpenAPIDefinitions)
	cs = clientset.NewForConfigOrDie(config)

	shutdown = make(chan struct{})
	si = sharedinformers.NewSharedInformers(config, shutdown)
	controller = stack.NewStackController(config, si)
	controller.Run(shutdown)
})

var _ = AfterSuite(func() {
	close(shutdown)
	testenv.Stop()
})
