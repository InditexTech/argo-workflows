package help

<<<<<<< HEAD
const (
	root       = "https://argo-workflows.readthedocs.io/en/release-3.5"
	ArgoServer = root + "/argo-server/"
	CLI        = root + "/cli/argo"

	WorkflowTemplates                          = root + "/workflow-templates/"
	WorkflowTemplatesReferencingOtherTemplates = WorkflowTemplates + "#referencing-other-workflowtemplates"

	Scaling                        = root + "/scaling/"
	ConfigureMaximumRecursionDepth = Scaling + "#maximum-recursion-depth"
=======
import (
	"fmt"

	"github.com/argoproj/argo-workflows/v3"
>>>>>>> draft-3.6.5
)

func root() string {
	version := `latest`
	if major, minor, _, err := argo.GetVersion().MajorMinorPatch(); err == nil {
		version = fmt.Sprintf("release-%s.%s", major, minor)
	}
	return fmt.Sprintf("https://argo-workflows.readthedocs.io/en/%s", version)
}

// ArgoServer returns a URL to the argo-server documentation
func ArgoServer() string {
	return root() + "/argo-server/"
}

// CLI returns a URL to the cli documentation
func CLI() string {
	return root() + "/cli/argo"
}

// scaling returns a URL to the scaling documentation
func scaling() string {
	return root() + "/scaling/"
}

// ConfigureMaximumRecursionDepth returns a URL to the maximum recursion depth documentation
func ConfigureMaximumRecursionDepth() string {
	return scaling() + "#maximum-recursion-depth"
}

func metrics() string {
	return root() + "/metrics/"
}

func MetricHelp(metricName string) string {
	return fmt.Sprintf("%s#%s", metrics(), metricName)
}
