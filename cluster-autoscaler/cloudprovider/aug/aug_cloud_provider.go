package aug

import (
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	"k8s.io/autoscaler/cluster-autoscaler/utils/errors"
	"github.com/golang/glog"
	"fmt"
	"k8s.io/kubernetes/plugin/pkg/scheduler/schedulercache"
)

type AugNodeGroup struct {
	Name          string
	minSize       int
	maxSize       int
}

func (nodeGroup *AugNodeGroup) MaxSize() int {
	return nodeGroup.maxSize
}

func (nodeGroup *AugNodeGroup) MinSize() int {
	return nodeGroup.minSize
}

func (nodeGroup *AugNodeGroup) TargetSize() (int, error) {
	return 2, nil
}

func (nodeGroup *AugNodeGroup) IncreaseSize(delta int) error {
	return cloudprovider.ErrNotImplemented
}

func (nodeGroup *AugNodeGroup) DeleteNodes(nodes []*apiv1.Node) error {
	return nil
}

func (nodeGroup *AugNodeGroup) DecreaseTargetSize(delta int) error {
	return nil
}

func (nodeGroup *AugNodeGroup) Id() string {
	return nodeGroup.Name
}
func (nodeGroup *AugNodeGroup) Debug() string {
	return fmt.Sprintf("%s (%d:%d)", nodeGroup.Id(), nodeGroup.MinSize(), nodeGroup.MaxSize())
}

func (nodeGroup *AugNodeGroup) Nodes() ([]string, error) {
	ids := make([]string, 0)
	ids = append(ids, "192.168.161.64")
	return ids, nil
}

func (nodeGroup *AugNodeGroup) TemplateNodeInfo() (*schedulercache.NodeInfo, error) {
	return nil, nil
}

func (nodeGroup *AugNodeGroup) Exist() bool {
	return true
}

func (nodeGroup *AugNodeGroup) Create() error {
	return cloudprovider.ErrAlreadyExist
}

func (nodeGroup *AugNodeGroup) Delete() error {
	return cloudprovider.ErrNotImplemented
}

func (nodeGroup *AugNodeGroup) Autoprovisioned() bool {
	return false
}


type augCloudProvider struct {
	nodeGroups    []*AugNodeGroup
}

func BuildAugCloudProvider(discoveryOpts cloudprovider.NodeGroupDiscoveryOptions) (cloudprovider.CloudProvider, error) {
	aug := &augCloudProvider{
	}
	return aug, nil
}

func (aug *augCloudProvider) Name() string {
	return "aug"
}

// NodeGroups returns all node groups configured for this cloud provider.
func (aug *augCloudProvider) NodeGroups() []cloudprovider.NodeGroup {
	nodeGrop := &AugNodeGroup{
		Name:    "aug",
		maxSize: 2,
		minSize: 1,
	}
	return []cloudprovider.NodeGroup{
		nodeGrop,
	}
}

// NodeGroupForNode returns the node group for the given node.
func (aug *augCloudProvider) NodeGroupForNode(node *apiv1.Node) (cloudprovider.NodeGroup, error) {
	glog.V(4).Infof("NodeGroupForNode %s - name", node.Name)
	nodeGrop := &AugNodeGroup{
		Name:    "aug",
		maxSize: 2,
		minSize: 1,
	}
	return nodeGrop, nil
}

// Pricing returns pricing model for this cloud provider or error if not available.
func (aug *augCloudProvider) Pricing() (cloudprovider.PricingModel, errors.AutoscalerError) {
	return nil, cloudprovider.ErrNotImplemented
}

// GetAvailableMachineTypes get all machine types that can be requested from the cloud provider.
func (aug *augCloudProvider) GetAvailableMachineTypes() ([]string, error) {
	return []string{}, cloudprovider.ErrNotImplemented
}

// NewNodeGroup builds a theoretical node group based on the node definition provided. The node group is not automatically
// created on the cloud provider side. The node group is not returned by NodeGroups() until it is created.
func (aug *augCloudProvider) NewNodeGroup(machineType string, labels map[string]string, extraResources map[string]resource.Quantity) (cloudprovider.NodeGroup, error) {
	return nil, cloudprovider.ErrNotImplemented
}
