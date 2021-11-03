# go-clean-arch-sample

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

### Full list has been used:

- [gin] - Web framework
- [sqlboiler](https://github.com/volatiletech/sqlboiler) - Go ORM
- [godotenv](https://github.com/joho/godotenv) - Dotenv manager
- [gomock](https://github.com/golang/mock) - Mocking framework
- [testify](https://github.com/stretchr/testify) - Testing toolkit
