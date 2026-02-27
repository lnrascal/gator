package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/lnrascal/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2 // default

	if len(cmd.Args) > 1 {
		return fmt.Errorf("usage: %s [limit]", cmd.Name)
	}

	if len(cmd.Args) == 1 {
		parsedLimit, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("invalid limit")
		}
		limit = parsedLimit
	}

	posts, err := s.db.GetPostsForUser(
		context.Background(),
		database.GetPostsForUserParams{
			UserID: user.ID,
			Limit:  int32(limit),
		},
	)
	if err != nil {
		return fmt.Errorf("couldn't get posts: %w", err)
	}

	if len(posts) == 0 {
		fmt.Println("No posts found.")
		return nil
	}

	for _, post := range posts {
		fmt.Println("===================================")
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("URL: %s\n", post.Url)
		fmt.Printf("Description: %s\n", post.Description)
		fmt.Printf("Published: %s\n", post.PublishedAt.Time.Format(time.RFC1123))
	}

	return nil
}
