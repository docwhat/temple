Contributing
============

I love pull requests from everyone!

Getting Started
---------------

### Install Go

First you'll need to make sure you have go version 1.6 or later. golang.org has
some [good instructions on installing Go](https://golang.org/doc/install).

Since I use Mac OS, I can just do:

    brew install go

Fork, then clone the repo:

    git clone git@github.com:your-username/temple.git

Fetch the required dependencies:

    script/bootstrap

Make sure the tests pass:

    script/test

Make your change. Add tests for your change. Make the tests pass:

    script/test

Push to your fork and [submit a pull
request](https://github.com/docwhat/temple/compare/).

At this point you're waiting on me. I try to be responsive to pull requests,
but you know life can get in the way. I may suggest some changes or
improvements or alternatives.

Some things that will increase the chance that your pull request is accepted:

-   Write tests.
-   Write a [good commit
    message](http://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html).
