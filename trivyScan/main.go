package main

import (
	"context"
	"fmt"
)

type Trivy struct {}

func (m *Trivy) Base(ctx context.Context) (*Container, error) {
	return dag.Container().From("aquasec/trivy:latest").Sync(ctx)
}

func (c *Container) TrivyScan(ctx context.Context) (string, error) {
	success, err := c.Export(ctx, "/tmp/image.tar")
	if success != true || err != nil {
		fmt.Println("failure!")
	}
	trivy, err := (&Trivy{}).Base(ctx)
	if err != nil {
		panic(err)
	}
	return trivy.
		WithMountedFile("/opt/trivy/image.tar", dag.Host().File("/tmp/image.tar")).
		WithExec([]string{"image", "-q", "-f", "json", "--input", "/opt/trivy/image.tar" }).Stdout(ctx)
}


