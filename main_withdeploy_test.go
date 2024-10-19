// Copyright 2024 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build withdeploy
// +build withdeploy

package main

import (
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func TestWithdeploy(t *testing.T) {
	p := commonTestScriptsParam
	p.Dir = "testscripts/withdeploy"
	testscript.Run(t, p)
}