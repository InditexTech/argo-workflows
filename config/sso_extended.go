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
	Label       string   `json:"label,omitempty"`
	WriteGroups []string `json:"writeGroups,omitempty"`
	// The AdminGroup does not filter by label gets all the objects
	AdminGroup string `json:"adminGroup,omitempty"`
}

type ResourcesToFilter struct {
	ArrayLabels  []string
	Group        string
	LabelsFilter string
}

func CanDelegateByLabel() bool {
	return os.Getenv("SSO_DELEGATE_RBAC_TO_LABEL") == "true"
}

func RbacDelegateToLabel(ctx context.Context, mail string, apiUrl, apiEndpoint, apiPassword, label string, writeGroups []string) (*ResourcesToFilter, error) {
	resourcesToFilterPopulated := &ResourcesToFilter{}
	devhubClient := devhub.NewClient()
	mailParts := strings.Split(mail, "@")
	servicesAndGroup, err := devhub.GetServicesAndGroup(devhubClient, apiUrl, apiEndpoint, apiPassword, mailParts[0], writeGroups)
	if err != nil {
		fmt.Printf("Can't Procces the petition on devhub to get roles %+v", err)
	}
	resourcesToFilterPopulated.Group = servicesAndGroup.Group
	if servicesAndGroup.Services != nil {
		resourcesToFilterPopulated.ArrayLabels = append(resourcesToFilterPopulated.ArrayLabels, servicesAndGroup.Services...)
		resourcesToFilterPopulated.LabelsFilter = fmt.Sprintf("%s in (%s)", label, strings.Join(resourcesToFilterPopulated.ArrayLabels[:], ","))
	}
	return resourcesToFilterPopulated, nil
}
