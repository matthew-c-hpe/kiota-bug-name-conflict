# kiota-bug-name-conflict
Bug involving name conflict causing getter and setter ( {name}_{property_name}able ) interfaces to not be generated.

Note that in the description below the {name} variable of the two schemas is equal:

`````
Given a schema named "{name}" referenced in a path, 
which has a property named "config",
and another schema named "{name}Config" referenced in the same path,
Kiota will fail to generate the interface {name}_configable for the schema which has the "config" property.
`````

This minimal reproduction mimics the environment in which the bug was discovered - i.e. a schema that when bundled contained a schema named "{name}Config", which appears to conflict with another schema named "{name}" which has a property called "config".

The OpenAPI specs provided in this reproducer have been verified as acceptabler using `redocly lint` in its default configuration.

I have provided two minimal reproducer schemas - one bad, one OK.

The bad schema contains a conflict between a schema named "namedInstancePropertyConfig" and the "config" property of a schema named "namedInstanceProperty".

The OK schema resolves the conflict by renaming the "namedInstancePropertyConfig" schema to "namedInstancePropertyOKName".

### Example of name conflict:


OK output - no name conflict.
`````
# find output-ok/ -type f -exec grep -Ei 'type \w+_configable interface' {} \;
type NamedInstanceProperty_configable interface {
`````

Bad output - name conflict.

Note the absence of a `NamedInstanceProperty_configable` interface.
`````
# find output-bad/ -type f -exec grep -Ei 'type \w+_configable interface' {} \;
`````

Note that this behaviour is not limited to the word "Config" and a property named "config". It seems to affect any clashing name/property.

Note the absence of a `NamedInstanceProperty_fooable` interface.
`````
# find output-foo -type f -exec grep -REi 'type \w+_fooable interface' {} \;
`````
