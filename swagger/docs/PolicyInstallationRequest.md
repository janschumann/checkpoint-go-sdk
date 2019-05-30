# PolicyInstallationRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | **bool** | Set to be true in order to install the Access Control policy. By default, the value is true if Access Control policy is enabled on the input policy package, otherwise false. | [optional] [default to null]
**DesktopSecurity** | **bool** | Set to be true in order to install the Desktop Security policy. By default, the value is true if desktop security policy is enabled on the input policy package, otherwise false. | [optional] [default to null]
**InstallOnAllClusterMembersOrFail** | **bool** | Relevant for the gateway clusters. If true, the policy is installed on all the cluster members. If the installation on a cluster member fails, don&#39;t install on that cluster. | [optional] [default to null]
**PolicyPackage** | **string** | The name of the Policy Package to be installed. | [default to null]
**PrepareOnly** | **bool** | If true, prepares the policy for the installation, but doesn&#39;t install it on an installation target. | [optional] [default to null]
**Qos** | **bool** | Set to be true in order to install the QoS policy. By default, the value is true if Quality-of-Service policy is enabled on the input policy package, otherwise false. | [optional] [default to null]
**Revision** | **string** | The UID of the revision of the policy to install. | [optional] [default to null]
**Targets** | **string** | On what targets to execute this command. Targets may be identified by their name, or object unique identifier. | [default to null]
**ThreatPrevention** | **bool** | Set to be true in order to install the Threat Prevention policy. By default, the value is true if Threat Prevention policy is enabled on the input policy package, otherwise false. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


