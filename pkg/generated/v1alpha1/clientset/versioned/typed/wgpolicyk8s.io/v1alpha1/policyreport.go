/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/adeniyistephen/imgvuln/pkg/apis/wgpolicyk8s.io/v1alpha1"
	scheme "github.com/adeniyistephen/imgvuln/pkg/generated/v1alpha1/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PolicyReportsGetter has a method to return a PolicyReportInterface.
// A group's client should implement this interface.
type PolicyReportsGetter interface {
	PolicyReports(namespace string) PolicyReportInterface
}

// PolicyReportInterface has methods to work with PolicyReport resources.
type PolicyReportInterface interface {
	Create(ctx context.Context, policyReport *v1alpha1.PolicyReport, opts v1.CreateOptions) (*v1alpha1.PolicyReport, error)
	Update(ctx context.Context, policyReport *v1alpha1.PolicyReport, opts v1.UpdateOptions) (*v1alpha1.PolicyReport, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.PolicyReport, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.PolicyReportList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PolicyReport, err error)
	PolicyReportExpansion
}

// policyReports implements PolicyReportInterface
type policyReports struct {
	client rest.Interface
	ns     string
}

// newPolicyReports returns a PolicyReports
func newPolicyReports(c *Wgpolicyk8sV1alpha1Client, namespace string) *policyReports {
	return &policyReports{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the policyReport, and returns the corresponding policyReport object, and an error if there is any.
func (c *policyReports) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PolicyReport, err error) {
	result = &v1alpha1.PolicyReport{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("policyreports").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PolicyReports that match those selectors.
func (c *policyReports) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PolicyReportList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PolicyReportList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("policyreports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested policyReports.
func (c *policyReports) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("policyreports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a policyReport and creates it.  Returns the server's representation of the policyReport, and an error, if there is any.
func (c *policyReports) Create(ctx context.Context, policyReport *v1alpha1.PolicyReport, opts v1.CreateOptions) (result *v1alpha1.PolicyReport, err error) {
	result = &v1alpha1.PolicyReport{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("policyreports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(policyReport).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a policyReport and updates it. Returns the server's representation of the policyReport, and an error, if there is any.
func (c *policyReports) Update(ctx context.Context, policyReport *v1alpha1.PolicyReport, opts v1.UpdateOptions) (result *v1alpha1.PolicyReport, err error) {
	result = &v1alpha1.PolicyReport{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("policyreports").
		Name(policyReport.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(policyReport).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the policyReport and deletes it. Returns an error if one occurs.
func (c *policyReports) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("policyreports").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *policyReports) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("policyreports").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched policyReport.
func (c *policyReports) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PolicyReport, err error) {
	result = &v1alpha1.PolicyReport{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("policyreports").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
