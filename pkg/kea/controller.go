// Code generated by internal/generate/api/main.go; DO NOT EDIT.

package kea

import "github.com/browningluke/opnsense-go/pkg/api"

const keaReconfigureEndpoint = "/kea/service/reconfigure"

// Controller for kea
type Controller struct {
	Api *api.Client
}

func (c *Controller) Client() *api.Client {
	return c.Api
}