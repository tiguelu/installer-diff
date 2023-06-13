The first commit has a 2-directory diff produced by `git difftool`

Each other commit, removes files related to a logical change that has no impact for ARO. The log message summarizes why.

The helper script lists the commits that discard changes and can run a side-by-side diff (it requires `meld`).

```sh
$ ./diffview.sh
fce6295  0  Two directory diff from difftool
782c877  1  Support for image galleries (added to support arm64 images) openshift/installer@2612078
...

$ ./diffview.sh 1
# Opens the diff of the diff in meld
```

