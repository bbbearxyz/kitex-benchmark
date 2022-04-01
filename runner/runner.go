/*
 * Copyright 2021 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package runner

import (
	"sync"
	"time"
)

// 为了流量更均匀, 时间间隔设置为 10ms
const window = 10 * time.Millisecond

// 单次测试
type RunOnce func() error

type Runner struct {
	counter *Counter // 计数器
	timer   *Timer   // 计时器
}

func NewRunner() *Runner {
	r := &Runner{
		counter: NewCounter(),
		timer:   NewTimer(time.Microsecond),
	}
	return r
}

func (r *Runner) benching(onceFn RunOnce, concurrent int, total int64) {
	var wg sync.WaitGroup
	wg.Add(concurrent)
	r.counter.Reset(total)
	for i := 0; i < concurrent; i++ {
		go func() {
			defer wg.Done()
			for {
				idx := r.counter.Idx()
				if idx >= total {
					return
				}
				begin := r.timer.Now()
				err := onceFn()
				end := r.timer.Now()
				cost := end - begin
				r.counter.AddRecord(idx, err, cost)
			}
		}()
	}
	wg.Wait()
	r.counter.Total = total
}

func (r *Runner) Warmup(onceFn RunOnce, concurrent int, total int64) {
	r.benching(onceFn, concurrent, total)
}

func (r *Runner) StreamingWarmup(onceFn RunOnce, total int64) {
	r.benching(onceFn, 1, total)
}

// 并发测试
func (r *Runner) Run(title string, onceFn RunOnce, concurrent int, total int64, echoSize, sleepTime int, field, latency int64) {
	logInfo(
		"%s start benching [%s], concurrent: %d, total: %d, sleep: %d, field: %d, latency: %d, size: %d",
		"["+title+"]", time.Now().String(), concurrent, total, sleepTime, field, latency, echoSize,
	)

	start := r.timer.Now()
	r.benching(onceFn, concurrent, total)
	stop := r.timer.Now()
	r.counter.Report(title, stop-start, concurrent, total, echoSize)
}


// 流式测试
func (r *Runner) RunStream(title string, onceFn RunOnce, total int64) {
	logInfo(
		"use single thread to test stream, %s start benching [%s], total: %d",
		"["+title+"]", time.Now().String(), total,
	)
	start := r.timer.Now()
	r.benching(onceFn, 1, total)
	stop := r.timer.Now()
	r.counter.Report(title, stop-start, 1, total, 1024 * 1024 * 1024)
}
