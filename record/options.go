package record

type Options struct {
	TopicNameMapFunc TopicNameMapFunc
}

type Option func(*Options)
type TopicNameMapFunc func(name string) string

func WithTopicNameMapFunc(fn TopicNameMapFunc) Option {
	return func(o *Options) {
		o.TopicNameMapFunc = fn
	}
}
