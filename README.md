# Datadog CLI

A humble beginning. PR's welcome.

I admittedly didn't look too hard, but a quick search yielded clients with
lots of features but no ability to downtime based on tag. I tend to (ab)use
labels, so this lets me wrap automation and pipelines around those to schedule
and cancel downtime.

## Setup

Only configuration is the usual `DD_*` env vars.

```console
❯ cat .envrc
export DD_SITE="datadoghq.com"
export DD_API_KEY="your-api-key"
export DD_APP_KEY="your-appkey"
export DD_USERNAME="you@example.com"
```

## Usage

Currently only manages downtime:

```console
❯ ./dd downtime -h
List & modify downtime

Usage:
  dd downtime [command]

Available Commands:
  cancel      cancel downtime
  list        list downtime
  schedule    schedule downtime

Flags:
  -h, --help   help for downtime

Use "dd downtime [command] --help" for more information about a command.

❯ ./dd downtime schedule -m "cli test" -t env:prod,testing:true
{
  "active": true,
  "canceled": null,
  "creator_id": 1556103,
  "disabled": false,
  "downtime_type": 0,
  "end": null,
  "id": 1516563902,
  "message": "cli test",
  "monitor_id": null,
  "monitor_tags": [
    "env:prod",
    "testing:true"
  ],
  "parent_id": null,
  "recurrence": null,
  "scope": [
    "*"
  ],
  "start": 1634873065,
  "timezone": "UTC",
  "updater_id": null
}

❯ ./dd downtime cancel -t env:prod,testing:true
cancelled downtime 1516563902
```

## References

- https://docs.datadoghq.com/api/latest/downtimes
- https://github.com/DataDog/datadog-api-client-go
