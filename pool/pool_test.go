package pool

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestGetPool(t *testing.T) {
	type args struct {
		cpu        uint8
		goroutines uint32
		queueSize  uint32
		expireTime int
	}
	tests := []struct {
		name string
		args args
		want *pool
	}{
		{
			name: "TestCase 1",
			args: args{
				cpu:        2,
				goroutines: 4,
				queueSize:  100,
				expireTime: 10,
			},
			want: &pool{
				procs:         2,
				maxGoRoutines: 4,
				queueSize:     100,
				ttl:           time.Duration(10) * time.Second,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPool(tt.args.cpu, tt.args.goroutines, tt.args.queueSize, tt.args.expireTime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pool_Start(t *testing.T) {
	p := &pool{
		procs:         1,
		maxGoRoutines: 2,
		queueSize:     100,
		workChan:      make(chan work),
		ttl:           10,
	}
	tests := []struct {
		name string
	}{
		{
			name: "TestCase - 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p.Start()
		})
	}
}

func Test_pool_Submit(t *testing.T) {
	p1 := &pool{
		procs:         1,
		maxGoRoutines: 2,
		queueSize:     100,
		workChan:      make(chan work, 100),
		ttl:           10,
	}
	p2 := &pool{
		procs:         1,
		maxGoRoutines: 2,
		queueSize:     100,
		ttl:           10,
	}
	type args struct {
		f   workableUnit
		arg []interface{}
	}
	tests := []struct {
		name    string
		args    args
		p       *pool
		wantErr bool
	}{
		{
			name: "TestCase - 1",
			args: args{
				f: func(args ...interface{}) {
					fmt.Println(args[0])
				},
				arg: []interface{}{0},
			},
			p:       p1,
			wantErr: false,
		},
		{
			name: "TestCase - 1",
			args: args{
				f: func(args ...interface{}) {
					fmt.Println(args[0])
				},
				arg: []interface{}{0},
			},
			p:       p2,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.p
			if err := p.Submit(tt.args.f, tt.args.arg...); (err != nil) != tt.wantErr {
				t.Errorf("Submit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_pool_enableMaxGoRoutines(t *testing.T) {
	p := &pool{
		procs:         1,
		maxGoRoutines: 2,
		queueSize:     100,
		workChan:      make(chan work, 100),
		ttl:           10,
	}
	tests := []struct {
		name string
	}{
		{
			name: "TestCase - 1",
		},
	}
	p.workChan <- work{
		f: func(args ...interface{}) {
			fmt.Println(args[0])
		},
		createdAt: time.Now(),
		args:      []interface{}{"go"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p.enableMaxGoRoutines()
		})
	}
}

func Test_pool_enableMaxProcs(t *testing.T) {
	p := &pool{
		procs:         1,
		maxGoRoutines: 2,
		queueSize:     100,
		workChan:      make(chan work, 100),
		ttl:           10,
	}
	tests := []struct {
		name string
	}{
		{
			name: "TestCase - 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p.enableMaxProcs()
		})
	}
}

func Test_pool_enableQueue(t *testing.T) {
	p := &pool{
		procs:         1,
		maxGoRoutines: 2,
		queueSize:     100,
		ttl:           10,
	}
	tests := []struct {
		name string
	}{
		{
			name: "TestCase - 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p.enableQueue()
			if p.workChan == nil {
				t.Errorf("queue not created and initialized")
			}
		})
	}
}

//func Test_worker(t *testing.T) {
//	type arg struct {
//		id   int
//		jobs <-chan work
//		p    work
//	}
//	p := &pool{
//		procs:         1,
//		maxGoRoutines: 2,
//		queueSize:     100,
//		workChan:      make(chan work, 100),
//		ttl:           10,
//	}
//
//	tests := []struct {
//		name string
//		args arg
//	}{
//		{
//			name: "TestCase - 1",
//			args: arg{id: 1,
//				jobs: p.workChan},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p.enableMaxGoRoutines()
//			worker(tt.args.id, tt.args.jobs, p)
//		})
//	}
//}

func Test_pool_Stop(t *testing.T) {
	p := &pool{
		procs:         1,
		maxGoRoutines: 2,
		queueSize:     100,
		workChan:      make(chan work, 100),
		ttl:           10,
	}
	tests := []struct {
		name string
	}{
		{
			name: "TestCase - 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p.Stop()
		})
	}
}
