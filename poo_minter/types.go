package poo_minter

// Prop 道具
type Prop struct {
	Cost         float64 `json:"cost"`         // 升级花费
	NextLevel    uint    `json:"nextLevel"`    // 下一等级
	CurrentValue float64 `json:"currentValue"` // 当前数值
	NextValue    float64 `json:"nextValue"`    // 下一等级数值
	Cap          uint    `json:"cap"`          // 最大下一等级
}
