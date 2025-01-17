package controller

import (
	"github.com/sohankunkerkar/onprem-operator/pkg/controller/hubcluster"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, hubcluster.Add)
}
