package keeper_test

import (
	"reflect"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/ovrclk/akash/testutil"
	"github.com/ovrclk/akash/testutil/state"
	"github.com/ovrclk/akash/x/provider/keeper"
	types "github.com/ovrclk/akash/x/provider/types/v1beta2"
)

func TestProviderCreate(t *testing.T) {
	ctx, keeper := setupKeeper(t)
	prov := testutil.Provider(t)

	err := keeper.Create(ctx, prov)
	require.NoError(t, err)

	owner, err := sdk.AccAddressFromBech32(prov.Owner)
	require.NoError(t, err)

	foundProv, found := keeper.Get(ctx, owner)
	require.True(t, found)
	require.Equal(t, prov, foundProv)
}

func TestProviderDuplicate(t *testing.T) {
	ctx, keeper := setupKeeper(t)
	prov := testutil.Provider(t)

	err := keeper.Create(ctx, prov)
	require.NoError(t, err)

	err = keeper.Create(ctx, prov)
	require.EqualError(t, err, types.ErrProviderExists.Error())
}

func TestProviderGetNonExisting(t *testing.T) {
	ctx, keeper := setupKeeper(t)
	prov := testutil.Provider(t)

	owner, err := sdk.AccAddressFromBech32(prov.Owner)
	require.NoError(t, err)

	foundProv, found := keeper.Get(ctx, owner)
	require.False(t, found)
	require.Equal(t, types.Provider{}, foundProv)
}

func TestProviderDeleteExisting(t *testing.T) {
	ctx, keeper := setupKeeper(t)
	prov := testutil.Provider(t)

	err := keeper.Create(ctx, prov)
	require.NoError(t, err)

	owner, err := sdk.AccAddressFromBech32(prov.Owner)
	require.NoError(t, err)

	require.Panics(t, func() {
		keeper.Delete(ctx, owner)
	})

	foundProv, found := keeper.Get(ctx, owner)
	require.True(t, found)
	require.Equal(t, prov, foundProv)
}

func TestProviderUpdateNonExisting(t *testing.T) {
	ctx, keeper := setupKeeper(t)
	prov := testutil.Provider(t)

	err := keeper.Update(ctx, prov)
	require.EqualError(t, err, types.ErrProviderNotFound.Error())
}

func TestProviderUpdateExisting(t *testing.T) {
	ctx, keeper := setupKeeper(t)
	prov := testutil.Provider(t)

	err := keeper.Create(ctx, prov)
	require.NoError(t, err)

	prov.HostURI = "akash.domain.com"
	err = keeper.Update(ctx, prov)
	require.NoError(t, err)

	owner, err := sdk.AccAddressFromBech32(prov.Owner)
	require.NoError(t, err)

	foundProv, found := keeper.Get(ctx, owner)
	require.True(t, found)
	require.Equal(t, prov, foundProv)
}

func TestWithProviders(t *testing.T) {
	ctx, keeper := setupKeeper(t)
	prov := testutil.Provider(t)
	prov2 := testutil.Provider(t)

	err := keeper.Create(ctx, prov)
	require.NoError(t, err)

	err = keeper.Create(ctx, prov2)
	require.NoError(t, err)

	count := 0

	keeper.WithProviders(ctx, func(provider types.Provider) bool {
		if !reflect.DeepEqual(provider, prov) && !reflect.DeepEqual(provider, prov2) {
			require.Fail(t, "unknown provider")
		}
		count++
		return false
	})

	require.Equal(t, 2, count)
}

func TestWithProvidersBreak(t *testing.T) {
	ctx, keeper := setupKeeper(t)
	prov := testutil.Provider(t)
	prov2 := testutil.Provider(t)

	err := keeper.Create(ctx, prov)
	require.NoError(t, err)

	err = keeper.Create(ctx, prov2)
	require.NoError(t, err)

	count := 0

	keeper.WithProviders(ctx, func(provider types.Provider) bool {
		if !reflect.DeepEqual(provider, prov) && !reflect.DeepEqual(provider, prov2) {
			require.Fail(t, "unknown provider")
		}
		count++
		return true
	})

	require.Equal(t, 1, count)
}

func TestKeeperCoder(t *testing.T) {
	_, keeper := setupKeeper(t)
	codec := keeper.Codec()
	require.NotNil(t, codec)
}

func setupKeeper(t testing.TB) (sdk.Context, keeper.IKeeper) {
	t.Helper()

	suite := state.SetupTestSuite(t)

	return suite.Context(), suite.ProviderKeeper()
}
