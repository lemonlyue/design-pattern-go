package options

// Options
type Options struct {
	optionA string
	optionB int
	optionC bool
}

// Option
type Option func(opts *Options)

func InitOptions(opts ...Option) *Options {
	options := &Options{}
	for _, opt := range opts {
		opt(options)
	}
	return options
}

// SetOptionA set option a
func SetOptionA(a string) Option {
	return func(opts *Options) {
		opts.optionA = a
	}
}

// SetOptionB set option b
func SetOptionB(b int) Option {
	return func(opts *Options) {
		opts.optionB = b
	}
}

// SetOptionC set option c
func SetOptionC(c bool) Option {
	return func(opts *Options) {
		opts.optionC = c
	}
}
