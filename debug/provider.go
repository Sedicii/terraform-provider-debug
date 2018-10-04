package debug

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"os"
)

type ProviderConf struct {
	logFile *os.File
}

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"log_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "/tmp/tf-debug-provider.log",
				Description: "file to write logs",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"debug_log": LogDataSource(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	logFilePath := d.Get("log_file").(string)
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return ProviderConf{
		logFile,
	}, nil
}
