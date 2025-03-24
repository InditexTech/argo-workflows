package executor

import (
	"context"
	"encoding/json"
	"fmt"

	"k8s.io/utils/ptr"

	"github.com/argoproj/argo-workflows/v3/workflow/data"
)

func (we *WorkflowExecutor) Data(ctx context.Context) error {
	dataTemplate := we.Template.Data
	if dataTemplate == nil {
		return fmt.Errorf("no data template found")
	}

	transformedData, err := data.ProcessData(dataTemplate, newExecutorDataSourceProcessor(ctx, we))
	if err != nil {
		return fmt.Errorf("unable to process data template: %w", err)
	}

	out, err := json.Marshal(transformedData)
	if err != nil {
		return err
	}
<<<<<<< HEAD
	we.Template.Outputs.Result = pointer.StringPtr(string(out))
=======
	we.Template.Outputs.Result = ptr.To(string(out))
>>>>>>> draft-3.6.5
	err = we.ReportOutputs(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
