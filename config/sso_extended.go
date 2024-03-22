package config

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/argoproj/argo-workflows/v3/server/auth/devhub"
)

type SSOExtendedLabel struct {
	ApiPassword string `json:"apiPassword,omitempty"`
	ApiUrl      string `json:"apiUrl,omitempty"`
	Label       string `json:"label,omitempty"`
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

func RbacDelegateToLabel(ctx context.Context, mail string, apiUrl, apiPassword, label string) (*ResourcesToFilter, error) {
	resourcesToFilterPopulated := &ResourcesToFilter{}
	devhubClient := devhub.NewClient()
	mailParts := strings.Split(mail, "@")
	servicesAndRoles, err := devhub.GetServicesAndRoles(devhubClient, apiUrl, apiPassword, mailParts[0])
	if err != nil {
		fmt.Printf("Can't Procces the petition on devhub to get roles %+v", err)
	}
	if len(servicesAndRoles.Roles) != 0 {
		resourcesToFilterPopulated.Group = getUserGroup(servicesAndRoles.Roles)
	}
	if servicesAndRoles.Services != nil {
		for service := range servicesAndRoles.Services {
			resourcesToFilterPopulated.ArrayLabels = append(resourcesToFilterPopulated.ArrayLabels, service)
		}
		resourcesToFilterPopulated.LabelsFilter = fmt.Sprintf("%s in (%s)", label, strings.Join(resourcesToFilterPopulated.ArrayLabels[:], ","))
	}
	return resourcesToFilterPopulated, nil
}

func getUserGroup(roles map[string]string) string {
	sabyRole := "reader"
	for role := range roles {
		if role == "Authorized Deployer" || role == "Product Owner" || role == "Technical Lead" {
			sabyRole = "writer"
		}
	}
	return sabyRole
}
