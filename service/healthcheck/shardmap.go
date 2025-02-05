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

package healthcheck

import (
	"sync"
	"sync/atomic"
)

// A concurrent safe shardMap for healthCheckInstances
// To avoid lock bottlenecks this map is dived to several (shardSize) concurrent map.
type shardMap struct {
	shardSize uint32
	shards    []*shard
	size      int32
}

type shard struct {
	healthCheckInstances map[string]*InstanceWithChecker
	healthCheckMutex     *sync.RWMutex
}

// NewShardMap creates a new shardMap and init shardSize.
func NewShardMap(size uint32) *shardMap {
	m := &shardMap{
		shardSize: size,
		shards:    make([]*shard, size),
		size:      0,
	}
	for i := range m.shards {
		m.shards[i] = &shard{
			healthCheckInstances: make(map[string]*InstanceWithChecker),
			healthCheckMutex:     &sync.RWMutex{},
		}
	}
	return m
}

// getShard returns shard under given instanceId.
func (m *shardMap) getShard(instanceId string) *shard {
	return m.shards[fnv32(instanceId)%m.shardSize]
}

// Store stores healthCheckInstances under given instanceId.
func (m *shardMap) Store(instanceId string, healthCheckInstance *InstanceWithChecker) {
	if len(instanceId) == 0 {
		return
	}
	shard := m.getShard(instanceId)
	shard.healthCheckMutex.Lock()
	_, ok := shard.healthCheckInstances[instanceId]
	if ok {
		shard.healthCheckInstances[instanceId] = healthCheckInstance
	} else {
		shard.healthCheckInstances[instanceId] = healthCheckInstance
		atomic.AddInt32(&m.size, 1)
	}
	shard.healthCheckMutex.Unlock()
}

// PutIfAbsent to avoid storing twice when key is the same in the concurrent scenario.
func (m *shardMap) PutIfAbsent(instanceId string, healthCheckInstance *InstanceWithChecker) (*InstanceWithChecker, bool) {
	if len(instanceId) == 0 {
		return nil, false
	}
	shard := m.getShard(instanceId)
	shard.healthCheckMutex.Lock()
	value, has := shard.healthCheckInstances[instanceId]
	if !has {
		shard.healthCheckInstances[instanceId] = healthCheckInstance
		shard.healthCheckMutex.Unlock()
		atomic.AddInt32(&m.size, 1)
		return healthCheckInstance, true
	}
	shard.healthCheckMutex.Unlock()
	return value, false
}

// Load loads the healthCheckInstances under the instanceId.
func (m *shardMap) Load(instanceId string) (healthCheckInstance *InstanceWithChecker, ok bool) {
	if len(instanceId) == 0 {
		return nil, false
	}
	shard := m.getShard(instanceId)
	shard.healthCheckMutex.RLock()
	healthCheckInstance, ok = shard.healthCheckInstances[instanceId]
	shard.healthCheckMutex.RUnlock()
	return healthCheckInstance, ok
}

// Delete deletes the healthCheckInstances under the given instanceId.
func (m *shardMap) Delete(instanceId string) {
	if len(instanceId) == 0 {
		return
	}
	shard := m.getShard(instanceId)
	shard.healthCheckMutex.Lock()
	_, ok := shard.healthCheckInstances[instanceId]
	if ok {
		delete(shard.healthCheckInstances, instanceId)
		atomic.AddInt32(&m.size, -1)
	}
	shard.healthCheckMutex.Unlock()
}

// DeleteIfExist to avoid deleting twice when key is the same in the concurrent scenario.
func (m *shardMap) DeleteIfExist(instanceId string) bool {
	if len(instanceId) == 0 {
		return false
	}
	shard := m.getShard(instanceId)
	shard.healthCheckMutex.Lock()
	_, ok := shard.healthCheckInstances[instanceId]
	if ok {
		delete(shard.healthCheckInstances, instanceId)
		atomic.AddInt32(&m.size, -1)
		shard.healthCheckMutex.Unlock()
		return true
	}
	shard.healthCheckMutex.Unlock()
	return false
}

// Range iterates over the shardMap.
func (m *shardMap) Range(fn func(instanceId string, healthCheckInstance *InstanceWithChecker)) {
	for _, shard := range m.shards {
		shard.healthCheckMutex.RLock()
		for k, v := range shard.healthCheckInstances {
			fn(k, v)
		}
		shard.healthCheckMutex.RUnlock()
	}
}

// Count returns the number of elements within the map.
func (m *shardMap) Count() int32 {
	return atomic.LoadInt32(&m.size)
}

// FNV hash.
func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}
