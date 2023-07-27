package reddit

import (
	"fmt"
	"testing"
)

func TestGallery(t *testing.T) {
	cli, err := NewReadonlyClient()

	if err != nil {
		t.Fatalf("%v", err)
	}

	p, _, err := cli.Post.Get(ctx, "15am2k8")
	if err != nil {
		t.Fatalf("%v", err)
	}

	fmt.Printf("p.Post.IsGallery = %+v\n", p.Post.IsGallery)
	fmt.Printf("p.Post.GalleryData.Items = %#v\n", p.Post.GalleryData.Items)
}
