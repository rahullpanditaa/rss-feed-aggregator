package rss

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/rahullpanditaa/rssfeedaggregator/internal/cli"
	"github.com/rahullpanditaa/rssfeedaggregator/internal/database"
)

func ScrapeFeeds(s *cli.State) {
	// get the next feed to fetch
	feed, err := s.DbQueries.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not get feed to fetch: %v\n", err)
		os.Exit(1)
	}

	// mark it as fetched
	err = s.DbQueries.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not mark feed as fetched: %v\n", err)
		os.Exit(1)
	}

	// fetch feed using feed url
	feedStruct, err := FetchFeed(context.Background(), feed.Url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "coud not fetch feed: %v\n", err)
		return
	}

	// iterate over ITEMS in feed
	// for _, item := range feedStruct.Channel.Item {
	// 	fmt.Printf("Item title: %s\n", item.Title)
	// }

	// instead of printing titles, save posts to db
	for _, post := range feedStruct.Channel.Item {
		postPublicationdate, err := time.Parse(time.RFC1123, post.PubDate)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to convert given post's (%s) pubDate %v\n", post.Title, post.PubDate)
			continue
		}
		err = s.DbQueries.CreatePost(
			context.Background(),
			database.CreatePostParams{
				ID:          uuid.New(),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				Title:       post.Title,
				Url:         post.Link,
				Description: sql.NullString{String: post.Description, Valid: true},
				PublishedAt: sql.NullTime{Time: postPublicationdate, Valid: false},
				FeedID:      feed.ID,
			},
		)
		if err != nil {
			var pgErr *pq.Error
			if errors.As(err, &pgErr) && pgErr.Code == "23505" {
				fmt.Printf("URL: %s - already exists\n", post.Link)
				continue
			} else {
				fmt.Fprintf(os.Stderr, "could not save post: %v", err)
				continue
			}
		}
	}
}
