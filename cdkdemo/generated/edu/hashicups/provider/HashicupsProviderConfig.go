package provider


type HashicupsProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://hashicorp.com/providers/edu/hashicups/latest/docs#alias HashicupsProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://hashicorp.com/providers/edu/hashicups/latest/docs#host HashicupsProvider#host}.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://hashicorp.com/providers/edu/hashicups/latest/docs#password HashicupsProvider#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://hashicorp.com/providers/edu/hashicups/latest/docs#username HashicupsProvider#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

