# 4chan General [![Go](https://github.com/bingsoo420/4ch-general/actions/workflows/go.yml/badge.svg)](https://github.com/bingsoo420/4ch-general/actions/workflows/go.yml)

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

### As a source of truth
As a bash script

```sh
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

### As a program

```sh
git clone https://github.com/bingsoo420/4ch-general
cd 4ch-general
go run ./
```

See the output JSON file in `output/mappings.json`

### As a fork

Set up a personal access token in GitHub by visiting 

`Profile / Settings / Developer settings / Personal access tokens / Fine-grained tokens / generate new token`

Set the token's permission to have read/write to `Contents`

Copy the output and put that into your forked repo settings

`Repo / Settings / Secrets and variables / Actions / New repository secret`

Add your access token with the name `PAT` and the value is your copied token

Set workflow permission to have read/write permissions

`Repo / Settings / Actions / General / Workflow Permissions / Read and Write Permissions`

Update the `.github/workflows/cron-go.yml` the final block containing the git configs

```sh
git config --global user.name $USERNAME
git config --global user.email $EMAIL

git remote set-url --push origin https://$USERNAME:$TOKEN@github.com/$USERNAME/4ch-general

git add -A
git commit -m "[CRON] mappings.json updated"
git push
```
