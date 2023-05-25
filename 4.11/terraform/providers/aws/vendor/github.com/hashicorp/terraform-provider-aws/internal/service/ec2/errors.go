package ec2

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	multierror "github.com/hashicorp/go-multierror"
)

const (
	ErrCodeAuthFailure                                    = "AuthFailure"
	ErrCodeClientInvalidHostIDNotFound                    = "Client.InvalidHostID.NotFound"
	ErrCodeDefaultSubnetAlreadyExistsInAvailabilityZone   = "DefaultSubnetAlreadyExistsInAvailabilityZone"
	ErrCodeDependencyViolation                            = "DependencyViolation"
	ErrCodeGatewayNotAttached                             = "Gateway.NotAttached"
	ErrCodeIncorrectState                                 = "IncorrectState"
	ErrCodeInvalidAddressNotFound                         = "InvalidAddress.NotFound"
	ErrCodeInvalidAllocationIDNotFound                    = "InvalidAllocationID.NotFound"
	ErrCodeInvalidAssociationIDNotFound                   = "InvalidAssociationID.NotFound"
	ErrCodeInvalidAttachmentIDNotFound                    = "InvalidAttachmentID.NotFound"
	ErrCodeInvalidCarrierGatewayIDNotFound                = "InvalidCarrierGatewayID.NotFound"
	ErrCodeInvalidClientVpnActiveAssociationNotFound      = "InvalidClientVpnActiveAssociationNotFound"
	ErrCodeInvalidClientVpnAssociationIdNotFound          = "InvalidClientVpnAssociationIdNotFound"
	ErrCodeInvalidClientVpnAuthorizationRuleNotFound      = "InvalidClientVpnEndpointAuthorizationRuleNotFound"
	ErrCodeInvalidClientVpnEndpointIdNotFound             = "InvalidClientVpnEndpointId.NotFound"
	ErrCodeInvalidClientVpnRouteNotFound                  = "InvalidClientVpnRouteNotFound"
	ErrCodeInvalidCustomerGatewayIDNotFound               = "InvalidCustomerGatewayID.NotFound"
	ErrCodeInvalidDhcpOptionIDNotFound                    = "InvalidDhcpOptionID.NotFound"
	ErrCodeInvalidFlowLogIdNotFound                       = "InvalidFlowLogId.NotFound"
	ErrCodeInvalidGatewayIDNotFound                       = "InvalidGatewayID.NotFound"
	ErrCodeInvalidGroupNotFound                           = "InvalidGroup.NotFound"
	ErrCodeInvalidHostIDNotFound                          = "InvalidHostID.NotFound"
	ErrCodeInvalidInstanceIDNotFound                      = "InvalidInstanceID.NotFound"
	ErrCodeInvalidInternetGatewayIDNotFound               = "InvalidInternetGatewayID.NotFound"
	ErrCodeInvalidKeyPairNotFound                         = "InvalidKeyPair.NotFound"
	ErrCodeInvalidNetworkAclEntryNotFound                 = "InvalidNetworkAclEntry.NotFound"
	ErrCodeInvalidNetworkAclIDNotFound                    = "InvalidNetworkAclID.NotFound"
	ErrCodeInvalidNetworkInterfaceIDNotFound              = "InvalidNetworkInterfaceID.NotFound"
	ErrCodeInvalidNetworkInsightsPathIdNotFound           = "InvalidNetworkInsightsPathId.NotFound"
	ErrCodeInvalidParameter                               = "InvalidParameter"
	ErrCodeInvalidParameterException                      = "InvalidParameterException"
	ErrCodeInvalidParameterValue                          = "InvalidParameterValue"
	ErrCodeInvalidPermissionDuplicate                     = "InvalidPermission.Duplicate"
	ErrCodeInvalidPermissionMalformed                     = "InvalidPermission.Malformed"
	ErrCodeInvalidPermissionNotFound                      = "InvalidPermission.NotFound"
	ErrCodeInvalidPlacementGroupUnknown                   = "InvalidPlacementGroup.Unknown"
	ErrCodeInvalidPoolIDNotFound                          = "InvalidPoolID.NotFound"
	ErrCodeInvalidPrefixListIDNotFound                    = "InvalidPrefixListID.NotFound"
	ErrCodeInvalidRouteNotFound                           = "InvalidRoute.NotFound"
	ErrCodeInvalidRouteTableIDNotFound                    = "InvalidRouteTableID.NotFound"
	ErrCodeInvalidRouteTableIdNotFound                    = "InvalidRouteTableId.NotFound"
	ErrCodeInvalidSecurityGroupIDNotFound                 = "InvalidSecurityGroupID.NotFound"
	ErrCodeInvalidSnapshotInUse                           = "InvalidSnapshot.InUse"
	ErrCodeInvalidSnapshotNotFound                        = "InvalidSnapshot.NotFound"
	ErrCodeInvalidSpotDatafeedNotFound                    = "InvalidSpotDatafeed.NotFound"
	ErrCodeInvalidSpotInstanceRequestIDNotFound           = "InvalidSpotInstanceRequestID.NotFound"
	ErrCodeInvalidSubnetCidrReservationIDNotFound         = "InvalidSubnetCidrReservationID.NotFound"
	ErrCodeInvalidSubnetIDNotFound                        = "InvalidSubnetID.NotFound"
	ErrCodeInvalidSubnetIdNotFound                        = "InvalidSubnetId.NotFound"
	ErrCodeInvalidTransitGatewayAttachmentIDNotFound      = "InvalidTransitGatewayAttachmentID.NotFound"
	ErrCodeInvalidTransitGatewayConnectPeerIDNotFound     = "InvalidTransitGatewayConnectPeerID.NotFound"
	ErrCodeInvalidTransitGatewayIDNotFound                = "InvalidTransitGatewayID.NotFound"
	ErrCodeInvalidTransitGatewayMulticastDomainIdNotFound = "InvalidTransitGatewayMulticastDomainId.NotFound"
	ErrCodeInvalidVolumeNotFound                          = "InvalidVolume.NotFound"
	ErrCodeInvalidVpcCidrBlockAssociationIDNotFound       = "InvalidVpcCidrBlockAssociationID.NotFound"
	ErrCodeInvalidVpcEndpointIdNotFound                   = "InvalidVpcEndpointId.NotFound"
	ErrCodeInvalidVpcEndpointNotFound                     = "InvalidVpcEndpoint.NotFound"
	ErrCodeInvalidVpcEndpointServiceIdNotFound            = "InvalidVpcEndpointServiceId.NotFound"
	ErrCodeInvalidVpcIDNotFound                           = "InvalidVpcID.NotFound"
	ErrCodeInvalidVpcPeeringConnectionIDNotFound          = "InvalidVpcPeeringConnectionID.NotFound"
	ErrCodeInvalidVpnConnectionIDNotFound                 = "InvalidVpnConnectionID.NotFound"
	ErrCodeInvalidVpnGatewayAttachmentNotFound            = "InvalidVpnGatewayAttachment.NotFound"
	ErrCodeInvalidVpnGatewayIDNotFound                    = "InvalidVpnGatewayID.NotFound"
	ErrCodeNatGatewayNotFound                             = "NatGatewayNotFound"
	ErrCodeUnsupportedOperation                           = "UnsupportedOperation"
)

func UnsuccessfulItemError(apiObject *ec2.UnsuccessfulItemError) error {
	if apiObject == nil {
		return nil
	}

	return awserr.New(aws.StringValue(apiObject.Code), aws.StringValue(apiObject.Message), nil)
}

func UnsuccessfulItemsError(apiObjects []*ec2.UnsuccessfulItem) error {
	var errors *multierror.Error

	for _, apiObject := range apiObjects {
		if apiObject == nil {
			continue
		}

		err := UnsuccessfulItemError(apiObject.Error)

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("%s: %w", aws.StringValue(apiObject.ResourceId), err))
		}
	}

	return errors.ErrorOrNil()
}
