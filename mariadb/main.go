package main

import "maria-db/internal/dagger"

type MariaDb struct{}

func (m *MariaDb) AsService() *dagger.Service {
	return dag.Container().
		From("mariadb:10.11.4").
		WithEnvVariable("MARIADB_USER", "user").
		WithEnvVariable("MARIADB_PASSWORD", "password").
		WithEnvVariable("MARIADB_DATABASE", "drupal").
		WithEnvVariable("MARIADB_ROOT_PASSWORD", "root").
		WithExposedPort(3306).
		AsService()
}
