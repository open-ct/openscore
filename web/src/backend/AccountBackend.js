/*
 * @Author: Junlang
 * @Date: 2021-07-22 18:39:51
 * @LastEditTime: 2021-07-23 16:01:20
 * @LastEditors: Junlang
 * @FilePath: /casbin-oa/web/src/backend/AccountBackend.js
 */
// Copyright 2021 The casbin Authors. All Rights Reserved.
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

import * as Setting from "../Setting";
import {authConfig} from "../auth/Auth";
let ServerUrl=Setting.ServerUrl.slice(0,-7)

export function getAccount() {
  return fetch(`${ServerUrl}/api/get-account`, {
    method: 'GET',
    credentials: 'include'
  }).then(res => res.json());
}

export function getUsers(owner) {
  return fetch(`${ServerUrl}/api/get-users?owner=${owner}`, {
    method: "GET",
    credentials: "include"
  }).then(res => res.json());
}

export function login(code, state) {
  return fetch(`${ServerUrl}/api/login?code=${code}&state=${state}`, {
    method: 'POST',
    credentials: 'include',
  }).then(res => res.json());
}

export function logout() {
  return fetch(`${ServerUrl}/api/logout`, {
    method: 'POST',
    credentials: "include",
  }).then(res => res.json());
}
