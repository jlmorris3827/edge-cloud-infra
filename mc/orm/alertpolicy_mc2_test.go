// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: alertpolicy.proto

package orm

import (
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	"github.com/mobiledgex/edge-cloud-infra/mc/mcctl/mctestclient"
	"github.com/mobiledgex/edge-cloud-infra/mc/orm/testutil"
	edgeproto "github.com/mobiledgex/edge-cloud/edgeproto"
	_ "github.com/mobiledgex/edge-cloud/protogen"
	"github.com/stretchr/testify/require"
	math "math"
	"net/http"
	"testing"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Auto-generated code: DO NOT EDIT

var _ = edgeproto.GetFields

func badPermCreateAlertPolicy(t *testing.T, mcClient *mctestclient.Client, uri, token, region, org string, modFuncs ...func(*edgeproto.AlertPolicy)) {
	_, status, err := testutil.TestPermCreateAlertPolicy(mcClient, uri, token, region, org, modFuncs...)
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "Forbidden")
	require.Equal(t, http.StatusForbidden, status)
}

func badCreateAlertPolicy(t *testing.T, mcClient *mctestclient.Client, uri, token, region, org string, status int, modFuncs ...func(*edgeproto.AlertPolicy)) {
	_, st, err := testutil.TestPermCreateAlertPolicy(mcClient, uri, token, region, org, modFuncs...)
	require.NotNil(t, err)
	require.Equal(t, status, st)
}

func goodPermCreateAlertPolicy(t *testing.T, mcClient *mctestclient.Client, uri, token, region, org string, modFuncs ...func(*edgeproto.AlertPolicy)) {
	_, status, err := testutil.TestPermCreateAlertPolicy(mcClient, uri, token, region, org, modFuncs...)
	require.Nil(t, err)
	require.Equal(t, http.StatusOK, status)
}

func badRegionCreateAlertPolicy(t *testing.T, mcClient *mctestclient.Client, uri, token, org string, modFuncs ...func(*edgeproto.AlertPolicy)) {
	out, status, err := testutil.TestPermCreateAlertPolicy(mcClient, uri, token, "bad region", org, modFuncs...)
	require.NotNil(t, err)
	if err.Error() == "Forbidden" {
		require.Equal(t, http.StatusForbidden, status)
	} else {
		require.Contains(t, err.Error(), "\"bad region\" not found")
		require.Equal(t, http.StatusBadRequest, status)
	}
	_ = out
}

var _ = edgeproto.GetFields

func badPermDeleteAlertPolicy(t *testing.T, mcClient *mctestclient.Client, uri, token, region, org string, modFuncs ...func(*edgeproto.AlertPolicy)) {
	_, status, err := testutil.TestPermDeleteAlertPolicy(mcClient, uri, token, region, org, modFuncs...)
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "Forbidden")
	require.Equal(t, http.StatusForbidden, status)
}

func badDeleteAlertPolicy(t *testing.T, mcClient *mctestclient.Client, uri, token, region, org string, status int, modFuncs ...func(*edgeproto.AlertPolicy)) {
	_, st, err := testutil.TestPermDeleteAlertPolicy(mcClient, uri, token, region, org, modFuncs...)
	require.NotNil(t, err)
	require.Equal(t, status, st)
}

func goodPermDeleteAlertPolicy(t *testing.T, mcClient *mctestclient.Client, uri, token, region, org string, modFuncs ...func(*edgeproto.AlertPolicy)) {
	_, status, err := testutil.TestPermDeleteAlertPolicy(mcClient, uri, token, region, org, modFuncs...)
	require.Nil(t, err)
	require.Equal(t, http.StatusOK, status)
}

func badRegionDeleteAlertPolicy(t *testing.T, mcClient *mctestclient.Client, uri, token, org string, modFuncs ...func(*edgeproto.AlertPolicy)) {
	out, status, err := testutil.TestPermDeleteAlertPolicy(mcClient, uri, token, "bad region", org, modFuncs...)
	require.NotNil(t, err)
	if err.Error() == "Forbidden" {
		require.Equal(t, http.StatusForbidden, status)
	} else {
		require.Contains(t, err.Error(), "\"bad region\" not found")
		require.Equal(t, http.StatusBadRequest, status)
	}
	_ = out
}

var _ = edgeproto.GetFields

func badPermUpdateAlertPolicy(t *testing.T, mcClient *mctestclient.Client, uri, token, region, org string, modFuncs ...func(*edgeproto.AlertPolicy)) {
	_, status, err := testutil.TestPermUpdateAlertPolicy(mcClient, uri, token, region, org, modFuncs...)
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "Forbidden")
	require.Equal(t, http.StatusForbidden, status)
}

func badUpdateAlertPolicy(t *testing.T, mcClient *mctestclient.Client, uri, token, region, org string, status int, modFuncs ...func(*edgeproto.AlertPolicy)) {
	_, st, err := testutil.TestPermUpdateAlertPolicy(mcClient, uri, token, region, org, modFuncs...)
	require.NotNil(t, err)
	require.Equal(t, status, st)
}

func goodPermUpdateAlertPolicy(t *testing.T, mcClient *mctestclient.Client, uri, token, region, org string, modFuncs ...func(*edgeproto.AlertPolicy)) {
	_, status, err := testutil.TestPermUpdateAlertPolicy(mcClient, uri, token, region, org, modFuncs...)
	require.Nil(t, err)
	require.Equal(t, http.StatusOK, status)
}

