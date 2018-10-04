package debug

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"time"
)

func LogDataSource() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRead,

		Schema: map[string]*schema.Schema{
			"data": &schema.Schema{
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Data to log",
				Default:     map[string]string{},
			},
			"line": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Data to log",
			},
			"tag": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Data to log",
				Default:     "default",
			},
		},
	}
}

func dataSourceRead(d *schema.ResourceData, meta interface{}) error {
	conf := meta.(ProviderConf)

	data := d.Get("data").(map[string]interface{})
	line := d.Get("line").(string)
	tag := d.Get("tag").(string)

	now := time.Now().String()

	if line == "" && len(data) == 0 {
		return errors.New("either data or line are required in debug_log resource")
	}

	if line != "" {
		_, err := conf.logFile.WriteString(fmt.Sprintf("[%s][%s] %s", now, tag, line))
		if err != nil {
			return err
		}
	}

	if len(data) == 0 {
		_, err := conf.logFile.WriteString(fmt.Sprintf("[%s][%s] %+v", now, tag, data))
		if err != nil {
			return err
		}
	}

	d.SetId(hash(line))
	return nil
}

func hash(s string) string {
	sha := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sha[:])
}
