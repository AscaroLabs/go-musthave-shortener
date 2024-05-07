package rest_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/AscaroLabs/go-musthave-shortener/internal/adapters/api/rest"
	"github.com/AscaroLabs/go-musthave-shortener/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var mIDURL = map[string]string{}

var h = rest.NewLinkHandler()

func TestLinkHandlerLink(t *testing.T) {
	t.Run("Test Short", testLinkHandlerShort)

	t.Run("Test RedirectOriginal", testLinkHandlerRedirectOriginal)
}

func testLinkHandlerShort(t *testing.T) {
	type want struct {
		code        int
		contentType string
		response    *regexp.Regexp
		saveResult  bool
	}
	tests := []struct {
		name        string
		requestBody string
		want        want
	}{
		{
			name:        "positive test #1",
			requestBody: "https://practicum.yandex.ru/",
			want: want{
				code:        http.StatusCreated,
				contentType: "text/plain",
				response: regexp.MustCompile(
					config.NetProtocol + "://" +
						config.HTTPHost + config.HTTPPort +
						fmt.Sprintf("/[a-zA-Z0-9]{%d}", config.IDLength)),
				saveResult: true,
			},
		},
		{
			name:        "positive test #2 (same link)",
			requestBody: "https://practicum.yandex.ru/",
			want: want{
				code:        http.StatusCreated,
				contentType: "text/plain",
				response: regexp.MustCompile(
					config.NetProtocol + "://" +
						config.HTTPHost + config.HTTPPort +
						fmt.Sprintf("/[a-zA-Z0-9]{%d}", config.IDLength)),
				saveResult: true,
			},
		},
		{
			name:        "positive test #3",
			requestBody: "https://play.golang.com/",
			want: want{
				code:        http.StatusCreated,
				contentType: "text/plain",
				response: regexp.MustCompile(
					config.NetProtocol + "://" +
						config.HTTPHost + config.HTTPPort +
						fmt.Sprintf("/[a-zA-Z0-9]{%d}", config.IDLength)),
				saveResult: true,
			},
		},
		{
			name:        "positive test #4",
			requestBody: "youtube.com",
			want: want{
				code:        http.StatusCreated,
				contentType: "text/plain",
				response: regexp.MustCompile(
					config.NetProtocol + "://" +
						config.HTTPHost + config.HTTPPort +
						fmt.Sprintf("/[a-zA-Z0-9]{%d}", config.IDLength)),
				saveResult: true,
			},
		},
		{
			name:        "negative test #1 (empty body)",
			requestBody: "",
			want: want{
				code:        http.StatusBadRequest,
				contentType: "",
				response:    regexp.MustCompile(""),
			},
		},
		{
			name:        "negative test #2 (bad URL)",
			requestBody: "https://play....golang.com/",
			want: want{
				code:        http.StatusBadRequest,
				contentType: "",
				response:    regexp.MustCompile(""),
			},
		},
		{
			name:        "negative test #3 (bad URL)",
			requestBody: "/a/b/c",
			want: want{
				code:        http.StatusBadRequest,
				contentType: "",
				response:    regexp.MustCompile(""),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.requestBody))
			w := httptest.NewRecorder()
			h.Short(w, request)

			res := w.Result()
			assert.Equal(t, tt.want.code, res.StatusCode)
			assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"))

			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)
			require.NoError(t, err)
			require.Regexp(t, tt.want.response, string(resBody))

			if tt.want.saveResult {
				id := strings.Split(string(resBody), "/")[3]
				mIDURL[id] = tt.requestBody
			}
		})
	}
}

func testLinkHandlerRedirectOriginal(t *testing.T) {
	type want struct {
		code     int
		location string
	}
	type test struct {
		name string
		id   string
		want want
	}
	tests := make([]test, 0, len(mIDURL))
	i := 1
	for id, url := range mIDURL {
		tests = append(tests, test{
			name: fmt.Sprintf("positive test #%d", i),
			id:   id,
			want: want{
				code:     http.StatusTemporaryRedirect,
				location: url,
			},
		})
		i += 1
	}
	tests = append(tests, []test{
		{
			name: "negative test #1 (no id)",
			id:   "",
			want: want{
				code:     http.StatusBadRequest,
				location: "",
			},
		},
		{
			name: "negative test #2 (id does not exist)",
			id:   "id",
			want: want{
				code:     http.StatusBadRequest,
				location: "",
			},
		},
	}...)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/"+tt.id, nil)
			w := httptest.NewRecorder()
			h.RedirectOriginal(w, request)

			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.want.code, res.StatusCode)
			assert.Equal(t, tt.want.location, res.Header.Get("Location"))
		})
	}
}
