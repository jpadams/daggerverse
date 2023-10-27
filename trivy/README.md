### Try me
```sh
dagger call -m github.com/jpadams/daggerverse/trivy scan-image --image-ref alpine:latest

dagger call -m github.com/jpadams/daggerverse/trivy scan-image --severity MEDIUM --image-ref alpine/git:latest

dagger call -m github.com/jpadams/daggerverse/trivy scan-image --severity HIGH,CRITICAL --exit-code 1 --format json --image-ref alpine/git:latest
```
