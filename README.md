# create-aws-profile
`create-aws-profile` is a command line tool designed with the intent to have the output of `aws sts assume-role` piped out to it. It takes this output and creates a new set of AWS credentials in `~/.aws/credentials` under the profile name specified with the `--profile` flag. More instructions follow below.


## How to use this

### Build Only
Creates `bin/create-aws-profile_<os>_<arch>`.

`./build.sh`
### Build and Install
Adding the `--install` flag, as shown below, will install the binary to `/usr/local/bin`.
`./build.sh --install`

### How to use
Call the standard AWS CLI and pipe the output into `create-aws-profile`. Be sure to add a profile name with `--profile`. This is the profile name that you'll need to use to access the AWS CLI.

`aws sts assume-role --role-arn arn:aws:iam::<account>:role/OrganizationAccountAccessRole --role-session-name build-cicd | create-aws-profile --profile <local role name>`

Next, you can specify the `AWS_PROFILE` you want your command to use, as shown below:
`AWS_PROFILE=<local role name> aws s3 ls`
