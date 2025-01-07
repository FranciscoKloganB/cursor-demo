package modules

// ModuleInjector represents the generic configuration for modules
type ModuleInjector[TImports any, TProviders any, TOptions any] struct {
	Import  TImports
	Provide TProviders
	Options TOptions
}
