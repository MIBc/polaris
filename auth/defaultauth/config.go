/**
 * Tencent is pleased to support the open source community by making Polaris available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package defaultauth

import "errors"

var (
	// AuthOption 鉴权的配置信息
	AuthOption *AuthConfig = DefaultAuthConfig()
)

// AuthConfig 鉴权配置
type AuthConfig struct {
	// Open 是否开启鉴权
	Open bool `json:"open" xml:"open"`
	// Salt 相关密码、token加密的salt
	Salt string `json:"salt" xml:"salt"`
}

// Verify 检查配置是否合法
func (cfg *AuthConfig) Verify() error {
	k := len(cfg.Salt)
	switch k {
	default:
		return errors.New("[Auth][Config] salt len must 16 | 24 | 32")
	case 16, 24, 32:
		break
	}

	return nil
}

// DefaultAuthConfig 返回一个默认的鉴权配置
func DefaultAuthConfig() *AuthConfig {
	return &AuthConfig{
		Open: false,
		Salt: "polarismesh@2021",
	}
}
