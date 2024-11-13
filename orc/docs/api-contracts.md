# ORC API Contracts

## General

* Do not define API fields which are not implemented. We will define them when we implement them.
* All strings must have a maximum length, even if we have to guess.
* All lists must have a maximum length, even if we have to guess.
* Do not reference `tenant_id` anywhere, either in spec or status. Use `project_id` instead.

## ProjectID

In general, use of ProjectID in a spec requires admin credentials unless the ProjectID matches the project of the current session. This is also the default, so in general non-admin users do not require ProjectID. CAPO does not focus on use cases which require admin credentials, so CAPO does not allow ProjectID to be specified. However, ORC does intend to support use cases which require admin credentials. Therefore, consider allowing ProjectID to be specified in resource specs. It should always be reported in resource statuses when it is available.

## Resource-specific conventions

After scaffolding, each resource will require 3 custom structs:

* `CustomFilter`
* `CustomResourceSpec`
* `CustomResourceStatus`

where `Custom` is the name of the specific resource.

### Filter

This is located at `spec.import.filter` in the base object. It is used when importing a pre-existing OpenStack resource into ORC when the resource's ID is not already known.

* Filter must not contain an ID field. This is handled separately by `spec.import.id`.
* Where an equivalent filter exists in CAPO, consider copying it where possible.
* Neutron types should include FilterByNeutronTags inline

### ResourceSpec

This is located at `spec.resource` is the base object. It is only defined for managed objects (`spec.managementPolicy == 'managed'`).

* Where relevant, the ResourceSpec should include a `name` field to allow object name to be overridden
* All fields should use pre-defined validated types where possible, e.g. `OpenStackName`, `OpenStackDescription`, `IPvAny`.
* Lists should have type `set` or `map` where possible, but `atomic` lists may be necessary where a struct has no merge key.

### ResourceStatus

This is located at `status.resource` in the base object. It contains the observed state of the OpenStack object.

* ID must not be included. It is stored separately at `status.ID`.
* ResourceStatus fields should not be validated: we should store any value returned by OpenStack, even invalid ones.
  * This may require implementing separate `Spec` and `Status` variants of structs.
* Lists should be `atomic`.