package main

import (
	"context"
	"embed"
)

type vdflib struct {
	Path                     string
	Label                    string
	ContentID                string
	TotalSize                string
	UpdateCleanBytesTally    string
	TimeLastUpdateCorruption string
	Apps                     []libapp
}

type libapp struct {
	AppID   string
	BuildID string
}

type operation func(ctx context.Context) error

//go:embed blank.mp3
var embeds embed.FS
