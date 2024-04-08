# SendTeamsNotifications Package

The `sendteamsnotifications` package is a Go library designed to facilitate sending custom notifications to Microsoft Teams channels using incoming webhooks. It allows users to craft messages with a summary, activity titles, subtitles, and key-value fact pairs to deliver rich and informative updates directly to Teams.

## Installation

To install the `sendteamsnotifications` package, you can use `go get`:

```
go get -u github.com/domenicomastrangelo/sendteamsnotifications
```

## Usage
To use the sendteamsnotifications package in your Go application, you need to follow these steps:

### Import the Package:
```
import "github.com/yourusername/sendteamsnotifications"
```

### Initialize with Webhook URL:
You need to initialize the SendTeamsNotification struct with the webhook URL of your Microsoft Teams channel.

```
notifier := sendteamsnotifications.New("YOUR_WEBHOOK_URL")
```

### Send a Notification:
Create and send a notification by specifying a summary, activity title, activity subtitle, and any number of facts.

```
notifier.Send(
    "Notification Summary",
    "Activity Title",
    "Activity Subtitle",
    map[string]string{"Fact Name": "Fact Value"},
)
```

## Example
Here's a full example that sends a notification to a Microsoft Teams channel:

```
package main

import (
    "github.com/yourusername/sendteamsnotifications"
)

func main() {
    notifier := sendteamsnotifications.New("YOUR_WEBHOOK_URL")
    notifier.Send(
        "Deployment Notification",
        "Deployed Successfully",
        "Backend Service v1.2",
        map[string]string{"Environment": "Production", "Version": "1.2"},
    )
}
```

# Contributing
Contributions to the sendteamsnotifications package are welcome! Here's how you can contribute:

- Report Issues: If you find a bug or have a suggestion for an improvement, please open an issue.
- Submit Pull Requests: Feel free to fork the repository and submit pull requests with bug fixes or new features. Please make sure to write clear commit messages and provide a description of your changes.

Thank you for considering contributing to the sendteamsnotifications package.
