#! /bin/bash

echo -e "Easy\n====\n" > Easy/README.md
for i in $(ls Easy)
do
  echo "- [$i](Easy/${i})"
done >> Easy/README.md

echo -e "Medium\n====\n" > Medium/README.md
for i in $(ls Medium)
do
  echo "- [$i](Hard/${i})"
done >> Medium/README.md

echo -e "Hard\n====\n" > Hard/README.md
for i in $(ls Hard)
do
  echo "- [$i](Hard/${i})"
done >> Hard/README.md
