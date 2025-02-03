## README:
A GO script to automate the process of sending emails on Saturdays. The script uses net/smtp libray in-order to send emails with required configuration. This script along with Jenkin configuration is designed to send automated emails every Saturday around 8 AM.

# Pre-requistes:
- Go version 1.22.4
- Local jenkins setup

# Jenkin job configuration:

- In order to send email following properties needs to be provided,
    - Email_sender
    - Email_recipients
    - Email_message
    - Email_server_configurations

**Please NOTE:** 
1. That the properties needs to be provided with values as described in the description of the jenkins file. Ensure to use double quotes while providing the values.
Example: **"value"**
2. Input **'Email_recipients'** and **'Email_server_configurations'** which takes multiple values needs to be separated with ','.
3. The default value for **'Email_server_configurations'** (for gmail specific) and **'Email_message'** is provided. It can be modified to the desired values.

# Unit testing:
- In order to add unit tests please use *main_test.go* file
- To run unit test use,
`go test` 

# References:

- [smtp](https://pkg.go.dev/net/smtp)
- [jenkin_job](https://www.jenkins.io/doc/book/pipeline/)
