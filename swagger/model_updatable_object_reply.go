/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type UpdatableObjectReply struct {
	AdditionalProperties *UpdatableObjectAdditionalProperties `json:"additional-properties,omitempty" xml:"additional-properties"`
	// Color of the object. Should be one of existing colors.
	Color string `json:"color,omitempty" xml:"color"`
	// Comments string.
	Comments string `json:"comments,omitempty" xml:"comments"`
	Domain *ApiDomainIdentifier `json:"domain,omitempty" xml:"domain"`
	// Object icon.
	Icon string `json:"icon,omitempty" xml:"icon"`
	MetaInfo *MetaInfoForTopLevelReply `json:"meta-info,omitempty" xml:"meta-info"`
	// Object name. Should be unique in the domain.
	Name string `json:"name,omitempty" xml:"name"`
	// Object name in the Updatable Objects Repository.
	NameInUpdatableObjectsRepository string `json:"name-in-updatable-objects-repository,omitempty" xml:"name-in-updatable-objects-repository"`
	// Indicates whether the object is read-only.
	ReadOnly bool `json:"read-only,omitempty" xml:"read-only"`
	// Collection of tag objects identified by the name or UID. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard.
	Tags []ApiObjectStandardIdentifier `json:"tags,omitempty" xml:"tags"`
	// Type of the object.
	Type_ string `json:"type,omitempty" xml:"type"`
	// Object unique identifier.
	Uid string `json:"uid,omitempty" xml:"uid"`
	// Unique identifier of the object in the Updatable Objects Repository.
	UidInUpdatableObjectsRepository string `json:"uid-in-updatable-objects-repository,omitempty" xml:"uid-in-updatable-objects-repository"`
	UpdatableObjectMetaInfo *RemoteUpdatableObjectMetaData `json:"updatable-object-meta-info,omitempty" xml:"updatable-object-meta-info"`
}
