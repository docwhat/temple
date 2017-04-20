Contributing
============

I love pull requests from everyone!

Getting Started
---------------

### Install Go

First you'll need to make sure you have go version 1.7 or later. golang.org has some [good instructions on installing Go](https://golang.org/doc/install).

### Getting the source

If you will be contributing, then you'll want to [fork the repository](https://help.github.com/articles/fork-a-repo/).

Once you've forked it, then you can clone the source:

``` sh
$ git clone git@github.com:<your-username>/<repository-name>.git
```

Fetch the required dependencies:

``` sh
$ script/bootstrap
```

Before you do any changes, make sure the tests pass:

``` sh
$ script/test
```

Make your change. Add tests for your change. Make the tests pass:

``` sh
$ script/test
```

Push to your fork and [submit a pull request](https://help.github.com/articles/creating-a-pull-request/).

At this point you're waiting on me. I try to be responsive to pull requests, but you know life can get in the way. I may suggest some changes or improvements or alternatives.

Some things that will increase the chance that your pull request is accepted:

-   Write tests.
-   Write a [good commit message](http://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html).

Releases
--------

1.  Update the `CHANGELOG.md`. See [keepachangelog.com](http://keepachangelog.com/) for info.
2.  Commit `CHANGELOG.md`
3.  Copy the current change log entry for the version to be released.
4.  Tag the new version. Example:

    ``` sh
    $ git tag v0.1.0
    ```

5.  Push up the commit and tag:

    ``` sh
    $ git push
    ...
    $ git push --tags
    ...
    ```

6.  Go to [the releases page](https://github.com/docwhat/temple/releases)
    1.  Edit the new release.
    2.  Paste in the change log entry.
    3.  Save.
