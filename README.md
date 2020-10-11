# gocron

gocron is a simple cronjob executor. The goal is to be used in Docker to launch cronjob easily under no root user.
The celery goes to the second, some examples below :

```
# second minute hour day month
0 * * * * date
* * * * * date -u
*/10 * * * * /bin/bash -c 'echo "************$(date -u)"'
* * * * * date -d "1 day ago"
@every 1h30 date -d "1 day ago"
@daily date
```

I don't have a lot of time so I will add some improvements soon or not soon... And yes, I know supercronic and sometimes, I use it also in Docker images.

# Launch it

```
gocron -f crontab-example
```

# Build it

```
go get
go build gocron.go
```

# Release

A remind for me :

```
git tag -a v0.1.0 -m "First release"
git push origin v0.1.0
goreleaser --snapshot  # Check
goreleaser 
```