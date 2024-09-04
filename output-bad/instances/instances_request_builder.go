package instances

import (
    "context"
    i87dbae2d8a883aec7b251177237de1546d9c9ef4aca8acb03d77a0ffd56cfc08 "ApiSdk/models"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// InstancesRequestBuilder builds and executes requests for operations under \instances
type InstancesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// InstancesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InstancesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewInstancesRequestBuilderInternal instantiates a new InstancesRequestBuilder and sets the default values.
func NewInstancesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InstancesRequestBuilder) {
    m := &InstancesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/instances", pathParameters),
    }
    return m
}
// NewInstancesRequestBuilder instantiates a new InstancesRequestBuilder and sets the default values.
func NewInstancesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InstancesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInstancesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create an Instance
// Deprecated: This method is obsolete. Use PostAsInstancesPostResponse instead.
// returns a InstancesResponseable when successful
func (m *InstancesRequestBuilder) Post(ctx context.Context, body i87dbae2d8a883aec7b251177237de1546d9c9ef4aca8acb03d77a0ffd56cfc08.NamedInstancePropertyConfigable, requestConfiguration *InstancesRequestBuilderPostRequestConfiguration)(InstancesResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateInstancesResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InstancesResponseable), nil
}
// PostAsInstancesPostResponse create an Instance
// returns a InstancesPostResponseable when successful
func (m *InstancesRequestBuilder) PostAsInstancesPostResponse(ctx context.Context, body i87dbae2d8a883aec7b251177237de1546d9c9ef4aca8acb03d77a0ffd56cfc08.NamedInstancePropertyConfigable, requestConfiguration *InstancesRequestBuilderPostRequestConfiguration)(InstancesPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateInstancesPostResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InstancesPostResponseable), nil
}
// ToPostRequestInformation create an Instance
// returns a *RequestInformation when successful
func (m *InstancesRequestBuilder) ToPostRequestInformation(ctx context.Context, body i87dbae2d8a883aec7b251177237de1546d9c9ef4aca8acb03d77a0ffd56cfc08.NamedInstancePropertyConfigable, requestConfiguration *InstancesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *InstancesRequestBuilder when successful
func (m *InstancesRequestBuilder) WithUrl(rawUrl string)(*InstancesRequestBuilder) {
    return NewInstancesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
