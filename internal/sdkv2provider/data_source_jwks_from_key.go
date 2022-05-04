package sdkv2provider

import (
	"context"
	"crypto"
	"encoding/hex"
	"encoding/json"

	"github.com/lestrrat-go/jwx/jwk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceJwksFromKey() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceJwksFromKeyRead,
		Schema:      dataSourceJwksFromKeySchema(),
		Description: `Calculates a JSON Web Key Set from a given public or private key.`,
	}
}

func dataSourceJwksFromKeySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"key": {
			Type:        schema.TypeString,
			Required:    true,
			Description: `Requires either a pem encoded or base64 der encoded public or private key.`,
		},
		"kid": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: `Used to populate the kid field of the JWK.`,
		},
		"jwks": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: `The calculated JSON Web Key Sets.`,
		},
	}
}

func dataSourceJwksFromKeyRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	data := d.Get("key").(string)
	keyData, err := readKeyData(data)
	if err != nil {
		return diag.FromErr(err)
	}

	key, err := jwk.New(keyData)
	if err != nil {
		return diag.FromErr(err)
	}
	kid, ok := d.GetOk("kid")
	if ok {
		err = key.Set(jwk.KeyIDKey, kid.(string))
		if err != nil {
			return diag.FromErr(err)
		}
	}
	b, err := json.Marshal(key)
	if err != nil {
		return diag.FromErr(err)
	}
	tb, err := key.Thumbprint(crypto.SHA256)
	if err != nil {
		return diag.Errorf("unable to generate fingerprint: %s", err)
	}
	d.SetId(hex.EncodeToString(tb))
	return diag.FromErr(d.Set("jwks", string(b)))
}
