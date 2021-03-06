package visualstudio

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// AccountResource the response to an account resource GET request.
type AccountResource struct {
	autorest.Response `json:"-"`
	// Properties - Resource properties.
	Properties map[string]*string `json:"properties"`
	// ID - Unique identifier of the resource.
	ID *string `json:"id,omitempty"`
	// Location - Resource location.
	Location *string `json:"location,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for AccountResource.
func (ar AccountResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ar.Properties != nil {
		objectMap["properties"] = ar.Properties
	}
	if ar.ID != nil {
		objectMap["id"] = ar.ID
	}
	if ar.Location != nil {
		objectMap["location"] = ar.Location
	}
	if ar.Name != nil {
		objectMap["name"] = ar.Name
	}
	if ar.Tags != nil {
		objectMap["tags"] = ar.Tags
	}
	if ar.Type != nil {
		objectMap["type"] = ar.Type
	}
	return json.Marshal(objectMap)
}

// AccountResourceListResult the response to an account resource list GET request.
type AccountResourceListResult struct {
	autorest.Response `json:"-"`
	// Value - Array of resource details.
	Value *[]AccountResource `json:"value,omitempty"`
}

// AccountResourceRequest the body of a PUT request to modify a Visual Studio account resource.
type AccountResourceRequest struct {
	// AccountName - The account name.
	AccountName *string `json:"accountName,omitempty"`
	// Location - The Azure instance location.
	Location *string `json:"location,omitempty"`
	// OperationType - The type of the operation.
	OperationType interface{} `json:"operationType,omitempty"`
	// Properties - The custom properties of the resource.
	Properties map[string]*string `json:"properties"`
	// Tags - The custom tags of the resource.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for AccountResourceRequest.
func (arr AccountResourceRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if arr.AccountName != nil {
		objectMap["accountName"] = arr.AccountName
	}
	if arr.Location != nil {
		objectMap["location"] = arr.Location
	}
	objectMap["operationType"] = arr.OperationType
	if arr.Properties != nil {
		objectMap["properties"] = arr.Properties
	}
	if arr.Tags != nil {
		objectMap["tags"] = arr.Tags
	}
	return json.Marshal(objectMap)
}

// CheckNameAvailabilityParameter the body of a POST request to check name availability.
type CheckNameAvailabilityParameter struct {
	// ResourceName - The name of the resource to check availability for.
	ResourceName *string `json:"resourceName,omitempty"`
	// ResourceType - The type of resource to check availability for.
	ResourceType *string `json:"resourceType,omitempty"`
}

// CheckNameAvailabilityResult the response to a name availability request.
type CheckNameAvailabilityResult struct {
	autorest.Response `json:"-"`
	// Message - The message describing the detailed reason.
	Message *string `json:"message,omitempty"`
	// NameAvailable - The value which indicates whether the provided name is available.
	NameAvailable *bool `json:"nameAvailable,omitempty"`
}

// ExtensionResource the response to an extension resource GET request.
type ExtensionResource struct {
	autorest.Response `json:"-"`
	// Plan - The extension plan that was purchased.
	Plan *ExtensionResourcePlan `json:"plan,omitempty"`
	// Properties - Resource properties.
	Properties map[string]*string `json:"properties"`
	// ID - Unique identifier of the resource.
	ID *string `json:"id,omitempty"`
	// Location - Resource location.
	Location *string `json:"location,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for ExtensionResource.
func (er ExtensionResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if er.Plan != nil {
		objectMap["plan"] = er.Plan
	}
	if er.Properties != nil {
		objectMap["properties"] = er.Properties
	}
	if er.ID != nil {
		objectMap["id"] = er.ID
	}
	if er.Location != nil {
		objectMap["location"] = er.Location
	}
	if er.Name != nil {
		objectMap["name"] = er.Name
	}
	if er.Tags != nil {
		objectMap["tags"] = er.Tags
	}
	if er.Type != nil {
		objectMap["type"] = er.Type
	}
	return json.Marshal(objectMap)
}

// ExtensionResourceListResult the response to an extension resource list GET request.
type ExtensionResourceListResult struct {
	autorest.Response `json:"-"`
	// Value - Array of extension resource details.
	Value *[]ExtensionResource `json:"value,omitempty"`
}

// ExtensionResourcePlan plan data for an extension resource.
type ExtensionResourcePlan struct {
	// Name - Name of the plan.
	Name *string `json:"name,omitempty"`
	// Product - Product name.
	Product *string `json:"product,omitempty"`
	// PromotionCode - Optional: the promotion code associated with the plan.
	PromotionCode *string `json:"promotionCode,omitempty"`
	// Publisher - Name of the extension publisher.
	Publisher *string `json:"publisher,omitempty"`
	// Version - A string that uniquely identifies the plan version.
	Version *string `json:"version,omitempty"`
}

// ExtensionResourceRequest the body of an extension resource PUT request.
type ExtensionResourceRequest struct {
	// Location - The Azure region of the Visual Studio account associated with this request (i.e 'southcentralus'.)
	Location *string `json:"location,omitempty"`
	// Plan - Extended information about the plan being purchased for this extension resource.
	Plan *ExtensionResourcePlan `json:"plan,omitempty"`
	// Properties - A dictionary of extended properties. This property is currently unused.
	Properties map[string]*string `json:"properties"`
	// Tags - A dictionary of user-defined tags to be stored with the extension resource.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for ExtensionResourceRequest.
func (errVar ExtensionResourceRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if errVar.Location != nil {
		objectMap["location"] = errVar.Location
	}
	if errVar.Plan != nil {
		objectMap["plan"] = errVar.Plan
	}
	if errVar.Properties != nil {
		objectMap["properties"] = errVar.Properties
	}
	if errVar.Tags != nil {
		objectMap["tags"] = errVar.Tags
	}
	return json.Marshal(objectMap)
}

// Operation properties of an operation supported by the resource provider.
type Operation struct {
	// Display - The properties of the resource operation.
	Display *OperationProperties `json:"display,omitempty"`
	// Name - The name of the resource operation.
	Name *string `json:"name,omitempty"`
}

// OperationListResult container for a list of operations supported by a resource provider.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - A list of operations supported by a resource provider.
	Value *[]Operation `json:"value,omitempty"`
}

// OperationProperties properties of an operation supported by the resource provider.
type OperationProperties struct {
	// Description - The description of the resource operation.
	Description *string `json:"description,omitempty"`
	// Operation - The operation name.
	Operation *string `json:"operation,omitempty"`
	// Provider - The provider name.
	Provider *string `json:"provider,omitempty"`
	// Resource - The resource name.
	Resource *string `json:"resource,omitempty"`
}

// ProjectResource a Visual Studio Team Services project resource.
type ProjectResource struct {
	autorest.Response `json:"-"`
	// Properties - Key/value pair of resource properties.
	Properties map[string]*string `json:"properties"`
	// ID - Unique identifier of the resource.
	ID *string `json:"id,omitempty"`
	// Location - Resource location.
	Location *string `json:"location,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for ProjectResource.
func (pr ProjectResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if pr.Properties != nil {
		objectMap["properties"] = pr.Properties
	}
	if pr.ID != nil {
		objectMap["id"] = pr.ID
	}
	if pr.Location != nil {
		objectMap["location"] = pr.Location
	}
	if pr.Name != nil {
		objectMap["name"] = pr.Name
	}
	if pr.Tags != nil {
		objectMap["tags"] = pr.Tags
	}
	if pr.Type != nil {
		objectMap["type"] = pr.Type
	}
	return json.Marshal(objectMap)
}

// ProjectResourceListResult the response to a request to list Team Services project resources in a resource
// group/account.
type ProjectResourceListResult struct {
	autorest.Response `json:"-"`
	// Value - List of project resource details.
	Value *[]ProjectResource `json:"value,omitempty"`
}

// ProjectsCreateFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type ProjectsCreateFuture struct {
	azure.Future
	req *http.Request
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future ProjectsCreateFuture) Result(client ProjectsClient) (pr ProjectResource, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "visualstudio.ProjectsCreateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		return pr, azure.NewAsyncOpIncompleteError("visualstudio.ProjectsCreateFuture")
	}
	if future.PollingMethod() == azure.PollingLocation {
		pr, err = client.CreateResponder(future.Response())
		if err != nil {
			err = autorest.NewErrorWithError(err, "visualstudio.ProjectsCreateFuture", "Result", future.Response(), "Failure responding to request")
		}
		return
	}
	var req *http.Request
	var resp *http.Response
	if future.PollingURL() != "" {
		req, err = http.NewRequest(http.MethodGet, future.PollingURL(), nil)
		if err != nil {
			return
		}
	} else {
		req = autorest.ChangeToGet(future.req)
	}
	resp, err = autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		err = autorest.NewErrorWithError(err, "visualstudio.ProjectsCreateFuture", "Result", resp, "Failure sending request")
		return
	}
	pr, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "visualstudio.ProjectsCreateFuture", "Result", resp, "Failure responding to request")
	}
	return
}

// Resource a generic Azure Resource Manager resource.
type Resource struct {
	// ID - Unique identifier of the resource.
	ID *string `json:"id,omitempty"`
	// Location - Resource location.
	Location *string `json:"location,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.ID != nil {
		objectMap["id"] = r.ID
	}
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	if r.Name != nil {
		objectMap["name"] = r.Name
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	if r.Type != nil {
		objectMap["type"] = r.Type
	}
	return json.Marshal(objectMap)
}
