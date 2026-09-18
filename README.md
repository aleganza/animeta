<p align="center">
    <img height="70px" src="https://github.com/aleganza/animeta/blob/main/assets/logo.webp?raw=true"/>
    <h1 align="center">animeta</h1>
</p>

<p align="center">
  <strong>Anime episodes metadata API server.</strong>
</p>

<p align="center">
    <img alt="GitHub release (with filter)" src="https://img.shields.io/github/languages/top/aleganza/animeta?style=for-the-badge&labelColor=1a1d24&color=00acd7">
</p>

## API

- `GET /health`
- `GET /meta/{provider}/{id}` with `provider` = `anilist` | `mal`

## Metadata

A request to `/meta` returns, for a single anime entry:

- **titles** — multilingual titles (main, official, synonyms, short), one per language
- **episodes** — per episode: number, multilingual titles, overview, aired date, runtime, thumbnail, TVDB/AniDB ids, **filler/recap flags**
- **episodeCount** — total episodes for the requested entry
- **artworks** — posters, banners, logos, backgrounds, one per type
- **mappings** — the same anime id across providers (anilist, mal, tvdb, anidb, imdb, kitsu and more)

Filler/recap flags and multilingual titles come from Tenrai and AniDB
respectively, and gracefully degrade to TVDB data when unavailable.

## Data sources

- anime-lists (Fribb) for cross-provider id mappings
- TheTVDB for series, episodes and artworks
- AniDB for canonical multilingual titles
- Tenrai for filler/recap per episode

## Example

```bash
curl localhost:8080/meta/anilist/11061
```

```json
{
  "success": true,
  "data": {
    "titles": [
      { "name": "Hunter x Hunter (2011)", "language": "x-jat" },
      { "name": "ハンター×ハンター (2011)", "language": "ja" },
      [...]
    ],
    "episodes": [
      {
        "tvdbId": 4180539,
        "anidbId": 132161,
        "number": 1,
        "thumbnail": "https://artworks.thetvdb.com/banners/episodes/252322/4180539.jpg",
        "titles": [
          { "name": "タビダチ×ト×ナカマタチ", "language": "ja" },
          { "name": "Departure x and x Friends", "language": "en" }
        ],
        "overview": "The story begins with a young boy named Gon...",
        "aired": "2011-10-02",
        "runtime": 25,
        "year": "2011",
        "isFiller": false,
        "isRecap": false
      },
      [...]
    ],
    "episodeCount": 148,
    "artworks": [
      {
        "type":	"banner"
        "url":	"https://artworks.thetvdb.com/banners/graphical/252322-g2.jpg"
      },
      [...]
    ],
    "mappings": {
      "anilistId": 11061,
      "malId": 11061,
      "theTvdbId": 252322,
      "anidbId": 8550
    }
  }
}
```

## Run locally

Requirements: Go 1.25+, and an internet connection to the data sources.

```bash
git clone git@github.com:aleganza/animeta.git
cd animeta

cp .env.example .env    # then fill the values
go run ./cmd/server     # or: just dev
```

The server listens on `http://localhost:8080` by default.

| Variable | Required | Description |
|---|---|---|
| `SERVER_PORT` | no | HTTP port (default `8080`) |
| `TVDB_APIKEY` | **yes** | TheTVDB v4 API key (subscription needed) |
| `ANIDB_CLIENT_NAME` / `ANIDB_CLIENT_VER` | no | AniDB registered client (enables canonical titles) |
| `TENRAI_SERVER_KEY` | no | Tenrai server key (optional, works publicly without) |

Then: `curl localhost:8080/meta/anilist/11061`

TVDB_APIKEY is required. AniDB and Tenrai are optional enrichments.