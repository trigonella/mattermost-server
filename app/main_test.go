// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package app

import (
	"testing"

	"github.com/tetrafolium/mattermost-server/v5/mlog"
	"github.com/tetrafolium/mattermost-server/v5/testlib"
)

var mainHelper *testlib.MainHelper

func TestMain(m *testing.M) {
	var options = testlib.HelperOptions{
		EnableStore:     true,
		EnableResources: true,
	}

	mlog.DisableZap()

	mainHelper = testlib.NewMainHelperWithOptions(&options)
	defer mainHelper.Close()

	mainHelper.Main(m)
}
