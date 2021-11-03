## setup

### DB

1. create database

```mysql
CREATE DATABASE goca;
```

2. create tables by sqldef

```
cd schema
sh ./hack.sh sqldef_dry
sh ./hack.sh sqldef_apply
```
