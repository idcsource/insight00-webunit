// Copyright 2016-2017
// CoderG the 2016 project
// Insight 0+0 [ 洞悉 0+0 ]
// InDimensions Construct Source [ 忆黛蒙逝·建造源 ]
// Stephen Fire Meditation Qin [ 火志溟 ] -> firemeditation@gmail.com
// Use of this source code is governed by GNU LGPL v3 license

package admin

import (
	"github.com/idcsource/insight00-lib/webs2"
)

// admin组件入口
type AdminDoor struct {
}

// webs2的door接口实现
func (d *AdminDoor) FloorList() (floors webs2.FloorDoor) {
	floors = make(map[string]webs2.FloorInterface)
	floors["/"] = &Index{}
	floors["login"] = &LoginFloor{}
	floors["logindo"] = &Logindo{}
	return
}
