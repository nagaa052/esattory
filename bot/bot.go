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
}

// Options ...
type Options struct {
	EsaToken         string
	EsaTeamName      string
	SlackToken       string
	SlackPostChannel string
	SummaryDays      int
}

// New ...
func New(opt Options) (bot *Bot) {
	b := &Bot{
		esaClient:        esa.NewClient(opt.EsaToken, opt.EsaTeamName),
		slackClient:      slack.New(opt.SlackToken),
		slackPostChannel: opt.SlackPostChannel,
		summaryDays:      opt.SummaryDays,
	}

	return b
}

func (bot *Bot) getHotentry(page uint, perPage uint) ([]esa.Post, error) {
	res, err := bot.esaClient.ListPosts(
		context.Background(),
		esa.ListPostsParam{
			Q:     "updated:>" + getAgoDays(bot.summaryDays),
			Sort:  "best_match",
			Order: esa.DESC,
		},
		page,
		perPage,
	)

	return res.Posts, err
}

func (bot *Bot) postSlack(posts []esa.Post) (string, string, string, error) {
	attachments := bot.toSlackAttachments(posts)
	return bot.slackClient.SendMessage(
		bot.slackPostChannel,
		slack.MsgOptionText("Hottent esa post", false),
		slack.MsgOptionAttachments(attachments...))
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

func getAgoDays(days int) string {
	return time.Now().Add(time.Duration(-24*days) * time.Hour).Format("2006-01-02")
}

// Run ...
func (bot *Bot) Run() {
	log.Println("run bot")
	posts, err := bot.getHotentry(1, 10)
	if err != nil {
		log.Fatal(err)
		return
	}
	_, _, _, error := bot.postSlack(posts)
	if error != nil {
		log.Fatalf("Expected error: channel_not_found; instead succeeded")
		return
	}
}
