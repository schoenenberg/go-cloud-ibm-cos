package bucketop

import (
	"context"
	"io"

	"gocloud.dev/blob"
)

// ListObjects lists all objects saved in the provided bucket
func ListObjects(
	ctx context.Context,
	bkt *blob.Bucket,
) (*[]blob.ListObject, error) {
	// Init result array of Objects
	objects := make([]blob.ListObject, 0)
	// Do a list request on the bucket
	iter := bkt.List(nil)
	// Iterate through each element of the iterator
	for {
		obj, err := iter.Next(ctx)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		objects = append(objects, *obj)
	}

	return &objects, nil
}
