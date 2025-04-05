package auth

import (
	"net/http"
	"strings"
	"testing"
)


func TestGetAPIKey(t *testing.T) {
    tests := []struct {
        name string
        headerKey string
        headerValue string
        wantErr string
    }{
        {
            name: "Missing Header",
            headerKey: "x-test",
            headerValue: "",
            wantErr: "no authorization header included",
        },
        {
            name: "Valid API key",
            headerKey: "Authorization",
            headerValue: "ApiKey my-test-key",
            wantErr: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            headers := make(http.Header, 1)
            headers.Add(tt.headerKey, tt.headerValue)
            apiKey, err := GetAPIKey(headers)
            if err != nil && err.Error() != tt.wantErr {
                t.Fatalf("Wanted error %v, got error %v", err.Error(), tt.wantErr)
            } else {
                if !strings.Contains(tt.headerValue, apiKey) {
                    t.Fatalf("Wanted API key %v, got API key %v", apiKey, tt.headerValue)
                }
            }
        })
    }
}
