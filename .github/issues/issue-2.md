---
id: 2
database_id: 414459327
node_id: MDU6SXNzdWU0MTQ0NTkzMjc=
status: open
title: "remove FlagSet"
labels: ["type:refactoring","help wanted"]
url: https://github.com/octolab/parallel/issues/2
created_at: 2019-02-26T06:55:57Z
updated_at: 2019-02-26T07:03:54Z
---

# remove FlagSet

- remove the method from interface https://github.com/kamilsk/semaphore/blob/1e58f96d738b1c6c0a130c7764cac3c67b70ea89/cmd/semaphore/cli.go#L28-L29
- extend the method to accept arguments https://github.com/kamilsk/semaphore/blob/1e58f96d738b1c6c0a130c7764cac3c67b70ea89/cmd/semaphore/cli.go#L35
- extend the method to return parsed arguments https://github.com/kamilsk/semaphore/blob/1e58f96d738b1c6c0a130c7764cac3c67b70ea89/cmd/semaphore/cli.go#L42
- these parts https://github.com/kamilsk/semaphore/blob/1e58f96d738b1c6c0a130c7764cac3c67b70ea89/cmd/semaphore/cli.go#L80 and https://github.com/kamilsk/semaphore/blob/1e58f96d738b1c6c0a130c7764cac3c67b70ea89/cmd/semaphore/cli.go#L105 should be removed
- these parts
```go
	if c.Flags == nil {
		c.Flags = c.BaseCommand.FlagSet(c.CmdName)
		c.Flags.BoolVar(&c.Edit, "edit", false, "switch to edit mode to read arguments from input (not implemented yet)")
	}
	return c.Flags
```
should be refactored to stateless
