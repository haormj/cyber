package component

import "google.golang.org/protobuf/proto"

type BaseTimerComponent struct{}

func (c *BaseTimerComponent) Init() bool { return false }
func (c *BaseTimerComponent) Proc() bool { return false }

type BaseComponent struct{}

func (c *BaseComponent) Init() bool { return false }
func (c *BaseComponent) Proc() bool { return false }

type BaseComponent1[M proto.Message] struct{}

func (c *BaseComponent1[M]) Init() bool { return false }
func (c *BaseComponent1[M]) Proc() bool { return false }

type BaseComponent2[M0, M1 proto.Message] struct{}

func (c *BaseComponent2[M0, M1]) Init() bool             { return false }
func (c *BaseComponent2[M0, M1]) Proc(m0 M0, m1 M1) bool { return false }

type BaseComponent3[M0, M1, M2 proto.Message] struct{}

func (c *BaseComponent3[M0, M1, M2]) Init() bool                    { return false }
func (c *BaseComponent3[M0, M1, M2]) Proc(m0 M0, m1 M1, m2 M2) bool { return false }

type BaseComponent4[M0, M1, M2, M3 proto.Message] struct{}

func (c *BaseComponent4[M0, M1, M2, M3]) Init() bool                           { return false }
func (c *BaseComponent4[M0, M1, M2, M3]) Proc(m0 M0, m1 M1, m2 M2, m3 M3) bool { return false }
