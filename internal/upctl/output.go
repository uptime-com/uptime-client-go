package upctl

import (
	"errors"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/mattn/go-colorable"
	"github.com/neilotoole/jsoncolor"
)

func output(v any, err error) error {
	if err != nil {
		return err
	}
	color := cmdArgs.Color && jsoncolor.IsColorTerminal(os.Stdout)
	if cmdArgs.Output == "json" && color {
		return outputColorJson(v)
	} else if cmdArgs.Output == "json" {
		return outputJson(v)
	} else if cmdArgs.Output == "spew" {
		return outputSpew(v)
	} else {
		return errors.New("invalid output format")
	}
}

func outputColorJson(v any) error {
	enc := jsoncolor.NewEncoder(colorable.NewColorable(os.Stdout))
	enc.SetIndent("", "  ")
	enc.SetColors(jsoncolor.DefaultColors())
	return enc.Encode(v)
}

func outputJson(v any) error {
	enc := jsoncolor.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}

func outputSpew(v any) error {
	spew.Fdump(os.Stdout, v)
	return nil
}
