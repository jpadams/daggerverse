package main

import (
	"context"
	"golang.org/x/mod/modfile"
)

type GoVersion struct {}

func (m *GoVersion) Version(ctx context.Context, srcDir *Directory) (string) {
    goMod, err := srcDir.File("go.mod").Contents(ctx)
    if err != nil {
        panic(err)
    }
    f, err := modfile.Parse("go.mod", []byte(goMod), nil)
    if err != nil {
        panic(err)
    }
    return f.Go.Version
}

