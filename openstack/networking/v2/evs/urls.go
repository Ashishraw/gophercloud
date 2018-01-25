package evs

import (
	"github.com/gophercloud/gophercloud"
)

const resourcePathlist = "cloudvolumes/detail"
const resourcePath = "cloudvolumes"
const resourcePathSizeUpdate = "cloudvolumes"
const resourcePathvolume = "volumes"

// for List down all EVS
func EvsListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePathlist)
}

// for creating EVS
func EvsURLCreateEvs(c *gophercloud.ServiceClient) string {
		return c.ServiceURL(resourcePath)
}

// for EVS information updation
func EvsURLupdate(c *gophercloud.ServiceClient,id string) string {
    	return c.ServiceURL(resourcePath,id)
}

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(resourcePathvolume, id)
}
// for EVS size updation
func EvsURLSizeUpdate(c *gophercloud.ServiceClient,id  string ) string {
	action := "action"
	return c.ServiceURL(resourcePathSizeUpdate,id,action)
}