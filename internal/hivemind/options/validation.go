package options

func (o *Options) Validate() []error {
	var errs []error
	errs = append(errs, o.GenericServerRunOptions.Validate()...)
	errs = append(errs, o.GRPCOptions.Validate()...)
	return errs
}
