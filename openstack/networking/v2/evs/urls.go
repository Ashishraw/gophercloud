package evs

import (
	"github.com/gophercloud/gophercloud"
	)

const resourcePathlist = "cloudvolumes/detail"
const resourcePathvolume = "volumes"

func EvsListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePathlist)
}



func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(resourcePathvolume, id)
}
