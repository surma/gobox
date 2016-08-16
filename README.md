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

GoBox is now `go get`-able.

Developing applets
------------------
Write your applet as a standalone Go application. When done, you need to execute the following steps:

1. Change the package name from `main` to something sensible (i.e. the exectuables name)
2. Change `import "flag"` to `import flag "appletflag"`
3. Change `func main() {...` to `func Main() {...`

Now move your code into it’s own folder unter `applets` and add your applet to the `cmd/gobox/applets.go`.

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
<surma@surmair.de>

Thanks
------
- Thanks to [Andreas Krennmair](https://github.com/akrennmair) for `grep`, `gzip` and `gunzip`.
- Thanks to [ukaszg](https://github.com/ukaszg) for `head`.
- Thanks to [vbatts](https://github.com/vbatts) for making GoBox `go get`-able.

Credits
-------
(c) 2011-2014 Alexander "Surma" Surma <surma@surmair.de>
