package uploadbills

import (
	"net/http/httptest"
	"strings"
	"testing"
)

// TEST=true bash -c ' go test -timeout=1200s'
func TestUploadCreditNoteAPI(t *testing.T) {
	// go panicOnTimeout(10 * time.Hour)

	// RB - WnwBuNNbGQLHNRAOStli
	// bizom - bizom
	// ITC - ksBQR6afWnyaP7G9lv22
	tests := []struct {
		body string
		want string
	}{
		{body: `{"filePath":"rb/2020-08-27/Demo_bills_10_08_34.xlsx", "brand":"ksBQR6afWnyaP7G9lv22", "table":"demo", "prod": false, "customerId":"ksBQR6afWnyaP7G9lv22", "warehouseId":"rDQeYeaiFtAIFRRVnU6I", "storeType":"TT"}`, want: "null"},
	}

	for _, test := range tests {
		req := httptest.NewRequest("POST", "/", strings.NewReader(test.body))
		req.Header.Add("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		UploadCreditNoteAPI(rr, req)

		if got := rr.Body.String(); got != test.want {
			t.Errorf("HelloHTTP(%q) = %q, want %q", test.body, got, test.want)
		}
	}
}
