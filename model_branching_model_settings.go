/*
 * Bitbucket API
 *
 * Code against the Bitbucket API to automate simple tasks, embed Bitbucket data into your own site, build mobile or desktop apps, or even add custom UI add-ons into Bitbucket itself using the Connect framework.
 *
 * API version: 2.0
 * Contact: support@bitbucket.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package bitbucket

type BranchingModelSettings struct {
	Type_ string `json:"type"`
	Links *MilestoneLinks `json:"links,omitempty"`
	BranchTypes []BranchingModelSettingsBranchTypes `json:"branch_types,omitempty"`
	Development *BranchingModelSettingsDevelopment `json:"development,omitempty"`
	Production *BranchingModelSettingsProduction `json:"production,omitempty"`
}
