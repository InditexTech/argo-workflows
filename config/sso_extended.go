package config

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/argoproj/argo-workflows/v3/server/auth/devhub"
)

type SSOExtendedLabel struct {
	ApiPassword string   `json:"apiPassword,omitempty"`
	ApiUrl      string   `json:"apiUrl,omitempty"`
	ApiEndpoint string   `json:"apiEndpoint,omitempty"`
	EnvToFilter string   `json:"env,omitempty"`
	Label       string   `json:"label,omitempty"`
	WriteGroups []string `json:"writeGroups,omitempty"`
	// The AdminGroup does not filter by label gets all the objects
	AdminGroup string `json:"adminGroup,omitempty"`
}

type ResourcesToFilter struct {
	ArrayLabels  map[string]string
	Group        string
	LabelsFilter string
}

func CanDelegateByLabel() bool {
	return os.Getenv("SSO_DELEGATE_RBAC_TO_LABEL") == "true"
}

func RbacDelegateToLabel(ctx context.Context, mail string, apiUrl, apiEndpoint, apiPassword, label string, writeGroups []string) (*ResourcesToFilter, error) {
	resourcesToFilterPopulated := &ResourcesToFilter{}
	servicesToFilter := []string{}
	devhubClient := devhub.NewClient()
	mailParts := strings.Split(mail, "@")
	servicesAndGroup, err := devhub.GetServicesAndGroup(devhubClient, apiUrl, apiEndpoint, apiPassword, mailParts[0], writeGroups)
	if err != nil {
		fmt.Printf("Can't Procces the petition on devhub to get roles %+v", err)
	}
	resourcesToFilterPopulated.Group = servicesAndGroup.Group
	if servicesAndGroup.Services != nil {
		for service := range servicesAndGroup.Services {
			servicesToFilter = append(servicesToFilter, service)
		}
		resourcesToFilterPopulated.ArrayLabels = servicesAndGroup.Services
		resourcesToFilterPopulated.LabelsFilter = fmt.Sprintf("%s in (%s)", label, strings.Join(servicesToFilter[:], ","))
	}
	return resourcesToFilterPopulated, nil
}
