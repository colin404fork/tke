/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the “License”); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an “AS IS” BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

package manifest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"tkestack.io/tke/pkg/kubectl"
	"tkestack.io/tke/pkg/util/template"
)

func TestManifest(t *testing.T) {
	data, err := template.ParseFile("tke-gateway.yaml",
		map[string]interface{}{
			"Image":            "Image",
			"OIDCClientSecret": "OIDCClientSecret",
			"SelfSigned":       true,
			"EnableRegistry":   true,
			"EnableAuth":       true,
			"EnableMonitor":    true,
			"EnableBusiness":   true,
			"EnableLogagent":   true,
			"EnableAudit":      true,
		})
	if !assert.Nil(t, err) {
		t.FailNow()
	}

	data, err = kubectl.Validate(data)
	if !assert.Nil(t, err) {
		t.Fatal(string(data))
	}
}
