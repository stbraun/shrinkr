# shrinkr

Helps with shrinking the size of HTML files clipped from Medium by dropping the stuff below the article.

With the following call `theSourceToShrink.html` will be read and shrinked. The shrinked file will be written to the directory `/path/to/put/the/created/file` with the file name taken from the HTML title tag.

``` sh
$ shrinkr --outpath /path/to/put/the/created/file theSourceToShrink.html
```

If the target file name shall be given via the command line the `--outfile` parameter can be used.
``` sh
$ shrinkr --outpath /path/to/put/the/created/file --outfile myTarget.txt theSourceToShrink.html
````

