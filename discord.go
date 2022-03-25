package main

import (
	"log"

	"github.com/diamondburned/arikawa/v3/gateway"
)

func handleMessage(c *gateway.MessageCreateEvent) {
	// see if the message matches the pattern
	submatches := memePattern.FindStringSubmatch(c.Content)

	if len(submatches) != 0 {
		// make the text memed
		memeText := memeify(submatches[1])

		// generate a meme and get a link to it
		log.Println(memeText, "from", c.Author.Username+"#"+c.Author.Discriminator)
		memeURL, err := generateMemeURL(memeText)
		if err != nil {
			// ping me if it breaks
			s.SendMessage(c.ChannelID, "The bot encountered an error :(\nLet <@291384073459597313> know")
			log.Println("error getting URL: ", err)
			return
		}

		// reply to the original message with the meme
		s.SendMessageReply(c.ChannelID, memeURL, c.Message.ID)
	}
}
