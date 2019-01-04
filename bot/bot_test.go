package bot

import (
	"testing"
)

var (
	esaToken    = "123123"
	esaTeamName = "hogehoge"
	slackToken  = "xoxb-123123"
	postChannel = "ABCABC"
	summaryDays = 7
	bot         = New(Options{
		EsaToken:         esaToken,
		EsaTeamName:      esaTeamName,
		SlackToken:       slackToken,
		SlackPostChannel: postChannel,
		SummaryDays:      summaryDays,
	})
)

func TestBot_GetHotentry(t *testing.T) {
	// TODO assert api call
}

func TestBot_PostSlack(t *testing.T) {
	// TODO assert api call
}

func TestBot_ToSlackAttachments(t *testing.T) {
	// TODO assert Attachment list
}
