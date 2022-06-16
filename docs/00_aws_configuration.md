## AWS CLI configuration

1. edit your AWS config file:

  - in Mac OS and Linux it's `~/.aws/config`
  - In Windows it's `%UserProfile%\.aws\config`

2. add following contents to this file

    ```ini
    [profile workshop-base]
    aws_access_key_id = <key provided by Stepan Vrany>
    aws_secret_access_key = <key provided by Stepan Vrany>
    region = eu-west-1

    [profile workshop-student]
    source_profile = workshop-base
    role_arn = arn:aws:iam::314595822951:role/eks_access_eu-west-1_student
    region = eu-west-1
    ```

3. set `AWS_PROFILE` environment variable to this path

    ```bash
    export AWS_PROFILE=workshop-student
    ```

    or

    ```powershell
    $env:AWS_PROFILE = "workshop-student"
    ```

4. verify your setup

    ```bash
    aws sts get-caller-identity
    ```

    You should get a JSON object similar to this one:

    ```json
    {
        "UserId": "AROAUSP27JVTXH66QKWFQ:botocore-session-1652684483",
        "Account": "314595822951",
        "Arn": "arn:aws:sts::314595822951:assumed-role/eks_access_eu-west-1_student/botocore-session-1652684483"
    }
    ```
