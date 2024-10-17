package devhub

import (
	"encoding/json"
	"fmt"
	"slices"
)

type GroupAndServices struct {
	Services map[string]string
	Group    string
}

func GetServicesAndGroup(devhubclient *Client, apiUrl, apiEndpoint, apiPassword, userToIdentify string, writeGroups []string) (*GroupAndServices, error) {
	var result map[string]interface{}
	roles := make(map[string]string)
	services := make(map[string]string)
	servicesAndGroup := &GroupAndServices{}
	apiDevhub := fmt.Sprintf("%s%s%s", apiUrl, apiEndpoint, userToIdentify)
	res, err := HandleRequestApiInditex(devhubclient, apiDevhub, "GET", apiPassword, map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}
	if teams, ok := result["teams"].([]interface{}); ok {
		roles, services = GetRolesAndServices(teams, services, roles)
	}

	servicesAndGroup.Group = GetGroupByRole(writeGroups, roles)
	servicesAndGroup.Services = services
	return servicesAndGroup, nil
}

func GetRolesAndServices(teams []interface{}, services, roles map[string]string) (map[string]string, map[string]string) {
	for _, team := range teams {
		if len(team.(map[string]interface{})["projects"].([]interface{})) > 0 {
			for _, project := range team.(map[string]interface{})["projects"].([]interface{}) {
				if relationshipType, ok := project.(map[string]interface{})["relationshipType"].(map[string]interface{}); ok && relationshipType["name"] == "Owner" {
					services[project.(map[string]interface{})["key"].(string)] = "service"
					for _, profile := range team.(map[string]interface{})["profiles"].([]interface{}) {
						roles[profile.(map[string]interface{})["name"].(string)] = "role"
					}
					if len(team.(map[string]interface{})["effectiveCrossProfiles"].([]string)) > 0 {
						for _, effectiveCrossProfile := range team.(map[string]interface{})["effectiveCrossProfiles"].([]string) {
							roles[effectiveCrossProfile] = "role"
						}
					}
				}
			}
		}
	}
	return roles, services
}

func GetGroupByRole(writeGroups []string, roles map[string]string) string {
	groupByRole := "reader"
	for role := range roles {
		if slices.Contains(writeGroups, role) {
			groupByRole = "writer"
		}
	}
	return groupByRole
}
