// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// PublicIPPrefixesClient contains the methods for the PublicIPPrefixes group.
// Don't use this type directly, use NewPublicIPPrefixesClient() instead.
type PublicIPPrefixesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewPublicIPPrefixesClient creates a new instance of PublicIPPrefixesClient with the specified values.
func NewPublicIPPrefixesClient(con *armcore.Connection, subscriptionID string) *PublicIPPrefixesClient {
	return &PublicIPPrefixesClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a static or dynamic public IP prefix.
// If the operation fails it returns the *CloudError error type.
func (client *PublicIPPrefixesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, publicIPPrefixName string, parameters PublicIPPrefix, options *PublicIPPrefixesBeginCreateOrUpdateOptions) (PublicIPPrefixPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, publicIPPrefixName, parameters, options)
	if err != nil {
		return PublicIPPrefixPollerResponse{}, err
	}
	result := PublicIPPrefixPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("PublicIPPrefixesClient.CreateOrUpdate", "location", resp, client.createOrUpdateHandleError)
	if err != nil {
		return PublicIPPrefixPollerResponse{}, err
	}
	poller := &publicIPPrefixPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (PublicIPPrefixResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new PublicIPPrefixPoller from the specified resume token.
// token - The value must come from a previous call to PublicIPPrefixPoller.ResumeToken().
func (client *PublicIPPrefixesClient) ResumeCreateOrUpdate(ctx context.Context, token string) (PublicIPPrefixPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("PublicIPPrefixesClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return PublicIPPrefixPollerResponse{}, err
	}
	poller := &publicIPPrefixPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return PublicIPPrefixPollerResponse{}, err
	}
	result := PublicIPPrefixPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (PublicIPPrefixResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a static or dynamic public IP prefix.
// If the operation fails it returns the *CloudError error type.
func (client *PublicIPPrefixesClient) createOrUpdate(ctx context.Context, resourceGroupName string, publicIPPrefixName string, parameters PublicIPPrefix, options *PublicIPPrefixesBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, publicIPPrefixName, parameters, options)
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
func (client *PublicIPPrefixesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, publicIPPrefixName string, parameters PublicIPPrefix, options *PublicIPPrefixesBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publicIPPrefixName == "" {
		return nil, errors.New("parameter publicIPPrefixName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publicIpPrefixName}", url.PathEscape(publicIPPrefixName))
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
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *PublicIPPrefixesClient) createOrUpdateHandleError(resp *azcore.Response) error {
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

// BeginDelete - Deletes the specified public IP prefix.
// If the operation fails it returns the *CloudError error type.
func (client *PublicIPPrefixesClient) BeginDelete(ctx context.Context, resourceGroupName string, publicIPPrefixName string, options *PublicIPPrefixesBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, publicIPPrefixName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("PublicIPPrefixesClient.Delete", "location", resp, client.deleteHandleError)
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
func (client *PublicIPPrefixesClient) ResumeDelete(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("PublicIPPrefixesClient.Delete", token, client.deleteHandleError)
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

// Delete - Deletes the specified public IP prefix.
// If the operation fails it returns the *CloudError error type.
func (client *PublicIPPrefixesClient) deleteOperation(ctx context.Context, resourceGroupName string, publicIPPrefixName string, options *PublicIPPrefixesBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, publicIPPrefixName, options)
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
func (client *PublicIPPrefixesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, publicIPPrefixName string, options *PublicIPPrefixesBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publicIPPrefixName == "" {
		return nil, errors.New("parameter publicIPPrefixName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publicIpPrefixName}", url.PathEscape(publicIPPrefixName))
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
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *PublicIPPrefixesClient) deleteHandleError(resp *azcore.Response) error {
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

// Get - Gets the specified public IP prefix in a specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *PublicIPPrefixesClient) Get(ctx context.Context, resourceGroupName string, publicIPPrefixName string, options *PublicIPPrefixesGetOptions) (PublicIPPrefixResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, publicIPPrefixName, options)
	if err != nil {
		return PublicIPPrefixResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PublicIPPrefixResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PublicIPPrefixResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PublicIPPrefixesClient) getCreateRequest(ctx context.Context, resourceGroupName string, publicIPPrefixName string, options *PublicIPPrefixesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publicIPPrefixName == "" {
		return nil, errors.New("parameter publicIPPrefixName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publicIpPrefixName}", url.PathEscape(publicIPPrefixName))
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
	reqQP.Set("api-version", "2021-02-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PublicIPPrefixesClient) getHandleResponse(resp *azcore.Response) (PublicIPPrefixResponse, error) {
	var val *PublicIPPrefix
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return PublicIPPrefixResponse{}, err
	}
	return PublicIPPrefixResponse{RawResponse: resp.Response, PublicIPPrefix: val}, nil
}

// getHandleError handles the Get error response.
func (client *PublicIPPrefixesClient) getHandleError(resp *azcore.Response) error {
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

// List - Gets all public IP prefixes in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *PublicIPPrefixesClient) List(resourceGroupName string, options *PublicIPPrefixesListOptions) PublicIPPrefixListResultPager {
	return &publicIPPrefixListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp PublicIPPrefixListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.PublicIPPrefixListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *PublicIPPrefixesClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *PublicIPPrefixesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PublicIPPrefixesClient) listHandleResponse(resp *azcore.Response) (PublicIPPrefixListResultResponse, error) {
	var val *PublicIPPrefixListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return PublicIPPrefixListResultResponse{}, err
	}
	return PublicIPPrefixListResultResponse{RawResponse: resp.Response, PublicIPPrefixListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *PublicIPPrefixesClient) listHandleError(resp *azcore.Response) error {
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

// ListAll - Gets all the public IP prefixes in a subscription.
// If the operation fails it returns the *CloudError error type.
func (client *PublicIPPrefixesClient) ListAll(options *PublicIPPrefixesListAllOptions) PublicIPPrefixListResultPager {
	return &publicIPPrefixListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		responder: client.listAllHandleResponse,
		errorer:   client.listAllHandleError,
		advancer: func(ctx context.Context, resp PublicIPPrefixListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.PublicIPPrefixListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client *PublicIPPrefixesClient) listAllCreateRequest(ctx context.Context, options *PublicIPPrefixesListAllOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/publicIPPrefixes"
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
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *PublicIPPrefixesClient) listAllHandleResponse(resp *azcore.Response) (PublicIPPrefixListResultResponse, error) {
	var val *PublicIPPrefixListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return PublicIPPrefixListResultResponse{}, err
	}
	return PublicIPPrefixListResultResponse{RawResponse: resp.Response, PublicIPPrefixListResult: val}, nil
}

// listAllHandleError handles the ListAll error response.
func (client *PublicIPPrefixesClient) listAllHandleError(resp *azcore.Response) error {
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

// UpdateTags - Updates public IP prefix tags.
// If the operation fails it returns the *CloudError error type.
func (client *PublicIPPrefixesClient) UpdateTags(ctx context.Context, resourceGroupName string, publicIPPrefixName string, parameters TagsObject, options *PublicIPPrefixesUpdateTagsOptions) (PublicIPPrefixResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, publicIPPrefixName, parameters, options)
	if err != nil {
		return PublicIPPrefixResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PublicIPPrefixResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PublicIPPrefixResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *PublicIPPrefixesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, publicIPPrefixName string, parameters TagsObject, options *PublicIPPrefixesUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publicIPPrefixName == "" {
		return nil, errors.New("parameter publicIPPrefixName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publicIpPrefixName}", url.PathEscape(publicIPPrefixName))
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
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *PublicIPPrefixesClient) updateTagsHandleResponse(resp *azcore.Response) (PublicIPPrefixResponse, error) {
	var val *PublicIPPrefix
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return PublicIPPrefixResponse{}, err
	}
	return PublicIPPrefixResponse{RawResponse: resp.Response, PublicIPPrefix: val}, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *PublicIPPrefixesClient) updateTagsHandleError(resp *azcore.Response) error {
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
