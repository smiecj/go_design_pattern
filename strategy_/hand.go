package strategy_

const (
	handCount = 3

	scissorName = "剪刀"
	rockName    = "石头"
	paperName   = "布"
)

var (
	scissor = &hand{
		name:  scissorName,
		index: 0,
	}
	rock = &hand{
		name:  rockName,
		index: 1,
	}
	paper = &hand{
		name:  paperName,
		index: 2,
	}

	handArr = []*hand{scissor, rock, paper}
	handMap = map[string]*hand{
		scissorName: scissor,
		rockName:    rock,
		paperName:   paper,
	}
)

// hand: 出拳
type hand struct {
	name  string
	index int
}

// stronger: 判断胜负关系
func (hand *hand) Stronger(anoHand *hand) bool {
	indexDiff := anoHand.index - hand.index
	if indexDiff == 1 || indexDiff == -2 {
		return true
	}
	return false
}

func (hand *hand) GetIndex(anoHand *hand) bool {
	indexDiff := anoHand.index - hand.index
	if indexDiff == 1 || indexDiff == -2 {
		return true
	}
	return false
}

// 根据顺序下标获取出拳
func getHandByIndex(index int) *hand {
	if index < 0 || index >= handCount {
		return nil
	}
	return handArr[index]
}

// 根据名字获取出拳
func getHandByName(name string) *hand {
	return handMap[name]
}
