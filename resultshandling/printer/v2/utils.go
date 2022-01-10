package v2

import (
	"github.com/armosec/k8s-interface/workloadinterface"
	"github.com/armosec/kubescape/cautils"
	"github.com/armosec/opa-utils/reporthandling/results/v1/resourcesresults"
	reporthandlingv2 "github.com/armosec/opa-utils/reporthandling/v2"
)

// finalizeV2Report finalize the results objects by copying data from map to lists
func finalizeReport(opaSessionObj *cautils.OPASessionObj) {
	if len(opaSessionObj.Report.Results) == 0 {
		opaSessionObj.Report.Results = make([]resourcesresults.Result, len(opaSessionObj.ResourcesResult))
		finalizeResults(opaSessionObj.Report.Results, opaSessionObj.ResourcesResult)
		opaSessionObj.ResourcesResult = nil
	}

	if len(opaSessionObj.Report.Resources) == 0 {
		opaSessionObj.Report.Resources = make([]reporthandlingv2.Resource, len(opaSessionObj.AllResources))
		finalizeResources(opaSessionObj.Report.Resources, opaSessionObj.AllResources)
		opaSessionObj.AllResources = nil
	}

}
func finalizeResults(results []resourcesresults.Result, resourcesResult map[string]resourcesresults.Result) {
	index := 0
	for resourceID := range resourcesResult {
		results[index] = resourcesResult[resourceID]
		index++
	}
}

func finalizeResources(resources []reporthandlingv2.Resource, allResources map[string]workloadinterface.IMetadata) {
	index := 0
	for resourceID := range allResources {
		resources[index] = reporthandlingv2.Resource{
			ResourceID: resourceID,
			Object:     allResources[resourceID],
		}
		index++
	}
}