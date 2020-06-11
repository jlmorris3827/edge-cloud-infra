// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: autoprovpolicy.proto

package ormctl

import edgeproto "github.com/mobiledgex/edge-cloud/edgeproto"
import "strings"
import "github.com/mobiledgex/edge-cloud-infra/mc/ormapi"
import "github.com/mobiledgex/edge-cloud/cli"
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/googleapis/google/api"
import _ "github.com/mobiledgex/edge-cloud/protogen"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/mobiledgex/edge-cloud/d-match-engine/dme-proto"
import _ "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Auto-generated code: DO NOT EDIT

var CreateAutoProvPolicyCmd = &cli.Command{
	Use:          "CreateAutoProvPolicy",
	RequiredArgs: "region " + strings.Join(CreateAutoProvPolicyRequiredArgs, " "),
	OptionalArgs: strings.Join(CreateAutoProvPolicyOptionalArgs, " "),
	AliasArgs:    strings.Join(AutoProvPolicyAliasArgs, " "),
	SpecialArgs:  &AutoProvPolicySpecialArgs,
	Comments:     addRegionComment(AutoProvPolicyComments),
	ReqData:      &ormapi.RegionAutoProvPolicy{},
	ReplyData:    &edgeproto.Result{},
	Run:          runRest("/auth/ctrl/CreateAutoProvPolicy"),
}

var DeleteAutoProvPolicyCmd = &cli.Command{
	Use:          "DeleteAutoProvPolicy",
	RequiredArgs: "region " + strings.Join(AutoProvPolicyRequiredArgs, " "),
	OptionalArgs: strings.Join(AutoProvPolicyOptionalArgs, " "),
	AliasArgs:    strings.Join(AutoProvPolicyAliasArgs, " "),
	SpecialArgs:  &AutoProvPolicySpecialArgs,
	Comments:     addRegionComment(AutoProvPolicyComments),
	ReqData:      &ormapi.RegionAutoProvPolicy{},
	ReplyData:    &edgeproto.Result{},
	Run:          runRest("/auth/ctrl/DeleteAutoProvPolicy"),
}

var UpdateAutoProvPolicyCmd = &cli.Command{
	Use:          "UpdateAutoProvPolicy",
	RequiredArgs: "region " + strings.Join(AutoProvPolicyRequiredArgs, " "),
	OptionalArgs: strings.Join(AutoProvPolicyOptionalArgs, " "),
	AliasArgs:    strings.Join(AutoProvPolicyAliasArgs, " "),
	SpecialArgs:  &AutoProvPolicySpecialArgs,
	Comments:     addRegionComment(AutoProvPolicyComments),
	ReqData:      &ormapi.RegionAutoProvPolicy{},
	ReplyData:    &edgeproto.Result{},
	Run: runRest("/auth/ctrl/UpdateAutoProvPolicy",
		withSetFieldsFunc(setUpdateAutoProvPolicyFields),
	),
}

func setUpdateAutoProvPolicyFields(in map[string]interface{}) {
	// get map for edgeproto object in region struct
	obj := in[strings.ToLower("AutoProvPolicy")]
	if obj == nil {
		return
	}
	objmap, ok := obj.(map[string]interface{})
	if !ok {
		return
	}
	objmap["fields"] = cli.GetSpecifiedFields(objmap, &edgeproto.AutoProvPolicy{}, cli.JsonNamespace)
}

var ShowAutoProvPolicyCmd = &cli.Command{
	Use:          "ShowAutoProvPolicy",
	RequiredArgs: "region",
	OptionalArgs: strings.Join(append(AutoProvPolicyRequiredArgs, AutoProvPolicyOptionalArgs...), " "),
	AliasArgs:    strings.Join(AutoProvPolicyAliasArgs, " "),
	SpecialArgs:  &AutoProvPolicySpecialArgs,
	Comments:     addRegionComment(AutoProvPolicyComments),
	ReqData:      &ormapi.RegionAutoProvPolicy{},
	ReplyData:    &edgeproto.AutoProvPolicy{},
	Run:          runRest("/auth/ctrl/ShowAutoProvPolicy"),
	StreamOut:    true,
}

var AddAutoProvPolicyCloudletCmd = &cli.Command{
	Use:          "AddAutoProvPolicyCloudlet",
	RequiredArgs: "region " + strings.Join(AutoProvPolicyCloudletRequiredArgs, " "),
	OptionalArgs: strings.Join(AutoProvPolicyCloudletOptionalArgs, " "),
	AliasArgs:    strings.Join(AutoProvPolicyCloudletAliasArgs, " "),
	SpecialArgs:  &AutoProvPolicyCloudletSpecialArgs,
	Comments:     addRegionComment(AutoProvPolicyCloudletComments),
	ReqData:      &ormapi.RegionAutoProvPolicyCloudlet{},
	ReplyData:    &edgeproto.Result{},
	Run:          runRest("/auth/ctrl/AddAutoProvPolicyCloudlet"),
}

