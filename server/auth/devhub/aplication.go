package devhub

import (
	"encoding/json"
	"fmt"
)

type RolesAndServices struct {
	Roles    map[string]string
	Services map[string]string
}

func GetServicesAndRoles(devhubclient *Client, apiUrl, apiPassword, userToIdentify string) (*RolesAndServices, error) {
	roles := make(map[string]string)
	services := make(map[string]string)
	servicesAndRoles := &RolesAndServices{}
	apiDevhub := fmt.Sprintf("%s/api/identity/%s", apiUrl, userToIdentify)
	res, err := HandleRequestApiInditex(devhubclient, apiDevhub, "GET", apiPassword, map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}
	if teams, ok := result["teams"].([]interface{}); ok {
		for _, team := range teams {
			if len(team.(map[string]interface{})["projects"].([]interface{})) > 0 {
				for _, project := range team.(map[string]interface{})["projects"].([]interface{}) {
					if relationshipType, ok := project.(map[string]interface{})["relationshipType"].(map[string]interface{}); ok && relationshipType["name"] == "Owner" {
						services[project.(map[string]interface{})["key"].(string)] = "service"
						for _, profile := range team.(map[string]interface{})["profiles"].([]interface{}) {
							roles[profile.(map[string]interface{})["name"].(string)] = "role"
						}
					}
				}
			}
		}
	}
	servicesAndRoles.Roles = roles
	servicesAndRoles.Services = services
	return servicesAndRoles, nil
}
