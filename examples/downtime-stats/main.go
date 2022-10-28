package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/uptime-com/uptime-client-go"
)

func main() {

	var pk uint
	flag.UintVar(&pk, "pk", 0, "check pk")

	flag.Parse()

	if pk == 0 {
		flag.Usage()
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := run(ctx, int(pk)); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, pk int) error {
	cfg := uptime.Config{
		Token: os.Getenv("UPTIME_TOKEN"),
	}
	api, err := uptime.NewClient(&cfg)
	if err != nil {
		return err
	}

	var (
		end   = time.Now().UTC().Truncate(time.Hour * 24)
		start = end.Add(-time.Hour * 24 * 30)
	)

	opts := &uptime.CheckStatsOptions{
		StartDate:              start.Format(time.RFC3339),
		EndDate:                end.Format(time.RFC3339),
		Location:               "US East",
		LocationsResponseTimes: true,
		IncludeAlerts:          true,
		Download:               false,
		PDF:                    false,
	}

	stats, _, err := api.Checks.Stats(ctx, pk, opts)
	if err != nil {
		return err
	}

	fmt.Println("PK:       ", pk)
	fmt.Println("Range:    ", start.Format("2006-01-02"), end.Format("2006-01-02"))
	fmt.Println("Outages:  ", stats.Totals.Outages)
	fmt.Println("Downtime: ", time.Duration(stats.Totals.DowntimeSecs)*time.Second)
	fmt.Println("")

	for _, item := range stats.Statistics {
		if item.Outages == 0 {
			continue
		}
		fmt.Println(item.Date, item.Outages, time.Duration(item.DowntimeSecs)*time.Second)
	}

	return nil
}
