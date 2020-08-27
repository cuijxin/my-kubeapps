/*
Copyright 2017 Bitnami.
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
	clientset "github.com/cuijxin/my-kubeapps/cmd/apprepository-controller/pkg/client/clientset/versioned"
	listers "github.com/cuijxin/my-kubeapps/cmd/apprepository-controller/pkg/client/listers/apprepository/v1alpha1"
	"k8s.io/client-go/informers"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	batchlisters "k8s.io/client-go/listers/batch/v1beta1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
)

const controllerAgentName = "apprepository-controller"

const (
	// SuccessSynced is used as part of the Event 'reason' when an AppRepository
	// is synced
	SuccessSynced = "Synced"
	// ErrResourceExists is used as part of the Event 'reason' when an
	// AppRepository fails to sync due to a CronJob of the same name already
	// existing.
	ErrResourceExists = "ErrResourceExists"

	// LabelRepoName is the label used to identify the repository name.
	LabelRepoName = "apprepositories.kubeapps.com/repo-name"
	// LabelRepoNamespace is the label used to identify the repository namespace.
	LabelRepoNamespace = "apprepositories.kubeapps.com/repo-namespace"

	// MessageResourceExists is the message used for Events when a resource
	// fails to sync due to a CronJob already existing
	MessageResourceExists = "Resource %q already exists and is not managed by AppRepository"
	// MessageResourceSynced is the message used for an Event fired when an
	// AppRepsitory is synced successfully
	MessageResourceSynced = "AppRepository synced successfully"
)

// Controller is the controller implementation for AppRepository resources
type Controller struct {
	// kubeclientset is a standard kubernetes clientset
	kubeclientset kubernetes.Interface
	// apprepoclientset is the clientset for AppRepository resources
	apprepoclientset clientset.Interface

	cronjobsLister batchlisters.CronJobLister
	cronjobsSynced cache.InformerSynced
	appreposLister listers.AppRepositoryLister
	appreposSynced cache.InformerSynced

	// workqueue is a rate limited work queue. This is used to queue work to be
	// processed instead of performing it as soon as a change happens. This
	// means we can ensure we only process a fixed amount of resources at a
	// time, and makes it easy to ensure we are never processing the same item
	// simultaneously in two different workers.
	workqueue workqueue.RateLimitingInterface
	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	recorder record.EventRecorder

	kubeappsNamespace string
}

// NewController returns a new sample controller
func NewController(
	kubeclientset kubernetes.Interface,
	apprepoclientset clientset.Interface,
	kubeInformerFactory kubeinformers.SharedInformerFactory,
	apprepoInformerFactory informers.SharedInformerFactory,
	kubeappsNamespace string) *Controller {

	// obtain references to shared index informers for the CronJob and
	// AppRepository types.
	cronjobInformer := kubeInformerFactory.Batch().V1beta1().CronJobs()

}
