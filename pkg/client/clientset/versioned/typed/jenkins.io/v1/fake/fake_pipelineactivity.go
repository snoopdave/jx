package fake

import (
	jenkins_io_v1 "github.com/jenkins-x/jx/pkg/apis/jenkins.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePipelineActivities implements PipelineActivityInterface
type FakePipelineActivities struct {
	Fake *FakeJenkinsV1
	ns   string
}

var pipelineactivitiesResource = schema.GroupVersionResource{Group: "jenkins.io", Version: "v1", Resource: "pipelineactivities"}

var pipelineactivitiesKind = schema.GroupVersionKind{Group: "jenkins.io", Version: "v1", Kind: "PipelineActivity"}

// Get takes name of the pipelineActivity, and returns the corresponding pipelineActivity object, and an error if there is any.
func (c *FakePipelineActivities) Get(name string, options v1.GetOptions) (result *jenkins_io_v1.PipelineActivity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pipelineactivitiesResource, c.ns, name), &jenkins_io_v1.PipelineActivity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkins_io_v1.PipelineActivity), err
}

// List takes label and field selectors, and returns the list of PipelineActivities that match those selectors.
func (c *FakePipelineActivities) List(opts v1.ListOptions) (result *jenkins_io_v1.PipelineActivityList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pipelineactivitiesResource, pipelineactivitiesKind, c.ns, opts), &jenkins_io_v1.PipelineActivityList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &jenkins_io_v1.PipelineActivityList{}
	for _, item := range obj.(*jenkins_io_v1.PipelineActivityList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pipelineActivities.
func (c *FakePipelineActivities) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pipelineactivitiesResource, c.ns, opts))

}

// Create takes the representation of a pipelineActivity and creates it.  Returns the server's representation of the pipelineActivity, and an error, if there is any.
func (c *FakePipelineActivities) Create(pipelineActivity *jenkins_io_v1.PipelineActivity) (result *jenkins_io_v1.PipelineActivity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pipelineactivitiesResource, c.ns, pipelineActivity), &jenkins_io_v1.PipelineActivity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkins_io_v1.PipelineActivity), err
}

// Update takes the representation of a pipelineActivity and updates it. Returns the server's representation of the pipelineActivity, and an error, if there is any.
func (c *FakePipelineActivities) Update(pipelineActivity *jenkins_io_v1.PipelineActivity) (result *jenkins_io_v1.PipelineActivity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pipelineactivitiesResource, c.ns, pipelineActivity), &jenkins_io_v1.PipelineActivity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkins_io_v1.PipelineActivity), err
}

// Delete takes name of the pipelineActivity and deletes it. Returns an error if one occurs.
func (c *FakePipelineActivities) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pipelineactivitiesResource, c.ns, name), &jenkins_io_v1.PipelineActivity{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePipelineActivities) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pipelineactivitiesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &jenkins_io_v1.PipelineActivityList{})
	return err
}

// Patch applies the patch and returns the patched pipelineActivity.
func (c *FakePipelineActivities) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *jenkins_io_v1.PipelineActivity, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pipelineactivitiesResource, c.ns, name, data, subresources...), &jenkins_io_v1.PipelineActivity{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkins_io_v1.PipelineActivity), err
}