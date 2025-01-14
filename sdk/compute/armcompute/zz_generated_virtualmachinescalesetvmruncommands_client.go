// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// VirtualMachineScaleSetVMRunCommandsClient contains the methods for the VirtualMachineScaleSetVMRunCommands group.
// Don't use this type directly, use NewVirtualMachineScaleSetVMRunCommandsClient() instead.
type VirtualMachineScaleSetVMRunCommandsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVirtualMachineScaleSetVMRunCommandsClient creates a new instance of VirtualMachineScaleSetVMRunCommandsClient with the specified values.
func NewVirtualMachineScaleSetVMRunCommandsClient(con *armcore.Connection, subscriptionID string) *VirtualMachineScaleSetVMRunCommandsClient {
	return &VirtualMachineScaleSetVMRunCommandsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - The operation to create or update the VMSS VM run command.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineScaleSetVMRunCommandsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, runCommand VirtualMachineRunCommand, options *VirtualMachineScaleSetVMRunCommandsBeginCreateOrUpdateOptions) (VirtualMachineRunCommandPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, vmScaleSetName, instanceID, runCommandName, runCommand, options)
	if err != nil {
		return VirtualMachineRunCommandPollerResponse{}, err
	}
	result := VirtualMachineRunCommandPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualMachineScaleSetVMRunCommandsClient.CreateOrUpdate", "", resp, client.createOrUpdateHandleError)
	if err != nil {
		return VirtualMachineRunCommandPollerResponse{}, err
	}
	poller := &virtualMachineRunCommandPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VirtualMachineRunCommandResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new VirtualMachineRunCommandPoller from the specified resume token.
