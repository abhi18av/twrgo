package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServiceInfo(t *testing.T) {

	serviceInfo := ServiceInfo(TOWER_API_ENDPOINT)

	assert.Equal(t, "/login", serviceInfo.ServiceInfo.LoginPath, "check login path")

	assert.Equal(t, "Docs", serviceInfo.ServiceInfo.Navbar.Menus[0].Label, "check Docs label")

}
