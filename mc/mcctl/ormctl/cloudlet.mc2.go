// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cloudlet.proto

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
import _ "github.com/mobiledgex/edge-cloud/d-match-engine/dme-proto"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Auto-generated code: DO NOT EDIT

var CreateCloudletCmd = &cli.Command{
	Use:                  "CreateCloudlet",
	RequiredArgs:         "region " + strings.Join(CreateCloudletRequiredArgs, " "),
	OptionalArgs:         strings.Join(CreateCloudletOptionalArgs, " "),
	AliasArgs:            strings.Join(CloudletAliasArgs, " "),
	SpecialArgs:          &CloudletSpecialArgs,
	Comments:             addRegionComment(CloudletComments),
	ReqData:              &ormapi.RegionCloudlet{},
	ReplyData:            &edgeproto.Result{},
	Run:                  runRest("/auth/ctrl/CreateCloudlet"),
	StreamOut:            true,
	StreamOutIncremental: true,
}

var DeleteCloudletCmd = &cli.Command{
	Use:                  "DeleteCloudlet",
	RequiredArgs:         "region " + strings.Join(CloudletRequiredArgs, " "),
	OptionalArgs:         strings.Join(CloudletOptionalArgs, " "),
	AliasArgs:            strings.Join(CloudletAliasArgs, " "),
	SpecialArgs:          &CloudletSpecialArgs,
	Comments:             addRegionComment(CloudletComments),
	ReqData:              &ormapi.RegionCloudlet{},
	ReplyData:            &edgeproto.Result{},
	Run:                  runRest("/auth/ctrl/DeleteCloudlet"),
	StreamOut:            true,
	StreamOutIncremental: true,
}

var UpdateCloudletCmd = &cli.Command{
	Use:          "UpdateCloudlet",
	RequiredArgs: "region " + strings.Join(CloudletRequiredArgs, " "),
	OptionalArgs: strings.Join(CloudletOptionalArgs, " "),
	AliasArgs:    strings.Join(CloudletAliasArgs, " "),
	SpecialArgs:  &CloudletSpecialArgs,
	Comments:     addRegionComment(CloudletComments),
	ReqData:      &ormapi.RegionCloudlet{},
	ReplyData:    &edgeproto.Result{},
	Run: runRest("/auth/ctrl/UpdateCloudlet",
		withSetFieldsFunc(setUpdateCloudletFields),
	),
	StreamOut:            true,
	StreamOutIncremental: true,
}

func setUpdateCloudletFields(in map[string]interface{}) {
	// get map for edgeproto object in region struct
	obj := in[strings.ToLower("Cloudlet")]
	if obj == nil {
		return
	}
	objmap, ok := obj.(map[string]interface{})
	if !ok {
		return
	}
	objmap["fields"] = cli.GetSpecifiedFields(objmap, &edgeproto.Cloudlet{}, cli.JsonNamespace)
}

var ShowCloudletCmd = &cli.Command{
	Use:          "ShowCloudlet",
	RequiredArgs: "region",
	OptionalArgs: strings.Join(append(CloudletRequiredArgs, CloudletOptionalArgs...), " "),
	AliasArgs:    strings.Join(CloudletAliasArgs, " "),
	SpecialArgs:  &CloudletSpecialArgs,
	Comments:     addRegionComment(CloudletComments),
	ReqData:      &ormapi.RegionCloudlet{},
	ReplyData:    &edgeproto.Cloudlet{},
	Run:          runRest("/auth/ctrl/ShowCloudlet"),
	StreamOut:    true,
}

var AddCloudletResMappingCmd = &cli.Command{
	Use:          "AddCloudletResMapping",
	RequiredArgs: "region " + strings.Join(CloudletResMapRequiredArgs, " "),
	OptionalArgs: strings.Join(CloudletResMapOptionalArgs, " "),
	AliasArgs:    strings.Join(CloudletResMapAliasArgs, " "),
	SpecialArgs:  &CloudletResMapSpecialArgs,
	Comments:     addRegionComment(CloudletResMapComments),
	ReqData:      &ormapi.RegionCloudletResMap{},
	ReplyData:    &edgeproto.Result{},
	Run:          runRest("/auth/ctrl/AddCloudletResMapping"),
}

