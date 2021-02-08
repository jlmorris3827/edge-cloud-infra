// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: node.proto

package ormctl

import (
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	"github.com/mobiledgex/edge-cloud-infra/mc/ormapi"
	"github.com/mobiledgex/edge-cloud/cli"
	edgeproto "github.com/mobiledgex/edge-cloud/edgeproto"
	_ "github.com/mobiledgex/edge-cloud/protogen"
	math "math"
	"strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Auto-generated code: DO NOT EDIT

var ShowNodeCmd = &cli.Command{
	Use:          "ShowNode",
	OptionalArgs: strings.Join(append(NodeRequiredArgs, NodeOptionalArgs...), " "),
	AliasArgs:    strings.Join(NodeAliasArgs, " "),
	SpecialArgs:  &NodeSpecialArgs,
	Comments:     addRegionComment(NodeComments),
	ReqData:      &ormapi.RegionNode{},
	ReplyData:    &edgeproto.Node{},
	Run:          runRest("/auth/ctrl/ShowNode"),
	StreamOut:    true,
}

var NodeApiCmds = []*cli.Command{
	ShowNodeCmd,
}

var NodeKeyRequiredArgs = []string{}
var NodeKeyOptionalArgs = []string{
	"name",
	"type",
	"cloudletkey.organization",
	"cloudletkey.name",
	"region",
}
var NodeKeyAliasArgs = []string{
	"name=nodekey.name",
	"type=nodekey.type",
	"cloudletkey.organization=nodekey.cloudletkey.organization",
	"cloudletkey.name=nodekey.cloudletkey.name",
	"region=nodekey.region",
}
var NodeKeyComments = map[string]string{
	"name":                     "Name or hostname of node",
	"type":                     "Node type",
	"cloudletkey.organization": "Organization of the cloudlet site",
	"cloudletkey.name":         "Name of the cloudlet",
	"region":                   "Region the node is in",
}
var NodeKeySpecialArgs = map[string]string{}
var NodeRequiredArgs = []string{
	"name",
	"type",
	"cloudlet-org",
	"cloudlet",
	"region",
}
var NodeOptionalArgs = []string{
	"notifyid",
	"buildmaster",
	"buildhead",
	"buildauthor",
	"builddate",
	"hostname",
	"containerversion",
	"internalpki",
	"properties",
}
var NodeAliasArgs = []string{
	"fields=node.fields",
	"name=node.key.name",
	"type=node.key.type",
	"cloudlet-org=node.key.cloudletkey.organization",
	"cloudlet=node.key.cloudletkey.name",
	"region=node.key.region",
	"notifyid=node.notifyid",
	"buildmaster=node.buildmaster",
	"buildhead=node.buildhead",
	"buildauthor=node.buildauthor",
	"builddate=node.builddate",
	"hostname=node.hostname",
	"containerversion=node.containerversion",
	"internalpki=node.internalpki",
	"properties=node.properties",
}
var NodeComments = map[string]string{
	"fields":           "Fields are used for the Update API to specify which fields to apply",
	"name":             "Name or hostname of node",
	"type":             "Node type",
	"cloudlet-org":     "Organization of the cloudlet site",
	"cloudlet":         "Name of the cloudlet",
	"region":           "Region the node is in",
	"notifyid":         "Id of client assigned by server (internal use only)",
	"buildmaster":      "Build Master Version",
	"buildhead":        "Build Head Version",
	"buildauthor":      "Build Author",
	"builddate":        "Build Date",
	"hostname":         "Hostname",
	"containerversion": "Docker edge-cloud container version which node instance use",
	"internalpki":      "Internal PKI Config",
	"properties":       "Additional properties",
}
var NodeSpecialArgs = map[string]string{
	"node.fields":     "StringArray",
	"node.properties": "StringToString",
}
var NodeDataRequiredArgs = []string{}
var NodeDataOptionalArgs = []string{
	"nodes:#.fields",
	"nodes:#.key.name",
	"nodes:#.key.type",
	"nodes:#.key.cloudletkey.organization",
	"nodes:#.key.cloudletkey.name",
	"nodes:#.key.region",
	"nodes:#.notifyid",
	"nodes:#.buildmaster",
	"nodes:#.buildhead",
	"nodes:#.buildauthor",
	"nodes:#.builddate",
	"nodes:#.hostname",
	"nodes:#.containerversion",
	"nodes:#.internalpki",
	"nodes:#.properties",
}
var NodeDataAliasArgs = []string{
	"nodes:#.fields=nodedata.nodes:#.fields",
	"nodes:#.key.name=nodedata.nodes:#.key.name",
	"nodes:#.key.type=nodedata.nodes:#.key.type",
	"nodes:#.key.cloudletkey.organization=nodedata.nodes:#.key.cloudletkey.organization",
	"nodes:#.key.cloudletkey.name=nodedata.nodes:#.key.cloudletkey.name",
	"nodes:#.key.region=nodedata.nodes:#.key.region",
	"nodes:#.notifyid=nodedata.nodes:#.notifyid",
	"nodes:#.buildmaster=nodedata.nodes:#.buildmaster",
	"nodes:#.buildhead=nodedata.nodes:#.buildhead",
	"nodes:#.buildauthor=nodedata.nodes:#.buildauthor",
	"nodes:#.builddate=nodedata.nodes:#.builddate",
	"nodes:#.hostname=nodedata.nodes:#.hostname",
	"nodes:#.containerversion=nodedata.nodes:#.containerversion",
	"nodes:#.internalpki=nodedata.nodes:#.internalpki",
	"nodes:#.properties=nodedata.nodes:#.properties",
}
var NodeDataComments = map[string]string{
	"nodes:#.fields":                       "Fields are used for the Update API to specify which fields to apply",
	"nodes:#.key.name":                     "Name or hostname of node",
	"nodes:#.key.type":                     "Node type",
	"nodes:#.key.cloudletkey.organization": "Organization of the cloudlet site",
	"nodes:#.key.cloudletkey.name":         "Name of the cloudlet",
	"nodes:#.key.region":                   "Region the node is in",
	"nodes:#.notifyid":                     "Id of client assigned by server (internal use only)",
	"nodes:#.buildmaster":                  "Build Master Version",
	"nodes:#.buildhead":                    "Build Head Version",
	"nodes:#.buildauthor":                  "Build Author",
	"nodes:#.builddate":                    "Build Date",
	"nodes:#.hostname":                     "Hostname",
	"nodes:#.containerversion":             "Docker edge-cloud container version which node instance use",
	"nodes:#.internalpki":                  "Internal PKI Config",
	"nodes:#.properties":                   "Additional properties",
}
var NodeDataSpecialArgs = map[string]string{
	"nodedata.nodes:#.fields":     "StringArray",
	"nodedata.nodes:#.properties": "StringToString",
}
