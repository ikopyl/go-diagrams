package openstack

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type deploymentContainer struct {
	path  string
	attrs []attr.Attribute
}

var Deployment = &deploymentContainer{path: "assets/openstack/deployment"}

func (c *deploymentContainer) Charms(opts ...attr.Attribute) *node.Node {
	return node.New("charms", attr.AssetImage("assets/openstack/deployment/charms.png"), attr.Shape(attr.None))
}

func (c *deploymentContainer) Chef(opts ...attr.Attribute) *node.Node {
	return node.New("chef", attr.AssetImage("assets/openstack/deployment/chef.png"), attr.Shape(attr.None))
}

func (c *deploymentContainer) Helm(opts ...attr.Attribute) *node.Node {
	return node.New("helm", attr.AssetImage("assets/openstack/deployment/helm.png"), attr.Shape(attr.None))
}

func (c *deploymentContainer) Kolla(opts ...attr.Attribute) *node.Node {
	return node.New("kolla", attr.AssetImage("assets/openstack/deployment/kolla.png"), attr.Shape(attr.None))
}

func (c *deploymentContainer) Tripleo(opts ...attr.Attribute) *node.Node {
	return node.New("tripleo", attr.AssetImage("assets/openstack/deployment/tripleo.png"), attr.Shape(attr.None))
}

func (c *deploymentContainer) Ansible(opts ...attr.Attribute) *node.Node {
	return node.New("ansible", attr.AssetImage("assets/openstack/deployment/ansible.png"), attr.Shape(attr.None))
}
