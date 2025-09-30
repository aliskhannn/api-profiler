package sum

import (
	"net/http/httptest"
	"testing"
)

func BenchmarkSumHandler(b *testing.B) {
	handler := NewHandler()

	req := httptest.NewRequest("GET", "/sum?a=1&b=2", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		handler.Sum(w, req)
	}
}
