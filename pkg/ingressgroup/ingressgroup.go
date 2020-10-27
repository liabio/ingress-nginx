package ingressgroup

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	extensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/ingress-nginx/pkg/apis/ingressgroup/v1"
)

func CreateIngressGroupCRD(extensionCRClient *extensionsclient.Clientset) error {
	crd := &v1beta1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: "ingressgroups." + v1.SchemeGroupVersion.Group,
		},
		Spec: v1beta1.CustomResourceDefinitionSpec{
			Group: v1.SchemeGroupVersion.Group,
			Version: v1.SchemeGroupVersion.Version,
			Scope: v1beta1.NamespaceScoped,
			Names: v1beta1.CustomResourceDefinitionNames{
				Kind:       "IngressGroup",
				ListKind:   "IngressGroupList",
				Plural:     "ingressgroups",
				Singular:   "ingressgroup",
				ShortNames: []string{"ig"},
				//1.9.2k8s集群上不支持
				//Categories: []string{"all"},
			},
			//1.9.2k8s集群上不支持
			/*Validation: &v1beta1.CustomResourceValidation{
				OpenAPIV3Schema: &v1beta1.JSONSchemaProps{
					Properties: map[string]v1beta1.JSONSchemaProps{
						"spec": {
							Properties: map[string]v1beta1.JSONSchemaProps{
								"services": {
									Type: "array",
									Items: &v1beta1.JSONSchemaPropsOrArray{
										Schema: &v1beta1.JSONSchemaProps{
											Type:     "object",
											Required: []string{"name", "namespace"},
											Properties: map[string]v1beta1.JSONSchemaProps{
												"name": {
													Type: "string",
												},
												"namespace": {
													Type: "string",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},*/
		},
	}
	_, err := extensionCRClient.ApiextensionsV1beta1().CustomResourceDefinitions().Create(crd)
	return err
}
