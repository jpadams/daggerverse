package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
)

type Trivy struct {}

// Pull the official trivy image.
func (t *Trivy) Base(trivyImageTag string) (*Container) {
	if trivyImageTag == "" {
        trivyImageTag = "latest"
    }
	return dag.Container().From(fmt.Sprintf("aquasec/trivy:%s", trivyImageTag))
}

// Use an image ref for the container image to scan.
func (t *Trivy) ScanImage(
	ctx context.Context,
	imageRef string,
	severity Optional[string],
	exitCode Optional[int],
	format Optional[string],
	trivyImageTag Optional[string]) (string, error) {

	sv := severity.GetOr("UNKNOWN,LOW,MEDIUM,HIGH,CRITICAL")
	ec := exitCode.GetOr(0)
	ft := format.GetOr("table") 
	tag := trivyImageTag.GetOr("latest") 
	return t.Base(tag).
        WithExec([]string{"image", "--quiet", "--severity", sv, "--exit-code", strconv.Itoa(ec), "--format", ft, imageRef}).Stdout(ctx)
}

// Scan a Dagger Container.
func (t *Trivy) ScanContainer(
	ctx context.Context,
	ctr *Container,
	severity Optional[string],
	exitCode Optional[int],
	format Optional[string],
	trivyImageTag Optional[string]) (string, error) {

	sv := severity.GetOr("UNKNOWN,LOW,MEDIUM,HIGH,CRITICAL")
	ec := exitCode.GetOr(0)
	ft := format.GetOr("table") 
	tmp, _ := os.CreateTemp("", "image-trivy-scan-dagger.*.tar")
	tar := tmp.Name()
	success, err := ctr.Export(ctx, tar)
	if success != true || err != nil {
		return "", err	
	}
	tag := trivyImageTag.GetOr("latest") 
	return t.Base(tag).
		WithMountedFile(tar, dag.Host().File(tar)).
		WithExec([]string{"image",  "--quiet", "--severity", sv, "--exit-code", strconv.Itoa(ec), "--format", ft, "--input", tar}).Stdout(ctx)
}
