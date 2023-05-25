package volumes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VolumeProperties struct {
	AvsDataStore              *AvsDataStore                    `json:"avsDataStore,omitempty"`
	BackupId                  *string                          `json:"backupId,omitempty"`
	BaremetalTenantId         *string                          `json:"baremetalTenantId,omitempty"`
	CapacityPoolResourceId    *string                          `json:"capacityPoolResourceId,omitempty"`
	CloneProgress             *int64                           `json:"cloneProgress,omitempty"`
	CoolAccess                *bool                            `json:"coolAccess,omitempty"`
	CoolnessPeriod            *int64                           `json:"coolnessPeriod,omitempty"`
	CreationToken             string                           `json:"creationToken"`
	DataProtection            *VolumePropertiesDataProtection  `json:"dataProtection,omitempty"`
	DefaultGroupQuotaInKiBs   *int64                           `json:"defaultGroupQuotaInKiBs,omitempty"`
	DefaultUserQuotaInKiBs    *int64                           `json:"defaultUserQuotaInKiBs,omitempty"`
	EnableSubvolumes          *EnableSubvolumes                `json:"enableSubvolumes,omitempty"`
	EncryptionKeySource       *string                          `json:"encryptionKeySource,omitempty"`
	ExportPolicy              *VolumePropertiesExportPolicy    `json:"exportPolicy,omitempty"`
	FileSystemId              *string                          `json:"fileSystemId,omitempty"`
	IsDefaultQuotaEnabled     *bool                            `json:"isDefaultQuotaEnabled,omitempty"`
	IsRestoring               *bool                            `json:"isRestoring,omitempty"`
	KerberosEnabled           *bool                            `json:"kerberosEnabled,omitempty"`
	LdapEnabled               *bool                            `json:"ldapEnabled,omitempty"`
	MaximumNumberOfFiles      *int64                           `json:"maximumNumberOfFiles,omitempty"`
	MountTargets              *[]MountTargetProperties         `json:"mountTargets,omitempty"`
	NetworkFeatures           *NetworkFeatures                 `json:"networkFeatures,omitempty"`
	NetworkSiblingSetId       *string                          `json:"networkSiblingSetId,omitempty"`
	PlacementRules            *[]PlacementKeyValuePairs        `json:"placementRules,omitempty"`
	ProtocolTypes             *[]string                        `json:"protocolTypes,omitempty"`
	ProvisioningState         *string                          `json:"provisioningState,omitempty"`
	ProximityPlacementGroup   *string                          `json:"proximityPlacementGroup,omitempty"`
	SecurityStyle             *SecurityStyle                   `json:"securityStyle,omitempty"`
	ServiceLevel              *ServiceLevel                    `json:"serviceLevel,omitempty"`
	SmbContinuouslyAvailable  *bool                            `json:"smbContinuouslyAvailable,omitempty"`
	SmbEncryption             *bool                            `json:"smbEncryption,omitempty"`
	SnapshotDirectoryVisible  *bool                            `json:"snapshotDirectoryVisible,omitempty"`
	SnapshotId                *string                          `json:"snapshotId,omitempty"`
	StorageToNetworkProximity *VolumeStorageToNetworkProximity `json:"storageToNetworkProximity,omitempty"`
	SubnetId                  string                           `json:"subnetId"`
	T2Network                 *string                          `json:"t2Network,omitempty"`
	ThroughputMibps           *float64                         `json:"throughputMibps,omitempty"`
	UnixPermissions           *string                          `json:"unixPermissions,omitempty"`
	UsageThreshold            int64                            `json:"usageThreshold"`
	VolumeGroupName           *string                          `json:"volumeGroupName,omitempty"`
	VolumeSpecName            *string                          `json:"volumeSpecName,omitempty"`
	VolumeType                *string                          `json:"volumeType,omitempty"`
}
