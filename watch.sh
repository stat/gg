#!/bin/sh

while inotifywait -qqre modify --include '.*\.go$' ./; do
  make test
done
