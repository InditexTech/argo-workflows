package filter

import (
	"context"
	"fmt"
	"strings"

	"github.com/argoproj/argo-workflows/v3/config"
	"github.com/argoproj/argo-workflows/v3/server/auth"
	argoTypes "github.com/argoproj/argo-workflows/v3/server/auth/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateListOptions(ctx context.Context, listOptions *metav1.ListOptions) *metav1.ListOptions {
	filterExpressionDecompress := ""
	if ctx.Value(auth.ClaimsKey) != nil && ctx.Value(auth.ClaimsKey).(*argoTypes.Claims).TeamFilterClaims.FilterExpresion != "" {
		filterExpressionDecompress = config.Decompress(ctx.Value(auth.ClaimsKey).(*argoTypes.Claims).TeamFilterClaims.FilterExpresion)
	}

	listOptionsFiltered := &metav1.ListOptions{}
	if listOptions == nil {
		if config.CanDelegateByLabel() {
			if filterExpressionDecompress != "" {
				listOptionsFiltered = &metav1.ListOptions{LabelSelector: filterExpressionDecompress}
			}
		}
	} else {
		if listOptions.LabelSelector != "" && config.CanDelegateByLabel() {
			if filterExpressionDecompress != "" {
				listOptionsFiltered = &metav1.ListOptions{LabelSelector: fmt.Sprintf("%s,%s", filterExpressionDecompress, listOptions.LabelSelector)}
			}
		} else if listOptions.LabelSelector == "" && config.CanDelegateByLabel() {
			if filterExpressionDecompress != "" {
				listOptionsFiltered = &metav1.ListOptions{LabelSelector: filterExpressionDecompress}
			}
		} else {
			listOptionsFiltered = listOptions
		}
	}
	return listOptionsFiltered
}

func ForbidActionsIfNeeded(ctx context.Context, labels map[string]string) bool {
	serviceDecompress := ""
	if ctx.Value(auth.ClaimsKey) != nil && ctx.Value(auth.ClaimsKey).(*argoTypes.Claims).TeamFilterClaims.ServiceToGroup != "" {
		serviceDecompress = config.Decompress(ctx.Value(auth.ClaimsKey).(*argoTypes.Claims).TeamFilterClaims.ServiceToGroup)
	}

	if !config.CanDelegateByLabel() || ctx.Value(auth.ClaimsKey).(*argoTypes.Claims).TeamFilterClaims.IsAdmin || ctx.Value(auth.ClaimsKey).(*argoTypes.Claims).Issuer == "kubernetes/serviceaccount" {
		return true
	}
	if len(strings.Split(string(serviceDecompress), ",")) <= 0 {
		return false
	}
	for _, service := range strings.Split(string(serviceDecompress), ",") {
		if strings.Split(service, ":")[0] == labels[ctx.Value(auth.ClaimsKey).(*argoTypes.Claims).TeamFilterClaims.Label] && strings.Split(service, ":")[1] == "writer" {
			return true
		}
	}
	return false
}
