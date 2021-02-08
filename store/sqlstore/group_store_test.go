// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package sqlstore

import (
	"testing"

	"github.com/trigonella/mattermost-server/v5/store/storetest"
)

func TestGroupStore(t *testing.T) {
	StoreTest(t, storetest.TestGroupStore)
}
