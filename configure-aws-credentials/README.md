# Module: "Configure AWS Credentials" Action for GitHub Actions

![dagger-min-version](https://img.shields.io/badge/dagger%20version-v0.9.1-green)

Configures AWS credentials for use in subsequent steps in a GitHub Action workflow

This module is automatically generated using [actions-generator](https://github.com/aweris/gale/tree/main/daggerverse/actions/generator). It is a Dagger-compatible adaptation of the original [aws-actions/configure-aws-credentials](https://github.com/aws-actions/configure-aws-credentials) action.

## How to Use

Run the following command run this action:

```shell
dagger call -m <module-path> run [flags]
```

Replace `<module-path>` with the local path or a git repo reference to the module

## Flags

### Action Inputs

| Name | Required | Description | Default | 
| ------| ------| ------| ------| 
| --with-role-to-assume | false | The Amazon Resource Name (ARN) of the role to assume. Use the provided credentials to assume an IAM role and configure the Actions environment with the assumed role credentials rather than with the provided credentials. |  |
| --with-audience | false | The audience to use for the OIDC provider | sts.amazonaws.com |
| --with-mask-aws-account-id | false | Whether to mask the AWS account ID for these credentials as a secret value. By default the account ID will not be masked |  |
| --with-role-external-id | false | The external ID of the role to assume. |  |
| --with-role-skip-session-tagging | false | Skip session tagging during role assumption |  |
| --with-aws-session-token | false | AWS Session Token. |  |
| --with-disable-retry | false | Whether to disable the retry and backoff mechanism when the assume role call fails. By default the retry mechanism is enabled |  |
| --with-aws-access-key-id | false | AWS Access Key ID. Provide this key if you want to assume a role using access keys rather than a web identity token. |  |
| --with-web-identity-token-file | false | Use the web identity token file from the provided file system path in order to assume an IAM role using a web identity, e.g. from within an Amazon EKS worker node. |  |
| --with-http-proxy | false | Proxy to use for the AWS SDK agent |  |
| --with-role-session-name | false | Role session name (default: GitHubActions) |  |
| --with-managed-session-policies | false | Define a list of managed session policies to use when assuming a role |  |
| --with-unset-current-credentials | false | Whether to unset the existing credentials in your runner. May be useful if you run this action multiple times in the same job |  |
| --with-retry-max-attempts | false | The maximum number of attempts it will attempt to retry the assume role call. By default it will retry 12 times |  |
| --with-special-characters-workaround | false | Some environments do not support special characters in AWS_SECRET_ACCESS_KEY. This option will retry fetching credentials until the secret access key does not contain special characters. This option overrides disable-retry and retry-max-attempts. This option is disabled by default |  |
| --with-aws-region | true | AWS Region, e.g. us-east-2 |  |
| --with-aws-secret-access-key | false | AWS Secret Access Key. Required if aws-access-key-id is provided. |  |
| --with-role-chaining | false | Use existing credentials from the environment to assume a new role, rather than providing credentials as input. |  |
| --with-role-duration-seconds | false | Role duration in seconds. Default is one hour. |  |
| --with-inline-session-policy | false | Define an inline session policy to use when assuming a role |  |
| --with-output-credentials | false | Whether to set credentials as step output |  |


### Action Runtime Inputs

| Flag | Required | Description | 
| ------| ------| ------| 
| --token | Optional | GitHub token is optional for running the action. However, be aware that certain custom actions may require a token and could fail if it's not provided. |
| --source | Conditional | The directory containing the repository source. Either `--source` or `--repo` must be provided; `--source` takes precedence. |
| --repo | Conditional | The name of the repository (owner/name). Either `--source` or `--repo` must be provided; `--source` takes precedence. |
| --tag | Conditional | Tag name to check out. Only works with `--repo`. Either `--tag` or `--branch` must be provided; `--tag` takes precedence. |
| --branch | Conditional | Branch name to check out. Only works with `--repo`. Either `--tag` or `--branch` must be provided; `--tag` takes precedence. |
| --runner-image | Optional | Image to use for the runner. |
| --runner-debug | Optional | Enables debug mode. |
