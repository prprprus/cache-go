// Copyright (c) 2019, prprprus All rights reserved.
// Use of this source code is governed by a BSD-style.
// license that can be found in the LICENSE file.

package queue

// Queue Interface
type Queue interface {
	Put(value interface{})
	Get() (interface{}, error)
	indexInRange(index int) bool
}
