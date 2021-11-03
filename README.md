```
touch .env
```

```
CREATE DATABASE goca;
```

```
cd schema
sh ./hack.sh sqldef_dry
sh ./hack.sh sqldef_apply
```

- sqlboiler
- sqldef
- gin
