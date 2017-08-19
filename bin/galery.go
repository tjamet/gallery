package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/docopt/docopt-go"
	"github.com/tjamet/gallery"
	"github.com/tjamet/gallery/algolia"
	"github.com/tjamet/gallery/filters"
	"github.com/tjamet/gallery/imgix/client"
	"github.com/tjamet/gallery/imgix/url"
	"github.com/tjamet/gallery/metadata"
	"github.com/tjamet/gallery/s3"
	"github.com/tjamet/gallery/visiter"
)

func main() {
	help := `
Usage:
	galery [options]
Options:
	--region=<region>  The region in which the s3 bucket where to upload files [default: eu-west-1]
	--bucket=<bucket>  The name of the s3 bucket where to upload files
	--src=<src>        The path to load files from
	--shards=<shards>  The number of shards file uploads should be spread on [default: 3]
	--dry              Print files to restore rather than actually restoring them
	`
	args, err := docopt.Parse(help, os.Args[1:], true, "0.0.0", false)
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to parse commandline: %s\n", err.Error())
		os.Exit(1)
	}

	shards, err := strconv.Atoi(args["--shards"].(string))
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to parse shard number: %s\n", err.Error())
		os.Exit(1)
	}
	if args["--bucket"] == nil || args["--src"] == nil {
		fmt.Println(help)
		os.Exit(1)
	}

	s3Builder := s3.With().SharedConfigEnable().
		Bucket(args["--bucket"].(string)).
		Region(args["--region"].(string))

	imgixClient := client.With().New()

	urlBuilder := url.With().Builder(imgixClient).Update(url.HighQuality, url.Jpeg, url.NoUpscale)

	uploader := galery.Pipe{
		UploaderBuilder: s3Builder,
		SearchClient:    algolia.With().Index("gallery"),
		BuildURLs: func(path string) map[string]string {
			return map[string]string{
				"small":  urlBuilder.Clone().Update(url.Small).ForImage(path),
				"medium": urlBuilder.Clone().Update(url.Medium).ForImage(path),
				"large":  urlBuilder.Clone().Update(url.Large).ForImage(path),
			}
		},
	}

	if args["--dry"].(bool) {
		uploader.UploaderBuilder = &s3.DryBuilder{}
		uploader.SearchClient = &algolia.DryRun{}
	}

	ds := visiter.New(args["--src"].(string), shards)
	ds = ds.Filter(filters.IsFile).Map(visiter.GetPath)
	ds = ds.Map(metadata.FromFile).Map(metadata.Load).Filter(filters.HasNoError)
	ds = ds.Map(uploader.Entry).Map(uploader.Upload).Filter(filters.HasNoError)
	ds = ds.Map(uploader.Index)
	ds.Run()
}
