// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ExperimentActionStatus string

// Enum values for ExperimentActionStatus
const (
	ExperimentActionStatusPending    ExperimentActionStatus = "pending"
	ExperimentActionStatusInitiating ExperimentActionStatus = "initiating"
	ExperimentActionStatusRunning    ExperimentActionStatus = "running"
	ExperimentActionStatusCompleted  ExperimentActionStatus = "completed"
	ExperimentActionStatusCancelled  ExperimentActionStatus = "cancelled"
	ExperimentActionStatusStopping   ExperimentActionStatus = "stopping"
	ExperimentActionStatusStopped    ExperimentActionStatus = "stopped"
	ExperimentActionStatusFailed     ExperimentActionStatus = "failed"
)

// Values returns all known values for ExperimentActionStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ExperimentActionStatus) Values() []ExperimentActionStatus {
	return []ExperimentActionStatus{
		"pending",
		"initiating",
		"running",
		"completed",
		"cancelled",
		"stopping",
		"stopped",
		"failed",
	}
}

type ExperimentStatus string

// Enum values for ExperimentStatus
const (
	ExperimentStatusPending    ExperimentStatus = "pending"
	ExperimentStatusInitiating ExperimentStatus = "initiating"
	ExperimentStatusRunning    ExperimentStatus = "running"
	ExperimentStatusCompleted  ExperimentStatus = "completed"
	ExperimentStatusStopping   ExperimentStatus = "stopping"
	ExperimentStatusStopped    ExperimentStatus = "stopped"
	ExperimentStatusFailed     ExperimentStatus = "failed"
)

// Values returns all known values for ExperimentStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ExperimentStatus) Values() []ExperimentStatus {
	return []ExperimentStatus{
		"pending",
		"initiating",
		"running",
		"completed",
		"stopping",
		"stopped",
		"failed",
	}
}
