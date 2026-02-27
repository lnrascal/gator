package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lnrascal/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf(" the login handler expects 2 arguments: name and url of the feed")
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		Name:      name,
		Url:       url,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	fmt.Printf("Fetched feed: %s\n", feed)

	followCmd := command{
		Name: "follow",
		Args: []string{url},
	}
	_ = handlerFollow(s, followCmd, user)

	return nil
}

func handlerFeed(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get all feeds: %w", err)
	}

	for _, feed := range feeds {
		url := feed.Url
		user, err := s.db.GetUserById(context.Background(), feed.UserID)

		if err != nil {
			continue
		}

		feedMsg := fmt.Sprintf("User: %s \n Name: %s \n Url: %s", user.Name, feed.Name, url)
		fmt.Println(feedMsg)
	}

	return nil
}
