package saga

import "google.golang.org/grpc"

type Controller struct {
	connLicense *grpc.ClientConn
	connOrg     *grpc.ClientConn
	connProfile *grpc.ClientConn
}

func NewController(
	connLicense *grpc.ClientConn,
	connOrg *grpc.ClientConn,
	connProfile *grpc.ClientConn) *Controller {
	return &Controller{
		connLicense: connLicense,
		connOrg:     connOrg,
		connProfile: connProfile,
	}
}
