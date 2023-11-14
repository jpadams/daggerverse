// Code generated by actions-generator. DO NOT EDIT.

package main

// Dagger module for executing the aws-actions/configure-aws-credentials GitHub Action.
type ConfigureAwsCredentials struct{}

// Runs the aws-actions/configure-aws-credentials GitHub Action.
func (m ConfigureAwsCredentials) Run(
	// Define a list of managed session policies to use when assuming a role
	withManagedSessionPolicies Optional[string],
	// Whether to unset the existing credentials in your runner. May be useful if you run this action multiple times in the same job
	withUnsetCurrentCredentials Optional[string],
	// Whether to disable the retry and backoff mechanism when the assume role call fails. By default the retry mechanism is enabled
	withDisableRetry Optional[string],
	// AWS Access Key ID. Provide this key if you want to assume a role using access keys rather than a web identity token.
	withAwsAccessKeyId Optional[string],
	// Use the web identity token file from the provided file system path in order to assume an IAM role using a web identity, e.g. from within an Amazon EKS worker node.
	withWebIdentityTokenFile Optional[string],
	// Proxy to use for the AWS SDK agent
	withHttpProxy Optional[string],
	// Role session name (default: GitHubActions)
	withRoleSessionName Optional[string],
	// Define an inline session policy to use when assuming a role
	withInlineSessionPolicy Optional[string],
	// Whether to set credentials as step output
	withOutputCredentials Optional[string],
	// The maximum number of attempts it will attempt to retry the assume role call. By default it will retry 12 times
	withRetryMaxAttempts Optional[string],
	// Some environments do not support special characters in AWS_SECRET_ACCESS_KEY. This option will retry fetching credentials until the secret access key does not contain special characters. This option overrides disable-retry and retry-max-attempts. This option is disabled by default
	withSpecialCharactersWorkaround Optional[string],
	// AWS Region, e.g. us-east-2
	withAwsRegion string,
	// AWS Secret Access Key. Required if aws-access-key-id is provided.
	withAwsSecretAccessKey Optional[string],
	// Use existing credentials from the environment to assume a new role, rather than providing credentials as input.
	withRoleChaining Optional[string],
	// Role duration in seconds. Default is one hour.
	withRoleDurationSeconds Optional[string],
	// Skip session tagging during role assumption
	withRoleSkipSessionTagging Optional[string],
	// The Amazon Resource Name (ARN) of the role to assume. Use the provided credentials to assume an IAM role and configure the Actions environment with the assumed role credentials rather than with the provided credentials.
	withRoleToAssume Optional[string],
	// The audience to use for the OIDC provider
	withAudience Optional[string],
	// Whether to mask the AWS account ID for these credentials as a secret value. By default the account ID will not be masked
	withMaskAwsAccountId Optional[string],
	// The external ID of the role to assume.
	withRoleExternalId Optional[string],
	// AWS Session Token.
	withAwsSessionToken Optional[string],
	// Directory containing the repository source. Takes precedence over `--repo`.
	source Optional[*Directory],
	// Repository name, format: owner/name. Takes precedence over `--source`.
	repo Optional[string],
	// Tag name to check out. Only works with `--repo`. Takes precedence over `--branch`.
	tag Optional[string],
	// Branch name to check out. Only works with `--repo`.
	branch Optional[string],
	// Image for the runner.
	runnerImage Optional[string],
	// Enables debug mode.
	runnerDebug Optional[bool],
	// GitHub token. May be required for certain actions.
	token Optional[*Secret],
) *Container {
	// initializing runtime options
	opts := ActionsRuntimeRunOpts{
		Branch:      branch.GetOr(""),
		Repo:        repo.GetOr(""),
		RunnerDebug: runnerDebug.GetOr(false),
		RunnerImage: runnerImage.GetOr(""),
		Source:      source.GetOr(nil),
		Tag:         tag.GetOr(""),
		Token:       token.GetOr(nil),
	}

	return dag.ActionsRuntime().
		Run("aws-actions/configure-aws-credentials@v4", opts).
		WithInput("role-to-assume", withRoleToAssume.GetOr("")).
		WithInput("audience", withAudience.GetOr("sts.amazonaws.com")).
		WithInput("mask-aws-account-id", withMaskAwsAccountId.GetOr("")).
		WithInput("role-external-id", withRoleExternalId.GetOr("")).
		WithInput("role-skip-session-tagging", withRoleSkipSessionTagging.GetOr("")).
		WithInput("aws-session-token", withAwsSessionToken.GetOr("")).
		WithInput("aws-access-key-id", withAwsAccessKeyId.GetOr("")).
		WithInput("web-identity-token-file", withWebIdentityTokenFile.GetOr("")).
		WithInput("http-proxy", withHttpProxy.GetOr("")).
		WithInput("role-session-name", withRoleSessionName.GetOr("")).
		WithInput("managed-session-policies", withManagedSessionPolicies.GetOr("")).
		WithInput("unset-current-credentials", withUnsetCurrentCredentials.GetOr("")).
		WithInput("disable-retry", withDisableRetry.GetOr("")).
		WithInput("aws-region", withAwsRegion).
		WithInput("aws-secret-access-key", withAwsSecretAccessKey.GetOr("")).
		WithInput("role-chaining", withRoleChaining.GetOr("")).
		WithInput("role-duration-seconds", withRoleDurationSeconds.GetOr("")).
		WithInput("inline-session-policy", withInlineSessionPolicy.GetOr("")).
		WithInput("output-credentials", withOutputCredentials.GetOr("")).
		WithInput("retry-max-attempts", withRetryMaxAttempts.GetOr("")).
		WithInput("special-characters-workaround", withSpecialCharactersWorkaround.GetOr("")).
		Sync()
}
