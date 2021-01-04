// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package sqlstore_test

import (
	"testing"

	"github.com/tetrafolium/mattermost-server/mlog"
	"github.com/tetrafolium/mattermost-server/store/sqlstore"

	"github.com/tetrafolium/mattermost-server/testlib"
)

var mainHelper *testlib.MainHelper

func TestMain(m *testing.M) {
	mlog.DisableZap()
	mainHelper = testlib.NewMainHelperWithOptions(nil)
	defer mainHelper.Close()

	sqlstore.InitTest()

	mainHelper.Main(m)
	sqlstore.TearDownTest()
}
