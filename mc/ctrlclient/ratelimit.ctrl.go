// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ratelimit.proto

package ctrlclient

import (
	"context"
	fmt "fmt"
	"github.com/edgexr/edge-cloud-infra/mc/ormutil"
	edgeproto "github.com/edgexr/edge-cloud/edgeproto"
	"github.com/edgexr/edge-cloud/log"
	_ "github.com/edgexr/edge-cloud/protogen"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	"io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Auto-generated code: DO NOT EDIT

func ShowRateLimitSettingsStream(ctx context.Context, rc *ormutil.RegionContext, obj *edgeproto.RateLimitSettings, connObj ClientConnMgr, authz authzShow, cb func(res *edgeproto.RateLimitSettings) error) error {
	conn, err := connObj.GetRegionConn(ctx, rc.Region)
	if err != nil {
		return err
	}
	api := edgeproto.NewRateLimitSettingsApiClient(conn)
	log.SpanLog(ctx, log.DebugLevelApi, "start controller api")
	defer log.SpanLog(ctx, log.DebugLevelApi, "finish controller api")
	stream, err := api.ShowRateLimitSettings(ctx, obj)
	if err != nil {
		return err
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			return err
		}
		if !rc.SkipAuthz {
			if authz != nil {
				if !authz.Ok("") {
					continue
				}
			}
		}
		err = cb(res)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateFlowRateLimitSettingsObj(ctx context.Context, rc *ormutil.RegionContext, obj *edgeproto.FlowRateLimitSettings, connObj ClientConnMgr) (*edgeproto.Result, error) {
	conn, err := connObj.GetRegionConn(ctx, rc.Region)
	if err != nil {
		return nil, err
	}
	api := edgeproto.NewRateLimitSettingsApiClient(conn)
	log.SpanLog(ctx, log.DebugLevelApi, "start controller api")
	defer log.SpanLog(ctx, log.DebugLevelApi, "finish controller api")
	return api.CreateFlowRateLimitSettings(ctx, obj)
}

func UpdateFlowRateLimitSettingsObj(ctx context.Context, rc *ormutil.RegionContext, obj *edgeproto.FlowRateLimitSettings, connObj ClientConnMgr) (*edgeproto.Result, error) {
	conn, err := connObj.GetRegionConn(ctx, rc.Region)
	if err != nil {
		return nil, err
	}
	api := edgeproto.NewRateLimitSettingsApiClient(conn)
	log.SpanLog(ctx, log.DebugLevelApi, "start controller api")
	defer log.SpanLog(ctx, log.DebugLevelApi, "finish controller api")
	return api.UpdateFlowRateLimitSettings(ctx, obj)
}

func DeleteFlowRateLimitSettingsObj(ctx context.Context, rc *ormutil.RegionContext, obj *edgeproto.FlowRateLimitSettings, connObj ClientConnMgr) (*edgeproto.Result, error) {
	conn, err := connObj.GetRegionConn(ctx, rc.Region)
	if err != nil {
		return nil, err
	}
	api := edgeproto.NewRateLimitSettingsApiClient(conn)
	log.SpanLog(ctx, log.DebugLevelApi, "start controller api")
	defer log.SpanLog(ctx, log.DebugLevelApi, "finish controller api")
	return api.DeleteFlowRateLimitSettings(ctx, obj)
}

func ShowFlowRateLimitSettingsStream(ctx context.Context, rc *ormutil.RegionContext, obj *edgeproto.FlowRateLimitSettings, connObj ClientConnMgr, authz authzShow, cb func(res *edgeproto.FlowRateLimitSettings) error) error {
	conn, err := connObj.GetRegionConn(ctx, rc.Region)
	if err != nil {
		return err
	}
	api := edgeproto.NewRateLimitSettingsApiClient(conn)
	log.SpanLog(ctx, log.DebugLevelApi, "start controller api")
	defer log.SpanLog(ctx, log.DebugLevelApi, "finish controller api")
	stream, err := api.ShowFlowRateLimitSettings(ctx, obj)
	if err != nil {
		return err
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			return err
		}
		if !rc.SkipAuthz {
			if authz != nil {
				if !authz.Ok("") {
					continue
				}
			}
		}
		err = cb(res)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateMaxReqsRateLimitSettingsObj(ctx context.Context, rc *ormutil.RegionContext, obj *edgeproto.MaxReqsRateLimitSettings, connObj ClientConnMgr) (*edgeproto.Result, error) {
	conn, err := connObj.GetRegionConn(ctx, rc.Region)
	if err != nil {
		return nil, err
	}
	api := edgeproto.NewRateLimitSettingsApiClient(conn)
	log.SpanLog(ctx, log.DebugLevelApi, "start controller api")
	defer log.SpanLog(ctx, log.DebugLevelApi, "finish controller api")
	return api.CreateMaxReqsRateLimitSettings(ctx, obj)
}

func UpdateMaxReqsRateLimitSettingsObj(ctx context.Context, rc *ormutil.RegionContext, obj *edgeproto.MaxReqsRateLimitSettings, connObj ClientConnMgr) (*edgeproto.Result, error) {
	conn, err := connObj.GetRegionConn(ctx, rc.Region)
	if err != nil {
		return nil, err
	}
	api := edgeproto.NewRateLimitSettingsApiClient(conn)
	log.SpanLog(ctx, log.DebugLevelApi, "start controller api")
	defer log.SpanLog(ctx, log.DebugLevelApi, "finish controller api")
	return api.UpdateMaxReqsRateLimitSettings(ctx, obj)
}

func DeleteMaxReqsRateLimitSettingsObj(ctx context.Context, rc *ormutil.RegionContext, obj *edgeproto.MaxReqsRateLimitSettings, connObj ClientConnMgr) (*edgeproto.Result, error) {
	conn, err := connObj.GetRegionConn(ctx, rc.Region)
	if err != nil {
		return nil, err
	}
	api := edgeproto.NewRateLimitSettingsApiClient(conn)
	log.SpanLog(ctx, log.DebugLevelApi, "start controller api")
	defer log.SpanLog(ctx, log.DebugLevelApi, "finish controller api")
	return api.DeleteMaxReqsRateLimitSettings(ctx, obj)
}

func ShowMaxReqsRateLimitSettingsStream(ctx context.Context, rc *ormutil.RegionContext, obj *edgeproto.MaxReqsRateLimitSettings, connObj ClientConnMgr, authz authzShow, cb func(res *edgeproto.MaxReqsRateLimitSettings) error) error {
	conn, err := connObj.GetRegionConn(ctx, rc.Region)
	if err != nil {
		return err
	}
	api := edgeproto.NewRateLimitSettingsApiClient(conn)
	log.SpanLog(ctx, log.DebugLevelApi, "start controller api")
	defer log.SpanLog(ctx, log.DebugLevelApi, "finish controller api")
	stream, err := api.ShowMaxReqsRateLimitSettings(ctx, obj)
	if err != nil {
		return err
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			return err
		}
		if !rc.SkipAuthz {
			if authz != nil {
				if !authz.Ok("") {
					continue
				}
			}
		}
		err = cb(res)
		if err != nil {
			return err
		}
	}
	return nil
}
