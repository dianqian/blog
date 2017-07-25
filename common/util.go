package common

import "time"

/**
 判断一个Time是否为0
 */
func TimeIsZero(t time.Time) bool {
    isZero := false
    if t.Second() == 0 && t.Nanosecond() == 0 {
        isZero = true
    }

    return isZero
}

/**
 创建Zero Time的时间
 */
func CreateZeroTime() time.Time {
    return time.Unix(0, 0)
}