// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

// Contains the parameters for DetachVpnGateway.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DetachVpnGatewayRequest
type DetachVpnGatewayInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the VPC.
	//
	// VpcId is a required field
	VpcId *string `type:"string" required:"true"`

	// The ID of the virtual private gateway.
	//
	// VpnGatewayId is a required field
	VpnGatewayId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DetachVpnGatewayInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetachVpnGatewayInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DetachVpnGatewayInput"}

	if s.VpcId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcId"))
	}

	if s.VpnGatewayId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpnGatewayId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DetachVpnGatewayOutput
type DetachVpnGatewayOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DetachVpnGatewayOutput) String() string {
	return awsutil.Prettify(s)
}

const opDetachVpnGateway = "DetachVpnGateway"

// DetachVpnGatewayRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Detaches a virtual private gateway from a VPC. You do this if you're planning
// to turn off the VPC and not use it anymore. You can confirm a virtual private
// gateway has been completely detached from a VPC by describing the virtual
// private gateway (any attachments to the virtual private gateway are also
// described).
//
// You must wait for the attachment's state to switch to detached before you
// can delete the VPC or attach a different VPC to the virtual private gateway.
//
//    // Example sending a request using DetachVpnGatewayRequest.
//    req := client.DetachVpnGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DetachVpnGateway
func (c *Client) DetachVpnGatewayRequest(input *DetachVpnGatewayInput) DetachVpnGatewayRequest {
	op := &aws.Operation{
		Name:       opDetachVpnGateway,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetachVpnGatewayInput{}
	}

	req := c.newRequest(op, input, &DetachVpnGatewayOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DetachVpnGatewayRequest{Request: req, Input: input, Copy: c.DetachVpnGatewayRequest}
}

// DetachVpnGatewayRequest is the request type for the
// DetachVpnGateway API operation.
type DetachVpnGatewayRequest struct {
	*aws.Request
	Input *DetachVpnGatewayInput
	Copy  func(*DetachVpnGatewayInput) DetachVpnGatewayRequest
}

// Send marshals and sends the DetachVpnGateway API request.
func (r DetachVpnGatewayRequest) Send(ctx context.Context) (*DetachVpnGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DetachVpnGatewayResponse{
		DetachVpnGatewayOutput: r.Request.Data.(*DetachVpnGatewayOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DetachVpnGatewayResponse is the response type for the
// DetachVpnGateway API operation.
type DetachVpnGatewayResponse struct {
	*DetachVpnGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DetachVpnGateway request.
func (r *DetachVpnGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