func badRegionUpdateAlertPolicy(t *testing.T, mcClient *mctestclient.Client, uri, token, org string, modFuncs ...func(*edgeproto.AlertPolicy)) {
	out, status, err := testutil.TestPermUpdateAlertPolicy(mcClient, uri, token, "bad region", org, modFuncs...)
	require.NotNil(t, err)
	if err.Error() == "Forbidden" {
		require.Equal(t, http.StatusForbidden, status)
	} else {
		require.Contains(t, err.Error(), "\"bad region\" not found")
		require.Equal(t, http.StatusBadRequest, status)
	}
	_ = out
}

var _ = edgeproto.GetFields

func badPermShowAlertPolicy(t *testing.T, mcClient *mctestclient.Client, uri, token, region, org string, modFuncs ...func(*edgeproto.AlertPolicy)) {
	_, status, err := testutil.TestPermShowAlertPolicy(mcClient, uri, token, region, org, modFuncs...)
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "Forbidden")
	require.Equal(t, http.StatusForbidden, status)
}

func badShowAlertPolicy(t *testing.T, mcClient *mctestclient.Client, uri, token, region, org string, status int, modFuncs ...func(*edgeproto.AlertPolicy)) {
	_, st, err := testutil.TestPermShowAlertPolicy(mcClient, uri, token, region, org, modFuncs...)
	require.NotNil(t, err)
	require.Equal(t, status, st)
}

func goodPermShowAlertPolicy(t *testing.T, mcClient *mctestclient.Client, uri, token, region, org string, modFuncs ...func(*edgeproto.AlertPolicy)) {
	_, status, err := testutil.TestPermShowAlertPolicy(mcClient, uri, token, region, org, modFuncs...)
	require.Nil(t, err)
	require.Equal(t, http.StatusOK, status)
}

func badRegionShowAlertPolicy(t *testing.T, mcClient *mctestclient.Client, uri, token, org string, modFuncs ...func(*edgeproto.AlertPolicy)) {
	out, status, err := testutil.TestPermShowAlertPolicy(mcClient, uri, token, "bad region", org, modFuncs...)
	require.NotNil(t, err)
	if err.Error() == "Forbidden" {
		require.Equal(t, http.StatusForbidden, status)
	} else {
		require.Contains(t, err.Error(), "\"bad region\" not found")
		require.Equal(t, http.StatusBadRequest, status)
	}
	require.Equal(t, 0, len(out))
}

// This tests the user cannot modify the object because the obj belongs to
// an organization that the user does not have permissions for.
func badPermTestAlertPolicy(t *testing.T, mcClient *mctestclient.Client, uri, token, region, org string, modFuncs ...func(*edgeproto.AlertPolicy)) {
	badPermCreateAlertPolicy(t, mcClient, uri, token, region, org, modFuncs...)
	badPermUpdateAlertPolicy(t, mcClient, uri, token, region, org, modFuncs...)
	badPermDeleteAlertPolicy(t, mcClient, uri, token, region, org, modFuncs...)
}
func badPermTestShowAlertPolicy(t *testing.T, mcClient *mctestclient.Client, uri, token, region, org string) {
	// show is allowed but won't show anything
	var status int
	var err error
	list0, status, err := testutil.TestPermShowAlertPolicy(mcClient, uri, token, region, org)
	require.Nil(t, err)
	require.Equal(t, http.StatusOK, status)
	require.Equal(t, 0, len(list0))
}

// This tests the user can modify the object because the obj belongs to
// an organization that the user has permissions for.
func goodPermTestAlertPolicy(t *testing.T, mcClient *mctestclient.Client, uri, token, region, org string, showcount int, modFuncs ...func(*edgeproto.AlertPolicy)) {
	goodPermCreateAlertPolicy(t, mcClient, uri, token, region, org, modFuncs...)
	goodPermUpdateAlertPolicy(t, mcClient, uri, token, region, org, modFuncs...)
	goodPermDeleteAlertPolicy(t, mcClient, uri, token, region, org, modFuncs...)
	goodPermTestShowAlertPolicy(t, mcClient, uri, token, region, org, showcount)
	// make sure region check works
	badRegionCreateAlertPolicy(t, mcClient, uri, token, org, modFuncs...)
	badRegionUpdateAlertPolicy(t, mcClient, uri, token, org, modFuncs...)
	badRegionDeleteAlertPolicy(t, mcClient, uri, token, org, modFuncs...)
}
func goodPermTestShowAlertPolicy(t *testing.T, mcClient *mctestclient.Client, uri, token, region, org string, count int) {
	var status int
	var err error
	list0, status, err := testutil.TestPermShowAlertPolicy(mcClient, uri, token, region, org)
	require.Nil(t, err)
	require.Equal(t, http.StatusOK, status)
	require.Equal(t, count, len(list0))

	badRegionShowAlertPolicy(t, mcClient, uri, token, org)
}

// Test permissions for user with token1 who should have permissions for
// modifying obj1, and user with token2 who should have permissions for obj2.
// They should not have permissions to modify each other's objects.
func permTestAlertPolicy(t *testing.T, mcClient *mctestclient.Client, uri, token1, token2, region, org1, org2 string, showcount int, modFuncs ...func(*edgeproto.AlertPolicy)) {
	badPermTestAlertPolicy(t, mcClient, uri, token1, region, org2, modFuncs...)
	badPermTestAlertPolicy(t, mcClient, uri, token2, region, org1, modFuncs...)
	badPermTestShowAlertPolicy(t, mcClient, uri, token1, region, org2)
	badPermTestShowAlertPolicy(t, mcClient, uri, token2, region, org1)
	goodPermTestAlertPolicy(t, mcClient, uri, token1, region, org1, showcount, modFuncs...)
	goodPermTestAlertPolicy(t, mcClient, uri, token2, region, org2, showcount, modFuncs...)
}
