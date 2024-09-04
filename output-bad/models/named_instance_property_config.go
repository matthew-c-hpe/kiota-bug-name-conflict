package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NamedInstancePropertyConfig struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The information property
    information NamedInstancePropertyConfig_informationable
}
// NewNamedInstancePropertyConfig instantiates a new NamedInstancePropertyConfig and sets the default values.
func NewNamedInstancePropertyConfig()(*NamedInstancePropertyConfig) {
    m := &NamedInstancePropertyConfig{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNamedInstancePropertyConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNamedInstancePropertyConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNamedInstancePropertyConfig(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NamedInstancePropertyConfig) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NamedInstancePropertyConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["information"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNamedInstancePropertyConfig_informationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInformation(val.(NamedInstancePropertyConfig_informationable))
        }
        return nil
    }
    return res
}
// GetInformation gets the information property value. The information property
// returns a NamedInstancePropertyConfig_informationable when successful
func (m *NamedInstancePropertyConfig) GetInformation()(NamedInstancePropertyConfig_informationable) {
    return m.information
}
// Serialize serializes information the current object
func (m *NamedInstancePropertyConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("information", m.GetInformation())
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
func (m *NamedInstancePropertyConfig) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetInformation sets the information property value. The information property
func (m *NamedInstancePropertyConfig) SetInformation(value NamedInstancePropertyConfig_informationable)() {
    m.information = value
}
type NamedInstancePropertyConfigable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInformation()(NamedInstancePropertyConfig_informationable)
    SetInformation(value NamedInstancePropertyConfig_informationable)()
}
