// Copyright 2020 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0

package deploy

import (
	"fmt"

	"github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/kappctrl/v1alpha1"
	"k8s.io/client-go/kubernetes"
)

type Factory struct {
	coreClient kubernetes.Interface

	kubeconfigSecrets *KubeconfigSecrets
	serviceAccounts   *ServiceAccounts
}

func NewFactory(coreClient kubernetes.Interface) Factory {
	return Factory{coreClient, NewKubeconfigSecrets(coreClient), NewServiceAccounts(coreClient)}
}

func (f Factory) NewKapp(opts v1alpha1.AppDeployKapp, saName string,
	clusterOpts *v1alpha1.AppCluster, genericOpts GenericOpts, cancelCh chan struct{}) (*Kapp, error) {

	var err error
	var processedGenericOpts ProcessedGenericOpts

	switch {
	case len(saName) > 0:
		processedGenericOpts, err = f.serviceAccounts.Find(genericOpts, saName)
		if err != nil {
			return nil, err
		}

	case clusterOpts != nil:
		processedGenericOpts, err = f.kubeconfigSecrets.Find(genericOpts, clusterOpts)
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Expected service account or cluster specified")
	}

	return NewKapp(opts, processedGenericOpts, cancelCh), nil
}

// NewKappPrivileged is used for package repositories where users aren't required to provide
// a service account, so it will install resources using its own privileges.
func (f Factory) NewKappPrivileged(opts v1alpha1.AppDeployKapp,
	genericOpts GenericOpts, cancelCh chan struct{}) (*Kapp, error) {

	pgo := ProcessedGenericOpts{
		Name:      genericOpts.Name,
		Namespace: genericOpts.Namespace,
		// Just use the default service account. Mainly
		// used for PackageRepos now so users do not need
		// to specify serviceaccount via PackageRepo CR.
		Kubeconfig:                    nil,
		DangerousUsePodServiceAccount: true,
	}

	return NewKapp(opts, pgo, cancelCh), nil
}