// token - The value must come from a previous call to VirtualMachineRunCommandPoller.ResumeToken().
func (client *VirtualMachineScaleSetVMRunCommandsClient) ResumeCreateOrUpdate(ctx context.Context, token string) (VirtualMachineRunCommandPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualMachineScaleSetVMRunCommandsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return VirtualMachineRunCommandPollerResponse{}, err
	}
	poller := &virtualMachineRunCommandPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return VirtualMachineRunCommandPollerResponse{}, err
	}
	result := VirtualMachineRunCommandPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VirtualMachineRunCommandResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - The operation to create or update the VMSS VM run command.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineScaleSetVMRunCommandsClient) createOrUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, runCommand VirtualMachineRunCommand, options *VirtualMachineScaleSetVMRunCommandsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceID, runCommandName, runCommand, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualMachineScaleSetVMRunCommandsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, runCommand VirtualMachineRunCommand, options *VirtualMachineScaleSetVMRunCommandsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/runCommands/{runCommandName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	if runCommandName == "" {
		return nil, errors.New("parameter runCommandName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runCommandName}", url.PathEscape(runCommandName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json, text/json")
	return req, req.MarshalAsJSON(runCommand)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VirtualMachineScaleSetVMRunCommandsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginDelete - The operation to delete the VMSS VM run command.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineScaleSetVMRunCommandsClient) BeginDelete(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, options *VirtualMachineScaleSetVMRunCommandsBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, vmScaleSetName, instanceID, runCommandName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualMachineScaleSetVMRunCommandsClient.Delete", "", resp, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *VirtualMachineScaleSetVMRunCommandsClient) ResumeDelete(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualMachineScaleSetVMRunCommandsClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - The operation to delete the VMSS VM run command.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineScaleSetVMRunCommandsClient) deleteOperation(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, options *VirtualMachineScaleSetVMRunCommandsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceID, runCommandName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualMachineScaleSetVMRunCommandsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, options *VirtualMachineScaleSetVMRunCommandsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/runCommands/{runCommandName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	if runCommandName == "" {
		return nil, errors.New("parameter runCommandName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runCommandName}", url.PathEscape(runCommandName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *VirtualMachineScaleSetVMRunCommandsClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - The operation to get the VMSS VM run command.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineScaleSetVMRunCommandsClient) Get(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, options *VirtualMachineScaleSetVMRunCommandsGetOptions) (VirtualMachineRunCommandResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceID, runCommandName, options)
	if err != nil {
		return VirtualMachineRunCommandResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineRunCommandResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualMachineRunCommandResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualMachineScaleSetVMRunCommandsClient) getCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, options *VirtualMachineScaleSetVMRunCommandsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/runCommands/{runCommandName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	if runCommandName == "" {
		return nil, errors.New("parameter runCommandName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runCommandName}", url.PathEscape(runCommandName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualMachineScaleSetVMRunCommandsClient) getHandleResponse(resp *azcore.Response) (VirtualMachineRunCommandResponse, error) {
	var val *VirtualMachineRunCommand
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VirtualMachineRunCommandResponse{}, err
	}
	return VirtualMachineRunCommandResponse{RawResponse: resp.Response, VirtualMachineRunCommand: val}, nil
}

// getHandleError handles the Get error response.
func (client *VirtualMachineScaleSetVMRunCommandsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - The operation to get all run commands of an instance in Virtual Machine Scaleset.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineScaleSetVMRunCommandsClient) List(resourceGroupName string, vmScaleSetName string, instanceID string, options *VirtualMachineScaleSetVMRunCommandsListOptions) VirtualMachineRunCommandsListResultPager {
	return &virtualMachineRunCommandsListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceID, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp VirtualMachineRunCommandsListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.VirtualMachineRunCommandsListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *VirtualMachineScaleSetVMRunCommandsClient) listCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, options *VirtualMachineScaleSetVMRunCommandsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/runCommands"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualMachineScaleSetVMRunCommandsClient) listHandleResponse(resp *azcore.Response) (VirtualMachineRunCommandsListResultResponse, error) {
	var val *VirtualMachineRunCommandsListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VirtualMachineRunCommandsListResultResponse{}, err
	}
	return VirtualMachineRunCommandsListResultResponse{RawResponse: resp.Response, VirtualMachineRunCommandsListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *VirtualMachineScaleSetVMRunCommandsClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginUpdate - The operation to update the VMSS VM run command.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineScaleSetVMRunCommandsClient) BeginUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, runCommand VirtualMachineRunCommandUpdate, options *VirtualMachineScaleSetVMRunCommandsBeginUpdateOptions) (VirtualMachineRunCommandPollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, vmScaleSetName, instanceID, runCommandName, runCommand, options)
	if err != nil {
		return VirtualMachineRunCommandPollerResponse{}, err
	}
	result := VirtualMachineRunCommandPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualMachineScaleSetVMRunCommandsClient.Update", "", resp, client.updateHandleError)
	if err != nil {
		return VirtualMachineRunCommandPollerResponse{}, err
	}
	poller := &virtualMachineRunCommandPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VirtualMachineRunCommandResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUpdate creates a new VirtualMachineRunCommandPoller from the specified resume token.
// token - The value must come from a previous call to VirtualMachineRunCommandPoller.ResumeToken().
func (client *VirtualMachineScaleSetVMRunCommandsClient) ResumeUpdate(ctx context.Context, token string) (VirtualMachineRunCommandPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualMachineScaleSetVMRunCommandsClient.Update", token, client.updateHandleError)
	if err != nil {
		return VirtualMachineRunCommandPollerResponse{}, err
	}
	poller := &virtualMachineRunCommandPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return VirtualMachineRunCommandPollerResponse{}, err
	}
	result := VirtualMachineRunCommandPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VirtualMachineRunCommandResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Update - The operation to update the VMSS VM run command.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineScaleSetVMRunCommandsClient) update(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, runCommand VirtualMachineRunCommandUpdate, options *VirtualMachineScaleSetVMRunCommandsBeginUpdateOptions) (*azcore.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceID, runCommandName, runCommand, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *VirtualMachineScaleSetVMRunCommandsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, runCommand VirtualMachineRunCommandUpdate, options *VirtualMachineScaleSetVMRunCommandsBeginUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/runCommands/{runCommandName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	if runCommandName == "" {
		return nil, errors.New("parameter runCommandName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runCommandName}", url.PathEscape(runCommandName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json, text/json")
	return req, req.MarshalAsJSON(runCommand)
}

// updateHandleError handles the Update error response.
func (client *VirtualMachineScaleSetVMRunCommandsClient) updateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
