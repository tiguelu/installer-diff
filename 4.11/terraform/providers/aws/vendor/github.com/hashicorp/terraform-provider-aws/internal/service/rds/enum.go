package rds

const (
	ClusterRoleStatusActive  = "ACTIVE"
	ClusterRoleStatusDeleted = "DELETED"
	ClusterRoleStatusPending = "PENDING"
)

const (
	StorageTypeStandard = "standard"
	StorageTypeGp2      = "gp2"
	StorageTypeIo1      = "io1"
)

func StorageType_Values() []string {
	return []string{
		StorageTypeStandard,
		StorageTypeGp2,
		StorageTypeIo1,
	}
}

// https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/accessing-monitoring.html#Overview.DBInstance.Status.
const (
	InstanceStatusAvailable                     = "available"
	InstanceStatusBackingUp                     = "backing-up"
	InstanceStatusConfiguringEnhancedMonitoring = "configuring-enhanced-monitoring"
	InstanceStatusConfiguringLogExports         = "configuring-log-exports"
	InstanceStatusCreating                      = "creating"
	InstanceStatusDeleting                      = "deleting"
	InstanceStatusIncompatibleParameters        = "incompatible-parameters"
	InstanceStatusIncompatibleRestore           = "incompatible-restore"
	InstanceStatusModifying                     = "modifying"
	InstanceStatusStarting                      = "starting"
	InstanceStatusStopping                      = "stopping"
	InstanceStatusStorageFull                   = "storage-full"
	InstanceStatusStorageOptimization           = "storage-optimization"
)

const (
	EventSubscriptionStatusActive    = "active"
	EventSubscriptionStatusCreating  = "creating"
	EventSubscriptionStatusDeleting  = "deleting"
	EventSubscriptionStatusModifying = "modifying"
)

const (
	ExportableLogTypeAgent      = "agent"
	ExportableLogTypeAlert      = "alert"
	ExportableLogTypeAudit      = "audit"
	ExportableLogTypeError      = "error"
	ExportableLogTypeGeneral    = "general"
	ExportableLogTypeListener   = "listener"
	ExportableLogTypeOEMAgent   = "oemagent"
	ExportableLogTypePostgreSQL = "postgresql"
	ExportableLogTypeSlowQuery  = "slowquery"
	ExportableLogTypeTrace      = "trace"
	ExportableLogTypeUpgrade    = "upgrade"
)

func ExportableLogType_Values() []string {
	return []string{
		ExportableLogTypeAgent,
		ExportableLogTypeAlert,
		ExportableLogTypeAudit,
		ExportableLogTypeError,
		ExportableLogTypeGeneral,
		ExportableLogTypeListener,
		ExportableLogTypeOEMAgent,
		ExportableLogTypePostgreSQL,
		ExportableLogTypeSlowQuery,
		ExportableLogTypeTrace,
		ExportableLogTypeUpgrade,
	}
}
