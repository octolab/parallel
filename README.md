> # 🚦 cmd/semaphore
>
> Tool to execute terminal commands in parallel.

[![Awesome][icon_awesome]][page_awesome]

## 💡 Idea

```bash
$ semaphore create 2
$ semaphore add -- docker build
$ semaphore add -- vagrant up
$ semaphore add -- ansible-playbook
$ semaphore wait --timeout=1m --notify
```

Full description of the idea is available
[here](https://www.notion.so/octolab/semaphore-7d5ebf715d0141d1a8fa045c7966be3b?r=0b753cbf767346f5a6fd51194829a2f3).

## 🏆 Motivation

...

## 🤼‍♂️ How to

[![asciicast](https://asciinema.org/a/135943.png)](https://asciinema.org/a/135943)

```
Usage: semaphore COMMAND

Semaphore provides functionality to execute terminal commands in parallel.

Commands:

create	is a command to init a semaphore context
  -debug
    	show error stack trace
  -filename string
    	an absolute path to semaphore context (default "/tmp/semaphore.json")


add	is a command to add a job into a semaphore context
  -debug
    	show error stack trace
  -edit
    	switch to edit mode to read arguments from input (not implemented yet)
  -filename string
    	an absolute path to semaphore context (default "/tmp/semaphore.json")


wait	is a command to execute a semaphore task
  -debug
    	show error stack trace
  -filename string
    	an absolute path to semaphore context (default "/tmp/semaphore.json")
  -notify
    	show notification at the end (not implemented yet)
  -speed int
    	a velocity of report output (characters per second)
  -timeout duration
    	timeout for task execution (default 1m0s)

Version 4.0.0 (commit: ..., build date: ..., go version: go1.9, compiler: gc, platform: darwin/amd64)
```

### Complex example

```bash
$ semaphore create 2
$ semaphore add -- bash -c "cd /tmp; \
    git clone git@github.com:kamilsk/semaphore.git \
    && cd semaphore \
    && echo 'semaphore at revision' \$(git rev-parse HEAD) \
    && rm -rf /tmp/semaphore"
$ semaphore add -- bash -c "cd /tmp; \
    git clone git@github.com:kamilsk/retry.git \
    && cd retry \
    && echo 'retry at revision' \$(git rev-parse HEAD) \
    && rm -rf /tmp/retry"
$ semaphore wait
```

## 🧩 Installation

### Homebrew

```bash
$ brew install kamilsk/tap/semaphore
```

### Binary

```bash
$ export REQ_VER=4.0.0  # all available versions are on https://github.com/kamilsk/semaphore/releases
$ export REQ_OS=Linux   # macOS and Windows are also available
$ export REQ_ARCH=64bit # 32bit is also available
$ curl -sL -o semaphore.tar.gz \
       https://github.com/kamilsk/semaphore/releases/download/"${REQ_VER}/semaphore_${REQ_VER}_${REQ_OS}-${REQ_ARCH}".tar.gz
$ tar xf semaphore.tar.gz -C "${GOPATH}"/bin/ && rm semaphore.tar.gz
```

### From source code

```bash
$ egg github.com/kamilsk/semaphore@^4.0.0 -- make test install
$ # or use mirror
$ egg bitbucket.org/kamilsk/semaphore@^4.0.0 -- make test install
```

> [egg][page_egg]<sup id="anchor-egg">[1](#egg)</sup> is an `extended go get`.

### Bash and Zsh completions

```bash
$ semaphore completion bash > /path/to/bash_completion.d/semaphore.sh
$ semaphore completion zsh  > /path/to/zsh-completions/_semaphore.zsh
```

<sup id="egg">1</sup> The project is still in prototyping. [↩](#anchor-egg)

---

made with ❤️ for everyone

[icon_awesome]:    https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg
[icon_build]:      https://travis-ci.org/kamilsk/semaphore.cli.svg?branch=master

[page_awesome]:    https://github.com/avelino/awesome-go#goroutines
[page_build]:      https://travis-ci.org/kamilsk/semaphore.cli
[page_promo]:      https://github.com/kamilsk/semaphore.cli
[page_egg]:        https://github.com/kamilsk/egg
