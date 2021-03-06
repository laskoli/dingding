// Copyright 2020 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package edu_test

import (
	"fmt"

	"github.com/fastwego/dingding"
	"github.com/fastwego/dingding/apis/contact/edu"
)

func ExampleDeptList() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := edu.DeptList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDeptGet() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := edu.DeptGet(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUserList() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := edu.UserList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUserGet() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := edu.UserGet(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUserRelationList() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := edu.UserRelationList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUserRelationGet() {
	var ctx *dingding.App

	payload := []byte("{}")
	resp, err := edu.UserRelationGet(ctx, payload)

	fmt.Println(resp, err)
}
