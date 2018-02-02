# autoscaling-optimizer - AWS Autoscaling groups optimizer

## Description

Small toolkit/pieces of code facilitating the optimization of AWS Autoscaling groups 
(ASGs from now on) policies/alarms configuration, bearing in mind your
 current ASGs metrics and response times targets
 
## How?

* Fetching data from Cloudwatch
* Allowing you to define your perf/response time targets
* Provide an engine that, given a policy/alarm configuration, your 
perf/response time targets, and Cloudwatch data describing your workload for
a period of time, confirms the configuration does not affect 
service and the resulting cost
* Genetic algorithm that executes several simulations changing policies/alarms 
configurations 

## What's the status?

**WIP project**. Everything is pending. 

It will be built bearing in mind a very specific
scenario, so assume your conditions or targets may not be reflected 
in the current code base.

# Usage

## Install

To install, use `go get`:

```bash
$ go get -d github.com/dcaba/autoscaling-optimizer
```

## Execution

```bash
% go run *.go
Usage: autoscaling-optimizer [--version] [--help] <command> [<args>]

Available commands are:
    combine     
    fetch       
    simulate    
    version     Print autoscaling-optimizer version and quit

exit status 127
```

# Extra project data

## Why?

I'm a bit tired of the "trial and error" method to optimize existing ASGs configurations
without degrading service. Just wanting to demonstrate code can do this for you just applying 
some very basic AI algorithms. 

## Contribution

1. Fork ([https://github.com/dcaba/autoscaling-optimizer/fork](https://github.com/dcaba/autoscaling-optimizer/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[dcaba](https://github.com/dcaba)
