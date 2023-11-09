package bot

import (
	"context"
	"funnygomusic/databaser"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/meilisearch/meilisearch-go"
	"github.com/shkh/lastfm-go/lastfm"
	"gorm.io/gorm"
	"os"
)

type ComData int

type Botter struct {
	*state.State
	VoiceSes   *VoiceSessionHndlr
	Queue      *QueueManager
	Ctx        context.Context
	MyId       discord.UserID
	MyUsername string
	AllowList  []string
	SubChan    discord.ChannelID
	Db         *gorm.DB
	Meili      *meilisearch.Client
	LastFmApi  *lastfm.Api
}

func NewBotter(s *state.State, ctx *context.Context) *Botter {
	b := &Botter{
		State:     s,
		Ctx:       *ctx,
		AllowList: []string{},
		VoiceSes:  &VoiceSessionHndlr{},
		Db:        databaser.NewDatabase(),
		Meili:     databaser.NewMeili(),
		LastFmApi: lastfm.New(os.Getenv("LASTFM_API_KEY"), os.Getenv("LASTFM_API_SECRET")),
	}
	b.MeiliUpdate()
	b.Queue = NewQueueManager(b)
	return b

}
