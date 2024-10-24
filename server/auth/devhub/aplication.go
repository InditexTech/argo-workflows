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
type ApiStruct struct {
	Login      string `json:"login"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Mail       string `json:"mail"`
	Teams      []struct {
		Key                string `json:"key"`
		Name               string `json:"name"`
		CompleteDepartment struct {
			ID               string `json:"id"`
			Name             string `json:"name"`
			FullDepartmentID string `json:"fullDepartmentId"`
			FullDepartment   string `json:"fullDepartment"`
		} `json:"completeDepartment"`
		Applications []struct {
			Key              string `json:"key"`
			Name             string `json:"name"`
			RelationshipType string `json:"relationshipType"`
		} `json:"applications"`
		Profiles []struct {
			Key  int    `json:"key"`
			Name string `json:"name"`
		} `json:"profiles"`
		EffectiveCrossProfiles []string `json:"effectiveCrossProfiles"`
		IsMember               bool     `json:"isMember"`
	} `json:"teams"`
	CrossProfiles []struct {
		Key        string `json:"key"`
		Department string `json:"department"`
		Name       string `json:"name"`
	} `json:"crossProfiles"`
}

func GetServicesAndGroup(devhubclient *Client, apiUrl, apiEndpoint, apiPassword, userToIdentify string, writeGroups []string) (*GroupAndServices, error) {
	var roles []string
	var services []string
	servicesAndGroup := &GroupAndServices{}
	apiResponse := &ApiStruct{}
	apiDevhub := fmt.Sprintf("%s/%s/%s", apiUrl, apiEndpoint, userToIdentify)
	res, err := HandleRequestApiInditex(devhubclient, apiDevhub, "GET", apiPassword, map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	if err := json.NewDecoder(res.Body).Decode(apiResponse); err != nil {
		return nil, err
	}

	roles, services = GetRolesAndServices(apiResponse, services, roles)

	servicesAndGroup.Group = GetGroupByRole(writeGroups, roles)
	servicesAndGroup.Services = services
	return servicesAndGroup, nil
}

func GetRolesAndServices(result *ApiStruct, services, roles []string) ([]string, []string) {
	if result.Teams == nil {
		return services, roles
	}
	for _, team := range result.Teams {

		if len(team.Applications) <= 0 {
			continue
		}
		for _, project := range team.Applications {
			if project.RelationshipType != "Owner" {
				continue
			}
			if !slices.Contains(services, project.Key) {
				services = append(services, project.Key)
			}
			for _, profile := range team.Profiles {
				if !slices.Contains(roles, profile.Name) {
					roles = append(roles, profile.Name)
				}
			}
			crossprofiles := result.CrossProfiles
			if crossprofiles == nil {
				continue
			}
			for _, crossprofile := range crossprofiles {
				if !slices.Contains(roles, crossprofile.Name) {
					roles = append(roles, crossprofile.Name)
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
