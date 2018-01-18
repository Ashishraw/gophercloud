package v1

import (
	"testing"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/gophercloud/gophercloud/openstack/networking/v1/vpcs"
)

// this function will test creation of VPC with provided option
func CreateVpc(t *testing.T, client *gophercloud.ServiceClient) (*vpcs.Vpc, error) {

	vpcName := tools.RandomString("TESTACC-", 8)

	createOpts := vpcs.CreateOpts{
		Name: vpcName,
		CIDR: "192.168.0.0/16",
	}

	t.Logf("Attempting to create vpc: %s", vpcName)

	vpc, err := vpcs.Create(client, createOpts).Extract()
	if err != nil {
		return vpc, err
	}
	t.Logf("Created vpc: %s", vpcName)

	return vpc, nil
}

func DeleteVpc(t *testing.T, client *gophercloud.ServiceClient, vpcID string) {
	t.Logf("Attempting to delete vpc: %s", vpcID)

	err := vpcs.Delete(client, vpcID).ExtractErr()
	if err != nil {
		t.Fatalf("Error deleting vpc: %v", err)
	}

	t.Logf("Deleted vpc: %s", vpcID)
}
