package evs

import (
	"github.com/gophercloud/gophercloud"
)

const resourcePathlist = "cloudvolumes/detail"
const resourcePath = "cloudvolumes"
const resourcePathvolume = "volumes"

func EvsListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePathlist)
}

func EvsURLCreateEvs(c *gophercloud.ServiceClient) string {
		return c.ServiceURL(resourcePath)
}
func EvsURLupdate(c *gophercloud.ServiceClient,id string) string {
    	return c.ServiceURL(resourcePath,id)
}

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(resourcePathvolume, id)
}
