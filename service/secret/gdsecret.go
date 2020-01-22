package secret

//GDSecret is the representation of stored secrets uploaded with gonedrive
type GDSecret struct {
	Secrets []Secret `yaml:"secrets"`
}

//Secret is a property of GDSecret and represents a single secret
type Secret struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}
