package actions

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToActionListQuery() (string, error)
}

// ListOpts represents options used to list actions.
type ListOpts struct {
	Limit         int    `q:"limit"`
	Marker        string `q:"marker"`
	Sort          string `q:"sort"`
	GlobalProject *bool  `q:"global_project"`
	Name          string `q:"name"`
	Target        string `q:"target"`
	Action        string `q:"action"`
	Status        string `q:"status"`
}

// ToClusterListQuery builds a query string from ListOpts.
func (opts ListOpts) ToActionListQuery() (string, error) {
	q, err := ktvpcsdk.BuildQueryString(opts)
	return q.String(), err
}

// List instructs OpenStack to provide a list of actions.
func List(client *ktvpcsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToActionListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return ActionPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves details of a single action.
func Get(client *ktvpcsdk.ServiceClient, id string) (r GetResult) {
	resp, err := client.Get(getURL(client, id), &r.Body, &ktvpcsdk.RequestOpts{OkCodes: []int{200}})
	_, r.Header, r.Err = ktvpcsdk.ParseResponse(resp, err)
	return
}
