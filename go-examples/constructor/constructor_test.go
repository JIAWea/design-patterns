package constructor

import (
"context"
"fmt"
"testing"
"time"
)

func TestResponsibility(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"name"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 前期储备，设定相关参数
			commonsOpts := []Option{
				ConnectionTimeout(1 * time.Second),
				ReadTimeout(2 * time.Second),
				WriteTimeout(3 * time.Second),
				LogError(func(ctx context.Context, err error) {
				}),
			}

			// 终极操作，构造函数
			cluster := NewCluster(commonsOpts...)

			// 测试验证
			fmt.Println(cluster.opts.connectionTimeout)
			fmt.Println(cluster.opts.writeTimeout)
		})
	}
}
