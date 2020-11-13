Solution
========

```sh
# Read from the file words.txt and output the word frequency list to stdout.

sed 's/ /\n/g' words.txt | sort | grep -P "\w+" | uniq -c | sort -r | awk '{print $2, $1}'
```