var RemoveAutoProvPolicyCloudletCmd = &cli.Command{
	Use:          "RemoveAutoProvPolicyCloudlet",
	RequiredArgs: "region " + strings.Join(AutoProvPolicyCloudletRequiredArgs, " "),
	OptionalArgs: strings.Join(AutoProvPolicyCloudletOptionalArgs, " "),
	AliasArgs:    strings.Join(AutoProvPolicyCloudletAliasArgs, " "),
	SpecialArgs:  &AutoProvPolicyCloudletSpecialArgs,
	Comments:     addRegionComment(AutoProvPolicyCloudletComments),
	ReqData:      &ormapi.RegionAutoProvPolicyCloudlet{},
	ReplyData:    &edgeproto.Result{},
	Run:          runRest("/auth/ctrl/RemoveAutoProvPolicyCloudlet"),
}

var AutoProvPolicyApiCmds = []*cli.Command{
	CreateAutoProvPolicyCmd,
	DeleteAutoProvPolicyCmd,
	UpdateAutoProvPolicyCmd,
	ShowAutoProvPolicyCmd,
	AddAutoProvPolicyCloudletCmd,
	RemoveAutoProvPolicyCloudletCmd,
}

var CreateAutoProvPolicyRequiredArgs = []string{
	"app-org",
	"name",
}
var CreateAutoProvPolicyOptionalArgs = []string{
	"deployclientcount",
	"deployintervalcount",
	"cloudlets:#.key.organization",
	"cloudlets:#.key.name",
	"cloudlets:#.loc.latitude",
	"cloudlets:#.loc.longitude",
	"cloudlets:#.loc.horizontalaccuracy",
	"cloudlets:#.loc.verticalaccuracy",
	"cloudlets:#.loc.altitude",
	"cloudlets:#.loc.course",
	"cloudlets:#.loc.speed",
	"cloudlets:#.loc.timestamp.seconds",
	"cloudlets:#.loc.timestamp.nanos",
	"minactiveinstances",
	"maxinstances",
}
var AutoProvPolicyRequiredArgs = []string{
	"app-org",
	"name",
}
var AutoProvPolicyOptionalArgs = []string{
	"deployclientcount",
	"deployintervalcount",
	"cloudlets:#.key.organization",
	"cloudlets:#.key.name",
	"cloudlets:#.loc.latitude",
	"cloudlets:#.loc.longitude",
	"cloudlets:#.loc.horizontalaccuracy",
	"cloudlets:#.loc.verticalaccuracy",
	"cloudlets:#.loc.altitude",
	"cloudlets:#.loc.course",
	"cloudlets:#.loc.speed",
	"cloudlets:#.loc.timestamp.seconds",
	"cloudlets:#.loc.timestamp.nanos",
	"minactiveinstances",
	"maxinstances",
}
var AutoProvPolicyAliasArgs = []string{
	"fields=autoprovpolicy.fields",
	"app-org=autoprovpolicy.key.organization",
	"name=autoprovpolicy.key.name",
	"deployclientcount=autoprovpolicy.deployclientcount",
	"deployintervalcount=autoprovpolicy.deployintervalcount",
	"cloudlets:#.key.organization=autoprovpolicy.cloudlets:#.key.organization",
	"cloudlets:#.key.name=autoprovpolicy.cloudlets:#.key.name",
	"cloudlets:#.loc.latitude=autoprovpolicy.cloudlets:#.loc.latitude",
	"cloudlets:#.loc.longitude=autoprovpolicy.cloudlets:#.loc.longitude",
	"cloudlets:#.loc.horizontalaccuracy=autoprovpolicy.cloudlets:#.loc.horizontalaccuracy",
	"cloudlets:#.loc.verticalaccuracy=autoprovpolicy.cloudlets:#.loc.verticalaccuracy",
	"cloudlets:#.loc.altitude=autoprovpolicy.cloudlets:#.loc.altitude",
	"cloudlets:#.loc.course=autoprovpolicy.cloudlets:#.loc.course",
	"cloudlets:#.loc.speed=autoprovpolicy.cloudlets:#.loc.speed",
	"cloudlets:#.loc.timestamp.seconds=autoprovpolicy.cloudlets:#.loc.timestamp.seconds",
	"cloudlets:#.loc.timestamp.nanos=autoprovpolicy.cloudlets:#.loc.timestamp.nanos",
	"minactiveinstances=autoprovpolicy.minactiveinstances",
	"maxinstances=autoprovpolicy.maxinstances",
}
var AutoProvPolicyComments = map[string]string{
	"fields":                             "Fields are used for the Update API to specify which fields to apply",
	"app-org":                            "Name of the organization for the cluster that this policy will apply to",
	"name":                               "Policy name",
	"deployclientcount":                  "Minimum number of clients within the auto deploy interval to trigger deployment",
	"deployintervalcount":                "Number of intervals to check before triggering deployment",
	"cloudlets:#.key.organization":       "Organization of the cloudlet site",
	"cloudlets:#.key.name":               "Name of the cloudlet",
	"cloudlets:#.loc.latitude":           "latitude in WGS 84 coordinates",
	"cloudlets:#.loc.longitude":          "longitude in WGS 84 coordinates",
	"cloudlets:#.loc.horizontalaccuracy": "horizontal accuracy (radius in meters)",
	"cloudlets:#.loc.verticalaccuracy":   "vertical accuracy (meters)",
	"cloudlets:#.loc.altitude":           "On android only lat and long are guaranteed to be supplied altitude in meters",
	"cloudlets:#.loc.course":             "course (IOS) / bearing (Android) (degrees east relative to true north)",
	"cloudlets:#.loc.speed":              "speed (IOS) / velocity (Android) (meters/sec)",
	"minactiveinstances":                 "Minimum number of active instances for High-Availability",
	"maxinstances":                       "Maximum number of instances (active or not)",
}
var AutoProvPolicySpecialArgs = map[string]string{
	"autoprovpolicy.fields": "StringArray",
}
var AutoProvCloudletRequiredArgs = []string{
	"key.organization",
	"key.name",
}
var AutoProvCloudletOptionalArgs = []string{
	"loc.latitude",
	"loc.longitude",
	"loc.horizontalaccuracy",
	"loc.verticalaccuracy",
	"loc.altitude",
	"loc.course",
	"loc.speed",
	"loc.timestamp.seconds",
	"loc.timestamp.nanos",
}
var AutoProvCloudletAliasArgs = []string{
	"key.organization=autoprovcloudlet.key.organization",
	"key.name=autoprovcloudlet.key.name",
	"loc.latitude=autoprovcloudlet.loc.latitude",
	"loc.longitude=autoprovcloudlet.loc.longitude",
	"loc.horizontalaccuracy=autoprovcloudlet.loc.horizontalaccuracy",
	"loc.verticalaccuracy=autoprovcloudlet.loc.verticalaccuracy",
	"loc.altitude=autoprovcloudlet.loc.altitude",
	"loc.course=autoprovcloudlet.loc.course",
	"loc.speed=autoprovcloudlet.loc.speed",
	"loc.timestamp.seconds=autoprovcloudlet.loc.timestamp.seconds",
	"loc.timestamp.nanos=autoprovcloudlet.loc.timestamp.nanos",
}
var AutoProvCloudletComments = map[string]string{
	"key.organization":       "Organization of the cloudlet site",
	"key.name":               "Name of the cloudlet",
	"loc.latitude":           "latitude in WGS 84 coordinates",
	"loc.longitude":          "longitude in WGS 84 coordinates",
	"loc.horizontalaccuracy": "horizontal accuracy (radius in meters)",
	"loc.verticalaccuracy":   "vertical accuracy (meters)",
	"loc.altitude":           "On android only lat and long are guaranteed to be supplied altitude in meters",
	"loc.course":             "course (IOS) / bearing (Android) (degrees east relative to true north)",
	"loc.speed":              "speed (IOS) / velocity (Android) (meters/sec)",
}
var AutoProvCloudletSpecialArgs = map[string]string{}
var AutoProvCountRequiredArgs = []string{}
var AutoProvCountOptionalArgs = []string{
	"appkey.organization",
	"appkey.name",
	"appkey.version",
	"cloudletkey.organization",
	"cloudletkey.name",
	"count",
	"processnow",
	"deploynowkey.clusterkey.name",
	"deploynowkey.cloudletkey.organization",
	"deploynowkey.cloudletkey.name",
	"deploynowkey.organization",
}
var AutoProvCountAliasArgs = []string{
	"appkey.organization=autoprovcount.appkey.organization",
	"appkey.name=autoprovcount.appkey.name",
	"appkey.version=autoprovcount.appkey.version",
	"cloudletkey.organization=autoprovcount.cloudletkey.organization",
	"cloudletkey.name=autoprovcount.cloudletkey.name",
	"count=autoprovcount.count",
	"processnow=autoprovcount.processnow",
	"deploynowkey.clusterkey.name=autoprovcount.deploynowkey.clusterkey.name",
	"deploynowkey.cloudletkey.organization=autoprovcount.deploynowkey.cloudletkey.organization",
	"deploynowkey.cloudletkey.name=autoprovcount.deploynowkey.cloudletkey.name",
	"deploynowkey.organization=autoprovcount.deploynowkey.organization",
}
var AutoProvCountComments = map[string]string{
	"appkey.organization":                   "App developer organization",
	"appkey.name":                           "App name",
	"appkey.version":                        "App version",
	"cloudletkey.organization":              "Organization of the cloudlet site",
	"cloudletkey.name":                      "Name of the cloudlet",
	"count":                                 "FindCloudlet client count",
	"processnow":                            "Process count immediately",
	"deploynowkey.clusterkey.name":          "Cluster name",
	"deploynowkey.cloudletkey.organization": "Organization of the cloudlet site",
	"deploynowkey.cloudletkey.name":         "Name of the cloudlet",
	"deploynowkey.organization":             "Name of Developer organization that this cluster belongs to",
}
var AutoProvCountSpecialArgs = map[string]string{}
var AutoProvCountsRequiredArgs = []string{}
var AutoProvCountsOptionalArgs = []string{
	"dmenodename",
	"timestamp.seconds",
	"timestamp.nanos",
	"counts:#.appkey.organization",
	"counts:#.appkey.name",
	"counts:#.appkey.version",
	"counts:#.cloudletkey.organization",
	"counts:#.cloudletkey.name",
	"counts:#.count",
	"counts:#.processnow",
	"counts:#.deploynowkey.clusterkey.name",
	"counts:#.deploynowkey.cloudletkey.organization",
	"counts:#.deploynowkey.cloudletkey.name",
	"counts:#.deploynowkey.organization",
}
var AutoProvCountsAliasArgs = []string{
	"dmenodename=autoprovcounts.dmenodename",
	"timestamp.seconds=autoprovcounts.timestamp.seconds",
	"timestamp.nanos=autoprovcounts.timestamp.nanos",
	"counts:#.appkey.organization=autoprovcounts.counts:#.appkey.organization",
	"counts:#.appkey.name=autoprovcounts.counts:#.appkey.name",
	"counts:#.appkey.version=autoprovcounts.counts:#.appkey.version",
	"counts:#.cloudletkey.organization=autoprovcounts.counts:#.cloudletkey.organization",
	"counts:#.cloudletkey.name=autoprovcounts.counts:#.cloudletkey.name",
	"counts:#.count=autoprovcounts.counts:#.count",
	"counts:#.processnow=autoprovcounts.counts:#.processnow",
	"counts:#.deploynowkey.clusterkey.name=autoprovcounts.counts:#.deploynowkey.clusterkey.name",
	"counts:#.deploynowkey.cloudletkey.organization=autoprovcounts.counts:#.deploynowkey.cloudletkey.organization",
	"counts:#.deploynowkey.cloudletkey.name=autoprovcounts.counts:#.deploynowkey.cloudletkey.name",
	"counts:#.deploynowkey.organization=autoprovcounts.counts:#.deploynowkey.organization",
}
var AutoProvCountsComments = map[string]string{
	"dmenodename":                                    "DME node name",
	"timestamp.seconds":                              "Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive.",
	"timestamp.nanos":                                "Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive.",
	"counts:#.appkey.organization":                   "App developer organization",
	"counts:#.appkey.name":                           "App name",
	"counts:#.appkey.version":                        "App version",
	"counts:#.cloudletkey.organization":              "Organization of the cloudlet site",
	"counts:#.cloudletkey.name":                      "Name of the cloudlet",
	"counts:#.count":                                 "FindCloudlet client count",
	"counts:#.processnow":                            "Process count immediately",
	"counts:#.deploynowkey.clusterkey.name":          "Cluster name",
	"counts:#.deploynowkey.cloudletkey.organization": "Organization of the cloudlet site",
	"counts:#.deploynowkey.cloudletkey.name":         "Name of the cloudlet",
	"counts:#.deploynowkey.organization":             "Name of Developer organization that this cluster belongs to",
}
var AutoProvCountsSpecialArgs = map[string]string{}
var AutoProvPolicyCloudletRequiredArgs = []string{
	"app-org",
	"name",
}
var AutoProvPolicyCloudletOptionalArgs = []string{
	"cloudlet-org",
	"cloudlet",
}
var AutoProvPolicyCloudletAliasArgs = []string{
	"app-org=autoprovpolicycloudlet.key.organization",
	"name=autoprovpolicycloudlet.key.name",
	"cloudlet-org=autoprovpolicycloudlet.cloudletkey.organization",
	"cloudlet=autoprovpolicycloudlet.cloudletkey.name",
}
var AutoProvPolicyCloudletComments = map[string]string{
	"app-org":      "Name of the organization for the cluster that this policy will apply to",
	"name":         "Policy name",
	"cloudlet-org": "Organization of the cloudlet site",
	"cloudlet":     "Name of the cloudlet",
}
var AutoProvPolicyCloudletSpecialArgs = map[string]string{}
