branch=$1
git add . && git ci -m "update $branch" && git push --set-upstream origin $branch