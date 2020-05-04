/*
 * Author : xzf
 * Time    : 2020-05-05 01:43:07
 * Email   : xpoony@163.com
 */

package main

import (
	"testing"
)

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getGoroutineId()
	}
}
