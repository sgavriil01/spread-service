package spread

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetSymbolsHandlerOK(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/symbols", nil)
	rr := httptest.NewRecorder()

	GetSymbolsHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	var got []string
	if err := json.Unmarshal(rr.Body.Bytes(), &got); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(got) != len(Symbols) {
		t.Fatalf("expected %d symbols, got %d", len(Symbols), len(got))
	}
}

func TestGetSymbolsHandlerMethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/symbols", nil)
	rr := httptest.NewRecorder()

	GetSymbolsHandler(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status %d, got %d", http.StatusMethodNotAllowed, rr.Code)
	}
}

func TestGetSpreadHandlerNotFoundWhenNeverSet(t *testing.T) {
	resetSpreads()

	req := httptest.NewRequest(http.MethodGet, "/spreads/EURCAD", nil)
	rr := httptest.NewRecorder()

	GetSpreadHandler(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, rr.Code)
	}
}

func TestPatchSpreadHandlerBadRequestInvalidBody(t *testing.T) {
	req := httptest.NewRequest(http.MethodPatch, "/spreads/EURUSD", strings.NewReader("{bad json"))
	rr := httptest.NewRecorder()

	PatchSpreadHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, rr.Code)
	}
}

func TestPatchSpreadHandlerBadRequestInvalidSpread(t *testing.T) {
	req := httptest.NewRequest(http.MethodPatch, "/spreads/EURUSD", strings.NewReader(`{"spread":0}`))
	rr := httptest.NewRecorder()

	PatchSpreadHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, rr.Code)
	}
}

func TestPatchSpreadHandlerNotFoundInvalidSymbol(t *testing.T) {
	req := httptest.NewRequest(http.MethodPatch, "/spreads/INVALID", strings.NewReader(`{"spread":0.0001}`))
	rr := httptest.NewRecorder()

	PatchSpreadHandler(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, rr.Code)
	}
}

func TestPatchAndGetSpreadSuccess(t *testing.T) {
	resetSpreads()

	patchReq := httptest.NewRequest(http.MethodPatch, "/spreads/EURUSD", strings.NewReader(`{"spread":0.0001}`))
	patchRR := httptest.NewRecorder()
	PatchSpreadHandler(patchRR, patchReq)

	if patchRR.Code != http.StatusOK {
		t.Fatalf("expected patch status %d, got %d", http.StatusOK, patchRR.Code)
	}

	getReq := httptest.NewRequest(http.MethodGet, "/spreads/EURUSD", nil)
	getRR := httptest.NewRecorder()
	GetSpreadHandler(getRR, getReq)

	if getRR.Code != http.StatusOK {
		t.Fatalf("expected get status %d, got %d", http.StatusOK, getRR.Code)
	}

	var got Spread
	if err := json.Unmarshal(getRR.Body.Bytes(), &got); err != nil {
		t.Fatalf("failed to decode get response: %v", err)
	}

	if got.Symbol != "EURUSD" {
		t.Fatalf("unexpected symbol: %s", got.Symbol)
	}

	if got.Spread != 0.0001 {
		t.Fatalf("unexpected spread: %f", got.Spread)
	}

	if got.UpdatedAt.IsZero() {
		t.Fatalf("expected updated_at to be set")
	}
}

func TestSpreadHandlerMethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/spreads/EURUSD", nil)
	rr := httptest.NewRecorder()

	SpreadHandler(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status %d, got %d", http.StatusMethodNotAllowed, rr.Code)
	}
}
