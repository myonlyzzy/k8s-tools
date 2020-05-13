package main

import (
	"github.com/oam-dev/oam-go-sdk/apis/core.oam.dev/v1alpha1"
	"github.com/oam-dev/oam-go-sdk/pkg/client/clientset/versioned"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

const (
	DefaultNameSpace = "test"
)

func main() {
	oamCli := CreateOAMClient()
	appconfiSpec := v1alpha1.ApplicationConfigurationSpec{
		Variables:  nil,
		Scopes:     nil,
		Components: nil,
	}
	appConf := &v1alpha1.ApplicationConfiguration{
		TypeMeta:   nil,
		ObjectMeta: nil,
		Spec:       appconfiSpec,
		Status:     v1alpha1.ApplicationConfigurationStatus{},
	}
	appScopeSpec := v1alpha1.ApplicationScopeSpec{
		Type:                  "",
		AllowComponentOverlap: false,
		Parameters:            nil,
	}
	appScope := &v1alpha1.ApplicationScope{
		TypeMeta:   nil,
		ObjectMeta: nil,
		Spec:       appScopeSpec,
		Status:     v1alpha1.ApplicationScopeStatus{},
	}
	comp := &v1alpha1.ComponentSchematic{
		TypeMeta:   nil,
		ObjectMeta: nil,
		Spec: v1alpha1.ComponentSpec{
			Parameters:       nil,
			WorkloadType:     "",
			OsType:           "",
			Arch:             "",
			Containers:       nil,
			WorkloadSettings: nil,
		},
	}
	trait := &v1alpha1.Trait{
		TypeMeta:   nil,
		ObjectMeta: nil,
		Spec: v1alpha1.TraitSpec{
			Group:      "",
			Version:    "",
			Names:      v1alpha1.Names{},
			AppliesTo:  nil,
			Properties: "",
		},
		Status: v1alpha1.TraitStatus{},
	}
	oamCli.CoreV1alpha1().ApplicationConfigurations(DefaultNameSpace).Create(appConf)
	oamCli.CoreV1alpha1().ApplicationScopes(DefaultNameSpace).Create(appScope)
	oamCli.CoreV1alpha1().ComponentSchematics(DefaultNameSpace).Create(comp)
	oamCli.CoreV1alpha1().Traits(DefaultNameSpace).Create(trait)
}

