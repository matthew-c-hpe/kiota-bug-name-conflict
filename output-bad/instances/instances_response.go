package instances

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use InstancesPostResponseable instead.
type InstancesResponse struct {
    InstancesPostResponse
}
// NewInstancesResponse instantiates a new InstancesResponse and sets the default values.
func NewInstancesResponse()(*InstancesResponse) {
    m := &InstancesResponse{
        InstancesPostResponse: *NewInstancesPostResponse(),
    }
    return m
}
// CreateInstancesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateInstancesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInstancesResponse(), nil
}
// Deprecated: This class is obsolete. Use InstancesPostResponseable instead.
type InstancesResponseable interface {
    InstancesPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
