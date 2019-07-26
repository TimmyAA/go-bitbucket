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

import (
	"time"
)

type Comment struct {
	Type_ string `json:"type"`
	Id int32 `json:"id,omitempty"`
	CreatedOn time.Time `json:"created_on,omitempty"`
	UpdatedOn time.Time `json:"updated_on,omitempty"`
	Content *IssueContent `json:"content,omitempty"`
	User *User `json:"user,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
	Parent *Comment `json:"parent,omitempty"`
	Inline *CommentInline `json:"inline,omitempty"`
	Links *CommentLinks `json:"links,omitempty"`
}
