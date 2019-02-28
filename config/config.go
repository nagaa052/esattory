package config

import (
	"github.com/kelseyhightower/envconfig"
)

type configration struct {
	EsaToken           string `envconfig:"ESA_API_TOKEN" required:"true"`
	EsaTeamName        string `envconfig:"ESA_TEAM_NAME" required:"true"`
	SlackToken         string `envconfig:"SLACK_TOKEN" required:"true"`
	SlackPostChannelID string `envconfig:"SLACK_POST_CHANNEL_ID" required:"true"`
	SummaryDays        int    `envconfig:"SUMMARY_DAYS" default:"7"`
}

var conf = &configration{}

func init() {
	envconfig.Process("", conf)
}

// Reload ...
func Reload() {
	envconfig.Process("", conf)
}

// EsaToken ...
func EsaToken() string {
	return conf.EsaToken
}

// EsaTeamName ...
func EsaTeamName() string {
	return conf.EsaTeamName
}

// SlackToken ...
func SlackToken() string {
	return conf.SlackToken
}

// SlackPostChannelID ...
func SlackPostChannelID() string {
	return conf.SlackPostChannelID
}

// SummaryDays ...
func SummaryDays() int {
	return conf.SummaryDays
}
