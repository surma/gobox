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

The current development status can be seen [here](https://trello.com/board/gobox/4ed265f07e5ffd00002b0aed)

Pitfalls
--------
- The shell is *not* a bash, sh or zsh. It is something original, written by me and
  is fairly limited. It does the job of acting as a shell, it‘s hardly adequate for
  scripting, though.
- Telnetd has no authentication mechanism right now. It’s noting more than a
  network-capable pipe.

Installation
------------

### Using gb
For development, I recommend using [`gb`](http://code.google.com/p/go-gb/).
The hassle of updating makefiles and the dependencies just vanish.

    .../gobox $ gb -g

### Not using gb
If you just want to build GoBox, make sure you have `make` available and run:

 	.../gobox $ ./build goinstall && ./build

 **For both scenarios:** I recommend working within the development environment (“*DevEnv*”)
 which is provided via [Vagrant](http://www.vagrantup.com). After installing `vagrant`

    .../gobox $ gem install vagrant

set up the DevEnv with

    .../gobox $ vagrant up

and enter it

    .../gobox $ vagrant ssh

Developing applets
------------------
- Copy `applets/template` and name the copy like your applet
- Rename `template.go` and edit its contents to fit your applet
- Add your applet to `cmd/gobox/applets.go`

The template provides the basic framework you should stick to that.

Why is there not real shell?
----------------------------
I got this question a lot and I have 2 main reasons:

- I seriously did not want to implement the broken and god-awful syntax of bash
  or any other currently used shell!
- You have Go. Do you need anything *more* lightweight? The philosohpy behind this
  project is, that it is cheap to (re)build and deploy. So you don’t really use
  scripting anymore. If you need to automate some process, write an applet in Go and
  integrate it with GoBox and push it.

Tools-Folder
------------
All these scripts are supposed to be run from the root of the repository inside
the DevEnv. Most of them will work on the outside as well, though.

- `geninitramfs.sh`
  This script will build a kernel compatible initramfs containing just enough to be
  able to boot with it.
  You usually don’t need to run this script yourself as `run_qemu.sh` does it for you.
  For details on customization take a look at the script itself and [the kernel’s implementation](https://github.com/torvalds/linux/blob/master/usr/gen_init_cpio.c)

- `run_qemu.sh`
  This script builds an initramfs and starts a virtual machine booting the DevEnv’s
  kernel together with the newly build initramfs.

- `netpkg_fix.sh`
  You don’t need to run this script except if you updated Go!
  A while ago, the Go team has decided to use libc´s DNS lookup routines instead of
  their own´s which requires dynamic linking. This script will recompile the `net`
  package of the Go distribution to reenable static linking.

Bugs
----
Probably

Contact
-------
If you have ideas for missing applets, found a bug or have a suggestion, use
this [project’s issues](https://github.com/asdf-systems/gobox/issues).
If you want to participate, just fork and code away. For questions contact me:
<surma@asdf-systems.de>

Thanks
------
- Thanks to [Andreas Krennmair](https://github.com/akrennmair) for `grep`, `gzip` and `gunzip`.
- Thanks to [ukaszg](https://github.com/ukaszg) for `head`.

Credits
-------
(c) 2011 Alexander "Surma" Surma <surma@asdf-systems.de>
