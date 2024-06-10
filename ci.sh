parse_git_branch() {
            git branch 2> /dev/null | sed -e '/^[^*]/d' -e 's/* \(.*\)/ (\1)/'
    }
branch=$(parse_git_branch)
git add . && git ci -m "update $branch" && git push --set-upstream origin $branch