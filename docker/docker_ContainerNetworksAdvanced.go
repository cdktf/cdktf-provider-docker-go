// Prebuilt docker Provider for Terraform CDK (cdktf)
package docker


type ContainerNetworksAdvanced struct {
	// The name or id of the network to use.
	//
	// You can use `name` or `id` attribute from a `docker_network` resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/container#name Container#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The network aliases of the container in the specific network.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/container#aliases Container#aliases}
	Aliases *[]*string `field:"optional" json:"aliases" yaml:"aliases"`
	// The IPV4 address of the container in the specific network.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/container#ipv4_address Container#ipv4_address}
	Ipv4Address *string `field:"optional" json:"ipv4Address" yaml:"ipv4Address"`
	// The IPV6 address of the container in the specific network.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/docker/r/container#ipv6_address Container#ipv6_address}
	Ipv6Address *string `field:"optional" json:"ipv6Address" yaml:"ipv6Address"`
}
