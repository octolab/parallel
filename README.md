> # 🚦 semaphore
>
> Tool to execute terminal commands in parallel.

[![Build][build.icon]][build.page]
[![Template][template.icon]][template.page]

## 💡 Idea

```bash
$ semaphore create 2
$ semaphore add -- docker build
$ semaphore add -- vagrant up
$ semaphore add -- ansible-playbook
$ semaphore wait --timeout=1m --notify
```

Full description of the idea is available [here][design.page].

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
$ curl -sSL https://bit.ly/install-semaphore | sh
# or
$ wget -qO- https://bit.ly/install-semaphore | sh
```

### Source

```bash
# use standard go tools
$ go get -u github.com/kamilsk/semaphore
# or use egg tool
$ egg tools add github.com/kamilsk/semaphore
```

> [egg][egg.page]<sup id="anchor-egg">[1](#egg)</sup> is an `extended go get`.

### Bash and Zsh completions

```bash
$ semaphore completion bash > /path/to/bash_completion.d/semaphore.sh
$ semaphore completion zsh  > /path/to/zsh-completions/_semaphore.zsh
```

<sup id="egg">1</sup> The project is still in prototyping.[↩](#anchor-egg)

---

made with ❤️ for everyone

[build.icon]:       https://travis-ci.org/kamilsk/semaphore.cli.svg?branch=master
[build.page]:       https://travis-ci.org/kamilsk/semaphore.cli

[design.page]:      https://www.notion.so/octolab/semaphore-7d5ebf715d0141d1a8fa045c7966be3b?r=0b753cbf767346f5a6fd51194829a2f3

[promo.page]:       https://github.com/kamilsk/semaphore.cli

[template.page]:    https://github.com/octomation/go-tool
[template.icon]:    https://img.shields.io/badge/template-go--tool-blue

[egg.page]:         https://github.com/kamilsk/egg
