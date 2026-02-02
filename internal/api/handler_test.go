package api

import (
	"net/http/httptest"
	"testing"
)

func BenchmarkSumHandler(b *testing.B) {
	h := New()
	r := httptest.NewRequest("GET", "/sum?a=123&b=456", nil)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		h.Sum(w, r)
		_ = w.Result().Body.Close()
	}
}

func BenchmarkJSONHandler(b *testing.B) {
	h := New()
	r := httptest.NewRequest("GET", "/json", nil)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		h.JSON(w, r)
		_ = w.Result().Body.Close()
	}
}
