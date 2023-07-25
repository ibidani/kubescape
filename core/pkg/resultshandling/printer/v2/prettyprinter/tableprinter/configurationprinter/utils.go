package configurationprinter

import (
	"fmt"
	"strings"

	"github.com/kubescape/opa-utils/reporthandling/results/v1/reportsummary"
)

// returns map of category ID to category controls (name and controls)
// controls will be on the map only if the are in the mapClusterControlsToCategories map
func mapCategoryToSummary(controlSummaries []reportsummary.IControlSummary, mapDisplayCtrlIDToCategory map[string]string) map[string]CategoryControls {

	mapCategoriesToCtrlSummary := map[string][]reportsummary.IControlSummary{}
	// helper map to get the category name
	mapCategoryIDToName := make(map[string]string)

	for i := range controlSummaries {
		// check if we need to print this control
		category, ok := mapDisplayCtrlIDToCategory[controlSummaries[i].GetID()]
		if !ok {
			continue
		}

		// the category on the map can be either category or subcategory, so we need to check both
		if controlSummaries[i].GetCategory().ID == category {
			if _, ok := mapCategoriesToCtrlSummary[controlSummaries[i].GetCategory().ID]; !ok {
				mapCategoryIDToName[controlSummaries[i].GetCategory().ID] = controlSummaries[i].GetCategory().Name // set category name
				mapCategoriesToCtrlSummary[controlSummaries[i].GetCategory().ID] = []reportsummary.IControlSummary{}
			}
			mapCategoriesToCtrlSummary[controlSummaries[i].GetCategory().ID] = append(mapCategoriesToCtrlSummary[controlSummaries[i].GetCategory().ID], controlSummaries[i])
			continue
		}

		if controlSummaries[i].GetCategory().SubCategory.ID == category {
			if _, ok := mapCategoriesToCtrlSummary[controlSummaries[i].GetCategory().SubCategory.ID]; !ok {
				mapCategoryIDToName[controlSummaries[i].GetCategory().SubCategory.ID] = controlSummaries[i].GetCategory().SubCategory.Name // set category name
				mapCategoriesToCtrlSummary[controlSummaries[i].GetCategory().SubCategory.ID] = []reportsummary.IControlSummary{}
			}
			mapCategoriesToCtrlSummary[controlSummaries[i].GetCategory().SubCategory.ID] = append(mapCategoriesToCtrlSummary[controlSummaries[i].GetCategory().SubCategory.ID], controlSummaries[i])
			continue
		}
	}

	mapCategoryToControls := buildCategoryToControlsMap(mapCategoriesToCtrlSummary, mapCategoryIDToName)

	return mapCategoryToControls
}

// returns map of category ID to category controls (name and controls)
func buildCategoryToControlsMap(mapCategoriesToCtrlSummary map[string][]reportsummary.IControlSummary, mapCategoryIDToName map[string]string) map[string]CategoryControls {
	mapCategoryToControls := make(map[string]CategoryControls)
	for categoryID, ctrls := range mapCategoriesToCtrlSummary {
		categoryName := mapCategoryIDToName[categoryID]
		mapCategoryToControls[categoryID] = CategoryControls{
			CategoryName:     categoryName,
			controlSummaries: ctrls,
		}
	}
	return mapCategoryToControls
}

// returns doc link for control
func getDocsForControl(controlSummary reportsummary.IControlSummary) string {
	return fmt.Sprintf("%s/%s", docsPrefix, strings.ToLower(controlSummary.GetID()))
}

// returns run command with verbose for control
func getRunCommandForControl(controlSummary reportsummary.IControlSummary) string {
	return fmt.Sprintf("%s %s -v", scanControlPrefix, controlSummary.GetID())
}