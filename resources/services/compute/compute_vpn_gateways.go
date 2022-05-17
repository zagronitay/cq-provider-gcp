package compute

import (
	"context"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/compute/v1"
)

func ComputeVpnGateways() *schema.Table {
	return &schema.Table{
		Name:          "gcp_compute_vpn_gateways",
		Description:   "Represents a HA VPN gateway  HA VPN is a high-availability (HA) Cloud VPN solution that lets you securely connect your on-premises network to your Google Cloud Virtual Private Cloud network through an IPsec VPN connection in a single region.",
		Resolver:      fetchComputeVpnGateways,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"project_id", "id"}},
		Multiplex:     client.ProjectMultiplex,
		IgnoreError:   client.IgnoreErrorHandler,
		DeleteFilter:  client.DeleteProjectFilter,
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
			},
			{
				Name:        "creation_timestamp",
				Description: "Creation timestamp in RFC3339 text format",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "An optional description of this resource Provide this property when you create the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The unique identifier for the resource This identifier is defined by the server",
				Type:        schema.TypeString,
				Resolver:    client.ResolveResourceId,
			},
			{
				Name:        "kind",
				Description: "Type of resource Always compute#vpnGateway for VPN gateways",
				Type:        schema.TypeString,
			},
			{
				Name:        "label_fingerprint",
				Description: "A fingerprint for the labels being applied to this VpnGateway, which is essentially a hash of the labels set used for optimistic locking The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet  To see the latest fingerprint, make a get() request to retrieve an VpnGateway",
				Type:        schema.TypeString,
			},
			{
				Name:        "labels",
				Description: "Labels for this resource",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "name",
				Description: "Name of the resource Provided by the client when the resource is created",
				Type:        schema.TypeString,
			},
			{
				Name:        "network",
				Description: "URL of the network to which this VPN gateway is attached Provided by the client when the VPN gateway is created",
				Type:        schema.TypeString,
			},
			{
				Name:        "region",
				Description: "URL of the region where the VPN gateway resides",
				Type:        schema.TypeString,
			},
			{
				Name:        "self_link",
				Description: "Server-defined URL for the resource",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "gcp_compute_vpn_gateway_vpn_interfaces",
				Description: "A VPN gateway interface",
				Resolver:    fetchComputeVpnGatewayVpnInterfaces,
				Columns: []schema.Column{
					{
						Name:        "vpn_gateway_cq_id",
						Description: "Unique ID of gcp_compute_vpn_gateways table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "vpn_gateway_id",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "id",
						Description: "The numeric ID of this VPN gateway interface",
						Type:        schema.TypeString,
						Resolver:    client.ResolveResourceId,
					},
					{
						Name:        "interconnect_attachment",
						Description: "URL of the interconnect attachment resource When the value of this field is present, the VPN Gateway will be used for IPsec-encrypted Cloud Interconnect; all Egress or Ingress traffic for this VPN Gateway interface will go through the specified interconnect attachment resource Not currently available in all Interconnect locations",
						Type:        schema.TypeString,
					},
					{
						Name:        "ip_address",
						Description: "The external IP address for this VPN gateway interface",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeVpnGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.Compute.VpnGateways.AggregatedList(c.ProjectId).PageToken(nextPageToken)
		list, err := c.RetryingDo(ctx, call)
		if err != nil {
			return err
		}
		output := list.(*compute.VpnGatewayAggregatedList)

		var vpnGateways []*compute.VpnGateway
		for _, items := range output.Items {
			vpnGateways = append(vpnGateways, items.VpnGateways...)
		}
		res <- vpnGateways

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func fetchComputeVpnGatewayVpnInterfaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*compute.VpnGateway)
	res <- r.VpnInterfaces
	return nil
}
