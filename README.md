GoBox v0.3.1
============
GoBox is supposed to be something like [BusyBox](http://www.busybox.net). I.e.
a single, preferably small executable which bundles all important shell tools.
A swiss army knife for the command line, if you will.
It is being developed with a focus on [Amazon EC2](http://aws.amazon.com) or as
a small footprint basis for an [OpenVZ](http://www.openvz.org) template.

In order to keep the source code and executable small, I have cut a lot of options
you might be used to from [GNU Coreutils](http://www.gnu.org/software/coreutils/) or
similar. I might even have less options than BusyBox itself. I certainly have
fewer applets right now, and probably ever will. But I consider that a good thing.

Pitfalls
--------
- The shell is *not* a bash, sh or zsh. It is something original, written by me and
  is fairly limited. It does the job of acting as a shell, it‘s hardly adequate for
  scripting, though.
- Telnetd has no authentication mechanism right now. It’s noting more than a
  network-capable pipe.

Installation
------------

### With Modules - Go 1.11 or higher

    git clone https://github.com/surma/gobox ;# clone outside of GOPATH
    cd gobox
    go install

### Without Modules - Before Go 1.11

    go get github.com/surma/gobox

`go get` can also be used with Modules, but it will get you only an immutable copy of the source code.

Developing applets
------------------
- Copy `applets/template` and name the copy like your applet
- Rename `template.go` and edit its contents to fit your applet
- Add your applet to `cmd/gobox/applets.go`

The template provides the basic framework you should stick to.

Why is there not real shell?
----------------------------
I got this question a lot and I have 2 main reasons:

- I seriously did not want to implement the broken and god-awful syntax of bash
  or any other currently used shell!
- You have Go. Do you need anything *more* lightweight? The philosohpy behind this
  project is that it is cheap to (re)build and deploy. You don’t really use
  scripting anymore. If you need to automate some process, write an applet in Go and
  integrate it with GoBox and push it.

Bugs
----
Probably

Contact
-------
If you have ideas for missing applets, found a bug or have a suggestion, use
this [project’s issues](https://github.com/asdf-systems/gobox/issues).
If you want to participate, just fork and code away. For questions contact me:
<surma@surmair.de>

Thanks
------
- Thanks to [Andreas Krennmair](https://github.com/akrennmair) for `grep`, `gzip` and `gunzip`.
- Thanks to [ukaszg](https://github.com/ukaszg) for `head`.
- Thanks to [vbatts](https://github.com/vbatts) for making GoBox `go get`-able.

Credits
-------
(c) 2011-2014 Alexander "Surma" Surma <surma@surmair.de>
