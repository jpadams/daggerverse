package main

import (
	"errors"
	"fmt"
	"context"
)

type Staticcheck struct {
    BaseImage *Container
}

// optional for user to call, otherwise default is used, as set in Check
func (s *Staticcheck) WithBaseImage(ctx context.Context, ctr *Container) (*Staticcheck) {
	_, err := ctr.WithExec([]string{"go", "version"}).Sync(ctx)
	// `go version` throws error
	if err != nil {
		panic(errors.New("Base image must have golang: " + err.Error()))
		//return s, errors.New("Base image must have golang: " + err.Error())
	}
    s.BaseImage = ctr
    return s
}

func (s *Staticcheck) defaultBaseImage(ctx context.Context, srcDir *Directory) error {
	v, err := dag.GoVersion().Version(ctx, srcDir)
        if err != nil {
            return err
        }
        s.BaseImage, err = dag.Container().
            From("golang:"+v).
            Sync(ctx)
        if err != nil {
            return err
        }
	return nil
}

func (s *Staticcheck) installStaticcheck(ctx context.Context) error {
	_, err := s.BaseImage.WithExec([]string{"staticcheck", "-version"}).Sync(ctx)
	// `staticcheck -version` throws error
	if err != nil {
		fmt.Println("Base image needs staticcheck installed. Installing...")
 		ctr, err := s.BaseImage.
        	WithExec([]string{"go", "install", "honnef.co/go/tools/cmd/staticcheck@latest"}).
        	Sync(ctx)
		if err != nil {
			return errors.New("Could not install staticcheck: " + err.Error())
		}
		s.BaseImage = ctr
		return nil
	}
	return nil
}

// Run staticcheck on the given directory, error if it fails
// (error will show up in progress output for user)
func (s *Staticcheck) Check(ctx context.Context, srcDir *Directory) error {
    if s.BaseImage == nil {
        // set a default; if you want to be fancy, parse the go.mod from the given srcDir and figure out their go version
		err := s.defaultBaseImage(ctx, srcDir)
		if err != nil {
			return err
		}
    }
	err := s.installStaticcheck(ctx)  
	if err != nil {
		return err
	}
    _, err = s.BaseImage.
		WithDirectory("/src", srcDir).
        WithWorkdir("/src").
		//WithExec(append([]string{"staticcheck"}, args...)).
		WithExec([]string{"staticcheck"}).
		Sync(ctx)
    return err
}

func (s *Staticcheck) Test() (error) {
	err := s.Check(context.Background(), dag.Host().Directory("."))
	return err
}

func (s *Staticcheck) Test2() (error) {
	ctx := context.Background()
	return s.WithBaseImage(ctx, dag.Container().From("golang:1.20")).Check(ctx, dag.Host().Directory("."))
}

func (s *Staticcheck) Test3() (error) {
	ctx := context.Background()
	return s.WithBaseImage(ctx, dag.Container().From("alpine:latest")).Check(ctx, dag.Host().Directory("."))
}
// Run staticcheck on the given directory, return the result whether or not it fails
//func (s *Staticcheck) CheckResult(ctx context.Context, srcDir *Directory) (*Result, error) {
//    if s.BaseImage == nil {
//        // set a default; if you want to be fancy, parse the go.mod from the given srcDir and figure out their go version
//		err := s.defaultBaseImage(ctx, srcDir)
//		if err != nil {
//			panic(err)
//		}
//    }

   // output, err := s.BaseImage.
	//	WithExec(...).
//		Stdout(ctx)
    // this should work, but pretty sure it doesn't right now; making issue!
//    var execErr *ExecError
//    if errors.As(err, &execErr) {
//        return &Result{
//            Output: execErr.Stderr,
//            Passed: false,
//        }
//    }
//    return &Result{Output: output, Passed: true}, nil
//}

type Result struct {
    Output string
    Passed bool
}
