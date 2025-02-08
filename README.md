# create-aws-profile
[![Build](https://github.com/JoshuaSchlichting/create-aws-profile/actions/workflows/build.yml/badge.svg?branch=master)](https://github.com/JoshuaSchlichting/create-aws-profile/actions/workflows/build.yml)[![Unit Tests](https://github.com/JoshuaSchlichting/create-aws-profile/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/JoshuaSchlichting/create-aws-profile/actions/workflows/test.yml)

`create-aws-profile` is a command line tool designed for managing AWS profiles within the `~/.aws/credentials` file. It operates by having the output of `aws sts assume-role` piped into the program, like so: <pre><code>aws sts assume-role --role-arn arn:aws:iam::000000000000:role/role-name --role-session-name AWSCLI-Session <b>| create-aws-profile --profile desired-profile-name</b></code></pre>

The above command will create or modify the `~/.aws/credentials` file and add the requested role under the `--profile` name passed to it. Any existing records for the `--profile` will be overwritten in `~/.aws/credentials`.


## How to use
Call the standard AWS CLI and pipe the output into `create-aws-profile`. Be sure to add a profile name with `--profile`. This is the profile name that you'll need to use to access the AWS CLI.

`aws sts assume-role --role-arn arn:aws:iam::<account>:role/OrganizationAccountAccessRole --role-session-name build-cicd | create-aws-profile --profile <local role name>`
Next, you can specify the `AWS_PROFILE` you want your command to use, as shown below:
`AWS_PROFILE=<local role name> aws s3 ls`
> NOTE: Existing profiles with the same `--profile` name will be overwritten.

## Install from GitHub
### macOS
`curl -o "/usr/local/bin/create-aws-profile" -L "https://github.com/JoshuaSchlichting/create-aws-profile/releases/download/v0.1.1/create-aws-profile_macos_x86_64" && chmod +x "/usr/local/bin/create-aws-profile/create-aws-profile"`

### Building from source
create `bin/create-aws-profile_<os>_<arch>` by executing `./build.sh`
> Adding the `--install` flag, as shown below, will install the binary to `/usr/local/bin`.
`./build.sh --install`



