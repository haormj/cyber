package data

import (
	"github.com/haormj/cyber/common"
)

type FusionDataType2[T0, T1 any] struct {
	T0 T0
	T1 T1
}

type AllLatest2[T0, T1 any] struct {
	buffer0      *ChannelBuffer[T0]
	buffer1      *ChannelBuffer[T1]
	bufferFusion *ChannelBuffer[*FusionDataType2[T0, T1]]
}

func NewAllLatest2[T0, T1 any](buffer0 *ChannelBuffer[T0], buffer1 *ChannelBuffer[T1]) *AllLatest2[T0, T1] {
	bufferFusion := NewChannelBuffer(buffer0.ChannelID(),
		NewCacheBuffer[*FusionDataType2[T0, T1]](int(buffer0.Capacity()-1)))
	buffer0.SetFusionCallback(func(t0 T0) {
		t1 := common.Zero[T1]()
		if !buffer1.Latest(&t1) {
			return
		}

		bufferFusion.Mutex().Lock()
		defer bufferFusion.Mutex().Unlock()

		bufferFusion.Fill(&FusionDataType2[T0, T1]{
			T0: t0,
			T1: t1,
		})
	})

	return &AllLatest2[T0, T1]{
		buffer0:      buffer0,
		buffer1:      buffer1,
		bufferFusion: bufferFusion,
	}
}

func (f *AllLatest2[T0, T1]) Fusion(index *uint64, t0 *T0, t1 *T1) bool {
	var fusionData *FusionDataType2[T0, T1]
	if !f.bufferFusion.Fetch(index, &fusionData) {
		return false
	}

	*t0 = fusionData.T0
	*t1 = fusionData.T1
	return true
}

type FusionDataType3[T0, T1, T2 any] struct {
	T0 T0
	T1 T1
	T2 T2
}

type AllLatest3[T0, T1, T2 any] struct {
	buffer0      *ChannelBuffer[T0]
	buffer1      *ChannelBuffer[T1]
	buffer2      *ChannelBuffer[T2]
	bufferFusion *ChannelBuffer[*FusionDataType3[T0, T1, T2]]
}

func NewAllLatest3[T0, T1, T2 any](buffer0 *ChannelBuffer[T0], buffer1 *ChannelBuffer[T1],
	buffer2 *ChannelBuffer[T2]) *AllLatest3[T0, T1, T2] {
	bufferFusion := NewChannelBuffer(buffer0.ChannelID(),
		NewCacheBuffer[*FusionDataType3[T0, T1, T2]](int(buffer0.Capacity()-1)))
	buffer0.SetFusionCallback(func(t0 T0) {
		t1 := common.Zero[T1]()
		t2 := common.Zero[T2]()
		if !buffer1.Latest(&t1) || !buffer2.Latest(&t2) {
			return
		}

		bufferFusion.Mutex().Lock()
		defer bufferFusion.Mutex().Unlock()

		bufferFusion.Fill(&FusionDataType3[T0, T1, T2]{
			T0: t0,
			T1: t1,
			T2: t2,
		})
	})

	return &AllLatest3[T0, T1, T2]{
		buffer0:      buffer0,
		buffer1:      buffer1,
		buffer2:      buffer2,
		bufferFusion: bufferFusion,
	}
}

func (f *AllLatest3[T0, T1, T2]) Fusion(index *uint64, t0 *T0, t1 *T1, t2 *T2) bool {
	var fusionData *FusionDataType3[T0, T1, T2]
	if !f.bufferFusion.Fetch(index, &fusionData) {
		return false
	}

	*t0 = fusionData.T0
	*t1 = fusionData.T1
	*t2 = fusionData.T2
	return true
}

type FusionDataType4[T0, T1, T2, T3 any] struct {
	T0 T0
	T1 T1
	T2 T2
	T3 T3
}

type AllLatest4[T0, T1, T2, T3 any] struct {
	buffer0      *ChannelBuffer[T0]
	buffer1      *ChannelBuffer[T1]
	buffer2      *ChannelBuffer[T2]
	buffer3      *ChannelBuffer[T3]
	bufferFusion *ChannelBuffer[*FusionDataType4[T0, T1, T2, T3]]
}

func NewAllLatest4[T0, T1, T2, T3 any](buffer0 *ChannelBuffer[T0], buffer1 *ChannelBuffer[T1],
	buffer2 *ChannelBuffer[T2], buffer3 *ChannelBuffer[T3]) *AllLatest4[T0, T1, T2, T3] {
	bufferFusion := NewChannelBuffer(buffer0.ChannelID(),
		NewCacheBuffer[*FusionDataType4[T0, T1, T2, T3]](int(buffer0.Capacity()-1)))
	buffer0.SetFusionCallback(func(t0 T0) {
		t1 := common.Zero[T1]()
		t2 := common.Zero[T2]()
		t3 := common.Zero[T3]()
		if !buffer1.Latest(&t1) || !buffer2.Latest(&t2) || !buffer3.Latest(&t3) {
			return
		}

		bufferFusion.Mutex().Lock()
		defer bufferFusion.Mutex().Unlock()

		bufferFusion.Fill(&FusionDataType4[T0, T1, T2, T3]{
			T0: t0,
			T1: t1,
			T2: t2,
			T3: t3,
		})
	})

	return &AllLatest4[T0, T1, T2, T3]{
		buffer0:      buffer0,
		buffer1:      buffer1,
		buffer2:      buffer2,
		buffer3:      buffer3,
		bufferFusion: bufferFusion,
	}
}

func (f *AllLatest4[T0, T1, T2, T3]) Fusion(index *uint64, t0 *T0, t1 *T1, t2 *T2, t3 *T3) bool {
	var fusionData *FusionDataType4[T0, T1, T2, T3]
	if !f.bufferFusion.Fetch(index, &fusionData) {
		return false
	}

	*t0 = fusionData.T0
	*t1 = fusionData.T1
	*t2 = fusionData.T2
	*t3 = fusionData.T3
	return true
}
