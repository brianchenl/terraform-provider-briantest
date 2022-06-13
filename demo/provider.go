package demo

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"net/http"
)

// Provider init  block
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			// provider arguments and their specifications go here
			"endpoint": {
				Type:        schema.TypeString,
				Optional:  true,
				Default: "https://www.fastmock.site/mock/00131b8a844102450accb81d6774bf4e/_provider_0613/",
				DefaultFunc: schema.EnvDefaultFunc("GW_ENDPOINT", nil),
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					value := v.(string)
					if value == "" {
						errors = append(errors, fmt.Errorf("Endpoint must not be an empty string"))
					}

					return
				},
			},
		},
		// map terraform dsl resources to functions
		ResourcesMap: map[string]*schema.Resource{
			"brian_demo": resourceDemo(),

		},
		DataSourcesMap: map[string]*schema.Resource{
			"brian_account": dataSourceBrianAccount(),
		},
		// provider configuration function
		ConfigureFunc: configureProvider,
	}
}


// Configuration struct
type Configuration struct {
	endpoint string
}

// configure provider options
func configureProvider(data *schema.ResourceData) (interface{}, error) {
	// pass options from terraform DSL to the client
	endpoint := data.Get("endpoint").(string)

	// test endpoint
	_, err := http.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("Error connect to gateway ")
	}
	if err != nil {
		return nil, err
	}
	// code to error handle
	return &Configuration{
		endpoint: endpoint,
	}, nil
}
