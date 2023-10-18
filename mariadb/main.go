package main

type MariaDb struct {}

func (m *MariaDb) AsService() (*Service) {
	return dag.Container().
        From("mariadb:10.11.2").
        WithEnvVariable("MARIADB_USER", "user").
        WithEnvVariable("MARIADB_PASSWORD", "password").
        WithEnvVariable("MARIADB_DATABASE", "drupal").
        WithEnvVariable("MARIADB_ROOT_PASSWORD", "root").
        WithExposedPort(3306).
        AsService()
}
