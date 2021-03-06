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

type PullrequestRendered struct {
	Title *IssueContent `json:"title,omitempty"`
	Description *IssueContent `json:"description,omitempty"`
	Reason *IssueContent `json:"reason,omitempty"`
}
