package bot

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/hori-ryota/esa-go/esa"
	"github.com/nlopes/slack"
)

const location = "Asia/Tokyo"

// Bot ...
type Bot struct {
	esaClient        esa.Client
	slackClient      *slack.Client
	slackPostChannel string
	summaryDays      int
	includeWip       bool
}

// Options ...
type Options struct {
	EsaToken         string
	EsaTeamName      string
	SlackToken       string
	SlackPostChannel string
	SummaryDays      int
	IncludeWip       bool
}

// New ...
func New(opt Options) (bot *Bot) {
	b := &Bot{
		esaClient:        esa.NewClient(opt.EsaToken, opt.EsaTeamName),
		slackClient:      slack.New(opt.SlackToken),
		slackPostChannel: opt.SlackPostChannel,
		summaryDays:      opt.SummaryDays,
		includeWip:       opt.IncludeWip,
	}
	return b
}

func (bot *Bot) getHotentry() (*esa.PostsResp, error) {
	return bot.esaClient.ListPosts(
		context.Background(),
		esa.ListPostsParam{
			Q:     bot.getHotentryQuery(),
			Sort:  esa.ListPostsParamSortBestMatch,
			Order: esa.DESC,
		},
		1,
		10,
	)
}

func (bot *Bot) getHotentryQuery() string {
	q := "updated:>" + getAgoDaysFormat(bot.summaryDays)
	if !bot.includeWip {
		q += " wip:false"
	}
	return q
}

func (bot *Bot) postSlack(posts []esa.Post) (string, string, string, error) {
	attachments := bot.toSlackAttachments(posts)
	return bot.slackClient.SendMessage(
		bot.slackPostChannel,
		slack.MsgOptionText(bot.slackMessageTitle(), false),
		slack.MsgOptionAttachments(attachments...))
}

func (bot *Bot) slackMessageTitle() string {
	return "Hottest esa posts(Updated at : " +
		getAgoDaysFormat(bot.summaryDays) +
		" ~ " +
		getAgoDaysFormat(0) +
		")"
}

func (bot *Bot) toSlackAttachments(posts []esa.Post) []slack.Attachment {
	var attachments = make([]slack.Attachment, 0, len(posts))
	for i := range posts {
		post := posts[i]
		log.Printf("UpdatedAt:%s, Title: %s, star: %d, watch: %d, URL: %s", post.UpdatedAt, post.Name, *post.StargazersCount, *post.WatchersCount, post.URL)

		attachments = append(attachments, slack.Attachment{
			Color:      "good",
			Title:      post.FullName,
			TitleLink:  post.URL,
			AuthorName: post.UpdatedBy.ScreenName,
			AuthorIcon: post.UpdatedBy.Icon,
			Footer:     fmt.Sprintf(":star:%d  :eyes: %d  :speech_balloon: %d", *post.StargazersCount, *post.WatchersCount, *post.CommentsCount),
			Ts:         json.Number(strconv.FormatInt(post.UpdatedAt.Unix(), 10)),
		})
	}
	return attachments
}

func getAgoDaysFormat(days int) string {
	return time.Now().Add(time.Duration(-24*days) * time.Hour).Format("2006-01-02")
}

// Run ...
func (bot *Bot) Run() {
	log.Println("run bot")
	res, err := bot.getHotentry()
	if err != nil {
		log.Fatal(err)
		return
	}
	_, _, _, error := bot.postSlack(res.Posts)
	if error != nil {
		log.Fatalf("Expected error: channel_not_found; instead succeeded")
		return
	}
}
