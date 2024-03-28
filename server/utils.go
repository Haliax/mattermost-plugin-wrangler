package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

func makePostLink(siteURL, teamName, postID string) string {
	return fmt.Sprintf("%s/%s/pl/%s", siteURL, teamName, postID)
}

func makeBotDM(base, newPostLink, executor string) string {
	message := cleanMessageJSON(base)
	message = strings.Replace(message, "{executor}", executor, -1)
	message = strings.Replace(message, "{postLink}", newPostLink, -1)

	return message
}

func cleanPost(post *model.Post) {
	post.Id = ""
	post.CreateAt = 0
	post.UpdateAt = 0
	post.EditAt = 0
}

func cleanPostID(post *model.Post) {
	post.Id = ""
}

func cleanAndTrimMessage(message string, trimLength int) string {
	return trimMessage(cleanMessage(message), trimLength)
}

func cleanMessage(message string) string {
	// Remove any leading whitespace and header markdown.
	message = strings.TrimLeft(message, " ")
	message = strings.TrimLeft(message, "#")
	message = strings.TrimLeft(message, " ")

	// Remove all code block markdown.
	message = strings.Replace(message, "```", "", -1)

	// Replace all newlines to keep summary condensed.
	message = strings.Replace(message, "\n", " | ", -1)

	return message
}

func cleanMessageJSON(message string) string {
	message = strings.TrimLeft(message, " ")
	message = strings.ReplaceAll(message, "\\n", "\n")
	return message
}

func trimMessage(message string, trimLength int) string {
	if len(message) <= trimLength {
		return message
	}

	return fmt.Sprintf("%s...", message[:trimLength])
}

func prettyPrintJSON(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}

func jsonCodeBlock(in string) string {
	return fmt.Sprintf("``` json\n%s\n```", in)
}

func codeBlock(in string) string {
	return fmt.Sprintf("```\n%s\n```", in)
}

func quoteBlock(in string) string {
	return fmt.Sprintf("> %s", in)
}

func inlineCode(in string) string {
	return fmt.Sprintf("`%s`", in)
}

// NewBool returns a pointer to a given bool.
func NewBool(b bool) *bool { return &b }

// NewInt returns a pointer to a given int.
func NewInt(n int) *int { return &n }

// NewInt32 returns a pointer to a given int32.
func NewInt32(n int32) *int32 { return &n }

// NewInt64 returns a pointer to a given int64.
func NewInt64(n int64) *int64 { return &n }

// NewString returns a pointer to a given string.
func NewString(s string) *string { return &s }

func isInputMessageLink(input string) bool {
	match, _ := regexp.MatchString("(http|https)://[a-zA-Z0-9\\-_]+(:\\d+)?/([a-zA-Z0-9\\-_]+)/pl/[a-zA-Z0-9]{26}", input)
	return match
}

func getMessageIDFromLink(inputLink string) string {
	regex := regexp.MustCompile(`(http|https)://[a-zA-Z0-9\-_]+(:\d+)?/([a-zA-Z0-9\-_]+)/pl/([a-zA-Z0-9]{26})`)
	matches := regex.FindStringSubmatch(inputLink)
	if len(matches) >= 5 {
		return matches[4]
	}
	return ""
}

func cleanInputID(input string) string {
	if isInputMessageLink(input) {
		return getMessageIDFromLink(input)
	}
	return input
}
