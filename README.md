# 4chan General

## What is this?

The repo contains a self-updating bot running on a 6 hourly CRON job to build
board general to it's correct thread location.

Example payload

```json
{
  "g": {
    "wdg": "https://boards.4chan.org/g/thread/91798719"
  }
}
```

The access patterns will always remain as `board.general`, currently the only
boards being scraped are

- fit
- g
- biz

Which can be modified in `main.go` file

## How do I use?

As a bash script

```bash
curl https://raw.githubusercontent.com/bingsoo420/4ch-general/master/output/mappings.json | jq .g.wdg
```

As a simple JavaScript fetch request

```js
fetch("https://raw.githubusercontent.com/bingsoo420/4ch-general/master/output/mappings.json")
    .then((r) => r.json())
    .then((r) => {
        if (r.g?.wdg) {
            window.location.href = r.g.wdg;
        }
    });
```

The source of truth will refresh every 6 hours. Files may be updated to be 
granular later as `output/g.json` or `output/fit.json` if the 6 hours interval
was found to be too little for certain boards.
