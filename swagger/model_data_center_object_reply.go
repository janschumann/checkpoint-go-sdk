/*
 * Check Point Management API
 *
 * Generated by https://github.com/janschumann/checkpoint-go-sdk.
 *
 * API version: v1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type DataCenterObjectReply struct {
	// Additional properties on the object.
	AdditionalProperties []DataCenterObjectProperty `json:"additional-properties,omitempty" xml:"additional-properties"`
	// Color of the object. Should be one of existing colors.
	Color string `json:"color,omitempty" xml:"color"`
	// Comments string.
	Comments string `json:"comments,omitempty" xml:"comments"`
	DataCenter *DataCenterServerReply `json:"data-center,omitempty" xml:"data-center"`
	DataCenterObjectMetaInfo *RemoteDataCenterObjectMetaData `json:"data-center-object-meta-info,omitempty" xml:"data-center-object-meta-info"`
	// Indicates if the object is inaccessible or deleted on Data Center Server.
	Deleted bool `json:"deleted,omitempty" xml:"deleted"`
	Domain *ApiDomainIdentifier `json:"domain,omitempty" xml:"domain"`
	// How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard.
	Groups []ApiObjectStandardIdentifier `json:"groups,omitempty" xml:"groups"`
	// Object icon.
	Icon string `json:"icon,omitempty" xml:"icon"`
	MetaInfo *MetaInfoForTopLevelReply `json:"meta-info,omitempty" xml:"meta-info"`
	// Object management name.
	Name string `json:"name,omitempty" xml:"name"`
	// Object name in the Data Center.
	NameInDataCenter string `json:"name-in-data-center,omitempty" xml:"name-in-data-center"`
	// Indicates whether the object is read-only.
	ReadOnly bool `json:"read-only,omitempty" xml:"read-only"`
	// Collection of tag objects identified by the name or UID. How much details are returned depends on the details-level field of the request. This table shows the level of detail shown when details-level is set to standard.
	Tags []ApiObjectStandardIdentifier `json:"tags,omitempty" xml:"tags"`
	// Type of the object.
	Type_ string `json:"type,omitempty" xml:"type"`
	// Object type in Data Center.
	TypeInDataCenter string `json:"type-in-data-center,omitempty" xml:"type-in-data-center"`
	// Object unique identifier.
	Uid string `json:"uid,omitempty" xml:"uid"`
	// Unique identifier of the object in the Data Center.
	UidInDataCenter string `json:"uid-in-data-center,omitempty" xml:"uid-in-data-center"`
}
