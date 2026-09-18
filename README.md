# animeta

anime metadata API server, Go standard library only.

## API

- `GET /health`
- `GET /meta/{provider}/{id}` with `provider` = `anilist` | `mal`

Returns titles (all languages), episodes, artworks, mappings and
per-episode filler/recap.

## Data sources

- anime-lists (Fribb) for cross-provider id mappings
- TheTVDB for series, episodes and artworks
- AniDB for canonical multilingual titles
- Tenrai for filler/recap per episode

## Run locally

```bash
cp .env.example .env   # fill TVDB_APIKEY at minimum
go run ./cmd/server    # or: just dev
```

Then: `curl localhost:8080/meta/anilist/11061`

TVDB_APIKEY is required. AniDB and Tenrai are optional enrichments.