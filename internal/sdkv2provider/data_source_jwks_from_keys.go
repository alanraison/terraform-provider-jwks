package sdkv2provider

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/lestrrat-go/jwx/jwk"
)

func dataSourceJwksFromKeys() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceJwksFromKeysRead,
		Schema:      dataSourceJwksFromKeysSchema(),
		Description: `Calculates a JSON Web Key Set from given public or private keys.`,
	}
}

func dataSourceJwksFromKeysSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"key": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"data": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		"jwks": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func dataSourceJwksFromKeysRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	s := jwk.NewSet()

	ks := d.Get("key").(*schema.Set).List()
	for _, raw := range ks {
		k := raw.(map[string]interface{})

		if v, ok := k["data"]; ok {
			keyData, err := readKeyData(v.(string))
			if err != nil {
				return diag.FromErr(err)
			}
			jwk, err := jwk.New(keyData)
			if err != nil {
				return diag.FromErr(err)
			}
			s.Add(jwk)
		}
	}

	jwks, err := json.Marshal(s)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("jwks", string(jwks))

	sha := sha256.New()
	sha.Write(jwks)
	d.SetId(base64.StdEncoding.EncodeToString(sha.Sum(nil)))
	return diag.FromErr(nil)
}
