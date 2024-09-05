package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NamedInstancePropertyOKName struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The information property
    information NamedInstancePropertyOKName_informationable
}
// NewNamedInstancePropertyOKName instantiates a new NamedInstancePropertyOKName and sets the default values.
func NewNamedInstancePropertyOKName()(*NamedInstancePropertyOKName) {
    m := &NamedInstancePropertyOKName{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNamedInstancePropertyOKNameFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNamedInstancePropertyOKNameFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNamedInstancePropertyOKName(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NamedInstancePropertyOKName) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NamedInstancePropertyOKName) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["information"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNamedInstancePropertyOKName_informationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInformation(val.(NamedInstancePropertyOKName_informationable))
        }
        return nil
    }
    return res
}
// GetInformation gets the information property value. The information property
// returns a NamedInstancePropertyOKName_informationable when successful
func (m *NamedInstancePropertyOKName) GetInformation()(NamedInstancePropertyOKName_informationable) {
    return m.information
}
// Serialize serializes information the current object
func (m *NamedInstancePropertyOKName) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *NamedInstancePropertyOKName) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetInformation sets the information property value. The information property
func (m *NamedInstancePropertyOKName) SetInformation(value NamedInstancePropertyOKName_informationable)() {
    m.information = value
}
type NamedInstancePropertyOKNameable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInformation()(NamedInstancePropertyOKName_informationable)
    SetInformation(value NamedInstancePropertyOKName_informationable)()
}
