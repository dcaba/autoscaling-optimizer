# AWS Autoscaling groups optimizer

## Objective

Small toolkit allowing you to optimize your AWS Autoscaling groups 
(ASGs from now on) policies/alarms configuration bearing in mind your
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

## Why?

I'm a bit tired of the "trial and error" method to optimize existing ASGs configurations
without degrading service. Just wanting to demonstrate code can do this for you just applying 
some very basic AI algorithms. 


## What's pending?

WIP project. Everything. 

It will be built bearing in mind a very specific
scenario, so assume your conditions or targets may not be reflected 
in the current code base.
