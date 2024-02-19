// Finds vulnerabilities from container image ref or Dagger Container

package main

import (
	"context"
	"fmt"
	"strconv"
)

type Trivy struct {}

// Return a Container from the official trivy image.
func (t *Trivy) Base(
	// +optional
	// +default=latest
	trivyImageTag string,
) (*Container) {
	return dag.Container().
		From(fmt.Sprintf("aquasec/trivy:%s", trivyImageTag)).
		WithMountedCache("/root/.cache/trivy", dag.CacheVolume("trivy-db-cache"))
}

// Scan an image ref.
func (t *Trivy) ScanImage(
	ctx context.Context,
	imageRef string,
	// +optional
	// +default=UNKNOWN,LOW,MEDIUM,HIGH,CRITICAL
	severity string,
	// +optional
	// +default=0
	exitCode int,
	// +optional
	// +default=table
	format string,
	// +optional
	// +default=latest
	trivyImageTag string,
) (string, error) {
	return t.Base(trivyImageTag).
        WithExec([]string{"image", "--quiet", "--severity", severity, "--exit-code", strconv.Itoa(exitCode), "--format", format, imageRef}).Stdout(ctx)
}

// Scan a Dagger Container.
func (t *Trivy) ScanContainer(
	ctx context.Context,
	ctr *Container,
	// +optional
	// +default=UNKNOWN,LOW,MEDIUM,HIGH,CRITICAL
	severity string,
	// +optional
	// +default=0
	exitCode int,
	// +optional
	// +default=table
	format string,
	// +optional
	// +default=latest
	trivyImageTag string,
) (string, error) {
	ref, _ := ctr.ImageRef(ctx)
	return t.Base(trivyImageTag).
		WithMountedFile("/scan/"+ref, ctr.AsTarball()).
		WithExec([]string{"image",  "--quiet", "--severity", severity, "--exit-code", strconv.Itoa(exitCode), "--format", format, "--input", "/scan/"+ref}).Stdout(ctx)
}
