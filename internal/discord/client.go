package discord

import "github.com/bwmarrin/discordgo"

func NewClient(token string) (*discordgo.Session, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	return dg, nil
}
