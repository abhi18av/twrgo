package api

import (
	"encoding/json"
	"github.com/imroc/req"
)

type ServiceInfoType struct {
	ServiceInfo struct {
		Version   string   `json:"version"`
		CommitID  string   `json:"commitId"`
		AuthTypes []string `json:"authTypes"`
		LoginPath string   `json:"loginPath"`
		Navbar    struct {
			Menus []struct {
				Label string `json:"label"`
				URL   string `json:"url"`
			} `json:"menus"`
		} `json:"navbar"`
		HeartbeatInterval int `json:"heartbeatInterval"`
	} `json:"serviceInfo"`
}

func ServiceInfo(towerApiEndpoint string) ServiceInfoType {
	ENDPOINT := towerApiEndpoint + "service-info/"
	res, _ := req.Get(ENDPOINT)

	var serviceInfo ServiceInfoType
	json.Unmarshal(res.Bytes(), &serviceInfo)

	return serviceInfo
}
