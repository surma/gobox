GoBox v0.3
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
is fairly limited. It does the job of acting as a shell, it's hardly adequate for 
scripting, though.
- Telnetd has no authentication mechanism right now. It just makes a program available
over network.

Installation
------------
For development, I recommend using [`gb`](http://code.google.com/p/go-gb/).
The hassle of updating makefiles and the dependencies just vanish.
If you just want to build GoBox, make sure you have `make` available and run:

 	.../gobox user$ ./build

Developing applets
------------------
- Copy `applets/template` and name the copy like your applet
- Rename `template.go` and edit it's contents to fit your applet
- Add your applet to `cmd/gobox/applets.go`

The template provides the basic framework you should stick to that.

Missing applets
---------------
- `kill`
- `grep`
- `tee`
- `ping`
- ...

Static linking
--------------
A while ago, the Go team has decided to use libc´s DNS lookup routines instead of
their own´s which requires dynamic linking.
It is however quite easy to reactivate the native DNS routines of Go:

	.../gobox user$ ./tools/netpkg_fix.sh 

Bugs
----
- Probably

Contact
-------
If you have ideas for missing applets, found a bug, have a suggestion
or maybe you even have an implementation ready, please contact me: alexander.surma@gmail.com

Credits
-------
(c) 2011 Alexander "Surma" Surma <surma@asdf-systems.de>
