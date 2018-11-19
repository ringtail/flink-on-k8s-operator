package flinkapplication

import (
	"reflect"

	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/iSofiane/flink-on-k8s-operator/pkg/apis/flinkoperator.k8s.io"
	"github.com/iSofiane/flink-on-k8s-operator/pkg/apis/flinkoperator.k8s.io/v1alpha1"
)

// CRD metadata.
const (
	Plural    = "flinkapplications"
	Singular  = "flinkapplication"
	ShortName = "flinkapp"
	Group     = flinkoperator.GroupName
	Version   = v1alpha1.Version
	FullName  = Plural + "." + Group
)

func GetCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return &apiextensionsv1beta1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: FullName,
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   Group,
			Version: Version,
			Scope:   apiextensionsv1beta1.NamespaceScoped,
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Plural:     Plural,
				Singular:   Singular,
				ShortNames: []string{ShortName},
				Kind:       reflect.TypeOf(v1alpha1.FlinkApplication{}).Name(),
			},
			Validation: getCustomResourceValidation(),
		},
	}
}

func getCustomResourceValidation() *apiextensionsv1beta1.CustomResourceValidation {
	return &apiextensionsv1beta1.CustomResourceValidation{
		OpenAPIV3Schema: &apiextensionsv1beta1.JSONSchemaProps{
			Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
				"spec": {
					Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
						"taskmanager": {
							Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
								"instances": {
									Type:    "integer",
									Minimum: float64Ptr(1),
								},
							},
						},
					},
				},
			},
		},
	}
}

func float64Ptr(f float64) *float64 {
	return &f
}
