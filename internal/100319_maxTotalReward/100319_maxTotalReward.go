package leetcode_100319_maxTotalReward

import "sort"

type BitMask struct {
	vec []uint64
}

func (b *BitMask) Set(i int) {
	b.vec[i/64] |= 1 << uint(i%64)
}

func (b *BitMask) Get(i int) bool {
	return b.vec[i/64]&(1<<uint(i%64)) != 0
}

func (b *BitMask) Clear(i int) {
	b.vec[i/64] &^= 1 << uint(i%64)
}

func (b *BitMask) Init(n int) {
	b.vec = make([]uint64, n/64+1)
}

func maxTotalReward(rewardValues []int) int {
	sort.Ints(rewardValues)
	bs := new(BitMask)
	count := 2000 * 10
	bs.Init(count)
	for _, v := range rewardValues {
		bs.Set(v)
		for i := v - 1; i >= 0; i-- {
			if bs.Get(i) {
				bs.Set(v + i)
			}
		}
	}
	for i := count; i >= 0; i-- {
		if bs.Get(i) {
			return i
		}
	}
	return 0
}
