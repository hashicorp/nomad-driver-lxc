# How to release nomad-driver-lxc

1. Preflight check:

```
# check you have clean checkout
git status
make clean

# check you have latest go version
go version

# Check you have the right SHA (e.g. latest master)
git checkout master
git pull origin/master
```

2. Update version references in CHANGELOG.md and PluginVersion in lxc/driver.go and commit

3. Build binaries:

```
# on mac or on Linux but without LXC installed, use Vagrant
vagrant up
vagrant ssh
make release
exit

# On linux with LXC installed
make release
```

4. Publish binaries on https://releases.hashicorp.com

```
./scripts/dist.sh <VERSION_NUMBER>
```

And check that binaries are propagated properly. 

5. Update version references for next iteration with `-dev`
6. Push shas and tags to GitHub
