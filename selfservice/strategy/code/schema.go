// Copyright © 2022 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package code

import (
	_ "embed"
)

//go:embed .schema/recovery.schema.json
var recoveryMethodSchema []byte
