package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateLocation(t *testing.T) {
	arg := CreateLocationParams{
		Name:      "Sai Gon",
		Longitude: 123,
		Latitude:  456,
	}

	location, err := testQueries.CreateLocation(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, arg.Name, location.Name)
	require.Equal(t, arg.Longitude, location.Longitude)
	require.Equal(t, arg.Latitude, location.Latitude)

}
