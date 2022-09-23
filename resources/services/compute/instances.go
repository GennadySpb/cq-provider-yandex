// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
)

func Instances() *schema.Table {
	return &schema.Table{
		Name:      "yandex_compute_instances",
		Resolver:  fetchInstances,
		Multiplex: client.MultiplexBy(client.Folders),
		Columns: []schema.Column{
			{
				Name:        "id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Id"),
				Description: `Resource ID`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "folder_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FolderId"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("CreatedAt"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "zone_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ZoneId"),
			},
			{
				Name:     "platform_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PlatformId"),
			},
			{
				Name:     "resources",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Resources"),
			},
			{
				Name:     "status",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Metadata"),
			},
			{
				Name:     "metadata_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MetadataOptions"),
			},
			{
				Name:     "boot_disk",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BootDisk"),
			},
			{
				Name:     "secondary_disks",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SecondaryDisks"),
			},
			{
				Name:     "local_disks",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LocalDisks"),
			},
			{
				Name:     "filesystems",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Filesystems"),
			},
			{
				Name:     "network_interfaces",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NetworkInterfaces"),
			},
			{
				Name:     "fqdn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Fqdn"),
			},
			{
				Name:     "scheduling_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SchedulingPolicy"),
			},
			{
				Name:     "service_account_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceAccountId"),
			},
			{
				Name:     "network_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NetworkSettings"),
			},
			{
				Name:     "placement_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PlacementPolicy"),
			},
		},
	}
}
