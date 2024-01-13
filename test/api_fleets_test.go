/*
Torizon OTA

Testing FleetsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	openapiclient "github.com/commontorizon/torizon-openapi-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_FleetsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FleetsAPIService DeleteFleetsFleetid", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fleetId string

		httpRes, err := apiClient.FleetsAPI.DeleteFleetsFleetid(context.Background(), fleetId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetsAPIService DeleteFleetsFleetidDevices", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fleetId string

		httpRes, err := apiClient.FleetsAPI.DeleteFleetsFleetidDevices(context.Background(), fleetId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetsAPIService GetFleets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FleetsAPI.GetFleets(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetsAPIService GetFleetsFleetidDevices", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fleetId string

		resp, httpRes, err := apiClient.FleetsAPI.GetFleetsFleetidDevices(context.Background(), fleetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetsAPIService PostFleets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FleetsAPI.PostFleets(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FleetsAPIService PostFleetsFleetidDevices", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var fleetId string

		httpRes, err := apiClient.FleetsAPI.PostFleetsFleetidDevices(context.Background(), fleetId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
