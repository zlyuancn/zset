/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/1/9
   Description :  无序set
-------------------------------------------------
*/

package zset

type StartTraverseFunc func(func(a interface{}))
type CancelFunc func()

type Set struct {
    m    map[interface{}]struct{}
    size int
}

func New() *Set {
    return NewAndMakeSize(0)
}

func NewAndMakeSize(size int) *Set {
    return &Set{
        m:    make(map[interface{}]struct{}, size),
        size: size,
    }
}

func (m *Set) Add(a interface{}) bool {
    l := len(m.m)
    m.m[a] = struct{}{}
    return len(m.m) > l
}

func (m *Set) Remove(a interface{}) bool {
    l := len(m.m)
    delete(m.m, a)
    return len(m.m) < l
}

func (m *Set) Len() int {
    return len(m.m)
}

func (m *Set) Clear() {
    m.m = make(map[interface{}]struct{}, m.size)
}

func (m *Set) Traverse() (StartTraverseFunc, CancelFunc) {
    var stop bool
    return func(fn func(a interface{})) {
            if stop {
                return
            }
            for a := range m.m {
                fn(a)
                if stop {
                    return
                }
            }
        }, func() {
            stop = true
        }
}
