package devhub

import (
	"encoding/json"
	"fmt"
	"slices"
)

func GetDevhubServices(devhubclient *Client, apiUrl, apiPassword, userToIdentify string) ([]string, error) {
	var servicesArray []string
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
						ifInArray := slices.Contains(servicesArray, project.(map[string]interface{})["key"].(string))
						if !ifInArray {
							servicesArray = append(servicesArray, project.(map[string]interface{})["key"].(string))
						}
					}
				}
			}
		}
	}
	return servicesArray, nil
}

func GetDevhubRoles(devhubclient *Client, apiUrl, apiPassword, userToIdentify string) ([]string, error) {
	var rolesArray []string
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
						for _, profile := range team.(map[string]interface{})["profiles"].([]interface{}) {
							ifInArray := slices.Contains(rolesArray, profile.(map[string]interface{})["name"].(string))
							if !ifInArray {
								rolesArray = append(rolesArray, profile.(map[string]interface{})["name"].(string))
							}
						}
					}
				}
			}

		}
	}
	return rolesArray, nil
}
