# assume-role
`assume-role` is a command line tool designed with the intent to have the output of `aws sts assume-role` piped out to it. It takes this output and creates a new set of AWS credentials in `~/.aws/credentials` under the profile name specified with the `--profile` flag. More instructions follow below.


## How to use this

### Build Only
Creates `bin/assume-role_<os>_<arch>`.

`./build.sh`
### Build and Install
`./build.sh --install`

### How to use
Call the standard AWS CLI and pipe the output into `assume-role`. Be sure to add a profile name with `--profile`. This is the profile name that you'll need to use to access the AWS CLI.

`aws sts assume-role --role-arn arn:aws:iam::<account>:role/OrganizationAccountAccessRole --role-session-name build-cicd | assume-role --profile <local role name>`