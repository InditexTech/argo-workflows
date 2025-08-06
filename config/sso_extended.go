package config

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/argoproj/argo-workflows/v3/server/auth/devhub"
	log "github.com/sirupsen/logrus"
)

type SSOExtendedLabel struct {
	ApiPassword string `json:"apiPassword,omitempty"`
	ApiUrl      string `json:"apiUrl,omitempty"`
	ApiEndpoint string `json:"apiEndpoint,omitempty"`
	Label       string `json:"label,omitempty"`
	WriteGroups devhub.WriteGroupsList
	// The AdminGroup does not filter by label gets all the objects
	AdminGroup string `json:"adminGroup,omitempty"`
}

type ResourcesToFilter struct {
	ServiceToGroup string
}

func CanDelegateByLabel() bool {
	return os.Getenv("SSO_DELEGATE_RBAC_TO_LABEL") == "true"
}

func RbacDelegateToLabel(ctx context.Context, mail string, apiUrl, apiEndpoint, apiPassword, label string, writeGroups devhub.WriteGroupsList) (*ResourcesToFilter, error) {
	resourcesToFilterPopulated := &ResourcesToFilter{}
	devhubClient := devhub.NewClient()
	mailParts := strings.Split(mail, "@")
	servicesAndGroup, err := devhub.GetServicesAndGroup(devhubClient, apiUrl, apiEndpoint, apiPassword, mailParts[0], writeGroups)
	if err != nil {
		log.WithError(err).Error(fmt.Printf("Can't Procces the petition on devhub to get roles %+v", err))

	}
	resourcesToFilterPopulated.ServiceToGroup = Compress(strings.Join(servicesAndGroup.ServiceToGroup, ","))
	return resourcesToFilterPopulated, nil
}
