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

package test

import (
	"testing"
	"time"

	api "github.com/polarismesh/polaris-server/common/api/v1"
	"github.com/polarismesh/polaris-server/test/grpc"
	"github.com/polarismesh/polaris-server/test/http"
	"github.com/polarismesh/polaris-server/test/resource"
)

/**
 * @brief 测试客户端GRPC接口
 */
func TestClientGRPC(t *testing.T) {
	t.Log("test client grpc interface")

	clientHTTP := http.NewClient(httpserverAddress, httpserverVersion)

	namespaces := resource.CreateNamespaces()
	services := resource.CreateServices(namespaces[0])

	// 创建命名空间
	ret, err := clientHTTP.CreateNamespaces(namespaces)
	if err != nil {
		t.Fatalf("create namespaces fail")
	}
	for index, item := range ret.GetResponses() {
		namespaces[index].Token = item.GetNamespace().GetToken()
	}
	t.Log("create namespaces success")

	// 创建服务
	ret, err = clientHTTP.CreateServices(services)
	if err != nil {
		t.Fatalf("create services fail")
	}
	for index, item := range ret.GetResponses() {
		services[index].Token = item.GetService().GetToken()
	}
	t.Log("create services success")

	// -------------------------------------------------------

	clientGRPC, err := grpc.NewClient(grpcserverAddress)
	if err != nil {
		t.Fatalf("new grpc client fail")
	}
	defer clientGRPC.Close()

	client := resource.CreateClient(0)
	instances := resource.CreateInstances(services[0])

	// 上报客户端信息
	err = clientGRPC.ReportClient(client)
	if err != nil {
		t.Fatalf("report client fail")
	}
	t.Log("report client success")

	// 注册服务实例
	err = clientGRPC.RegisterInstance(instances[0])
	if err != nil {
		t.Fatalf("register instance fail")
	}
	t.Log("register instance success")

	time.Sleep(3 * time.Second) // 延迟

	// 上报心跳
	err = clientGRPC.Heartbeat(instances[0])
	if err != nil {
		t.Fatalf("instance heartbeat fail")
	}
	t.Log("instance heartbeat success")

	time.Sleep(3 * time.Second) // 延迟

	// 查询服务实例
	err = clientGRPC.Discover(api.DiscoverRequest_INSTANCE, services[0])
	if err != nil {
		t.Fatalf("discover instance fail")
	}
	t.Log("discover instance success")

	// 反注册服务实例
	err = clientGRPC.DeregisterInstance(instances[0])
	if err != nil {
		t.Fatalf("deregister instance fail")
	}
	t.Log("deregister instance success")

	// 删除服务
	err = clientHTTP.DeleteServices(services)
	if err != nil {
		t.Fatalf("delete services fail")
	}
	t.Log("delete services success")

	// 删除命名空间
	err = clientHTTP.DeleteNamespaces(namespaces)
	if err != nil {
		t.Fatalf("delete namespaces fail")
	}
	t.Log("delete namespaces success")
}
