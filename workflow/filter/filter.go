package filter

import (
	"context"
	"fmt"

	"github.com/argoproj/argo-workflows/v3/config"
	"github.com/argoproj/argo-workflows/v3/server/auth"
	argoTypes "github.com/argoproj/argo-workflows/v3/server/auth/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateListOptions(ctx context.Context, listOptions *metav1.ListOptions) *metav1.ListOptions {
	listOptionsFiltered := &metav1.ListOptions{}
	if listOptions == nil {
		if config.CanDelegateByLabel() {
			if ctx.Value(auth.ClaimsKey).(*argoTypes.Claims).TeamFilterClaims.FilterExpresion != "" {
				listOptionsFiltered = &metav1.ListOptions{LabelSelector: ctx.Value(auth.ClaimsKey).(*argoTypes.Claims).TeamFilterClaims.FilterExpresion}
			}
		}
	} else {
		if listOptions.LabelSelector != "" && config.CanDelegateByLabel() {
			if ctx.Value(auth.ClaimsKey).(*argoTypes.Claims).TeamFilterClaims.FilterExpresion != "" {
				listOptionsFiltered = &metav1.ListOptions{LabelSelector: fmt.Sprintf("%s,%s", ctx.Value(auth.ClaimsKey).(*argoTypes.Claims).TeamFilterClaims.FilterExpresion, listOptions.LabelSelector)}
			}
		} else if listOptions.LabelSelector == "" && config.CanDelegateByLabel() {
			if ctx.Value(auth.ClaimsKey).(*argoTypes.Claims).TeamFilterClaims.FilterExpresion != "" {
				listOptionsFiltered = &metav1.ListOptions{LabelSelector: ctx.Value(auth.ClaimsKey).(*argoTypes.Claims).TeamFilterClaims.FilterExpresion}
			}
		} else {
			listOptionsFiltered = listOptions
		}
	}
	return listOptionsFiltered
}

func ForbidActionsIfNeeded(ctx context.Context, labels map[string]string) bool {

	if config.CanDelegateByLabel() {
		if !ctx.Value(auth.ClaimsKey).(*argoTypes.Claims).TeamFilterClaims.IsAdmin {
			if len(ctx.Value(auth.ClaimsKey).(*argoTypes.Claims).TeamFilterClaims.Values) > 0 {
				for _, labelToIdentify := range ctx.Value(auth.ClaimsKey).(*argoTypes.Claims).TeamFilterClaims.Values {
					if labelToIdentify == labels[ctx.Value(auth.ClaimsKey).(*argoTypes.Claims).TeamFilterClaims.Label] {
						return true
					}

				}
			}
			return false
		}
	}
	return true
}
