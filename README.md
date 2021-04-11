# slackbot
slackbot 

# Creating token

If you are getting this error: "RTM API ERror: not_allowed_token_type" and your token starts with "xoxb", maybe your token was created using the new granular access roles.

In order to create the right token type, please create using a classic app from the following URL: https://api.slack.com/apps?new_classic_app=1

More info @ https://github.com/slackapi/node-slack-sdk/issues/953

There is an issue opened about this API key working on only classic apps https://github.com/shomali11/slacker/issues/63
