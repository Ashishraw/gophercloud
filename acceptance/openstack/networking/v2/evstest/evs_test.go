package v1

import (
	"github.com/gophercloud/gophercloud/acceptance/clients"
	"github.com/gophercloud/gophercloud/acceptance/tools"
	"testing"
	"github.com/Ashishraw/gophercloud/openstack/networking/v2/evs"
)


// this function code will test list of volume available

func TestEvsList(t *testing.T) {
	client, err := clients.NewEvsV2Client()
	if err != nil {
		t.Fatalf("Unable to create a Evs client: %v", err)
	}

	listOpts := evs.ListOpts{}
	allEvs, err := evs.List(client, listOpts)
	if err != nil {
		t.Fatalf("Unable to list routers: %v", err)
	}
	for _, router := range allEvs {
		tools.PrintResource(t, router)
	}
}

// this function code will test single volume who's id is provided if volume is available

func TestEvsVolume(t *testing.T) {
	client, err := clients.NewEvsV2Client()
	if err != nil {
		t.Fatalf("Unable to create a Evs client: %v", err)
	}

	// available Id taken
	volume_id  := "df712809-2492-4606-97fc-4e19c1a91bb0"
	//listOpts := evs.ListOpts{}
	Result  := evs.Get(client,volume_id)
	if err != nil {
		t.Fatalf("Unable to list routers: %v", err)
	}
	tools.PrintResource(t, Result)

}

/*
func TestVpcsCRUD(t *testing.T) {
	client, err := clients.NewVpcV1Client()
	if err != nil {
		t.Fatalf("Unable to create a vpc client: %v", err)
	}

	// Create a vpc
	vpc, err := CreateVpc(t, client)
	if err != nil {
		t.Fatalf("Unable to create create: %v", err)
	}
	defer DeleteVpc(t, client, vpc.ID)

	tools.PrintResource(t, vpc)

	newName := tools.RandomString("TESTACC-", 8)
	updateOpts := &vpcs.UpdateOpts{
		Name: newName,
	}

	_, err = vpcs.Update(client, vpc.ID, updateOpts).Extract()
	if err != nil {
		t.Fatalf("Unable to update vpc: %v", err)
	}

	newVpc, err := vpcs.Get(client, vpc.ID).Extract()
	if err != nil {
		t.Fatalf("Unable to retrieve vpc: %v", err)
	}

	tools.PrintResource(t, newVpc)
}
*/