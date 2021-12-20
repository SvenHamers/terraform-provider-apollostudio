package client

import (
	"context"

	"github.com/machinebox/graphql"
)

type Client struct {
	ApiKey            string
	EnterPriseEnabled bool
	GraphClient       *graphql.Client
}

func (cl *Client) Init() {
	cl.GraphClient = graphql.NewClient("https://graphql.api.apollographql.com/api/graphql")
}

func (cl *Client) Query(c context.Context, q string, response interface{}) error {

	graphqlRequest := graphql.NewRequest(q)
	graphqlRequest.Header.Add("X-API-Key", cl.ApiKey)

	return cl.GraphClient.Run(c, graphqlRequest, &response)

}
