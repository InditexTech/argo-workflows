package devhub

import (
	"encoding/json"
	"fmt"
	"slices"
)

type GroupAndServices struct {
	ServiceToGroup []string
}
type WriteGroupsParams struct {
	Relationship string   `yaml:"relationship"`
	Roles        []string `yaml:"roles"`
}

type WriteGroupsList []WriteGroupsParams
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

func GetServicesAndGroup(devhubclient *Client, apiUrl, apiEndpoint, apiPassword, userToIdentify string, writeGroups WriteGroupsList) (*GroupAndServices, error) {
	servicesToRoles := make(map[string][]string)
	servicesToRelationship := make(map[string]string)
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

	servicesToRelationship, servicesToRoles = GetRolesAndServices(apiResponse, servicesToRelationship, servicesToRoles)
	servicesAndGroup.ServiceToGroup = GetServiceToGroup(writeGroups, servicesToRelationship, servicesToRoles)
	return servicesAndGroup, nil
}

func GetRolesAndServices(result *ApiStruct, servicesToRelationship map[string]string, servicesToRoles map[string][]string) (map[string]string, map[string][]string) {
	var roles []string
	if result.Teams == nil {
		return servicesToRelationship, servicesToRoles
	}
	for _, team := range result.Teams {

		if len(team.Applications) <= 0 {
			continue
		}
		for _, project := range team.Applications {
			if existingRole, exists := servicesToRelationship[project.Key]; exists {
				if existingRole != "Owner" && project.RelationshipType == "Owner" {
					servicesToRelationship[project.Key] = project.RelationshipType
				}
			} else {
				servicesToRelationship[project.Key] = project.RelationshipType
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
			servicesToRoles[project.Key] = roles

		}
	}
	return servicesToRelationship, servicesToRoles
}

func GetServiceToGroup(writeGroups WriteGroupsList, servicesToRelationship map[string]string, servicesToRoles map[string][]string) []string {
	var appToRole []string
	isWriter := false
	for _, writeGroup := range writeGroups {
		for app, relationship := range servicesToRelationship {
			if relationship != writeGroup.Relationship {
				continue
			}
			for _, role := range servicesToRoles[app] {
				if slices.Contains(writeGroup.Roles, role) {
					isWriter = true
					break
				}
			}
			if isWriter {
				appToRole = append(appToRole, fmt.Sprintf("%s:w", app))
			} else {
				appToRole = append(appToRole, fmt.Sprintf("%s:r", app))
			}
		}

	}
	return appToRole
}
