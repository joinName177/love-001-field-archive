package fieldarchive

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBug01_CancelledRenderReturnsCancellationResponse(t *testing.T) {
	s := NewMemoryStore()
	clips := NewClipService(s)
	if e := clips.Register(context.Background(), Clip{ID: "clip-7", CaseID: "case-4", Title: "loading dock", Status: "queued"}); e != nil {
		t.Fatal(e)
	}
	a := NewAPI(NewRenderService(clips, InlineRenderer{}))
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	r := httptest.NewRequest(http.MethodPost, "/renders", bytes.NewBufferString(`{"clip_id":"clip-7","format":"mp4"}`)).WithContext(ctx)
	w := httptest.NewRecorder()
	a.Render(w, r)
	if w.Code != 499 {
		t.Fatalf("got %d, want 499", w.Code)
	}
}
