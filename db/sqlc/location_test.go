package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vukieuhaihoa/go-weather-api/util"
)

func createRandomLocation(t *testing.T) Location {
	arg := CreateLocationParams{
		Name:      util.RandomNameOfCity(),
		Longitude: util.RandomPosition(),
		Latitude:  util.RandomPosition(),
	}

	location, err := testQueries.CreateLocation(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, arg.Name, location.Name)
	require.Equal(t, arg.Longitude, location.Longitude)
	require.Equal(t, arg.Latitude, location.Latitude)
	return location
}

func TestCreateLocation(t *testing.T) {
	createRandomLocation(t)
}

func TestGetLocation(t *testing.T) {
	location1 := createRandomLocation(t)
	location2, err := testQueries.GetLocation(context.Background(), location1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, location2)

	require.Equal(t, location1.ID, location2.ID)
	require.Equal(t, location1.Name, location2.Name)
	require.Equal(t, location1.Longitude, location2.Longitude)
	require.Equal(t, location1.Latitude, location2.Latitude)
	require.Equal(t, location1.Count, location2.Count)
}

func TestUpdateLocation(t *testing.T) {
	location1 := createRandomLocation(t)

	location2, err := testQueries.UpdateLocation(context.Background(), location1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, location2)

	require.Equal(t, location1.Count+1, location2.Count)
}

func TestDeleteLocation(t *testing.T) {
	location1 := createRandomLocation(t)

	err := testQueries.DeleteLocation(context.Background(), location1.ID)
	require.NoError(t, err)

	location2, err := testQueries.GetLocation(context.Background(), location1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())

	require.Empty(t, location2)

}
