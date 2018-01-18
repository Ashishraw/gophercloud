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
// evc. It accepts a ListOpts struct, which allows you to filter and sort
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
