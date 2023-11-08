package factory

import (
	"errors"

	"github.com/cloudfstrife/licensed/pkg/plugin"
	"github.com/cloudfstrife/licensed/pkg/plugin/baseboard"
	"github.com/cloudfstrife/licensed/pkg/plugin/mac"
	"github.com/cloudfstrife/licensed/pkg/plugin/machineid"
	"github.com/cloudfstrife/licensed/pkg/plugin/plaintext"
)

// CreateClientIDFuncFactory CreateClientIDFunc factory
func CreateClientIDFuncFactory(name string) (plugin.CreateClientIDFunc, error) {
	switch name {
	case "machine-id":
		return machineid.CreateClientID, nil
	case "mac":
		return mac.CreateClientID, nil
	case "text":
		return plaintext.CreateClientID, nil
	case "baseboard":
		return baseboard.CreateClientID, nil
	default:
		return nil, errors.New("bad type")
	}
}
