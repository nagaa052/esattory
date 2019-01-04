package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/nagaa052/esattory/bot"
)

var (
	esaToken    = os.Getenv("ESA_API_TOKEN")
	esaTeamName = os.Getenv("ESA_TEAM_NAME")
	slackToken  = os.Getenv("SLACK_TOKEN")
	postChannel = os.Getenv("SLACK_POST_CHANNEL_ID")
	location    = os.Getenv("LOCATION")
	summaryDays = os.Getenv("SUMMARY_DAYS")
	includeWip  = os.Getenv("INCLUDE_WIP_FLAG")
)

func init() {
	if esaToken == "" {
		log.Fatalln("must be set in ESA_API_TOKEN")
	}

	if esaTeamName == "" {
		log.Fatalln("must be set in ESA_TEAM_NAME")
	}

	if slackToken == "" {
		log.Fatalln("must be set in SLACK_TOKEN")
	}

	if postChannel == "" {
		log.Fatalln("must be set in SLACK_POST_CHANNEL_ID")
	}

	if location == "" {
		location = "Asia/Tokyo"
	}

	if summaryDays == "" {
		summaryDays = "7"
	}

	// time zone fix
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

func main() {
	sumD, err := strconv.Atoi(summaryDays)
	if err != nil {
		log.Fatalln("cast error SUMMARY_DAYS")
	}

	inWip := false
	if includeWip == "1" {
		inWip = true
	}

	b := bot.New(bot.Options{
		EsaToken:         esaToken,
		EsaTeamName:      esaTeamName,
		SlackToken:       slackToken,
		SlackPostChannel: postChannel,
		SummaryDays:      sumD,
		IncludeWip:       inWip,
	})

	b.Run()
}
