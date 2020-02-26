#!/bin/sh
. ./settings.sh

github-release info -u $user -r $repo

github-release upload \
    --user $user \
    --repo $repo \
    --tag $tag \
    --name "$repo-Linux" \
    --file ../build/linux/$buildname

github-release upload \
    --user $user \
    --repo $repo \
    --tag $tag \
    --name "$repo-Windows.exe" \
    --file ../build/windows/$buildname.exe