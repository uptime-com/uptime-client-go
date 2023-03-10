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
		Token:            os.Getenv("UPTIME_TOKEN"),
		RateMilliseconds: 2000,
	}
	api, err := uptime.NewClient(&cfg)
	if err != nil {
		return err
	}

	check, res, err := api.Checks.Create(ctx, &uptime.Check{
		CheckType:     "WHOIS",
		Address:       "example.com",
		Interval:      5,
		Threshold:     15,
		Locations:     []string{"US East", "US West"},
		ContactGroups: []string{"examples"}, // must exist
		Tags:          []string{"examples"}, // must exist
	})
	if err != nil {
		return err
	}

	fmt.Println(res.Status, "PK:", check.PK)
	return nil
}
