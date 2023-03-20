package upapi_test

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/uptime-com/uptime-client-go/v2/pkg/upapi"
)

func ExampleAPI_Checks_create() {
	api, err := upapi.New(upapi.WithToken(os.Getenv("UPTIME_TOKEN")))
	if err != nil {
		log.Fatalln("api client initialization failed:", err)
	}

	check, err := api.Checks().Create(context.Background(), upapi.Check{
		Name:      "example",
		CheckType: "HTTP",
		Address:   "https://example.com",
		Locations: []string{
			"US East",
			"US West",
		},
		Interval:      5,
		ContactGroups: []string{"nobody"},
	})
	if err != nil {
		log.Fatalln("api call failed:", err)
	}

	fmt.Println(check.Name, check.CheckType, check.Address)

	// Output:
	// example HTTP https://example.com
}

func ExampleAPI_Checks_list_delete() {
	api, err := upapi.New(upapi.WithToken(os.Getenv("UPTIME_TOKEN")))
	if err != nil {
		log.Fatalln("api client initialization failed:", err)
	}

	checks, err := api.Checks().List(context.Background(), upapi.CheckListOptions{PageSize: 250})
	if err != nil {
		log.Fatalln("api call failed (list):", err)
	}
	for _, check := range checks {
		if check.Name == "example" {
			err = api.Checks().Delete(context.Background(), upapi.PrimaryKey(check.PK))
			if err != nil {
				log.Fatalln("api call failed (delete):", err)
			}
			fmt.Println(check.Name, check.CheckType, check.Address)
			break
		}
	}

	// Output:
	// example HTTP https://example.com
}

func ExampleAPI_Tags_create() {
	api, err := upapi.New(upapi.WithToken(os.Getenv("UPTIME_TOKEN")))
	if err != nil {
		log.Fatalln("api client initialization failed:", err)
	}

	tag, err := api.Tags().Create(context.Background(), upapi.Tag{Tag: "example", ColorHex: "#0000f0"})
	if err != nil {
		log.Fatalln("api call failed:", err)
	}

	fmt.Println(tag.Tag, tag.ColorHex)

	// Output:
	// example #0000f0
}

func ExampleAPI_Tags_list_delete() {
	api, err := upapi.New(upapi.WithToken(os.Getenv("UPTIME_TOKEN")))
	if err != nil {
		log.Fatalln("api client initialization failed:", err)
	}

	tags, err := api.Tags().List(context.Background(), upapi.TagListOptions{})
	if err != nil {
		log.Fatalln("api call failed (list):", err)
	}
	for _, tag := range tags {
		if tag.Tag == "example" {
			err = api.Tags().Delete(context.Background(), upapi.PrimaryKey(tag.PK))
			if err != nil {
				log.Fatalln("api call failed (delete):", err)
			}
			fmt.Println(tag.Tag, tag.ColorHex)
			break
		}
	}

	// Output:
	// example #0000f0
}

func ExampleAPI_Outages_list() {
	api, err := upapi.New(upapi.WithToken(os.Getenv("UPTIME_TOKEN")))
	if err != nil {
		log.Fatalln("api client initialization failed:", err)
	}

	outages, err := api.Outages().List(context.Background(), upapi.OutageListOptions{})
	for _, outage := range outages {
		fmt.Println(outage.PK, outage.CheckName)
	}

	// Output:
}
