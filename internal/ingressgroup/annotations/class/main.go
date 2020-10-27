/*
Copyright 2015 The Kubernetes Authors.

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

package class

import (
	customextensionsv1 "k8s.io/ingress-nginx/pkg/apis/ingressgroup/v1"
	ingclass "k8s.io/ingress-nginx/internal/ingress/annotations/class"
	"k8s.io/klog"
)

const (
	// IngressGroupKey picks a specific "class" for the Ingress.
	// The controller only processes Ingresses with this annotation either
	// unset, or set to either the configured value or the empty string.
	IngressGroupKey = "kubernetes.io/ingress.class"
)
// IsValid returns true if the given IngressGroup either doesn't specify
// the ingress.class annotation
func IsValid(ingGroup *customextensionsv1.IngressGroup) bool {
	ingGroupAnno, ok := ingGroup.GetAnnotations()[IngressGroupKey]
	if !ok {
		klog.V(3).Infof("annotation %v is not present in IngressGroup %v/%v", IngressGroupKey, ingGroup.Namespace, ingGroup.Name)
	}

	// we have 2 valid combinations
	// 1 - ingress with default class | blank annotation on ingress
	// 2 - ingress with specific class | same annotation on ingress
	//
	// and 2 invalid combinations
	// 3 - ingress with default class | fixed annotation on ingress
	// 4 - ingress with specific class | different annotation on ingress
	if ingGroupAnno == "" && ingclass.IngressClass == ingclass.DefaultClass {
		return true
	}

	return ingGroupAnno == ingclass.IngressClass
}
