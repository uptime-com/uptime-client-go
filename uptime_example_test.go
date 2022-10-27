package uptime_test

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/uptime-com/uptime-client-go"
)

func ExampleCheckService_Create() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	cfg := uptime.Config{
		Token:            os.Getenv("UPTIME_TOKEN"),
		RateMilliseconds: 2000,
	}
	api, err := uptime.NewClient(&cfg)
	if err != nil {
		panic(err)
	}

	_, res, err := api.Checks.Create(ctx, &uptime.Check{
		CheckType:     "HTTP",
		Address:       "https://uptime.com",
		Interval:      1,
		Threshold:     15,
		Locations:     []string{"US East", "US West"},
		ContactGroups: []string{"examples"}, // must exist
		Tags:          []string{"examples"}, // must exist
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Status)

	// Output: 200 OK
}

func ExampleCheckService_Delete() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	cfg := uptime.Config{
		Token:            os.Getenv("UPTIME_TOKEN"),
		RateMilliseconds: 2000,
	}
	api, err := uptime.NewClient(&cfg)
	if err != nil {
		panic(err)
	}

	checks, res, err := api.Checks.List(ctx, &uptime.CheckListOptions{
		Tag: []string{"examples"},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Status)

	for _, check := range checks {
		res, err := api.Checks.Delete(ctx, check.PK)
		if err != nil {
			panic(err)
		}
		fmt.Println(res.Status)
	}

	// Output:
	// 200 OK
	// 200 OK
}
