// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: trustpolicyexception.proto

package testutil

import (
	"context"
	fmt "fmt"
	"github.com/edgexr/edge-cloud-infra/mc/mcctl/mctestclient"
	"github.com/edgexr/edge-cloud-infra/mc/ormapi"
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

func TestCreateTrustPolicyException(mcClient *mctestclient.Client, uri, token, region string, in *edgeproto.TrustPolicyException, modFuncs ...func(*edgeproto.TrustPolicyException)) (*edgeproto.Result, int, error) {
	dat := &ormapi.RegionTrustPolicyException{}
	dat.Region = region
	dat.TrustPolicyException = *in
	for _, fn := range modFuncs {
		fn(&dat.TrustPolicyException)
	}
	return mcClient.CreateTrustPolicyException(uri, token, dat)
}
func TestPermCreateTrustPolicyException(mcClient *mctestclient.Client, uri, token, region, org string, modFuncs ...func(*edgeproto.TrustPolicyException)) (*edgeproto.Result, int, error) {
	in := &edgeproto.TrustPolicyException{}
	in.Key.AppKey.Organization = org
	return TestCreateTrustPolicyException(mcClient, uri, token, region, in, modFuncs...)
}

func TestUpdateTrustPolicyException(mcClient *mctestclient.Client, uri, token, region string, in *edgeproto.TrustPolicyException, modFuncs ...func(*edgeproto.TrustPolicyException)) (*edgeproto.Result, int, error) {
	dat := &ormapi.RegionTrustPolicyException{}
	dat.Region = region
	dat.TrustPolicyException = *in
	for _, fn := range modFuncs {
		fn(&dat.TrustPolicyException)
	}
	return mcClient.UpdateTrustPolicyException(uri, token, dat)
}
func TestPermUpdateTrustPolicyException(mcClient *mctestclient.Client, uri, token, region, org string, modFuncs ...func(*edgeproto.TrustPolicyException)) (*edgeproto.Result, int, error) {
	in := &edgeproto.TrustPolicyException{}
	in.Key.CloudletPoolKey.Organization = org
	in.Fields = append(in.Fields, edgeproto.TrustPolicyExceptionFieldKeyCloudletPoolKeyOrganization)
	return TestUpdateTrustPolicyException(mcClient, uri, token, region, in, modFuncs...)
}

func TestDeleteTrustPolicyException(mcClient *mctestclient.Client, uri, token, region string, in *edgeproto.TrustPolicyException, modFuncs ...func(*edgeproto.TrustPolicyException)) (*edgeproto.Result, int, error) {
	dat := &ormapi.RegionTrustPolicyException{}
	dat.Region = region
	dat.TrustPolicyException = *in
	for _, fn := range modFuncs {
		fn(&dat.TrustPolicyException)
	}
	return mcClient.DeleteTrustPolicyException(uri, token, dat)
}
func TestPermDeleteTrustPolicyException(mcClient *mctestclient.Client, uri, token, region, org string, modFuncs ...func(*edgeproto.TrustPolicyException)) (*edgeproto.Result, int, error) {
	in := &edgeproto.TrustPolicyException{}
	in.Key.AppKey.Organization = org
	return TestDeleteTrustPolicyException(mcClient, uri, token, region, in, modFuncs...)
}

func TestShowTrustPolicyException(mcClient *mctestclient.Client, uri, token, region string, in *edgeproto.TrustPolicyException, modFuncs ...func(*edgeproto.TrustPolicyException)) ([]edgeproto.TrustPolicyException, int, error) {
	dat := &ormapi.RegionTrustPolicyException{}
	dat.Region = region
	dat.TrustPolicyException = *in
	for _, fn := range modFuncs {
		fn(&dat.TrustPolicyException)
	}
	return mcClient.ShowTrustPolicyException(uri, token, dat)
}
func TestPermShowTrustPolicyException(mcClient *mctestclient.Client, uri, token, region, org string, modFuncs ...func(*edgeproto.TrustPolicyException)) ([]edgeproto.TrustPolicyException, int, error) {
	in := &edgeproto.TrustPolicyException{}
	in.Key.AppKey.Organization = org
	return TestShowTrustPolicyException(mcClient, uri, token, region, in, modFuncs...)
}

func (s *TestClient) CreateTrustPolicyException(ctx context.Context, in *edgeproto.TrustPolicyException) (*edgeproto.Result, error) {
	inR := &ormapi.RegionTrustPolicyException{
		Region:               s.Region,
		TrustPolicyException: *in,
	}
	out, status, err := s.McClient.CreateTrustPolicyException(s.Uri, s.Token, inR)
	if err == nil && status != 200 {
		err = fmt.Errorf("status: %d\n", status)
	}
	return out, err
}

func (s *TestClient) UpdateTrustPolicyException(ctx context.Context, in *edgeproto.TrustPolicyException) (*edgeproto.Result, error) {
	inR := &ormapi.RegionTrustPolicyException{
		Region:               s.Region,
		TrustPolicyException: *in,
	}
	out, status, err := s.McClient.UpdateTrustPolicyException(s.Uri, s.Token, inR)
	if err == nil && status != 200 {
		err = fmt.Errorf("status: %d\n", status)
	}
	return out, err
}

func (s *TestClient) DeleteTrustPolicyException(ctx context.Context, in *edgeproto.TrustPolicyException) (*edgeproto.Result, error) {
	inR := &ormapi.RegionTrustPolicyException{
		Region:               s.Region,
		TrustPolicyException: *in,
	}
	out, status, err := s.McClient.DeleteTrustPolicyException(s.Uri, s.Token, inR)
	if err == nil && status != 200 {
		err = fmt.Errorf("status: %d\n", status)
	}
	return out, err
}

func (s *TestClient) ShowTrustPolicyException(ctx context.Context, in *edgeproto.TrustPolicyException) ([]edgeproto.TrustPolicyException, error) {
	inR := &ormapi.RegionTrustPolicyException{
		Region:               s.Region,
		TrustPolicyException: *in,
	}
	out, status, err := s.McClient.ShowTrustPolicyException(s.Uri, s.Token, inR)
	if err == nil && status != 200 {
		err = fmt.Errorf("status: %d\n", status)
	}
	return out, err
}
