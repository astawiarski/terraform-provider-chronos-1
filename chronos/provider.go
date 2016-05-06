package chronos

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"time"
)

// Provider is the provider for terraform
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Chronos's Base HTTP URL",
			},
			"request_timeout": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     5,
				Description: "'Request Timeout",
			},
			"debug": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Debug option for chronos",
			},
			"deployment_timeout": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     600,
				Description: "'Deployment Timeout",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"chronos_job": resourceChronosJob(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := config{
		URL:                      d.Get("url").(string),
		RequestTimeout:           d.Get("request_timeout").(int),
		DefaultDeploymentTimeout: time.Duration(d.Get("deployment_timeout").(int)) * time.Second,
		Debug: d.Get("debug").(bool),
	}

	if err := config.loadAndValidate(); err != nil {
		return nil, err
	}

	return config, nil
}
