package devhub

import (
	"encoding/json"
	"fmt"
	"slices"
)

type GroupAndServices struct {
	Services []string
	Group    string
}

func GetServicesAndGroup(devhubclient *Client, apiUrl, apiEndpoint, apiPassword, userToIdentify string, writeGroups []string) (*GroupAndServices, error) {
	var result map[string]interface{}
	var roles []string
	var services []string
	servicesAndGroup := &GroupAndServices{}
	apiDevhub := fmt.Sprintf("%s/%s/%s", apiUrl, apiEndpoint, userToIdentify)
	res, err := HandleRequestApiInditex(devhubclient, apiDevhub, "GET", apiPassword, map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	roles, services = GetRolesAndServices(result, services, roles)

	servicesAndGroup.Group = GetGroupByRole(writeGroups, roles)
	servicesAndGroup.Services = services
	return servicesAndGroup, nil
}

func GetRolesAndServices(result map[string]interface{}, services, roles []string) ([]string, []string) {
	teams, ok := result["teams"].([]interface{})
	if !ok {
		return services, roles
	}
	for _, team := range teams {
		if len(team.(map[string]interface{})["applications"].([]interface{})) <= 0 {
			continue
		}
		for _, project := range team.(map[string]interface{})["applications"].([]interface{}) {
			if project.(map[string]interface{})["relationshipType"].(string) != "Owner" {
				continue
			}
			if !slices.Contains(services, project.(map[string]interface{})["key"].(string)) {
				services = append(services, project.(map[string]interface{})["key"].(string))
			}
			for _, profile := range team.(map[string]interface{})["profiles"].([]interface{}) {
				if !slices.Contains(roles, profile.(map[string]interface{})["name"].(string)) {
					roles = append(roles, profile.(map[string]interface{})["name"].(string))
				}
			}
			crossprofiles, ok := result["crossProfiles"].([]interface{})
			if !ok {
				continue
			}
			for _, crossprofile := range crossprofiles {
				if !slices.Contains(roles, crossprofile.(map[string]interface{})["name"].(string)) {
					roles = append(roles, crossprofile.(map[string]interface{})["name"].(string))
				}
			}
		}
	}
	return roles, services
}

func GetGroupByRole(writeGroups []string, roles []string) string {
	groupByRole := "reader"
	for _, role := range roles {
		if slices.Contains(writeGroups, role) {
			groupByRole = "writer"
		}
	}
	return groupByRole
}
