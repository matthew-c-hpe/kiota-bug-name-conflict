package instances

import (
    i87dbae2d8a883aec7b251177237de1546d9c9ef4aca8acb03d77a0ffd56cfc08 "ApiSdk/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type InstancesPostResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The instance property
    instance i87dbae2d8a883aec7b251177237de1546d9c9ef4aca8acb03d77a0ffd56cfc08.NamedInstancePropertyable
}
// NewInstancesPostResponse instantiates a new InstancesPostResponse and sets the default values.
func NewInstancesPostResponse()(*InstancesPostResponse) {
    m := &InstancesPostResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateInstancesPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateInstancesPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInstancesPostResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *InstancesPostResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *InstancesPostResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["instance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i87dbae2d8a883aec7b251177237de1546d9c9ef4aca8acb03d77a0ffd56cfc08.CreateNamedInstancePropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstance(val.(i87dbae2d8a883aec7b251177237de1546d9c9ef4aca8acb03d77a0ffd56cfc08.NamedInstancePropertyable))
        }
        return nil
    }
    return res
}
// GetInstance gets the instance property value. The instance property
// returns a NamedInstancePropertyable when successful
func (m *InstancesPostResponse) GetInstance()(i87dbae2d8a883aec7b251177237de1546d9c9ef4aca8acb03d77a0ffd56cfc08.NamedInstancePropertyable) {
    return m.instance
}
// Serialize serializes information the current object
func (m *InstancesPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("instance", m.GetInstance())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InstancesPostResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetInstance sets the instance property value. The instance property
func (m *InstancesPostResponse) SetInstance(value i87dbae2d8a883aec7b251177237de1546d9c9ef4aca8acb03d77a0ffd56cfc08.NamedInstancePropertyable)() {
    m.instance = value
}
type InstancesPostResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInstance()(i87dbae2d8a883aec7b251177237de1546d9c9ef4aca8acb03d77a0ffd56cfc08.NamedInstancePropertyable)
    SetInstance(value i87dbae2d8a883aec7b251177237de1546d9c9ef4aca8acb03d77a0ffd56cfc08.NamedInstancePropertyable)()
}
