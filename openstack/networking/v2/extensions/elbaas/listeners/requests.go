package listeners

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
	"fmt"
)

type Protocol string

// Supported attributes for create/update operations.
const (
	ProtocolTCP   Protocol = "TCP"
	ProtocolHTTP  Protocol = "HTTP"
	ProtocolHTTPS Protocol = "HTTPS"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToListenerListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the floating IP attributes you want to see returned. SortKey allows you to
// sort by a particular listener attribute. SortDir sets the direction, and is
// either `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {
	LoadbalancerId              string `q:"loadbalancer_id"`
}

// ToListenerListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToListenerListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// routers. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
//
// Default policy settings return only those routers that are owned by the
// tenant who submits the request, unless an admin user submits the request.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := rootURL(c)
	if opts != nil {
		query, err := opts.ToListenerListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return ListenerPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// CreateOptsBuilder is the interface options structs have to satisfy in order
// to be used in the main Create operation in this package. Since many
// extensions decorate or modify the common logic, it is useful for them to
// satisfy a basic interface in order for them to be used.
type CreateOptsBuilder interface {
	ToListenerCreateMap() (map[string]interface{}, error)
}

// CreateOpts is the common options struct used in this package's Create
// operation.
type CreateOpts struct {
	// Human-readable name for the Listener. Does not have to be unique.
	Name string `json:"name", required:"true"`
	// Human-readable description for the Listener.
	Description string `json:"description,omitempty"`
	// The load balancer on which to provision this listener.
	LoadbalancerID string `json:"loadbalancer_id" required:"true"`
	// The protocol - can either be TCP, HTTP or HTTPS.
	Protocol Protocol `json:"protocol" required:"true"`
	// The port on which to listen for client traffic.
	ProtocolPort int `json:"port" required:"true"`
	// The protocol - can either be TCP, HTTP or HTTPS.
	BackendProtocol Protocol `json:"backend_protocol" required:"true"`
	// The port on which to listen for client traffic.
	BackendProtocolPort int `json:"backend_port" required:"true"`
	Algorithm string `json:"lb_algorithm" required:"true"`
	SessionSticky bool `json:"session_sticky,omit_empty"`
	// Indicates the owner of the Listener. Required for admins.
	Tenant_ID string `json:"tenant_id,omitempty"`
	// The ID of the default pool with which the Listener is associated.
	DefaultPoolID string `json:"default_pool_id,omitempty"`
	// The maximum number of connections allowed for the Listener.
	ConnLimit *int `json:"connection_limit,omitempty"`
	// A reference to a container of TLS secrets.
	DefaultTlsContainerRef string `json:"default_tls_container_ref,omitempty"`
	// A list of references to TLS secrets.
	SniContainerRefs []string `json:"sni_container_refs,omitempty"`
	// The administrative state of the Listener. A valid value is true (UP)
	// or false (DOWN).
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
}

// ToListenerCreateMap casts a CreateOpts struct to a map.
func (opts CreateOpts) ToListenerCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Create is an operation which provisions a new Listeners based on the
// configuration defined in the CreateOpts struct. Once the request is
// validated and progress has started on the provisioning process, a
// CreateResult will be returned.
//
// Users with an admin role can create Listeners on behalf of other tenants by
// specifying a TenantID attribute different than their own.
func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToListenerCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(rootURL(c), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Get retrieves a particular Listeners based on its unique ID.
func Get(c *gophercloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(resourceURL(c, id), &r.Body, nil)
	return
}

// UpdateOptsBuilder is the interface options structs have to satisfy in order
// to be used in the main Update operation in this package. Since many
// extensions decorate or modify the common logic, it is useful for them to
// satisfy a basic interface in order for them to be used.
type UpdateOptsBuilder interface {
	ToListenerUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts is the common options struct used in this package's Update
// operation.
type UpdateOpts struct {
	// Human-readable name for the Listener. Does not have to be unique.
	Name string `json:"name,omitempty"`
	// Human-readable description for the Listener.
	Description string `json:"description,omitempty"`
	// The maximum number of connections allowed for the Listener.
	ConnLimit *int `json:"connection_limit,omitempty"`
	// A reference to a container of TLS secrets.
	DefaultTlsContainerRef string `json:"default_tls_container_ref,omitempty"`
	//  A list of references to TLS secrets.
	SniContainerRefs []string `json:"sni_container_refs,omitempty"`
	// The administrative state of the Listener. A valid value is true (UP)
	// or false (DOWN).
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
}

// ToListenerUpdateMap casts a UpdateOpts struct to a map.
func (opts UpdateOpts) ToListenerUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Update is an operation which modifies the attributes of the specified Listener.
func Update(c *gophercloud.ServiceClient, id string, opts UpdateOpts) (r UpdateResult) {
	b, err := opts.ToListenerUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(resourceURL(c, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return
}

// Delete will permanently delete a particular Listeners based on its unique ID.
func Delete(c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	url := resourceURL(c, id)
	fmt.Printf("Delete listener url: %s.\n", url)
	_, r.Err = c.Delete1(url, &gophercloud.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
