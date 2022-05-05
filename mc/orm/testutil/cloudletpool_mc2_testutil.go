// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cloudletpool.proto

package testutil

import (
	"context"
	fmt "fmt"
	"github.com/edgexr/edge-cloud-infra/mc/mcctl/mctestclient"
	"github.com/edgexr/edge-cloud-infra/mc/ormapi"
	_ "github.com/edgexr/edge-cloud/d-match-engine/dme-proto"
	edgeproto "github.com/edgexr/edge-cloud/edgeproto"
	_ "github.com/edgexr/edge-cloud/protogen"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Auto-generated code: DO NOT EDIT

func TestCreateCloudletPool(mcClient *mctestclient.Client, uri, token, region string, in *edgeproto.CloudletPool, modFuncs ...func(*edgeproto.CloudletPool)) (*edgeproto.Result, int, error) {
	dat := &ormapi.RegionCloudletPool{}
	dat.Region = region
	dat.CloudletPool = *in
	for _, fn := range modFuncs {
		fn(&dat.CloudletPool)
	}
	return mcClient.CreateCloudletPool(uri, token, dat)
}
func TestPermCreateCloudletPool(mcClient *mctestclient.Client, uri, token, region, org string, modFuncs ...func(*edgeproto.CloudletPool)) (*edgeproto.Result, int, error) {
	in := &edgeproto.CloudletPool{}
	in.Key.Organization = org
	return TestCreateCloudletPool(mcClient, uri, token, region, in, modFuncs...)
}

func TestDeleteCloudletPool(mcClient *mctestclient.Client, uri, token, region string, in *edgeproto.CloudletPool, modFuncs ...func(*edgeproto.CloudletPool)) (*edgeproto.Result, int, error) {
	dat := &ormapi.RegionCloudletPool{}
	dat.Region = region
	dat.CloudletPool = *in
	for _, fn := range modFuncs {
		fn(&dat.CloudletPool)
	}
	return mcClient.DeleteCloudletPool(uri, token, dat)
}
func TestPermDeleteCloudletPool(mcClient *mctestclient.Client, uri, token, region, org string, modFuncs ...func(*edgeproto.CloudletPool)) (*edgeproto.Result, int, error) {
	in := &edgeproto.CloudletPool{}
	in.Key.Organization = org
	return TestDeleteCloudletPool(mcClient, uri, token, region, in, modFuncs...)
}

func TestUpdateCloudletPool(mcClient *mctestclient.Client, uri, token, region string, in *edgeproto.CloudletPool, modFuncs ...func(*edgeproto.CloudletPool)) (*edgeproto.Result, int, error) {
	dat := &ormapi.RegionCloudletPool{}
	dat.Region = region
	dat.CloudletPool = *in
	for _, fn := range modFuncs {
		fn(&dat.CloudletPool)
	}
	return mcClient.UpdateCloudletPool(uri, token, dat)
}
func TestPermUpdateCloudletPool(mcClient *mctestclient.Client, uri, token, region, org string, modFuncs ...func(*edgeproto.CloudletPool)) (*edgeproto.Result, int, error) {
	in := &edgeproto.CloudletPool{}
	in.Key.Organization = org
	in.Fields = append(in.Fields, edgeproto.CloudletPoolFieldKeyOrganization)
	return TestUpdateCloudletPool(mcClient, uri, token, region, in, modFuncs...)
}

func TestShowCloudletPool(mcClient *mctestclient.Client, uri, token, region string, in *edgeproto.CloudletPool, modFuncs ...func(*edgeproto.CloudletPool)) ([]edgeproto.CloudletPool, int, error) {
	dat := &ormapi.RegionCloudletPool{}
	dat.Region = region
	dat.CloudletPool = *in
	for _, fn := range modFuncs {
		fn(&dat.CloudletPool)
	}
	return mcClient.ShowCloudletPool(uri, token, dat)
}
func TestPermShowCloudletPool(mcClient *mctestclient.Client, uri, token, region, org string, modFuncs ...func(*edgeproto.CloudletPool)) ([]edgeproto.CloudletPool, int, error) {
	in := &edgeproto.CloudletPool{}
	in.Key.Organization = org
	return TestShowCloudletPool(mcClient, uri, token, region, in, modFuncs...)
}

func TestAddCloudletPoolMember(mcClient *mctestclient.Client, uri, token, region string, in *edgeproto.CloudletPoolMember, modFuncs ...func(*edgeproto.CloudletPoolMember)) (*edgeproto.Result, int, error) {
	dat := &ormapi.RegionCloudletPoolMember{}
	dat.Region = region
	dat.CloudletPoolMember = *in
	for _, fn := range modFuncs {
		fn(&dat.CloudletPoolMember)
	}
	return mcClient.AddCloudletPoolMember(uri, token, dat)
}
func TestPermAddCloudletPoolMember(mcClient *mctestclient.Client, uri, token, region, org string, modFuncs ...func(*edgeproto.CloudletPoolMember)) (*edgeproto.Result, int, error) {
	in := &edgeproto.CloudletPoolMember{}
	in.Key.Organization = org
	return TestAddCloudletPoolMember(mcClient, uri, token, region, in, modFuncs...)
}

func TestRemoveCloudletPoolMember(mcClient *mctestclient.Client, uri, token, region string, in *edgeproto.CloudletPoolMember, modFuncs ...func(*edgeproto.CloudletPoolMember)) (*edgeproto.Result, int, error) {
	dat := &ormapi.RegionCloudletPoolMember{}
	dat.Region = region
	dat.CloudletPoolMember = *in
	for _, fn := range modFuncs {
		fn(&dat.CloudletPoolMember)
	}
	return mcClient.RemoveCloudletPoolMember(uri, token, dat)
}
func TestPermRemoveCloudletPoolMember(mcClient *mctestclient.Client, uri, token, region, org string, modFuncs ...func(*edgeproto.CloudletPoolMember)) (*edgeproto.Result, int, error) {
	in := &edgeproto.CloudletPoolMember{}
	in.Key.Organization = org
	return TestRemoveCloudletPoolMember(mcClient, uri, token, region, in, modFuncs...)
}

func (s *TestClient) CreateCloudletPool(ctx context.Context, in *edgeproto.CloudletPool) (*edgeproto.Result, error) {
	inR := &ormapi.RegionCloudletPool{
		Region:       s.Region,
		CloudletPool: *in,
	}
	out, status, err := s.McClient.CreateCloudletPool(s.Uri, s.Token, inR)
	if err == nil && status != 200 {
		err = fmt.Errorf("status: %d\n", status)
	}
	return out, err
}

func (s *TestClient) DeleteCloudletPool(ctx context.Context, in *edgeproto.CloudletPool) (*edgeproto.Result, error) {
	inR := &ormapi.RegionCloudletPool{
		Region:       s.Region,
		CloudletPool: *in,
	}
	out, status, err := s.McClient.DeleteCloudletPool(s.Uri, s.Token, inR)
	if err == nil && status != 200 {
		err = fmt.Errorf("status: %d\n", status)
	}
	return out, err
}

func (s *TestClient) UpdateCloudletPool(ctx context.Context, in *edgeproto.CloudletPool) (*edgeproto.Result, error) {
	inR := &ormapi.RegionCloudletPool{
		Region:       s.Region,
		CloudletPool: *in,
	}
	out, status, err := s.McClient.UpdateCloudletPool(s.Uri, s.Token, inR)
	if err == nil && status != 200 {
		err = fmt.Errorf("status: %d\n", status)
	}
	return out, err
}

func (s *TestClient) ShowCloudletPool(ctx context.Context, in *edgeproto.CloudletPool) ([]edgeproto.CloudletPool, error) {
	inR := &ormapi.RegionCloudletPool{
		Region:       s.Region,
		CloudletPool: *in,
	}
	out, status, err := s.McClient.ShowCloudletPool(s.Uri, s.Token, inR)
	if err == nil && status != 200 {
		err = fmt.Errorf("status: %d\n", status)
	}
	return out, err
}

func (s *TestClient) AddCloudletPoolMember(ctx context.Context, in *edgeproto.CloudletPoolMember) (*edgeproto.Result, error) {
	inR := &ormapi.RegionCloudletPoolMember{
		Region:             s.Region,
		CloudletPoolMember: *in,
	}
	out, status, err := s.McClient.AddCloudletPoolMember(s.Uri, s.Token, inR)
	if err == nil && status != 200 {
		err = fmt.Errorf("status: %d\n", status)
	}
	return out, err
}

func (s *TestClient) RemoveCloudletPoolMember(ctx context.Context, in *edgeproto.CloudletPoolMember) (*edgeproto.Result, error) {
	inR := &ormapi.RegionCloudletPoolMember{
		Region:             s.Region,
		CloudletPoolMember: *in,
	}
	out, status, err := s.McClient.RemoveCloudletPoolMember(s.Uri, s.Token, inR)
	if err == nil && status != 200 {
		err = fmt.Errorf("status: %d\n", status)
	}
	return out, err
}
