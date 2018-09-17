## autocapture

> Take screenshot of your display automatically and simply !

---

It's very simple program. It will help you if you have to take a screenshot of your display repeatedly.

### How to work with `autocapture` ?

```bash
./autocapture --interval="15s" --output="./captures"
```

```bash
./autocapture --interval="15s" --start-hour=15 --finish-hour=20
```

Condition for hours is hour >= `start-hour` and hour < `finish-hour`.

### How to compile for Windows ?

**amd64**:
```bash
GOOS=windows GOARCH=amd64 go build -i -o autocapture.exe
```

**i386**
```bash
GOOS=windows GOARCH=386 go build -i -o autocapture.exe
```

---

Note that dependency manager is `dep`. So if you want to make `vendor` directory use `dep ensure`.
