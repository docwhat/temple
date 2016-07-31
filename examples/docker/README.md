# Docker Example

This can be used to create a simple `Dockerfile` suitable for building go
programs in.

Since I one of the use cases I have for `temple` is to build or modify images
so that they'll work with Jenkins Pipelines, it was important I could get the
UID of the current user.

Jenkins Pipelines run contains with a command line like this:

```
$ docker run -u "$(id -u):$(id -g)" -v "${PWD}:${PWD}" ...
```

Which can be problematic, depending on what you want to do in the container.
Some programs try looking up the current user id (e.g. Python, Ruby, Java) and
will get cranky when the user doesn't exist in `/etc/passwd`.
