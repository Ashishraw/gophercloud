package evs

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
	"reflect"

	)


// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the floating IP attributes you want to see returned. SortKey allows you to
// sort by a particular network attribute. SortDir sets the direction, and is
// either `asc' or `desc'. Marker and Limit are used for pagination.

type ListOpts struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Availability_zone string `json:"availability_zone"`
	Status string `json:"status"`
	}


// List returns collection of
// Evs. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
//
// Default policy settings return only those evs that are owned by the
// tenant who submits the request, unless an admin user submits the request.


func List(c *gophercloud.ServiceClient,opts ListOpts)([]EVS, error){
	u :=EvsListURL(c)

	pages, err := pagination.NewPager(c, u, func(r pagination.PageResult) pagination.Page {
		return EVSPage{pagination.LinkedPageBase{PageResult: r}}
	}).AllPages()

	allEvs, err := ExtractEVS(pages)
	if err != nil {
		return nil, err
	}

	return FilterEVS(allEvs, opts)

}

//https://evs.eu-de.otc.t-systems.com/v2/87a56a48977e42068f70ad3280c50f0e/volumes/volume_id

func Get(c *gophercloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(resourceURL(c, id), &r.Body, nil)
	return
}


//func FilterVPCs(vpcs []Vpc, opts ListOpts) ([]Vpc, error) {
func FilterEVS(Evs []EVS ,opts ListOpts) ([]EVS, error) {

	var refinedEVS []EVS
	var matched bool
	m := map[string]interface{}{}

	if opts.Status != "" {
		m["Status"] = opts.Status
	}
	if opts.ID != "" {
		m["ID"] = opts.ID
	}
	if opts.Name != "" {
		m["Name"] = opts.Name
	}
	if opts.Availability_zone != "" {
		m["Availability_zone"] = opts.Availability_zone
	}


	if len(m) > 0 && len(Evs) > 0 {
		for _, evsrange := range Evs {
			matched = true

			for key, value := range m {
				if sVal := getStructField(&evsrange, key); !(sVal == value) {
					matched = false
				}
			}

			if matched {
				refinedEVS = append(refinedEVS, evsrange)
			}
		}

	} else {
		refinedEVS = Evs
	}

	return refinedEVS, nil
}


func getStructField(v *EVS, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return string(f.String())
}

/* {
"volume": {
"backup_id": null,
"count": 1,
"availability_zone": "eu-de-02",
"description": "test_volume_1",
"size": 120,
"name": "test_volume_1",
"imageRef": null,
"volume_type": "SSD"
}
}*/
// CreateOpts contains all the values needed to create a new Evs. There are
// no. of required values.
type CreateOpts struct {
	Backup_id string `json:"backup_id,omitempty"`
	Count int `json:"count,omitempty"`
	Availability_zone string `json:"availability_zone,omitempty"`
	Description string `json:"description,omitempty"`
	Size int `json:"size,omitempty"`
	Name string `json:"name,omitempty"`
	ImageRef string `json:"imageref_type,omitempty"`
	Volume_type string `json:"volume_type,omitempty"`
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToEvsCreateMap() (map[string]interface{}, error)
}
// ToEvsCreateMap builds a create request body from CreateOpts.
func (opts CreateOpts) ToEvsCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "volume")
}

// Create accepts a CreateOpts struct and uses the values to create a new
// logical Evs. When it is created, the Evs does not have an internal
// interface - it is not associated to any subnet.

func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToEvsCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	reqOpt := &gophercloud.RequestOpts{OkCodes: []int{200}}
	_, r.Err = c.Post(EvsURLCreateEvs(c), b, &r.Body, reqOpt)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToEvsUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts contains the values used when updating a evs.
type UpdateOpts struct {
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// ToEvsUpdateMap builds an update body based on UpdateOpts.
func (opts UpdateOpts) ToEvsUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "volume")
}

// Update allows vpcs to be updated. You can update the name, administrative
// state, and the external gateway. For more information about how to set the
// external gateway for a vpc, see Create. This operation does not enable
// the update of vpc interfaces. To do this, use the AddInterface and
// RemoveInterface functions.
func Update(c *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToEvsUpdateMap()
	   if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(EvsURLupdate(c,id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateSizeBuilder interface {
	ToEvsUpdateSizeMap() (map[string]interface{}, error)
}


// UpdateOpts contains the values used when updating a evs.
type UpdateSize struct {
	Size int `json:"new_size,omitempty"`
}

// ToEvsUpdateSizeMap builds an update body based on UpdateOpts.
func (opts UpdateSize) ToEvsUpdateSizeMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "os-extend")
}


func ExtendSize(c *gophercloud.ServiceClient, id string, opts UpdateSizeBuilder) (r UpdateResult) {

	b, err := opts.ToEvsUpdateSizeMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(EvsURLSizeUpdate(c,id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}



// Delete will permanently delete a particular Evs based on its unique ID.
func Delete(c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = c.Delete(resourceURL(c, id), nil)
	return
}
