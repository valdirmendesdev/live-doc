package entities_test

import (
	"github.com/stretchr/testify/require"
	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/entities"
	"testing"
)

func createIntegrationSystem() entities.IntegrationSystem {
	return entities.NewIntegrationSystem(entities.SAP, entities.EccHttp)
}

func Test_NewIntegrationSystem(t *testing.T) {
	is := createIntegrationSystem()
	require.NotNil(t, is)
	require.Equal(t, entities.SAP, is.SystemType)
	require.Equal(t, entities.EccHttp, is.CommunicationType)
}

func Test_IntegrationSystemIsValid(t *testing.T) {
	is := createIntegrationSystem()
	is.SystemType = ""
	is.CommunicationType = ""

	isValid, err := is.IsValid()
	require.False(t, isValid)
	require.NotNil(t, err)
	require.EqualError(t, err, "system type must be filled")

	is.SystemType = "test"
	isValid, err = is.IsValid()
	require.False(t, isValid)
	require.EqualError(t, err, "system type is not valid")

	is.SystemType = entities.SAP
	isValid, err = is.IsValid()
	require.False(t, isValid)
	require.EqualError(t, err, "communication type must be filled")

	is.CommunicationType = "test"
	isValid, err = is.IsValid()
	require.False(t, isValid)
	require.EqualError(t, err, "communication type is not valid")

	is.CommunicationType = entities.EccHttp
	isValid, err = is.IsValid()
	require.True(t, isValid)
	require.Nil(t, err)
}