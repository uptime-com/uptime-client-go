package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/uptime-com/uptime-client-go"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := run(ctx); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg := uptime.Config{
		Token: os.Getenv("UPTIME_TOKEN"),
	}
	api, err := uptime.NewClient(&cfg)
	if err != nil {
		return err
	}

	opts := uptime.OutageListOptions{
		Ordering: "-created_at",
	}

	outages, _, err := api.Outages.List(ctx, &opts)
	if err != nil {
		return err
	}

	for _, item := range outages {
		fmt.Println(
			item.CreatedAt.Format(time.RFC3339),
			item.ResolvedAt.Format(time.RFC3339),
			time.Duration(item.DurationSecs)*time.Second,
		)
	}

	return nil
}