var RemoveCloudletResMappingCmd = &cli.Command{
	Use:          "RemoveCloudletResMapping",
	RequiredArgs: "region " + strings.Join(CloudletResMapRequiredArgs, " "),
	OptionalArgs: strings.Join(CloudletResMapOptionalArgs, " "),
	AliasArgs:    strings.Join(CloudletResMapAliasArgs, " "),
	SpecialArgs:  &CloudletResMapSpecialArgs,
	Comments:     addRegionComment(CloudletResMapComments),
	ReqData:      &ormapi.RegionCloudletResMap{},
	ReplyData:    &edgeproto.Result{},
	Run:          runRest("/auth/ctrl/RemoveCloudletResMapping"),
}

var FindFlavorMatchCmd = &cli.Command{
	Use:          "FindFlavorMatch",
	RequiredArgs: "region " + strings.Join(FlavorMatchRequiredArgs, " "),
	OptionalArgs: strings.Join(FlavorMatchOptionalArgs, " "),
	AliasArgs:    strings.Join(FlavorMatchAliasArgs, " "),
	SpecialArgs:  &FlavorMatchSpecialArgs,
	Comments:     addRegionComment(FlavorMatchComments),
	ReqData:      &ormapi.RegionFlavorMatch{},
	ReplyData:    &edgeproto.FlavorMatch{},
	Run:          runRest("/auth/ctrl/FindFlavorMatch"),
}

var CloudletApiCmds = []*cli.Command{
	CreateCloudletCmd,
	DeleteCloudletCmd,
	UpdateCloudletCmd,
	ShowCloudletCmd,
	AddCloudletResMappingCmd,
	RemoveCloudletResMappingCmd,
	FindFlavorMatchCmd,
}

var CreateCloudletRequiredArgs = []string{
	"cloudlet-org",
	"cloudlet",
	"location.latitude",
	"location.longitude",
	"numdynamicips",
}
var CreateCloudletOptionalArgs = []string{
	"location.altitude",
	"location.timestamp.seconds",
	"location.timestamp.nanos",
	"ipsupport",
	"staticips",
	"timelimits.createclusterinsttimeout",
	"timelimits.updateclusterinsttimeout",
	"timelimits.deleteclusterinsttimeout",
	"timelimits.createappinsttimeout",
	"timelimits.updateappinsttimeout",
	"timelimits.deleteappinsttimeout",
	"errors",
	"state",
	"crmoverride",
	"deploymentlocal",
	"platformtype",
	"flavor.name",
	"physicalname",
	"envvar",
	"containerversion",
	"restagmap[#].key",
	"restagmap[#].value.name",
	"restagmap[#].value.organization",
	"accessvars",
	"vmimageversion",
	"packageversion",
}

var ShowCloudletInfoCmd = &cli.Command{
	Use:          "ShowCloudletInfo",
	RequiredArgs: "region",
	OptionalArgs: strings.Join(append(CloudletInfoRequiredArgs, CloudletInfoOptionalArgs...), " "),
	AliasArgs:    strings.Join(CloudletInfoAliasArgs, " "),
	SpecialArgs:  &CloudletInfoSpecialArgs,
	Comments:     addRegionComment(CloudletInfoComments),
	ReqData:      &ormapi.RegionCloudletInfo{},
	ReplyData:    &edgeproto.CloudletInfo{},
	Run:          runRest("/auth/ctrl/ShowCloudletInfo"),
	StreamOut:    true,
}

var CloudletInfoApiCmds = []*cli.Command{
	ShowCloudletInfoCmd,
}

