package main

import (
	"context"
)

type DrupalTest struct {}

func (m *DrupalTest) Run(ctx context.Context) (string, error) {
	drupal := dag.Drupal().Base()
	
	mariaSvc := dag.MariaDb().AsService()
	
	test, err := drupal.
       WithServiceBinding("db", mariaSvc).
        WithEnvVariable("SIMPLETEST_DB", "mysql://user:password@db/drupal").
        WithEnvVariable("SYMFONY_DEPRECATIONS_HELPER", "disabled").
        WithWorkdir("/opt/drupal/web/core").
        WithExec([]string{"../../vendor/bin/phpunit", "-v", "--group", "KernelTests"}).
        Stdout(ctx)

    if err != nil {
        return "", err
    }
    return test, err
}
