package anidb

import (
	"animeta/lib/core/fetch"
	"bytes"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestFetchAnime_Parsing(t *testing.T) {
	data, err := os.ReadFile("mock_anime.xml")
	if err != nil {
		t.Fatalf("failed to read mock: %v", err)
	}

	mock := &http.Response{
		Body: io.NopCloser(bytes.NewReader(data)),
	}

	var out AnidbAnime
	if err := fetch.ExtractResponseXmlBody(mock, &out); err != nil {
		t.Fatalf("failed to parse mock: %v", err)
	}

	if out.ID != 12681 {
		t.Errorf("expected id 12681, got %d", out.ID)
	}

	if len(out.Titles) != 9 {
		t.Fatalf("expected 9 titles, got %d", len(out.Titles))
	}

	if out.Titles[0].Type != "main" || out.Titles[0].Lang != "x-jat" || out.Titles[0].Name != "Made in Abyss" {
		t.Errorf("unexpected main title: %+v", out.Titles[0])
	}

	if out.Titles[3].Lang != "ja" || out.Titles[3].Name != "メイドインアビス" {
		t.Errorf("unexpected japanese title: %+v", out.Titles[3])
	}

	if out.Titles[6].Type != "syn" || out.Titles[6].Name != "Wyprodukowano w Otchłani" {
		t.Errorf("unexpected synonym: %+v", out.Titles[6])
	}

	if len(out.Episodes) != 3 {
		t.Fatalf("expected 3 episodes, got %d", len(out.Episodes))
	}

	ep := out.Episodes[0]
	if ep.ID != 138356 || ep.EpNo != "1" || ep.Length != 25 || len(ep.Titles) != 2 {
		t.Errorf("unexpected first episode: %+v", ep)
	}

	if ep.Titles[1].Lang != "ja" || ep.Titles[1].Name != "大穴の街" {
		t.Errorf("unexpected episode japanese title: %+v", ep.Titles[1])
	}

	special := out.Episodes[2]
	if special.EpNo != "S1" {
		t.Errorf("expected special epno S1, got %q", special.EpNo)
	}
}

func TestDecompress(t *testing.T) {
	plain := []byte("<anime id=\"1\"></anime>")

	out, err := decompress(plain)
	if err != nil {
		t.Fatalf("plain decompress failed: %v", err)
	}
	if string(out) != string(plain) {
		t.Errorf("plain roundtrip mismatch: %q", string(out))
	}
}
