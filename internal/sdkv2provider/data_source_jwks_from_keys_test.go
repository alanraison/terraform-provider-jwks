package sdkv2provider

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccJwksFromKeysDataSource(t *testing.T) {
	resourceName := "data.jwks_from_keys.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckJwksFromKeyDataSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccJwksFromKeys([]string{}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "jwks", `{"keys":[]}`),
				),
			},
			{
				Config: testAccJwksFromKeys([]string{PublicKey}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "jwks", `{"keys":[{"e":"AQAB","kty":"RSA","n":"gUElV5mwqkloIrM8ZNZ72gSCcnSJt7-_Usa5G-D15YQUAdf9c1zEekTfHgDP-04nw_uFNFaE5v1RbHaPxhZYVg5ZErNCa_hzn-x10xzcepeS3KPVXcxae4MR0BEegvqZqJzN9loXsNL_c3H_B-2Gle3hTxjlWFb3F5qLgR-4Mf4ruhER1v6eHQa_nchi03MBpT4UeJ7MrL92hTJYLdpSyCqmr8yjxkKJDVC2uRrr-sTSxfh7r6v24u_vp_QTmBIAlNPgadVAZw17iNNb7vjV7Gwl_5gHXonCUKURaV--dBNLrHIZpqcAM8wHRph8mD1EfL9hsz77pHewxolBATV-7Q"}]}`),
				),
			},
			{
				Config: testAccJwksFromKeys([]string{PublicKey, PrivateKey}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "jwks", `{"keys":[{"e":"AQAB","kty":"RSA","n":"gUElV5mwqkloIrM8ZNZ72gSCcnSJt7-_Usa5G-D15YQUAdf9c1zEekTfHgDP-04nw_uFNFaE5v1RbHaPxhZYVg5ZErNCa_hzn-x10xzcepeS3KPVXcxae4MR0BEegvqZqJzN9loXsNL_c3H_B-2Gle3hTxjlWFb3F5qLgR-4Mf4ruhER1v6eHQa_nchi03MBpT4UeJ7MrL92hTJYLdpSyCqmr8yjxkKJDVC2uRrr-sTSxfh7r6v24u_vp_QTmBIAlNPgadVAZw17iNNb7vjV7Gwl_5gHXonCUKURaV--dBNLrHIZpqcAM8wHRph8mD1EfL9hsz77pHewxolBATV-7Q"},{"d":"LWsr6QVbe-sBibf7z99CdBDDmd5zijr5yueDolVBct4ffapdKE1_yQssb8mPS9phxM88NnD0kIqRMxDSYqP1dH_kMtsEyxXp41de68knfHiIpbl4P1aQsPgUqU1qsFHI7uPyx89Opdbu5q03FZ0UftixJG82H_2nJActYDk6K26J0ythtzcRTHwHUY83In1bxDzOc1gt_Bg6NTha_VIsTqoZBAf1DRbMDZZR_yOSUJjIxhMjLACMZvyRsAmUqGaPa3YiCOaq3C3U9hykaWpu1AENSJZe3lLhhf_WZohD_j3J41RqOOlsmIfcNuElsfAVPM_eln0_0Yc7OUaRhH7BIQ","dp":"GeyHfwYiiQFkaak0QV5b8AYV6Cr4yNh42D9tc_xMmfjxT87q_5TWmYSFJTJ1H4m28BPaRYIEM83_LHDO3O0n_pi6FUnaEP0bo8-wqr0M4tpykrDO7dHLq1CG95SVnnuf1HhGsWXJiLn6DVK86IQFhwSTiDwzOLLFEy0au8GDPbk","dq":"cIFq7k4_yATqaMq5exldZShqGkUsrDDCYlSna9Ck4U8n_pZk6WqRNVPpA8Bw9RtOiEqsXzSGxmpnIN1J9Pel4XHhkuhwGtLnqbo2YdFRgcy8nytCbvpPV-vIdWJ3PzmDfNGZ8F5aDzKpPmgWDgtiB8uw-bWBGM6Gzwt8-8g4uVE","e":"AQAB","kty":"RSA","n":"gUElV5mwqkloIrM8ZNZ72gSCcnSJt7-_Usa5G-D15YQUAdf9c1zEekTfHgDP-04nw_uFNFaE5v1RbHaPxhZYVg5ZErNCa_hzn-x10xzcepeS3KPVXcxae4MR0BEegvqZqJzN9loXsNL_c3H_B-2Gle3hTxjlWFb3F5qLgR-4Mf4ruhER1v6eHQa_nchi03MBpT4UeJ7MrL92hTJYLdpSyCqmr8yjxkKJDVC2uRrr-sTSxfh7r6v24u_vp_QTmBIAlNPgadVAZw17iNNb7vjV7Gwl_5gHXonCUKURaV--dBNLrHIZpqcAM8wHRph8mD1EfL9hsz77pHewxolBATV-7Q","p":"w_Xida8kSv8n86V3qSY_I-fYQ5V-jDtXIE-JhRnS8xzbOzz3v0WSOo5H-o4nJx5eL3Ghb3Gcm0Jn46dHrxinHbm-3RjXv_X6tlbxIYjRSQfHOTSMCTvdqcliF5vC6RCLXuc7R-IWR1Ky6eDEZGtrvt3DyeYABsp9fRUFR_6NluU","q":"qNsw1aSl7WJa27F0DoJdlU9LWerpXcazlJcIdOz_S9QDmSK3RDQTdqfTxRmrxiYI9LEsmkOkvzlnnOBMpnZ3ZOU5qIRfprecRIi37KDAOHWGnlC0EWGgl46YLb7_jXiWf0AGY-DfJJNd9i6TbIDWu8254_erAS6bKMhW_3q7f2k","qi":"hnlS_hnezw7tYpddR7-WrqPu9-JWRHU1b9n9Yspdmmea8nsWDE3h7npFdPsbA1Achjx4OMOewfHOvF36to7E5l4pwwePdfKMJfc0d9OxyewE8AZilj9bHOFPxRicW7pHOSLESVN_aux4Rq_uuR0qQjcsB4hnLsHQe6UoqyapVDs"}]}`),
				),
			},
		},
	})
}

func testAccJwksFromKeys(key []string) string {
	builder := strings.Builder{}
	builder.WriteString(`data "jwks_from_keys" "test" {
`)
	for _, k := range key {
		builder.WriteString(fmt.Sprintf(`
		key {
			data = <<EOF
%s
EOF
		}
		`, k))
	}
	builder.WriteString("}")
	return builder.String()
}
