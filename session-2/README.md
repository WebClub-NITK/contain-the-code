# Containers from Scratch

This direcotory contains go program that builds a container from scratch with
minimal functionality.

Heavily inspired from: https://github.com/lizrice/containers-from-scratch

**NOTE:** For Chroot to work, you will need a copy of ubuntu root file system.
You can obtain that by running the below commands:

```
docker run -d --rm --name ubuntufs ubuntu:20.04 sleep 1000
docker export ubuntufs -o ubuntufs.tar
docker stop ubuntufs
mkdir -p ubuntufs
tar xf ubuntufs.tar -C ubuntufs
```