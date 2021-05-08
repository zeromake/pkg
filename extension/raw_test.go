package extension

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/testutil"
	"testing"
)

func TestRaw(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			NewRawExtension(RawOption{
				IsBlock:    true,
				Tag:        "code",
				ClassName:  "raw-block",
				PrefixMark: "$$",
				SuffixMark: "$$",
				Trigger:    "$",
			}),
			NewRawExtension(RawOption{
				Tag:        "code",
				ClassName:  "raw-inline",
				PrefixMark: "$",
				SuffixMark: "$",
				Trigger:    "$",
			}),
		),
	)
	testutil.DoTestCaseFile(
		markdown,
		"_test/raw.txt",
		t,
		testutil.ParseCliCaseArg()...,
	)
}

func TestRawKeepMark(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			NewRawExtension(RawOption{
				IsBlock:    true,
				KeepMark:   true,
				Tag:        "code",
				ClassName:  "raw-block",
				PrefixMark: "$$",
				SuffixMark: "$$",
				Trigger:    "$",
			}),
			NewRawExtension(RawOption{
				KeepMark:   true,
				Tag:        "code",
				ClassName:  "raw-inline",
				PrefixMark: "$",
				SuffixMark: "$",
				Trigger:    "$",
			}),
		),
	)
	testutil.DoTestCaseFile(
		markdown,
		"_test/raw_keep_mark.txt",
		t,
		testutil.ParseCliCaseArg()...,
	)
}
