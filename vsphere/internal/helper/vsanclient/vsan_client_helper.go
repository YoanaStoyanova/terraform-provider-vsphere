package vsanclient

import (
	"context"

	"github.com/hashicorp/terraform-provider-vsphere/vsphere/internal/helper/provider"
	vimtypes "github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vsan"
	vsantypes "github.com/vmware/govmomi/vsan/types"
)

func Reconfigure(vsanClient *vsan.Client, cluster vimtypes.ManagedObjectReference, spec vsantypes.VimVsanReconfigSpec) error {
	ctx, cancel := context.WithTimeout(context.Background(), provider.DefaultAPITimeout)
	defer cancel()

	task, err := vsanClient.VsanClusterReconfig(ctx, cluster.Reference(), spec)
	if err != nil {
		return err
	}
	return task.Wait(ctx)
}

func GetVsanConfig(vsanClient *vsan.Client, cluster vimtypes.ManagedObjectReference) (*vsantypes.VsanConfigInfoEx, error) {
	ctx, cancel := context.WithTimeout(context.Background(), provider.DefaultAPITimeout)
	defer cancel()

	vsanConfig, err := vsanClient.VsanClusterGetConfig(ctx, cluster.Reference())

	return vsanConfig, err
}
