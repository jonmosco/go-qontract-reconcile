// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package example

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	"github.com/app-sre/go-qontract-reconcile/pkg/gql"
)

// UsersResponse is returned by Users on success.
type UsersResponse struct {
	Users_v1 []UsersUsers_v1User_v1 `json:"users_v1"`
}

// GetUsers_v1 returns UsersResponse.Users_v1, and is useful for accessing the field via an interface.
func (v *UsersResponse) GetUsers_v1() []UsersUsers_v1User_v1 { return v.Users_v1 }

// UsersUsers_v1User_v1 includes the requested fields of the GraphQL type User_v1.
type UsersUsers_v1User_v1 struct {
	Path               string `json:"path"`
	Name               string `json:"name"`
	Org_username       string `json:"org_username"`
	Github_username    string `json:"github_username"`
	Slack_username     string `json:"slack_username"`
	Pagerduty_username string `json:"pagerduty_username"`
	Public_gpg_key     string `json:"public_gpg_key"`
}

// GetPath returns UsersUsers_v1User_v1.Path, and is useful for accessing the field via an interface.
func (v *UsersUsers_v1User_v1) GetPath() string { return v.Path }

// GetName returns UsersUsers_v1User_v1.Name, and is useful for accessing the field via an interface.
func (v *UsersUsers_v1User_v1) GetName() string { return v.Name }

// GetOrg_username returns UsersUsers_v1User_v1.Org_username, and is useful for accessing the field via an interface.
func (v *UsersUsers_v1User_v1) GetOrg_username() string { return v.Org_username }

// GetGithub_username returns UsersUsers_v1User_v1.Github_username, and is useful for accessing the field via an interface.
func (v *UsersUsers_v1User_v1) GetGithub_username() string { return v.Github_username }

// GetSlack_username returns UsersUsers_v1User_v1.Slack_username, and is useful for accessing the field via an interface.
func (v *UsersUsers_v1User_v1) GetSlack_username() string { return v.Slack_username }

// GetPagerduty_username returns UsersUsers_v1User_v1.Pagerduty_username, and is useful for accessing the field via an interface.
func (v *UsersUsers_v1User_v1) GetPagerduty_username() string { return v.Pagerduty_username }

// GetPublic_gpg_key returns UsersUsers_v1User_v1.Public_gpg_key, and is useful for accessing the field via an interface.
func (v *UsersUsers_v1User_v1) GetPublic_gpg_key() string { return v.Public_gpg_key }

func Users(
	ctx context.Context,
) (*UsersResponse, error) {
	req := &graphql.Request{
		OpName: "Users",
		Query: `
query Users {
	users_v1 {
		path
		name
		org_username
		github_username
		slack_username
		pagerduty_username
		public_gpg_key
	}
}
`,
	}
	var err error
	var client graphql.Client

	client, err = gql.NewQontractClient(ctx)
	if err != nil {
		return nil, err
	}

	var data UsersResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
