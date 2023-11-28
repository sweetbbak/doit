<p></p>
<p align="center">
  <img src="doit-once.png" />
</p>

## What? and Why?
a setuid binary that allows you to run a single command as root before it deletes itself
idk why I thought'd it be funny tho. I was thinking about security and stuff and this
came to mind.

Then I was wondering if there was a "disposable" sudo binary that was a one time use, or
one that could give you very very limited commands that are able to run. Mainly because
working around user space operations that need root are just annoying.

maybe you could make a custom binary with `just` and allow-list a single command, like
`chroot` or something, let a process use it for just what it needs and then it silently
deletes itself. Not practical probably but I kinda like the idea of a `disposable` and
limited scope binary that you constantly generate and use as needed.

Im also curious if you could download a `tar` ball with setuid binaries in it and have
root on that machine. Probably not right? but also why not? 

## Examples
```bash
  ./doit whoami # root
  ls            # . .. --> binary deleted itself lol
```

## Installation
Build from source
```bash
go build
chown root:root doit
chmod u+s doit
# or
just
```

## Features
doit as root, but only once lol
