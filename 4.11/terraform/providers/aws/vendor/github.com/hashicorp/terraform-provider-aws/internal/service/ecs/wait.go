package ecs

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const (
	capacityProviderDeleteTimeout = 20 * time.Minute
	capacityProviderUpdateTimeout = 10 * time.Minute

	serviceCreateTimeout      = 2 * time.Minute
	serviceInactiveTimeout    = 10 * time.Minute
	serviceInactiveTimeoutMin = 1 * time.Second
	serviceDescribeTimeout    = 2 * time.Minute
	serviceUpdateTimeout      = 2 * time.Minute

	clusterAvailableTimeout = 10 * time.Minute
	clusterDeleteTimeout    = 10 * time.Minute
	clusterAvailableDelay   = 10 * time.Second
	clusterReadTimeout      = 2 * time.Second

	taskSetCreateTimeout = 10 * time.Minute
	taskSetDeleteTimeout = 10 * time.Minute
)

func waitCapacityProviderDeleted(conn *ecs.ECS, arn string) (*ecs.CapacityProvider, error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{ecs.CapacityProviderStatusActive},
		Target:  []string{},
		Refresh: statusCapacityProvider(conn, arn),
		Timeout: capacityProviderDeleteTimeout,
	}

	outputRaw, err := stateConf.WaitForState()

	if v, ok := outputRaw.(*ecs.CapacityProvider); ok {
		return v, err
	}

	return nil, err
}

func waitCapacityProviderUpdated(conn *ecs.ECS, arn string) (*ecs.CapacityProvider, error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{ecs.CapacityProviderUpdateStatusUpdateInProgress},
		Target:  []string{ecs.CapacityProviderUpdateStatusUpdateComplete},
		Refresh: statusCapacityProviderUpdate(conn, arn),
		Timeout: capacityProviderUpdateTimeout,
	}

	outputRaw, err := stateConf.WaitForState()

	if v, ok := outputRaw.(*ecs.CapacityProvider); ok {
		return v, err
	}

	return nil, err
}

func waitServiceStable(conn *ecs.ECS, id, cluster string) error {
	input := &ecs.DescribeServicesInput{
		Services: aws.StringSlice([]string{id}),
	}

	if cluster != "" {
		input.Cluster = aws.String(cluster)
	}

	if err := conn.WaitUntilServicesStable(input); err != nil {
		return err
	}
	return nil
}

func waitServiceInactive(conn *ecs.ECS, id, cluster string) error {
	input := &ecs.DescribeServicesInput{
		Services: aws.StringSlice([]string{id}),
	}

	if cluster != "" {
		input.Cluster = aws.String(cluster)
	}

	if err := conn.WaitUntilServicesInactive(input); err != nil {
		return err
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{serviceStatusActive, serviceStatusDraining},
		Target:     []string{serviceStatusInactive, serviceStatusNone},
		Refresh:    statusService(conn, id, cluster),
		Timeout:    serviceInactiveTimeout,
		MinTimeout: serviceInactiveTimeoutMin,
	}

	_, err := stateConf.WaitForState()

	if err != nil {
		return err
	}

	return nil
}

func waitServiceDescribeReady(conn *ecs.ECS, id, cluster string) (*ecs.DescribeServicesOutput, error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{serviceStatusInactive, serviceStatusDraining, serviceStatusNone},
		Target:  []string{serviceStatusActive},
		Refresh: statusService(conn, id, cluster),
		Timeout: serviceDescribeTimeout,
	}

	outputRaw, err := stateConf.WaitForState()

	if v, ok := outputRaw.(*ecs.DescribeServicesOutput); ok {
		return v, err
	}

	return nil, err
}

func waitClusterAvailable(ctx context.Context, conn *ecs.ECS, arn string) (*ecs.Cluster, error) { //nolint:unparam
	stateConf := &resource.StateChangeConf{
		Pending: []string{"PROVISIONING"},
		Target:  []string{"ACTIVE"},
		Refresh: statusCluster(ctx, conn, arn),
		Timeout: clusterAvailableTimeout,
		Delay:   clusterAvailableDelay,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if v, ok := outputRaw.(*ecs.Cluster); ok {
		return v, err
	}

	return nil, err
}

func waitClusterDeleted(conn *ecs.ECS, arn string) (*ecs.Cluster, error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{"ACTIVE", "DEPROVISIONING"},
		Target:  []string{"INACTIVE"},
		Refresh: statusCluster(context.Background(), conn, arn),
		Timeout: clusterDeleteTimeout,
	}

	outputRaw, err := stateConf.WaitForState()

	if v, ok := outputRaw.(*ecs.Cluster); ok {
		return v, err
	}

	return nil, err
}

func waitTaskSetStable(conn *ecs.ECS, timeout time.Duration, taskSetID, service, cluster string) error {
	stateConf := &resource.StateChangeConf{
		Pending: []string{ecs.StabilityStatusStabilizing},
		Target:  []string{ecs.StabilityStatusSteadyState},
		Refresh: stabilityStatusTaskSet(conn, taskSetID, service, cluster),
		Timeout: timeout,
	}

	_, err := stateConf.WaitForState()

	return err
}

func waitTaskSetDeleted(conn *ecs.ECS, taskSetID, service, cluster string) error {
	stateConf := &resource.StateChangeConf{
		Pending: []string{taskSetStatusActive, taskSetStatusPrimary, taskSetStatusDraining},
		Target:  []string{},
		Refresh: statusTaskSet(conn, taskSetID, service, cluster),
		Timeout: taskSetDeleteTimeout,
	}

	_, err := stateConf.WaitForState()

	return err
}
