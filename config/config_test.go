package config

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	esaToken           = "XXXXX"
	esaTeamName        = "HogeHoge"
	slackToken         = "XXX_XXX"
	slackPostChannelID = "123123"
	summaryDays        = 10
)

func setup() {
	os.Setenv("ESA_API_TOKEN", esaToken)
	os.Setenv("ESA_TEAM_NAME", esaTeamName)
	os.Setenv("SLACK_TOKEN", slackToken)
	os.Setenv("SLACK_POST_CHANNEL_ID", slackPostChannelID)
	os.Setenv("SUMMARY_DAYS", strconv.Itoa(summaryDays))
}

func teardown() {}

func TestMain(m *testing.M) {
	setup()
	Reload()
	code := m.Run()
	if code == 0 {
		teardown()
	}
	os.Exit(code)
}

func TestEsaToken(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esaToken, EsaToken())
}

func TestEsaTeamName(t *testing.T) {
	t.Parallel()

	assert.Equal(t, esaTeamName, EsaTeamName())
}

func TestSlackToken(t *testing.T) {
	t.Parallel()

	assert.Equal(t, slackToken, SlackToken())
}

func TestSlackPostChannelID(t *testing.T) {
	t.Parallel()

	assert.Equal(t, slackPostChannelID, SlackPostChannelID())
}

func TestSummaryDays(t *testing.T) {
	t.Parallel()

	assert.Equal(t, summaryDays, SummaryDays())
}
