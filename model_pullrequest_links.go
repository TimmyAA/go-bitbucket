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

type PullrequestLinks struct {
	Self *SubjectTypesRepositoryEvents `json:"self,omitempty"`
	Html *SubjectTypesRepositoryEvents `json:"html,omitempty"`
	Commits *SubjectTypesRepositoryEvents `json:"commits,omitempty"`
	Approve *SubjectTypesRepositoryEvents `json:"approve,omitempty"`
	Diff *SubjectTypesRepositoryEvents `json:"diff,omitempty"`
	Diffstat *SubjectTypesRepositoryEvents `json:"diffstat,omitempty"`
	Comments *SubjectTypesRepositoryEvents `json:"comments,omitempty"`
	Activity *SubjectTypesRepositoryEvents `json:"activity,omitempty"`
	Merge *SubjectTypesRepositoryEvents `json:"merge,omitempty"`
	Decline *SubjectTypesRepositoryEvents `json:"decline,omitempty"`
}
