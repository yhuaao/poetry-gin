package util

import (
	"context"
	"poetry/global"
	"time"
)

// 实现lock
const releaseLockLuaScript = `
if redis.call("get",KEYS[1]) == ARGV[1] then
    return redis.call("del",KEYS[1])
else
    return 0
end
`

type lockInterface interface {
	Get() bool
    Block(seconds int64) bool
    Release() bool
    ForceRelease()
}

type LockClass struct{
	context context.Context
    name string // 锁名称
    owner string // 锁标识
    seconds int64 // 有效期
}

func Lock() *LockClass{
	return &LockClass{
        context.Background(),
        "test",
        RandString(16),
        5,
    }
}

func (l *LockClass) Get() bool{
	return global.Redis.SetNX(l.context,l.name,l.owner,time.Duration(l.seconds)*time.Second).Val();
}

func test() func(int)int{
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func (l *LockClass) Block(seconds int64) bool{
	starting := time.Now().Unix()
    for {
        if !l.Get() {
            time.Sleep(time.Duration(1) * time.Second)
            if time.Now().Unix()-seconds >= starting {
                return false
            }
        } else {
            return true
        }
    }
}

func (l *LockClass) Release() bool {
	result:=global.Redis.Eval(l.context,releaseLockLuaScript,[]string{l.name},l.owner).Val().(int64)
    return result != 0
}

func (l *LockClass) ForceRelease(){
	global.Redis.Del(l.context,l.name).Val();
	return
} 
