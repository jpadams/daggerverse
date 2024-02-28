package main

type Drupal struct {}

func (m *Drupal) Base() (*Container) {
	return dag.Container().
        From("drupal:10.2.3-php8.3-fpm").
        WithExec([]string{"composer", "require", "drupal/core-dev", "--dev", "--update-with-all-dependencies"})
}
