package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hewenyu/translate/auzer"
	"github.com/hewenyu/translate/srt"
)

func main() {
	// flags
	flags, err := NewFlags(filepath.Base(os.Args[0]), os.Args[1:])
	if err != nil {
		if err != flag.ErrHelp {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		} else {
			os.Exit(0)
		}
	}

	// init translator
	translateFactor := auzer.NewTranslator(&auzer.AuzerConfig{
		ResourceKey: flags.GetKey(),
		Endpoint:    flags.GetEndpoint(),
		Region:      flags.GetRegion(),
	})
	srtTranslator := srt.NewSrtTranslator(flags.GetFile(), flags.GetFrom(), flags.GetTo(), translateFactor, true)

	err = srtTranslator.Translate(context.Background())
	if err != nil {
		if err != flag.ErrHelp {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		} else {
			os.Exit(0)
		}
	}
}
