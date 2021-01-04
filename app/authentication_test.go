// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package app

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/tetrafolium/mattermost-server/v5/model"
	"github.com/stretchr/testify/require"
)

func TestParseAuthTokenFromRequest(t *testing.T) {
	cases := []struct {
		header           string
		cookie           string
		query            string
		expectedToken    string
		expectedLocation TokenLocation
	}{
		{"", "", "", "", TokenLocationNotFound},
		{"token mytoken", "", "", "mytoken", TokenLocationHeader},
		{"BEARER mytoken", "", "", "mytoken", TokenLocationHeader},
		{"", "mytoken", "", "mytoken", TokenLocationCookie},
		{"", "", "mytoken", "mytoken", TokenLocationQueryString},
		{"mytoken", "", "", "mytoken", TokenLocationCloudHeader},
	}

	for testnum, tc := range cases {
		pathname := "/test/here"
		if tc.query != "" {
			pathname += "?access_token=" + tc.query
		}
		req := httptest.NewRequest("GET", pathname, nil)
		switch tc.expectedLocation {
		case TokenLocationHeader:
			req.Header.Add(model.HEADER_AUTH, tc.header)
		case TokenLocationCloudHeader:
			req.Header.Add(model.HEADER_CLOUD_TOKEN, tc.header)
		case TokenLocationCookie:
			req.AddCookie(&http.Cookie{
				Name:  model.SESSION_COOKIE_TOKEN,
				Value: tc.cookie,
			})
		}

		token, location := ParseAuthTokenFromRequest(req)

		require.Equal(t, tc.expectedToken, token, "Wrong token on test "+strconv.Itoa(testnum))
		require.Equal(t, tc.expectedLocation, location, "Wrong location on test "+strconv.Itoa(testnum))
	}
}
