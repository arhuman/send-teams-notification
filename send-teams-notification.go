package sendteamsnotifications

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
)

var (
	slogger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	}))
	webhook_url = ""
)

type SendTeamsNotification struct {
	WebhookURL string
}

type TeamsMessage struct {
	Type       string    `json:"@type"`
	Context    string    `json:"@context"`
	ThemeColor string    `json:"themeColor"`
	Summary    string    `json:"summary"`
	Sections   []Section `json:"sections"`
}

type Section struct {
	ActivityTitle    string `json:"activityTitle"`
	ActivitySubtitle string `json:"activitySubtitle"`
	ActivityImage    string `json:"activityImage"`
	Facts            []Fact `json:"facts"`
}

type Fact struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func New(webhookURL string) *SendTeamsNotification {
	return &SendTeamsNotification{
		WebhookURL: webhookURL,
	}
}

func (s *SendTeamsNotification) Send(summary string, activity_title string, activity_subtitle string, facts ...map[string]string) {
	factsList := []Fact{}
	for _, fact := range facts {
		for k, v := range fact {
			factsList = append(factsList, Fact{
				Name:  k,
				Value: v,
			})
		}
	}

	// Send a notification to Microsoft Teams
	teamsMessage := TeamsMessage{
		Type:       "MessageCard",
		Context:    "http://schema.org/extensions",
		ThemeColor: "0072C6",
		Summary:    "Query Killer",
		Sections: []Section{
			{
				ActivityTitle:    activity_title,
				ActivitySubtitle: activity_subtitle,
				Facts:            factsList,
			},
		},
	}

	// Send the message to the webhook
	client := &http.Client{}
	jsonValue, err := json.Marshal(teamsMessage)
	if err != nil {
		slogger.Error(err.Error())
		return
	}

	req, err := http.NewRequest("POST", webhook_url, bytes.NewBuffer(jsonValue))
	if err != nil {
		slogger.Error(err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		slogger.Error(err.Error())
		return
	}
	defer resp.Body.Close()
}
