
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



package stack

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"

	"github.com/aerokite/kube-ext-api-server/pkg/apis/aerokite/v1alpha1"
	"github.com/aerokite/kube-ext-api-server/pkg/controller/sharedinformers"
	listers "github.com/aerokite/kube-ext-api-server/pkg/client/listers_generated/aerokite/v1alpha1"
)

// +controller:group=aerokite,version=v1alpha1,kind=Stack,resource=stacks
type StackControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about Stack
	lister listers.StackLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *StackControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing stacks labels
	c.lister = arguments.GetSharedInformers().Factory.Aerokite().V1alpha1().Stacks().Lister()
}

// Reconcile handles enqueued messages
func (c *StackControllerImpl) Reconcile(u *v1alpha1.Stack) error {
	// Implement controller logic here
	log.Printf("Running reconcile Stack for %s\n", u.Name)
	return nil
}

func (c *StackControllerImpl) Get(namespace, name string) (*v1alpha1.Stack, error) {
	return c.lister.Stacks(namespace).Get(name)
}
