---
id: 4
database_id: 414460001
node_id: MDU6SXNzdWU0MTQ0NjAwMDE=
status: open
title: "add streaming"
labels: ["type:feature","help wanted"]
url: https://github.com/octolab/parallel/issues/4
created_at: 2019-02-26T06:58:30Z
updated_at: 2019-02-26T07:03:54Z
---

# add streaming

```go
semaphore wait --stream-to=12345
```

semaphore starts listen 12345 TCP port to stream stdout/stderr of executed commands to it.

can be useful https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html

some inspiration:
- https://github.com/wagoodman/bashful
