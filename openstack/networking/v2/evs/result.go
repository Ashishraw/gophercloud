package evs

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud"

)

type EVS struct{


	ID string `json:"id"`

	//Links map[string]string `json:"links"`

	Name string `json:"name"`

	Status string `json:"status"`

	Availability_zone string `json:"availability_zone"`

	Source_volid string `json:"source_volid"`

	//Snapshot_id string `json:"snapshot_id"`

	Volumes string `json:"volumes"`

   // attachments list[map[string]string] `json:"attachments"`

	//Description string `json:"description"`   // not coming

	Os_vol_tenant_attr string `json:"os-vol-tenant-attr:tenant_id"`


	//Volume_image_metadata string `json:"volume_image_metadata"`


	Created_at string `json:"created_at"`


	//Volume_type string `json:"volume_type"`


	Size int `json:"size"`


	Bootable string `json:"bootable"`


	Metadata map[string]string `json:"metadata"`


	Os_vol_host_attr string `json:"os-vol-host-attr:host"`


	Shareable string `json:"shareable"`


	//Message string `json:"message"`


	//Code string `json:"code"`
}

func (r EVSPage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"volumes_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

func (r EVSPage) IsEmpty() (bool, error) {
	is, err := ExtractEVS(r)
	return len(is) == 0, err
}

// Extract is a function that accepts a result and extracts a evs.
func (r commonResult) Extract() (*EVS, error) {
	var s struct {
		Evs *EVS `json:"volumes"`
	}
	err := r.ExtractInto(&s)
	return s.Evs, err
}

// ExtractEVs accepts a Page struct, specifically a EVsPage struct,
// and extracts the elements into a slice of Vpc structs. In other words,
// a generic collection is mapped into a relevant slice.


func ExtractEVS(r pagination.Page) ([]EVS, error) {
	var s struct {
		Evs []EVS `json:"volumes"`
		//Evs []EVS `json:"evs"`
		}
	err := (r.(EVSPage)).ExtractInto(&s)
	return s.Evs, err
}

// VpcPage is the page returned by a pager when traversing over a
// collection of vpcs.
type EVSPage struct {
	pagination.LinkedPageBase
}

type commonResult struct {
	gophercloud.Result
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Vpc.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Vpc.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Vpc.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its ExtractErr
// method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}
