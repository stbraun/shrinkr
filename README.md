# shrinkr

Helps with shrinking the size of HTML files clipped from Medium by dropping the stuff below the article.

With the following call `theSourceToShrink.html` will be read and shrinked. The shrinked file will be written to the directory `/path/to/put/the/created/file` with the file name taken from the HTML title tag.

``` sh
$ shrinkr --outpath /path/to/put/the/created/file theSourceToShrink.html
```

If the target file name shall be given via the command line the `--outfile` parameter can be used.
``` sh
$ shrinkr --outpath /path/to/put/the/created/file --outfile myTarget.txt theSourceToShrink.html
```

Query the version number with:
``` sh
$ shrinkr --version
```

## Integration with DEVONthink
The integration with DEVONthink is implemented via the AppleScript file `Shrink.scpt`. Copy it into DEVONthink's script folder. 

To shrink a document select it in DEVONthink and run the script on it (Menu scripts/Shrink). It will shrink the document, import it into DEVONthink and delete the original document.
