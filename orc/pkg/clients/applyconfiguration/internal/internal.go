/*
Copyright 2024 The ORC Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package internal

import (
	"fmt"
	"sync"

	typed "sigs.k8s.io/structured-merge-diff/v4/typed"
)

func Parser() *typed.Parser {
	parserOnce.Do(func() {
		var err error
		parser, err = typed.NewParser(schemaYAML)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse schema: %v", err))
		}
	})
	return parser
}

var parserOnce sync.Once
var parser *typed.Parser
var schemaYAML = typed.YAMLObject(`types:
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.AllocationPool
  map:
    fields:
    - name: end
      type:
        scalar: string
      default: ""
    - name: start
      type:
        scalar: string
      default: ""
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.CloudCredentialsReference
  map:
    fields:
    - name: cloudName
      type:
        scalar: string
      default: ""
    - name: secretName
      type:
        scalar: string
      default: ""
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.HostRoute
  map:
    fields:
    - name: destination
      type:
        scalar: string
      default: ""
    - name: nextHop
      type:
        scalar: string
      default: ""
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.IPv6Options
  map:
    fields:
    - name: addressMode
      type:
        scalar: string
    - name: raMode
      type:
        scalar: string
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.Image
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.ImageSpec
      default: {}
    - name: status
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.ImageStatus
      default: {}
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.ImageContent
  map:
    fields:
    - name: containerFormat
      type:
        scalar: string
    - name: diskFormat
      type:
        scalar: string
      default: ""
    - name: download
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.ImageContentSourceDownload
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.ImageContentSourceDownload
  map:
    fields:
    - name: decompress
      type:
        scalar: string
    - name: hash
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.ImageHash
    - name: url
      type:
        scalar: string
      default: ""
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.ImageFilter
  map:
    fields:
    - name: name
      type:
        scalar: string
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.ImageHash
  map:
    fields:
    - name: algorithm
      type:
        scalar: string
      default: ""
    - name: value
      type:
        scalar: string
      default: ""
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.ImageImport
  map:
    fields:
    - name: filter
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.ImageFilter
    - name: id
      type:
        scalar: string
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.ImageProperties
  map:
    fields:
    - name: hardware
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.ImagePropertiesHardware
    - name: minDiskGB
      type:
        scalar: numeric
    - name: minMemoryMB
      type:
        scalar: numeric
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.ImagePropertiesHardware
  map:
    fields:
    - name: cdromBus
      type:
        scalar: string
    - name: cpuCores
      type:
        scalar: numeric
    - name: cpuPolicy
      type:
        scalar: string
    - name: cpuSockets
      type:
        scalar: numeric
    - name: cpuThreadPolicy
      type:
        scalar: string
    - name: cpuThreads
      type:
        scalar: numeric
    - name: diskBus
      type:
        scalar: string
    - name: scsiModel
      type:
        scalar: string
    - name: vifModel
      type:
        scalar: string
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.ImageResourceSpec
  map:
    fields:
    - name: content
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.ImageContent
    - name: name
      type:
        scalar: string
    - name: properties
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.ImageProperties
    - name: protected
      type:
        scalar: boolean
    - name: tags
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: visibility
      type:
        scalar: string
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.ImageResourceStatus
  map:
    fields:
    - name: hash
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.ImageHash
    - name: sizeB
      type:
        scalar: numeric
    - name: status
      type:
        scalar: string
    - name: virtualSizeB
      type:
        scalar: numeric
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.ImageSpec
  map:
    fields:
    - name: cloudCredentialsRef
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.CloudCredentialsReference
      default: {}
    - name: import
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.ImageImport
    - name: managedOptions
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.ManagedOptions
    - name: managementPolicy
      type:
        scalar: string
    - name: resource
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.ImageResourceSpec
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.ImageStatus
  map:
    fields:
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Condition
          elementRelationship: associative
          keys:
          - type
    - name: downloadAttempts
      type:
        scalar: numeric
    - name: id
      type:
        scalar: string
    - name: resource
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.ImageResourceStatus
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.ManagedOptions
  map:
    fields:
    - name: onDelete
      type:
        scalar: string
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.Network
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.NetworkSpec
      default: {}
    - name: status
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.NetworkStatus
      default: {}
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.NetworkFilter
  map:
    fields:
    - name: description
      type:
        scalar: string
    - name: external
      type:
        scalar: boolean
    - name: name
      type:
        scalar: string
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.NetworkImport
  map:
    fields:
    - name: filter
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.NetworkFilter
    - name: id
      type:
        scalar: string
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.NetworkResourceSpec
  map:
    fields:
    - name: adminStateUp
      type:
        scalar: boolean
    - name: availabilityZoneHints
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: description
      type:
        scalar: string
    - name: dnsDomain
      type:
        scalar: string
    - name: external
      type:
        scalar: boolean
    - name: mtu
      type:
        scalar: numeric
    - name: name
      type:
        scalar: string
    - name: portSecurityEnabled
      type:
        scalar: boolean
    - name: qosPolicyID
      type:
        scalar: string
    - name: shared
      type:
        scalar: boolean
    - name: vlanTransparent
      type:
        scalar: boolean
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.NetworkResourceStatus
  map:
    fields:
    - name: adminStateUp
      type:
        scalar: boolean
    - name: availabilityZoneHints
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: availabilityZones
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: createdAt
      type:
        scalar: string
    - name: description
      type:
        scalar: string
    - name: dnsDomain
      type:
        scalar: string
    - name: external
      type:
        scalar: boolean
    - name: id
      type:
        scalar: string
    - name: ipv4AddressScope
      type:
        scalar: string
    - name: ipv6AddressScope
      type:
        scalar: string
    - name: isDefault
      type:
        scalar: boolean
    - name: l2Adjacency
      type:
        scalar: boolean
    - name: mtu
      type:
        scalar: numeric
    - name: name
      type:
        scalar: string
    - name: portSecurityEnabled
      type:
        scalar: boolean
    - name: projectID
      type:
        scalar: string
    - name: providerNetworkType
      type:
        scalar: string
    - name: providerPhysicalNetwork
      type:
        scalar: string
    - name: providerSegmentationID
      type:
        scalar: numeric
    - name: qosPolicyID
      type:
        scalar: string
    - name: revisionNumber
      type:
        scalar: numeric
    - name: segments
      type:
        list:
          elementType:
            namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.NetworkSegment
          elementRelationship: atomic
    - name: shared
      type:
        scalar: boolean
    - name: status
      type:
        scalar: string
    - name: subnets
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: tags
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: tenantID
      type:
        scalar: string
    - name: updatedAt
      type:
        scalar: string
    - name: vlanTransparent
      type:
        scalar: boolean
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.NetworkSegment
  map:
    fields:
    - name: providerNetworkType
      type:
        scalar: string
    - name: providerPhysicalNetwork
      type:
        scalar: string
    - name: providerSegmentationID
      type:
        scalar: numeric
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.NetworkSpec
  map:
    fields:
    - name: cloudCredentialsRef
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.CloudCredentialsReference
      default: {}
    - name: import
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.NetworkImport
    - name: managedOptions
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.ManagedOptions
    - name: managementPolicy
      type:
        scalar: string
    - name: resource
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.NetworkResourceSpec
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.NetworkStatus
  map:
    fields:
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Condition
          elementRelationship: associative
          keys:
          - type
    - name: id
      type:
        scalar: string
    - name: resource
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.NetworkResourceStatus
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.Subnet
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.SubnetSpec
      default: {}
    - name: status
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.SubnetStatus
      default: {}
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.SubnetFilter
  map:
    fields:
    - name: cidr
      type:
        scalar: string
    - name: description
      type:
        scalar: string
    - name: gatewayIP
      type:
        scalar: string
    - name: ipVersion
      type:
        scalar: numeric
    - name: ipv6
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.IPv6Options
    - name: name
      type:
        scalar: string
    - name: notTags
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: notTagsAny
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: projectID
      type:
        scalar: string
    - name: tags
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: tagsAny
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.SubnetGateway
  map:
    fields:
    - name: ip
      type:
        scalar: string
    - name: type
      type:
        scalar: string
      default: ""
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.SubnetImport
  map:
    fields:
    - name: filter
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.SubnetFilter
    - name: id
      type:
        scalar: string
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.SubnetResourceSpec
  map:
    fields:
    - name: allocationPools
      type:
        list:
          elementType:
            namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.AllocationPool
          elementRelationship: associative
    - name: cidr
      type:
        scalar: string
      default: ""
    - name: description
      type:
        scalar: string
    - name: dnsNameservers
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: dnsPublishFixedIP
      type:
        scalar: boolean
    - name: enableDHCP
      type:
        scalar: boolean
    - name: gateway
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.SubnetGateway
    - name: hostRoutes
      type:
        list:
          elementType:
            namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.HostRoute
          elementRelationship: associative
    - name: ipVersion
      type:
        scalar: numeric
      default: 0
    - name: ipv6
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.IPv6Options
    - name: name
      type:
        scalar: string
    - name: networkRef
      type:
        scalar: string
      default: ""
    - name: projectID
      type:
        scalar: string
    - name: tags
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.SubnetResourceStatus
  map:
    fields:
    - name: allocationPools
      type:
        list:
          elementType:
            namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.AllocationPool
          elementRelationship: atomic
    - name: cidr
      type:
        scalar: string
      default: ""
    - name: createdAt
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: description
      type:
        scalar: string
    - name: dnsNameservers
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: dnsPublishFixedIP
      type:
        scalar: boolean
    - name: enableDHCP
      type:
        scalar: boolean
      default: false
    - name: gatewayIP
      type:
        scalar: string
    - name: hostRoutes
      type:
        list:
          elementType:
            namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.HostRoute
          elementRelationship: atomic
    - name: ipVersion
      type:
        scalar: numeric
      default: 0
    - name: ipv6AddressMode
      type:
        scalar: string
    - name: ipv6RAMode
      type:
        scalar: string
    - name: name
      type:
        scalar: string
      default: ""
    - name: networkID
      type:
        scalar: string
      default: ""
    - name: projectID
      type:
        scalar: string
      default: ""
    - name: revisionNumber
      type:
        scalar: numeric
    - name: subnetPoolID
      type:
        scalar: string
    - name: tags
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: updatedAt
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.SubnetSpec
  map:
    fields:
    - name: cloudCredentialsRef
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.CloudCredentialsReference
      default: {}
    - name: import
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.SubnetImport
    - name: managedOptions
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.ManagedOptions
    - name: managementPolicy
      type:
        scalar: string
    - name: resource
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.SubnetResourceSpec
- name: com.github.k-orc.openstack-resource-controller.api.v1alpha1.SubnetStatus
  map:
    fields:
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Condition
          elementRelationship: associative
          keys:
          - type
    - name: id
      type:
        scalar: string
    - name: resource
      type:
        namedType: com.github.k-orc.openstack-resource-controller.api.v1alpha1.SubnetResourceStatus
- name: io.k8s.apimachinery.pkg.apis.meta.v1.Condition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: message
      type:
        scalar: string
      default: ""
    - name: observedGeneration
      type:
        scalar: numeric
    - name: reason
      type:
        scalar: string
      default: ""
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.apimachinery.pkg.apis.meta.v1.FieldsV1
  map:
    elementType:
      scalar: untyped
      list:
        elementType:
          namedType: __untyped_atomic_
        elementRelationship: atomic
      map:
        elementType:
          namedType: __untyped_deduced_
        elementRelationship: separable
- name: io.k8s.apimachinery.pkg.apis.meta.v1.ManagedFieldsEntry
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: fieldsType
      type:
        scalar: string
    - name: fieldsV1
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.FieldsV1
    - name: manager
      type:
        scalar: string
    - name: operation
      type:
        scalar: string
    - name: subresource
      type:
        scalar: string
    - name: time
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
- name: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
  map:
    fields:
    - name: annotations
      type:
        map:
          elementType:
            scalar: string
    - name: creationTimestamp
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: deletionGracePeriodSeconds
      type:
        scalar: numeric
    - name: deletionTimestamp
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: finalizers
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: generateName
      type:
        scalar: string
    - name: generation
      type:
        scalar: numeric
    - name: labels
      type:
        map:
          elementType:
            scalar: string
    - name: managedFields
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ManagedFieldsEntry
          elementRelationship: atomic
    - name: name
      type:
        scalar: string
    - name: namespace
      type:
        scalar: string
    - name: ownerReferences
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference
          elementRelationship: associative
          keys:
          - uid
    - name: resourceVersion
      type:
        scalar: string
    - name: selfLink
      type:
        scalar: string
    - name: uid
      type:
        scalar: string
- name: io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
      default: ""
    - name: blockOwnerDeletion
      type:
        scalar: boolean
    - name: controller
      type:
        scalar: boolean
    - name: kind
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
    - name: uid
      type:
        scalar: string
      default: ""
    elementRelationship: atomic
- name: io.k8s.apimachinery.pkg.apis.meta.v1.Time
  scalar: untyped
- name: __untyped_atomic_
  scalar: untyped
  list:
    elementType:
      namedType: __untyped_atomic_
    elementRelationship: atomic
  map:
    elementType:
      namedType: __untyped_atomic_
    elementRelationship: atomic
- name: __untyped_deduced_
  scalar: untyped
  list:
    elementType:
      namedType: __untyped_atomic_
    elementRelationship: atomic
  map:
    elementType:
      namedType: __untyped_deduced_
    elementRelationship: separable
`)
