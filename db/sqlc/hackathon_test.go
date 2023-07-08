package db

import (
	"context"
	"testing"

	"github.com/hackhack-Geek-vol6/backend/util"
	"github.com/stretchr/testify/require"
)

func createHackathonTest(t *testing.T) Hackathons {

	arg := CreateHackathonParams{
		Name:        util.RandomString(8),
		Icon:        []byte(util.RandomString(8)),
		Description: util.RandomString(8),
		Link:        util.RandomString(8),
		Expired:     util.RandomTime(),
		StartDate:   util.RandomTime(),
		Term:        int32(util.Random(100)),
	}

	hackathon, err := testQueries.CreateHackathon(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, hackathon)

	require.Equal(t, arg.HackathonID, hackathon.HackathonID)
	require.Equal(t, arg.Name, hackathon.Name)
	require.Equal(t, arg.Icon, hackathon.Icon)
	require.Equal(t, arg.Description, hackathon.Description)
	require.Equal(t, arg.Link, hackathon.Link)
	require.Equal(t, arg.Expired, hackathon.Expired)
	require.Equal(t, arg.StartDate, hackathon.StartDate)
	require.Equal(t, arg.Term, hackathon.Term)

	require.NotZero(t, hackathon.HackathonID)

	return hackathon
}
