package config

import (
	"context"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/argoproj/argo-workflows/v3/server/auth/devhub"
)

type SSOExtendedLabel struct {
	ApiPassword string `json:"apiPassword,omitempty"`
	ApiUrl      string `json:"apiUrl,omitempty"`
	Label       string `json:"label,omitempty"`
	AdminGroup  string `json:"adminGroup,omitempty"`
}

func CanDelegateByLabel() bool {
	if os.Getenv("SSO_DELEGATE_RBAC_TO_LABEL") != "true" {
		return false
	}
	return true
}

func RbacDelegateToLabel(ctx context.Context, name string, apiUrl, apiPassword, label string) (string, []string, string, error) {
	group := ""
	labelsFilter := ""
	arrayLabels := []string{}
	devhubClient := devhub.NewClient()
	roles, err := devhub.GetDevhubRoles(devhubClient, apiUrl, apiPassword, name)
	if err != nil {
		fmt.Printf("Can't Procces the petition on devhub %+v", err)
	}
	if len(roles) != 0 {
		group = getUserGroup(roles)
	}
	services, err := devhub.GetDevhubServices(devhubClient, apiUrl, apiPassword, name)
	if err != nil {
		fmt.Printf("Can't Procces the petition on devhub %+v", err)
	}
	if services != nil {
		labelsFilter = fmt.Sprintf("%s in (%s)", label, strings.Join(services[:], ","))
		for _, service := range services {
			arrayLabels = append(arrayLabels, service)
		}
	}
	return labelsFilter, arrayLabels, group, nil
}

func getUserGroup(roles []string) string {
	sabyRole := "reader"
	if slices.Contains(roles, "Authorized Deployer") || slices.Contains(roles, "Product Owner") || slices.Contains(roles, "Technical Lead") {
		sabyRole = "writer"
	}
	return sabyRole
}