var CloudletKeyRequiredArgs = []string{}
var CloudletKeyOptionalArgs = []string{
	"organization",
	"name",
}
var CloudletKeyAliasArgs = []string{
	"organization=cloudletkey.organization",
	"name=cloudletkey.name",
}
var CloudletKeyComments = map[string]string{
	"organization": "Organization of the cloudlet site",
	"name":         "Name of the cloudlet",
}
var CloudletKeySpecialArgs = map[string]string{}
var OperationTimeLimitsRequiredArgs = []string{}
var OperationTimeLimitsOptionalArgs = []string{
	"createclusterinsttimeout",
	"updateclusterinsttimeout",
	"deleteclusterinsttimeout",
	"createappinsttimeout",
	"updateappinsttimeout",
	"deleteappinsttimeout",
}
var OperationTimeLimitsAliasArgs = []string{
	"createclusterinsttimeout=operationtimelimits.createclusterinsttimeout",
	"updateclusterinsttimeout=operationtimelimits.updateclusterinsttimeout",
	"deleteclusterinsttimeout=operationtimelimits.deleteclusterinsttimeout",
	"createappinsttimeout=operationtimelimits.createappinsttimeout",
	"updateappinsttimeout=operationtimelimits.updateappinsttimeout",
	"deleteappinsttimeout=operationtimelimits.deleteappinsttimeout",
}
var OperationTimeLimitsComments = map[string]string{
	"createclusterinsttimeout": "override default max time to create a cluster instance (duration)",
	"updateclusterinsttimeout": "override default max time to update a cluster instance (duration)",
	"deleteclusterinsttimeout": "override default max time to delete a cluster instance (duration)",
	"createappinsttimeout":     "override default max time to create an app instance (duration)",
	"updateappinsttimeout":     "override default max time to update an app instance (duration)",
	"deleteappinsttimeout":     "override default max time to delete an app instance (duration)",
}
var OperationTimeLimitsSpecialArgs = map[string]string{}
var CloudletInfraCommonRequiredArgs = []string{}
var CloudletInfraCommonOptionalArgs = []string{
	"dockerregistry",
	"dnszone",
	"registryfileserver",
	"cfkey",
	"cfuser",
	"dockerregpass",
	"networkscheme",
	"dockerregistrysecret",
}
var CloudletInfraCommonAliasArgs = []string{
	"dockerregistry=cloudletinfracommon.dockerregistry",
	"dnszone=cloudletinfracommon.dnszone",
	"registryfileserver=cloudletinfracommon.registryfileserver",
	"cfkey=cloudletinfracommon.cfkey",
	"cfuser=cloudletinfracommon.cfuser",
	"dockerregpass=cloudletinfracommon.dockerregpass",
	"networkscheme=cloudletinfracommon.networkscheme",
	"dockerregistrysecret=cloudletinfracommon.dockerregistrysecret",
}
var CloudletInfraCommonComments = map[string]string{
	"dockerregistry":       "the mex docker registry, e.g.  registry.mobiledgex.net:5000.",
	"dnszone":              "DNS Zone",
	"registryfileserver":   "registry file server contains files which get pulled on instantiation such as certs and images",
	"cfkey":                "Cloudflare key",
	"cfuser":               "Cloudflare key",
	"dockerregpass":        "Docker registry password",
	"networkscheme":        "network scheme",
	"dockerregistrysecret": "the name of the docker registry secret, e.g. mexgitlabsecret",
}
var CloudletInfraCommonSpecialArgs = map[string]string{}
var AzurePropertiesRequiredArgs = []string{}
var AzurePropertiesOptionalArgs = []string{
	"location",
	"resourcegroup",
	"username",
	"password",
}
var AzurePropertiesAliasArgs = []string{
	"location=azureproperties.location",
	"resourcegroup=azureproperties.resourcegroup",
	"username=azureproperties.username",
	"password=azureproperties.password",
}
var AzurePropertiesComments = map[string]string{
	"location":      "azure region e.g. uswest2",
	"resourcegroup": "azure resource group",
	"username":      "azure username",
	"password":      "azure password",
}
var AzurePropertiesSpecialArgs = map[string]string{}
var GcpPropertiesRequiredArgs = []string{}
var GcpPropertiesOptionalArgs = []string{
	"project",
	"zone",
	"serviceaccount",
	"gcpauthkeyurl",
}
var GcpPropertiesAliasArgs = []string{
	"project=gcpproperties.project",
	"zone=gcpproperties.zone",
	"serviceaccount=gcpproperties.serviceaccount",
	"gcpauthkeyurl=gcpproperties.gcpauthkeyurl",
}
var GcpPropertiesComments = map[string]string{
	"project":        "gcp project for billing",
	"zone":           "availability zone",
	"serviceaccount": "service account to login with",
	"gcpauthkeyurl":  "vault credentials link",
}
var GcpPropertiesSpecialArgs = map[string]string{}
var OpenStackPropertiesRequiredArgs = []string{}
var OpenStackPropertiesOptionalArgs = []string{
	"osexternalnetworkname",
	"osimagename",
	"osexternalroutername",
	"osmexnetwork",
	"openrcvars",
}
var OpenStackPropertiesAliasArgs = []string{
	"osexternalnetworkname=openstackproperties.osexternalnetworkname",
	"osimagename=openstackproperties.osimagename",
	"osexternalroutername=openstackproperties.osexternalroutername",
	"osmexnetwork=openstackproperties.osmexnetwork",
	"openrcvars=openstackproperties.openrcvars",
}
var OpenStackPropertiesComments = map[string]string{
	"osexternalnetworkname": "name of the external network, e.g. external-network-shared",
	"osimagename":           "openstack image , e.g. mobiledgex",
	"osexternalroutername":  "openstack router",
	"osmexnetwork":          "openstack internal network",
	"openrcvars":            "openrc env vars",
}
var OpenStackPropertiesSpecialArgs = map[string]string{
	"openstackproperties.openrcvars": "StringToString",
}
var CloudletInfraPropertiesRequiredArgs = []string{}
var CloudletInfraPropertiesOptionalArgs = []string{
	"cloudletkind",
	"mexoscontainerimagename",
	"openstackproperties.osexternalnetworkname",
	"openstackproperties.osimagename",
	"openstackproperties.osexternalroutername",
	"openstackproperties.osmexnetwork",
	"openstackproperties.openrcvars",
	"azureproperties.location",
	"azureproperties.resourcegroup",
	"azureproperties.username",
	"azureproperties.password",
	"gcpproperties.project",
	"gcpproperties.zone",
	"gcpproperties.serviceaccount",
	"gcpproperties.gcpauthkeyurl",
}
var CloudletInfraPropertiesAliasArgs = []string{
	"cloudletkind=cloudletinfraproperties.cloudletkind",
	"mexoscontainerimagename=cloudletinfraproperties.mexoscontainerimagename",
	"openstackproperties.osexternalnetworkname=cloudletinfraproperties.openstackproperties.osexternalnetworkname",
	"openstackproperties.osimagename=cloudletinfraproperties.openstackproperties.osimagename",
	"openstackproperties.osexternalroutername=cloudletinfraproperties.openstackproperties.osexternalroutername",
	"openstackproperties.osmexnetwork=cloudletinfraproperties.openstackproperties.osmexnetwork",
	"openstackproperties.openrcvars=cloudletinfraproperties.openstackproperties.openrcvars",
	"azureproperties.location=cloudletinfraproperties.azureproperties.location",
	"azureproperties.resourcegroup=cloudletinfraproperties.azureproperties.resourcegroup",
	"azureproperties.username=cloudletinfraproperties.azureproperties.username",
	"azureproperties.password=cloudletinfraproperties.azureproperties.password",
	"gcpproperties.project=cloudletinfraproperties.gcpproperties.project",
	"gcpproperties.zone=cloudletinfraproperties.gcpproperties.zone",
	"gcpproperties.serviceaccount=cloudletinfraproperties.gcpproperties.serviceaccount",
	"gcpproperties.gcpauthkeyurl=cloudletinfraproperties.gcpproperties.gcpauthkeyurl",
}
var CloudletInfraPropertiesComments = map[string]string{
	"cloudletkind":                              "what kind of infrastructure: Azure, GCP, Openstack",
	"mexoscontainerimagename":                   "name and version of the docker image container image that mexos runs in",
	"openstackproperties.osexternalnetworkname": "name of the external network, e.g. external-network-shared",
	"openstackproperties.osimagename":           "openstack image , e.g. mobiledgex",
	"openstackproperties.osexternalroutername":  "openstack router",
	"openstackproperties.osmexnetwork":          "openstack internal network",
	"openstackproperties.openrcvars":            "openrc env vars",
	"azureproperties.location":                  "azure region e.g. uswest2",
	"azureproperties.resourcegroup":             "azure resource group",
	"azureproperties.username":                  "azure username",
	"azureproperties.password":                  "azure password",
	"gcpproperties.project":                     "gcp project for billing",
	"gcpproperties.zone":                        "availability zone",
	"gcpproperties.serviceaccount":              "service account to login with",
	"gcpproperties.gcpauthkeyurl":               "vault credentials link",
}
var CloudletInfraPropertiesSpecialArgs = map[string]string{
	"cloudletinfraproperties.openstackproperties.openrcvars": "StringToString",
}
var PlatformConfigRequiredArgs = []string{}
var PlatformConfigOptionalArgs = []string{
	"containerregistrypath",
	"cloudletvmimagepath",
	"notifyctrladdrs",
	"vaultaddr",
	"tlscertfile",
	"envvar",
	"platformtag",
	"testmode",
	"span",
	"cleanupmode",
	"region",
	"commercialcerts",
	"usevaultcerts",
}
var PlatformConfigAliasArgs = []string{
	"containerregistrypath=platformconfig.containerregistrypath",
	"cloudletvmimagepath=platformconfig.cloudletvmimagepath",
	"notifyctrladdrs=platformconfig.notifyctrladdrs",
	"vaultaddr=platformconfig.vaultaddr",
	"tlscertfile=platformconfig.tlscertfile",
	"envvar=platformconfig.envvar",
	"platformtag=platformconfig.platformtag",
	"testmode=platformconfig.testmode",
	"span=platformconfig.span",
	"cleanupmode=platformconfig.cleanupmode",
	"region=platformconfig.region",
	"commercialcerts=platformconfig.commercialcerts",
	"usevaultcerts=platformconfig.usevaultcerts",
}
var PlatformConfigComments = map[string]string{
	"containerregistrypath": "Path to Docker registry holding edge-cloud image",
	"cloudletvmimagepath":   "Path to platform base image",
	"notifyctrladdrs":       "Address of controller notify port (can be multiple of these)",
	"vaultaddr":             "Vault address",
	"tlscertfile":           "TLS cert file",
	"envvar":                "Environment variables",
	"platformtag":           "Tag of edge-cloud image",
	"testmode":              "Internal Test flag",
	"span":                  "Span string",
	"cleanupmode":           "Internal cleanup flag",
	"region":                "Region",
	"commercialcerts":       "Get certs from vault or generate your own for the root load balancer",
	"usevaultcerts":         "Use Vault certs for internal TLS communication",
}
var PlatformConfigSpecialArgs = map[string]string{
	"platformconfig.envvar": "StringToString",
}
var CloudletResMapRequiredArgs = []string{
	"cloudlet-org",
	"cloudlet",
	"mapping",
}
var CloudletResMapOptionalArgs = []string{}
var CloudletResMapAliasArgs = []string{
	"cloudlet-org=cloudletresmap.key.organization",
	"cloudlet=cloudletresmap.key.name",
	"mapping=cloudletresmap.mapping",
}
var CloudletResMapComments = map[string]string{
	"cloudlet-org": "Organization of the cloudlet site",
	"cloudlet":     "Name of the cloudlet",
	"mapping":      "Resource mapping info",
}
var CloudletResMapSpecialArgs = map[string]string{
	"cloudletresmap.mapping": "StringToString",
}
var CloudletRequiredArgs = []string{
	"cloudlet-org",
	"cloudlet",
}
var CloudletOptionalArgs = []string{
	"location.latitude",
	"location.longitude",
	"location.altitude",
	"location.timestamp.seconds",
	"location.timestamp.nanos",
	"ipsupport",
	"staticips",
	"numdynamicips",
	"timelimits.createclusterinsttimeout",
	"timelimits.updateclusterinsttimeout",
	"timelimits.deleteclusterinsttimeout",
	"timelimits.createappinsttimeout",
	"timelimits.updateappinsttimeout",
	"timelimits.deleteappinsttimeout",
	"errors",
	"state",
	"crmoverride",
	"deploymentlocal",
	"platformtype",
	"flavor.name",
	"physicalname",
	"envvar",
	"containerversion",
	"restagmap[#].key",
	"restagmap[#].value.name",
	"restagmap[#].value.organization",
	"accessvars",
	"vmimageversion",
	"packageversion",
}
var CloudletAliasArgs = []string{
	"fields=cloudlet.fields",
	"cloudlet-org=cloudlet.key.organization",
	"cloudlet=cloudlet.key.name",
	"location.latitude=cloudlet.location.latitude",
	"location.longitude=cloudlet.location.longitude",
	"location.horizontalaccuracy=cloudlet.location.horizontalaccuracy",
	"location.verticalaccuracy=cloudlet.location.verticalaccuracy",
	"location.altitude=cloudlet.location.altitude",
	"location.course=cloudlet.location.course",
	"location.speed=cloudlet.location.speed",
	"location.timestamp.seconds=cloudlet.location.timestamp.seconds",
	"location.timestamp.nanos=cloudlet.location.timestamp.nanos",
	"ipsupport=cloudlet.ipsupport",
	"staticips=cloudlet.staticips",
	"numdynamicips=cloudlet.numdynamicips",
	"timelimits.createclusterinsttimeout=cloudlet.timelimits.createclusterinsttimeout",
	"timelimits.updateclusterinsttimeout=cloudlet.timelimits.updateclusterinsttimeout",
	"timelimits.deleteclusterinsttimeout=cloudlet.timelimits.deleteclusterinsttimeout",
	"timelimits.createappinsttimeout=cloudlet.timelimits.createappinsttimeout",
	"timelimits.updateappinsttimeout=cloudlet.timelimits.updateappinsttimeout",
	"timelimits.deleteappinsttimeout=cloudlet.timelimits.deleteappinsttimeout",
	"errors=cloudlet.errors",
	"status.tasknumber=cloudlet.status.tasknumber",
	"status.maxtasks=cloudlet.status.maxtasks",
	"status.taskname=cloudlet.status.taskname",
	"status.stepname=cloudlet.status.stepname",
	"state=cloudlet.state",
	"crmoverride=cloudlet.crmoverride",
	"deploymentlocal=cloudlet.deploymentlocal",
	"platformtype=cloudlet.platformtype",
	"notifysrvaddr=cloudlet.notifysrvaddr",
	"flavor.name=cloudlet.flavor.name",
	"physicalname=cloudlet.physicalname",
	"envvar=cloudlet.envvar",
	"containerversion=cloudlet.containerversion",
	"config.containerregistrypath=cloudlet.config.containerregistrypath",
	"config.cloudletvmimagepath=cloudlet.config.cloudletvmimagepath",
	"config.notifyctrladdrs=cloudlet.config.notifyctrladdrs",
	"config.vaultaddr=cloudlet.config.vaultaddr",
	"config.tlscertfile=cloudlet.config.tlscertfile",
	"config.envvar=cloudlet.config.envvar",
	"config.platformtag=cloudlet.config.platformtag",
	"config.testmode=cloudlet.config.testmode",
	"config.span=cloudlet.config.span",
	"config.cleanupmode=cloudlet.config.cleanupmode",
	"config.region=cloudlet.config.region",
	"config.commercialcerts=cloudlet.config.commercialcerts",
	"config.usevaultcerts=cloudlet.config.usevaultcerts",
	"restagmap[#].key=cloudlet.restagmap[#].key",
	"restagmap[#].value.name=cloudlet.restagmap[#].value.name",
	"restagmap[#].value.organization=cloudlet.restagmap[#].value.organization",
	"accessvars=cloudlet.accessvars",
	"vmimageversion=cloudlet.vmimageversion",
	"packageversion=cloudlet.packageversion",
}
var CloudletComments = map[string]string{
	"fields":                              "Fields are used for the Update API to specify which fields to apply",
	"cloudlet-org":                        "Organization of the cloudlet site",
	"cloudlet":                            "Name of the cloudlet",
	"location.latitude":                   "latitude in WGS 84 coordinates",
	"location.longitude":                  "longitude in WGS 84 coordinates",
	"location.horizontalaccuracy":         "horizontal accuracy (radius in meters)",
	"location.verticalaccuracy":           "vertical accuracy (meters)",
	"location.altitude":                   "On android only lat and long are guaranteed to be supplied altitude in meters",
	"location.course":                     "course (IOS) / bearing (Android) (degrees east relative to true north)",
	"location.speed":                      "speed (IOS) / velocity (Android) (meters/sec)",
	"ipsupport":                           "Type of IP support provided by Cloudlet (see IpSupport), one of IpSupportUnknown, IpSupportStatic, IpSupportDynamic",
	"staticips":                           "List of static IPs for static IP support",
	"numdynamicips":                       "Number of dynamic IPs available for dynamic IP support",
	"timelimits.createclusterinsttimeout": "override default max time to create a cluster instance (duration)",
	"timelimits.updateclusterinsttimeout": "override default max time to update a cluster instance (duration)",
	"timelimits.deleteclusterinsttimeout": "override default max time to delete a cluster instance (duration)",
	"timelimits.createappinsttimeout":     "override default max time to create an app instance (duration)",
	"timelimits.updateappinsttimeout":     "override default max time to update an app instance (duration)",
	"timelimits.deleteappinsttimeout":     "override default max time to delete an app instance (duration)",
	"errors":                              "Any errors trying to create, update, or delete the Cloudlet.",
	"state":                               "Current state of the cloudlet, one of TrackedStateUnknown, NotPresent, CreateRequested, Creating, CreateError, Ready, UpdateRequested, Updating, UpdateError, DeleteRequested, Deleting, DeleteError, DeletePrepare, CrmInitok, CreatingDependencies",
	"crmoverride":                         "Override actions to CRM, one of NoOverride, IgnoreCrmErrors, IgnoreCrm, IgnoreTransientState, IgnoreCrmAndTransientState",
	"deploymentlocal":                     "Deploy cloudlet services locally",
	"platformtype":                        "Platform type, one of PlatformTypeFake, PlatformTypeDind, PlatformTypeOpenstack, PlatformTypeAzure, PlatformTypeGcp, PlatformTypeEdgebox, PlatformTypeFakeinfra",
	"notifysrvaddr":                       "Address for the CRM notify listener to run on",
	"flavor.name":                         "Flavor name",
	"physicalname":                        "Physical infrastructure cloudlet name",
	"envvar":                              "Single Key-Value pair of env var to be passed to CRM",
	"containerversion":                    "Cloudlet container version",
	"config.containerregistrypath":        "Path to Docker registry holding edge-cloud image",
	"config.cloudletvmimagepath":          "Path to platform base image",
	"config.notifyctrladdrs":              "Address of controller notify port (can be multiple of these)",
	"config.vaultaddr":                    "Vault address",
	"config.tlscertfile":                  "TLS cert file",
	"config.envvar":                       "Environment variables",
	"config.platformtag":                  "Tag of edge-cloud image",
	"config.testmode":                     "Internal Test flag",
	"config.span":                         "Span string",
	"config.cleanupmode":                  "Internal cleanup flag",
	"config.region":                       "Region",
	"config.commercialcerts":              "Get certs from vault or generate your own for the root load balancer",
	"config.usevaultcerts":                "Use Vault certs for internal TLS communication",
	"restagmap[#].value.name":             "Resource Table Name",
	"restagmap[#].value.organization":     "Operator organization of the cloudlet site.",
	"accessvars":                          "Variables required to access cloudlet",
	"vmimageversion":                      "MobiledgeX baseimage version where CRM services reside",
	"packageversion":                      "MobiledgeX OS package version on baseimage where CRM services reside",
}
var CloudletSpecialArgs = map[string]string{
	"cloudlet.accessvars":    "StringToString",
	"cloudlet.config.envvar": "StringToString",
	"cloudlet.envvar":        "StringToString",
	"cloudlet.errors":        "StringArray",
	"cloudlet.fields":        "StringArray",
}
var FlavorMatchRequiredArgs = []string{
	"cloudlet-org",
	"cloudlet",
}
var FlavorMatchOptionalArgs = []string{
	"flavor",
	"availabilityzone",
}
var FlavorMatchAliasArgs = []string{
	"cloudlet-org=flavormatch.key.organization",
	"cloudlet=flavormatch.key.name",
	"flavor=flavormatch.flavorname",
	"availabilityzone=flavormatch.availabilityzone",
}
var FlavorMatchComments = map[string]string{
	"cloudlet-org": "Organization of the cloudlet site",
	"cloudlet":     "Name of the cloudlet",
}
var FlavorMatchSpecialArgs = map[string]string{}
var FlavorInfoRequiredArgs = []string{}
var FlavorInfoOptionalArgs = []string{
	"name",
	"vcpus",
	"ram",
	"disk",
	"propmap",
}
var FlavorInfoAliasArgs = []string{
	"name=flavorinfo.name",
	"vcpus=flavorinfo.vcpus",
	"ram=flavorinfo.ram",
	"disk=flavorinfo.disk",
	"propmap=flavorinfo.propmap",
}
var FlavorInfoComments = map[string]string{
	"name":    "Name of the flavor on the Cloudlet",
	"vcpus":   "Number of VCPU cores on the Cloudlet",
	"ram":     "Ram in MB on the Cloudlet",
	"disk":    "Amount of disk in GB on the Cloudlet",
	"propmap": "OS Flavor Properties, if any",
}
var FlavorInfoSpecialArgs = map[string]string{
	"flavorinfo.propmap": "StringToString",
}
var OSAZoneRequiredArgs = []string{}
var OSAZoneOptionalArgs = []string{
	"name",
	"status",
}
var OSAZoneAliasArgs = []string{
	"name=osazone.name",
	"status=osazone.status",
}
var OSAZoneComments = map[string]string{}
var OSAZoneSpecialArgs = map[string]string{}
var OSImageRequiredArgs = []string{}
var OSImageOptionalArgs = []string{
	"name",
	"tags",
	"properties",
	"diskformat",
}
var OSImageAliasArgs = []string{
	"name=osimage.name",
	"tags=osimage.tags",
	"properties=osimage.properties",
	"diskformat=osimage.diskformat",
}
var OSImageComments = map[string]string{
	"name":       "image name",
	"tags":       "optional tags present on image",
	"properties": "image properties/metadata",
	"diskformat": "format qcow2, img, etc",
}
var OSImageSpecialArgs = map[string]string{}
var CloudletInfoRequiredArgs = []string{
	"cloudlet-org",
	"cloudlet",
}
var CloudletInfoOptionalArgs = []string{
	"state",
	"notifyid",
	"controller",
	"osmaxram",
	"osmaxvcores",
	"osmaxvolgb",
	"errors",
	"flavors[#].name",
	"flavors[#].vcpus",
	"flavors[#].ram",
	"flavors[#].disk",
	"flavors[#].propmap",
	"status.tasknumber",
	"status.maxtasks",
	"status.taskname",
	"status.stepname",
	"containerversion",
	"availabilityzones[#].name",
	"availabilityzones[#].status",
	"osimages[#].name",
	"osimages[#].tags",
	"osimages[#].properties",
	"osimages[#].diskformat",
}
var CloudletInfoAliasArgs = []string{
	"fields=cloudletinfo.fields",
	"cloudlet-org=cloudletinfo.key.organization",
	"cloudlet=cloudletinfo.key.name",
	"state=cloudletinfo.state",
	"notifyid=cloudletinfo.notifyid",
	"controller=cloudletinfo.controller",
	"osmaxram=cloudletinfo.osmaxram",
	"osmaxvcores=cloudletinfo.osmaxvcores",
	"osmaxvolgb=cloudletinfo.osmaxvolgb",
	"errors=cloudletinfo.errors",
	"flavors[#].name=cloudletinfo.flavors[#].name",
	"flavors[#].vcpus=cloudletinfo.flavors[#].vcpus",
	"flavors[#].ram=cloudletinfo.flavors[#].ram",
	"flavors[#].disk=cloudletinfo.flavors[#].disk",
	"flavors[#].propmap=cloudletinfo.flavors[#].propmap",
	"status.tasknumber=cloudletinfo.status.tasknumber",
	"status.maxtasks=cloudletinfo.status.maxtasks",
	"status.taskname=cloudletinfo.status.taskname",
	"status.stepname=cloudletinfo.status.stepname",
	"containerversion=cloudletinfo.containerversion",
	"availabilityzones[#].name=cloudletinfo.availabilityzones[#].name",
	"availabilityzones[#].status=cloudletinfo.availabilityzones[#].status",
	"osimages[#].name=cloudletinfo.osimages[#].name",
	"osimages[#].tags=cloudletinfo.osimages[#].tags",
	"osimages[#].properties=cloudletinfo.osimages[#].properties",
	"osimages[#].diskformat=cloudletinfo.osimages[#].diskformat",
}
var CloudletInfoComments = map[string]string{
	"fields":                 "Fields are used for the Update API to specify which fields to apply",
	"cloudlet-org":           "Organization of the cloudlet site",
	"cloudlet":               "Name of the cloudlet",
	"state":                  "State of cloudlet, one of CloudletStateUnknown, CloudletStateErrors, CloudletStateReady, CloudletStateOffline, CloudletStateNotPresent, CloudletStateInit, CloudletStateUpgrade",
	"notifyid":               "Id of client assigned by server (internal use only)",
	"controller":             "Connected controller unique id",
	"osmaxram":               "Maximum Ram in MB on the Cloudlet",
	"osmaxvcores":            "Maximum number of VCPU cores on the Cloudlet",
	"osmaxvolgb":             "Maximum amount of disk in GB on the Cloudlet",
	"errors":                 "Any errors encountered while making changes to the Cloudlet",
	"flavors[#].name":        "Name of the flavor on the Cloudlet",
	"flavors[#].vcpus":       "Number of VCPU cores on the Cloudlet",
	"flavors[#].ram":         "Ram in MB on the Cloudlet",
	"flavors[#].disk":        "Amount of disk in GB on the Cloudlet",
	"flavors[#].propmap":     "OS Flavor Properties, if any",
	"containerversion":       "Cloudlet container version",
	"osimages[#].name":       "image name",
	"osimages[#].tags":       "optional tags present on image",
	"osimages[#].properties": "image properties/metadata",
	"osimages[#].diskformat": "format qcow2, img, etc",
}
var CloudletInfoSpecialArgs = map[string]string{
	"cloudletinfo.errors":             "StringArray",
	"cloudletinfo.fields":             "StringArray",
	"cloudletinfo.flavors[#].propmap": "StringToString",
}
var CloudletMetricsRequiredArgs = []string{}
var CloudletMetricsOptionalArgs = []string{
	"foo",
}
var CloudletMetricsAliasArgs = []string{
	"foo=cloudletmetrics.foo",
}
var CloudletMetricsComments = map[string]string{
	"foo": "what goes here?",
}
var CloudletMetricsSpecialArgs = map[string]string{}
