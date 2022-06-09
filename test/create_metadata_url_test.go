package test

import (
	"fmt"
	"testing"

	"github.com/deezdegens/solago/metadata"
)

func TestCreateMetadataURL(t *testing.T) {
	metadata := metadata.MetadataOffChain{
		Image:       "https://images.unsplash.com/photo-1529778873920-4da4926a72c2?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8Y3V0ZSUyMGNhdHxlbnwwfHwwfHw%3D&w=1000&q=80",
		Description: "What a cute kitty",
		Properties: metadata.Properties{
			Category: "image",
			Files: []metadata.File{
				{
					URI:  "https://images.unsplash.com/photo-1529778873920-4da4926a72c2?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8Y3V0ZSUyMGNhdHxlbnwwfHwwfHw%3D&w=1000&q=80",
					Type: "image/jpg",
				},
			},
		},
		Attributes: []metadata.Attribute{
			{
				TraitType: "Silly Goose",
				Value:     "You KNOW IT",
			},
			{
				TraitType: "Piggly Wiggly",
				Value:     "DUH!!!",
			},
		},
	}
	fmt.Println(metadata.CreateURL())
}
