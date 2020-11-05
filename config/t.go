package config

type T struct {
	ProjectPath   string `yaml:"projectPath"`
	GoPackagePath string `yaml:"goPackagePath"`
}

func (t *T) PackagePath() string {
	if t.GoPackagePath != "" {
		return t.GoPackagePath
	}
	return "example.com/myNetwork"
}
