Solution
========

```sh
# Read from the file file.txt and output all valid phone numbers to stdout.
cat file.txt | grep -P '^\(\d\d\d\) \d\d\d-\d\d\d\d$|^\d\d\d-\d\d\d-\d\d\d\d$'
```
