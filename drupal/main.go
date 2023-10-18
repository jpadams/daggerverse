package main

type Drupal struct {}

func (m *Drupal) Base() (*Container) {
	return dag.Container().
        From("drupal:10.0.7-php8.2-fpm").
        WithExec([]string{"composer", "require", "drupal/core-dev", "--dev", "--update-with-all-dependencies"})
}
