## git merge vs rebase incoming change vs current change

You are on a feature branch

1. git pull origin master
Current changes
Changes on your current feature branch.

Incoming changes
Changes you are pulling from i.e the master branch

2. git pull origin master --rebase
During rebase your feature branch changes are applied on top of the commits that are already there in master branch.

Current changes
Changes on the master branch.

Incoming changes
Changes on the feature branch.

* After a rebase, you need to force push your branch. Use --force-with-lease instead of --force

Reference
* [https://stackoverflow.com/a/71060122]
