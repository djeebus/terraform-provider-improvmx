package improvmx

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	improvmxApi "github.com/issyl0/go-improvmx"
)

type Meta struct {
	Token string
}

func (m *Meta) Client() *improvmxApi.Client {
	return improvmxApi.NewClient(m.Token)
}

func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("IMPROVMX_API_TOKEN", nil),
				Description: "The API token for API operations."},
		},
		ResourcesMap: map[string]*schema.Resource{
			"improvmx_domain":          resourceDomain(),
			"improvmx_email_forward":   resourceEmailForward(),
			"improvmx_smtp_credential": resourceSMTPCredential(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"improvmx_domain_check": DataSourceDomainCheck(),
		},
		ProviderMetaSchema: map[string]*schema.Schema{},
	}

	p.ConfigureFunc = providerConfigure

	return p
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return &Meta{
		Token: d.Get("token").(string),
	}, nil
}
