package main

import (
	"flag"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Flags struct {
	*flag.FlagSet
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewFlags(name string, args []string) (*Flags, error) {
	// Create flags
	tmpFlags := new(Flags)
	tmpFlags.FlagSet = flag.NewFlagSet(name, flag.ContinueOnError)

	// Register flags
	tmpFlags.String("key", "", "Azure resourceKey")
	tmpFlags.String("endpoint", "https://api.cognitive.microsofttranslator.com", "Azure endpoint")
	tmpFlags.String("region", "eastus", "Azure region")

	tmpFlags.String("file", "test.srt", "srt file path")
	tmpFlags.String("from", "en", "srt language from")
	tmpFlags.String("to", "zh", "srt language to")

	// Parse flags
	if err := tmpFlags.Parse(args); err != nil {
		return nil, err
	}
	return tmpFlags, nil
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (f *Flags) GetKey() string {
	return f.Lookup("key").Value.String()
}

func (f *Flags) GetEndpoint() string {
	return f.Lookup("endpoint").Value.String()
}

func (f *Flags) GetRegion() string {
	return f.Lookup("region").Value.String()
}

func (f *Flags) GetFile() string {
	return f.Lookup("file").Value.String()
}

func (f *Flags) GetFrom() string {
	return f.Lookup("from").Value.String()
}

func (f *Flags) GetTo() string {
	return f.Lookup("to").Value.String()
}

///////////////////////////////////////////////////////////////////////////////
