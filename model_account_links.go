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

type AccountLinks struct {
	Self *SubjectTypesRepositoryEvents `json:"self,omitempty"`
	Html *SubjectTypesRepositoryEvents `json:"html,omitempty"`
	Avatar *SubjectTypesRepositoryEvents `json:"avatar,omitempty"`
	Followers *SubjectTypesRepositoryEvents `json:"followers,omitempty"`
	Following *SubjectTypesRepositoryEvents `json:"following,omitempty"`
	Repositories *SubjectTypesRepositoryEvents `json:"repositories,omitempty"`
}
