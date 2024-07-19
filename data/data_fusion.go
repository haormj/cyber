package data

type DataFusion2[T0, T1 any] interface {
	Fusion(index *uint64, t0 *T0, t1 *T1) bool
}
type DataFusion3[T0, T1, T2 any] interface {
	Fusion(index *uint64, t0 *T0, t1 *T1, t2 *T2) bool
}

type DataFusion4[T0, T1, T2, T3 any] interface {
	Fusion(index *uint64, t0 *T0, t1 *T1, t2 *T2, t3 *T3) bool
}
