#!/bin/bash

readarray -t commits <<< $(git log --reverse --pretty=format:'%h')
readarray -t messages <<< $(git log --reverse --pretty=format:'%s')

if [ "$1" == "list" ]; then
	for i in "${!commits[@]}"; do
		echo -e "${commits[$i]}  $i  ${messages[$i]}"
	done
	exit 0
fi

commit=$1
[[ -z "$commit" ]] && echo "Missing commit" && exit 1
[[ $commit == ${commits[0]} || $commit == "0" ]] && echo "First commit, nothing to diff" && exit 1

# Instead of commit, accepts also the position of the commit in the history of changes
if [[ ${#commit} -lt 8 ]];
then
	commit=${commits[$commit]}
fi

tempdir=$(mktemp -d)
git diff --name-only "${commit}~..${commit}" | xargs git --work-tree $tempdir checkout "${commit}~" --
meld $tempdir/*
rm -rf $tempdir
