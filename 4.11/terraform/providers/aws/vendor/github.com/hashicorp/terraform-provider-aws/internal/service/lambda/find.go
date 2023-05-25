package lambda

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// FindEventSourceMappingConfigurationByID returns the event source mapping corresponding to the specified ID.
// Returns NotFoundError if no event source mapping is found.
func FindEventSourceMappingConfigurationByID(conn *lambda.Lambda, uuid string) (*lambda.EventSourceMappingConfiguration, error) {
	input := &lambda.GetEventSourceMappingInput{
		UUID: aws.String(uuid),
	}

	return FindEventSourceMappingConfiguration(conn, input)
}

// FindEventSourceMappingConfiguration returns the event source mapping corresponding to the specified input.
// Returns NotFoundError if no event source mapping is found.
func FindEventSourceMappingConfiguration(conn *lambda.Lambda, input *lambda.GetEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error) {
	output, err := conn.GetEventSourceMapping(input)

	if tfawserr.ErrCodeEquals(err, lambda.ErrCodeResourceNotFoundException) {
		return nil, &resource.NotFoundError{
			LastError:   err,
			LastRequest: input,
		}
	}

	if err != nil {
		return nil, err
	}

	// Handle any empty result.
	if output == nil {
		return nil, &resource.NotFoundError{
			Message:     "Empty result",
			LastRequest: input,
		}
	}

	return output, nil
}
