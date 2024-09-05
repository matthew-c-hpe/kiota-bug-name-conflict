package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NamedInstanceProperty struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The config property
    config NamedInstanceProperty_configable
    // The foo property
    foo NamedInstanceProperty_fooable
}
// NewNamedInstanceProperty instantiates a new NamedInstanceProperty and sets the default values.
func NewNamedInstanceProperty()(*NamedInstanceProperty) {
    m := &NamedInstanceProperty{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNamedInstancePropertyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNamedInstancePropertyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNamedInstanceProperty(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NamedInstanceProperty) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConfig gets the config property value. The config property
// returns a NamedInstanceProperty_configable when successful
func (m *NamedInstanceProperty) GetConfig()(NamedInstanceProperty_configable) {
    return m.config
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NamedInstanceProperty) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["config"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNamedInstanceProperty_configFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfig(val.(NamedInstanceProperty_configable))
        }
        return nil
    }
    res["foo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNamedInstanceProperty_fooFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFoo(val.(NamedInstanceProperty_fooable))
        }
        return nil
    }
    return res
}
// GetFoo gets the foo property value. The foo property
// returns a NamedInstanceProperty_fooable when successful
func (m *NamedInstanceProperty) GetFoo()(NamedInstanceProperty_fooable) {
    return m.foo
}
// Serialize serializes information the current object
func (m *NamedInstanceProperty) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("config", m.GetConfig())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("foo", m.GetFoo())
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
func (m *NamedInstanceProperty) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConfig sets the config property value. The config property
func (m *NamedInstanceProperty) SetConfig(value NamedInstanceProperty_configable)() {
    m.config = value
}
// SetFoo sets the foo property value. The foo property
func (m *NamedInstanceProperty) SetFoo(value NamedInstanceProperty_fooable)() {
    m.foo = value
}
type NamedInstancePropertyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfig()(NamedInstanceProperty_configable)
    GetFoo()(NamedInstanceProperty_fooable)
    SetConfig(value NamedInstanceProperty_configable)()
    SetFoo(value NamedInstanceProperty_fooable)()
}
