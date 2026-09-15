package anidb

import (
	"animeta/lib/core/env"
	"animeta/lib/core/fetch"
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"io"
	"net/http"
)

func NewClient() (Client, error) {
	name, err := env.GetVar("ANIDB_CLIENT_NAME")
	if err != nil {
		return Client{}, err
	}

	version, err := env.GetIntVar("ANIDB_CLIENT_VER")
	if err != nil {
		return Client{}, err
	}

	return Client{
		Name:    name,
		Version: version,
	}, nil
}

func (c *Client) FetchAnime(aid int) (AnidbAnime, error) {
	var out AnidbAnime

	req, err := http.NewRequest("GET", BaseURL, nil)
	if err != nil {
		return out, err
	}

	query := req.URL.Query()
	query.Set("request", "anime")
	query.Set("client", c.Name)
	query.Set("clientver", fmt.Sprintf("%d", c.Version))
	query.Set("protover", "1")
	query.Set("aid", fmt.Sprintf("%d", aid))
	req.URL.RawQuery = query.Encode()

	resp, err := fetch.Do(req)
	if err != nil {
		return out, err
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return out, err
	}

	decompressed, err := decompress(body)
	if err != nil {
		return out, err
	}

	err = fetch.ExtractResponseXmlBody(&http.Response{Body: io.NopCloser(bytes.NewReader(decompressed))}, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}

func decompress(body []byte) ([]byte, error) {
	if len(body) == 0 {
		return body, nil
	}

	switch {
	case body[0] == 0x1f && body[1] == 0x8b: // gzip
		reader, err := gzip.NewReader(bytes.NewReader(body))
		if err != nil {
			return nil, err
		}
		defer reader.Close()
		return io.ReadAll(reader)
	case body[0] == 0x78: // zlib
		reader, err := zlib.NewReader(bytes.NewReader(body))
		if err != nil {
			return nil, err
		}
		defer reader.Close()
		return io.ReadAll(reader)
	default:
		return body, nil
	}
}
