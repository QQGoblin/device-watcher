package grpc

import (
	"context"
	"github.com/jaypipes/ghw"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"k8s.io/klog"
	"net"
	"os"
	"strings"
)

const (
	// IP represents the IP address of the api service
	IP = "0.0.0.0"
	// Port represents the port of the API service
	Port = "9115"
	// DefaultAddress represents the DefaultAddress at which api service can be called on which is 0.0.0.0:9115
	DefaultAddress = IP + ":" + Port
)

func Start() {
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)
	RegisterScannerServer(grpcServer, &Scanner{})

	l, err := net.Listen("tcp", DefaultAddress)
	if err != nil {
		klog.Errorf("Unable to listen %v", err)
		os.Exit(1)
	}

	// Listen for requests
	klog.Infof("Starting server at : %v ", DefaultAddress)
	grpcServer.Serve(l)

}

type Scanner struct {
	UnimplementedScannerServer
}

func (*Scanner) ScanNicDevice(context.Context, *Null) (*NicDevices, error) {

	nicDevs := make([]*NicDevice, 0)

	nics, err := ghw.Network()
	if err != nil {
		klog.Errorf("Error getting network info: %v", err)
		return nil, status.Errorf(codes.Canceled, "Error getting network info ")
	}

	for _, nic := range nics.NICs {

		if nic.IsVirtual || strings.EqualFold(nic.MacAddress, "") {
			continue
		}

		n := NicDevice{
			DeviceName: nic.Name,
			MacAddress: nic.MacAddress,
			Ipaddress:  make([]string, 0),
		}
		if intf, err := net.InterfaceByName(nic.Name); err != nil {
			klog.Errorf("Error getting interface %s msg ", intf.Name)
			continue
		} else {
			addrs, _ := intf.Addrs()
			for _, addr := range addrs {
				n.Ipaddress = append(n.Ipaddress, addr.String())
			}

			if intf.Flags&(1<<uint(0)) != 0 {
				n.Status = NicDevice_UP
			} else {
				n.Status = NicDevice_Down
			}
		}
		nicDevs = append(nicDevs, &n)

	}

	return &NicDevices{Nicdevices: nicDevs}, status.Errorf(codes.OK, "")
}
