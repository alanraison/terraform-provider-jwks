package sdkv2provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccJwksFromKeyDataSource(t *testing.T) {
	resourceName := "data.jwks_from_key.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProviders,
		CheckDestroy:             testAccCheckJwksFromKeyDataSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccJwksFromKeyDataSourceConfig(PrivateKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "jwks", `{"d":"LWsr6QVbe-sBibf7z99CdBDDmd5zijr5yueDolVBct4ffapdKE1_yQssb8mPS9phxM88NnD0kIqRMxDSYqP1dH_kMtsEyxXp41de68knfHiIpbl4P1aQsPgUqU1qsFHI7uPyx89Opdbu5q03FZ0UftixJG82H_2nJActYDk6K26J0ythtzcRTHwHUY83In1bxDzOc1gt_Bg6NTha_VIsTqoZBAf1DRbMDZZR_yOSUJjIxhMjLACMZvyRsAmUqGaPa3YiCOaq3C3U9hykaWpu1AENSJZe3lLhhf_WZohD_j3J41RqOOlsmIfcNuElsfAVPM_eln0_0Yc7OUaRhH7BIQ","dp":"GeyHfwYiiQFkaak0QV5b8AYV6Cr4yNh42D9tc_xMmfjxT87q_5TWmYSFJTJ1H4m28BPaRYIEM83_LHDO3O0n_pi6FUnaEP0bo8-wqr0M4tpykrDO7dHLq1CG95SVnnuf1HhGsWXJiLn6DVK86IQFhwSTiDwzOLLFEy0au8GDPbk","dq":"cIFq7k4_yATqaMq5exldZShqGkUsrDDCYlSna9Ck4U8n_pZk6WqRNVPpA8Bw9RtOiEqsXzSGxmpnIN1J9Pel4XHhkuhwGtLnqbo2YdFRgcy8nytCbvpPV-vIdWJ3PzmDfNGZ8F5aDzKpPmgWDgtiB8uw-bWBGM6Gzwt8-8g4uVE","e":"AQAB","kty":"RSA","n":"gUElV5mwqkloIrM8ZNZ72gSCcnSJt7-_Usa5G-D15YQUAdf9c1zEekTfHgDP-04nw_uFNFaE5v1RbHaPxhZYVg5ZErNCa_hzn-x10xzcepeS3KPVXcxae4MR0BEegvqZqJzN9loXsNL_c3H_B-2Gle3hTxjlWFb3F5qLgR-4Mf4ruhER1v6eHQa_nchi03MBpT4UeJ7MrL92hTJYLdpSyCqmr8yjxkKJDVC2uRrr-sTSxfh7r6v24u_vp_QTmBIAlNPgadVAZw17iNNb7vjV7Gwl_5gHXonCUKURaV--dBNLrHIZpqcAM8wHRph8mD1EfL9hsz77pHewxolBATV-7Q","p":"w_Xida8kSv8n86V3qSY_I-fYQ5V-jDtXIE-JhRnS8xzbOzz3v0WSOo5H-o4nJx5eL3Ghb3Gcm0Jn46dHrxinHbm-3RjXv_X6tlbxIYjRSQfHOTSMCTvdqcliF5vC6RCLXuc7R-IWR1Ky6eDEZGtrvt3DyeYABsp9fRUFR_6NluU","q":"qNsw1aSl7WJa27F0DoJdlU9LWerpXcazlJcIdOz_S9QDmSK3RDQTdqfTxRmrxiYI9LEsmkOkvzlnnOBMpnZ3ZOU5qIRfprecRIi37KDAOHWGnlC0EWGgl46YLb7_jXiWf0AGY-DfJJNd9i6TbIDWu8254_erAS6bKMhW_3q7f2k","qi":"hnlS_hnezw7tYpddR7-WrqPu9-JWRHU1b9n9Yspdmmea8nsWDE3h7npFdPsbA1Achjx4OMOewfHOvF36to7E5l4pwwePdfKMJfc0d9OxyewE8AZilj9bHOFPxRicW7pHOSLESVN_aux4Rq_uuR0qQjcsB4hnLsHQe6UoqyapVDs"}`),
				),
			},
			{
				Config: testAccJwksFromKeyDataSourceConfig(PublicKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "jwks", `{"e":"AQAB","kty":"RSA","n":"gUElV5mwqkloIrM8ZNZ72gSCcnSJt7-_Usa5G-D15YQUAdf9c1zEekTfHgDP-04nw_uFNFaE5v1RbHaPxhZYVg5ZErNCa_hzn-x10xzcepeS3KPVXcxae4MR0BEegvqZqJzN9loXsNL_c3H_B-2Gle3hTxjlWFb3F5qLgR-4Mf4ruhER1v6eHQa_nchi03MBpT4UeJ7MrL92hTJYLdpSyCqmr8yjxkKJDVC2uRrr-sTSxfh7r6v24u_vp_QTmBIAlNPgadVAZw17iNNb7vjV7Gwl_5gHXonCUKURaV--dBNLrHIZpqcAM8wHRph8mD1EfL9hsz77pHewxolBATV-7Q"}`),
				),
			},
			{
				Config: testAccJwksFromKeyDataSourceConfig(ECPrivateKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "jwks", `{"crv":"P-384","d":"WL_isXL5n9Ux2yd2phRRMH_aPCZwVfghUEhEU9F0L0EKoi7MD0oRirlsqWuYsNwd","kty":"EC","x":"QnSqFbKGdEiz_g0CoZFXBAvsMMbMBqqnZ3jJ_zrLxuBShDTmrsRoU7MCOALbt04f","y":"FoRI8H06aS1Wq1XsAbMiMLUqpl5joOWeUChxyPb0v-7B86IE9299UHBpbAfcnthd"}`),
				),
			},
			{
				Config: testAccJwksFromKeyDataSourceConfig(ECPublicKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "jwks", `{"crv":"P-384","kty":"EC","x":"QnSqFbKGdEiz_g0CoZFXBAvsMMbMBqqnZ3jJ_zrLxuBShDTmrsRoU7MCOALbt04f","y":"FoRI8H06aS1Wq1XsAbMiMLUqpl5joOWeUChxyPb0v-7B86IE9299UHBpbAfcnthd"}`),
				),
			},

			{
				Config: testAccJwksFromKeyDataSourceConfig(privateKeyDer()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "jwks", `{"d":"LWsr6QVbe-sBibf7z99CdBDDmd5zijr5yueDolVBct4ffapdKE1_yQssb8mPS9phxM88NnD0kIqRMxDSYqP1dH_kMtsEyxXp41de68knfHiIpbl4P1aQsPgUqU1qsFHI7uPyx89Opdbu5q03FZ0UftixJG82H_2nJActYDk6K26J0ythtzcRTHwHUY83In1bxDzOc1gt_Bg6NTha_VIsTqoZBAf1DRbMDZZR_yOSUJjIxhMjLACMZvyRsAmUqGaPa3YiCOaq3C3U9hykaWpu1AENSJZe3lLhhf_WZohD_j3J41RqOOlsmIfcNuElsfAVPM_eln0_0Yc7OUaRhH7BIQ","dp":"GeyHfwYiiQFkaak0QV5b8AYV6Cr4yNh42D9tc_xMmfjxT87q_5TWmYSFJTJ1H4m28BPaRYIEM83_LHDO3O0n_pi6FUnaEP0bo8-wqr0M4tpykrDO7dHLq1CG95SVnnuf1HhGsWXJiLn6DVK86IQFhwSTiDwzOLLFEy0au8GDPbk","dq":"cIFq7k4_yATqaMq5exldZShqGkUsrDDCYlSna9Ck4U8n_pZk6WqRNVPpA8Bw9RtOiEqsXzSGxmpnIN1J9Pel4XHhkuhwGtLnqbo2YdFRgcy8nytCbvpPV-vIdWJ3PzmDfNGZ8F5aDzKpPmgWDgtiB8uw-bWBGM6Gzwt8-8g4uVE","e":"AQAB","kty":"RSA","n":"gUElV5mwqkloIrM8ZNZ72gSCcnSJt7-_Usa5G-D15YQUAdf9c1zEekTfHgDP-04nw_uFNFaE5v1RbHaPxhZYVg5ZErNCa_hzn-x10xzcepeS3KPVXcxae4MR0BEegvqZqJzN9loXsNL_c3H_B-2Gle3hTxjlWFb3F5qLgR-4Mf4ruhER1v6eHQa_nchi03MBpT4UeJ7MrL92hTJYLdpSyCqmr8yjxkKJDVC2uRrr-sTSxfh7r6v24u_vp_QTmBIAlNPgadVAZw17iNNb7vjV7Gwl_5gHXonCUKURaV--dBNLrHIZpqcAM8wHRph8mD1EfL9hsz77pHewxolBATV-7Q","p":"w_Xida8kSv8n86V3qSY_I-fYQ5V-jDtXIE-JhRnS8xzbOzz3v0WSOo5H-o4nJx5eL3Ghb3Gcm0Jn46dHrxinHbm-3RjXv_X6tlbxIYjRSQfHOTSMCTvdqcliF5vC6RCLXuc7R-IWR1Ky6eDEZGtrvt3DyeYABsp9fRUFR_6NluU","q":"qNsw1aSl7WJa27F0DoJdlU9LWerpXcazlJcIdOz_S9QDmSK3RDQTdqfTxRmrxiYI9LEsmkOkvzlnnOBMpnZ3ZOU5qIRfprecRIi37KDAOHWGnlC0EWGgl46YLb7_jXiWf0AGY-DfJJNd9i6TbIDWu8254_erAS6bKMhW_3q7f2k","qi":"hnlS_hnezw7tYpddR7-WrqPu9-JWRHU1b9n9Yspdmmea8nsWDE3h7npFdPsbA1Achjx4OMOewfHOvF36to7E5l4pwwePdfKMJfc0d9OxyewE8AZilj9bHOFPxRicW7pHOSLESVN_aux4Rq_uuR0qQjcsB4hnLsHQe6UoqyapVDs"}`),
				),
			},
			{
				Config: testAccJwksFromKeyDataSourceConfig(publicKeyDer()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "jwks", `{"e":"AQAB","kty":"RSA","n":"gUElV5mwqkloIrM8ZNZ72gSCcnSJt7-_Usa5G-D15YQUAdf9c1zEekTfHgDP-04nw_uFNFaE5v1RbHaPxhZYVg5ZErNCa_hzn-x10xzcepeS3KPVXcxae4MR0BEegvqZqJzN9loXsNL_c3H_B-2Gle3hTxjlWFb3F5qLgR-4Mf4ruhER1v6eHQa_nchi03MBpT4UeJ7MrL92hTJYLdpSyCqmr8yjxkKJDVC2uRrr-sTSxfh7r6v24u_vp_QTmBIAlNPgadVAZw17iNNb7vjV7Gwl_5gHXonCUKURaV--dBNLrHIZpqcAM8wHRph8mD1EfL9hsz77pHewxolBATV-7Q"}`),
				),
			},
			{
				Config: testAccJwksFromKeyDataSourceConfig(ecPrivateKeyDer()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "jwks", `{"crv":"P-384","d":"WL_isXL5n9Ux2yd2phRRMH_aPCZwVfghUEhEU9F0L0EKoi7MD0oRirlsqWuYsNwd","kty":"EC","x":"QnSqFbKGdEiz_g0CoZFXBAvsMMbMBqqnZ3jJ_zrLxuBShDTmrsRoU7MCOALbt04f","y":"FoRI8H06aS1Wq1XsAbMiMLUqpl5joOWeUChxyPb0v-7B86IE9299UHBpbAfcnthd"}`),
				),
			},
			{
				Config: testAccJwksFromKeyDataSourceConfig(ecPublicKeyDer()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "jwks", `{"crv":"P-384","kty":"EC","x":"QnSqFbKGdEiz_g0CoZFXBAvsMMbMBqqnZ3jJ_zrLxuBShDTmrsRoU7MCOALbt04f","y":"FoRI8H06aS1Wq1XsAbMiMLUqpl5joOWeUChxyPb0v-7B86IE9299UHBpbAfcnthd"}`),
				),
			},
			{
				Config: testAccJwksFromKeyWithKidDataSourceConfig(ecPrivateKeyDer(), "123"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "jwks", `{"crv":"P-384","d":"WL_isXL5n9Ux2yd2phRRMH_aPCZwVfghUEhEU9F0L0EKoi7MD0oRirlsqWuYsNwd","kid":"123","kty":"EC","x":"QnSqFbKGdEiz_g0CoZFXBAvsMMbMBqqnZ3jJ_zrLxuBShDTmrsRoU7MCOALbt04f","y":"FoRI8H06aS1Wq1XsAbMiMLUqpl5joOWeUChxyPb0v-7B86IE9299UHBpbAfcnthd"}`),
				),
			},
			{
				Config: testAccJwksFromKeyWithKidDataSourceConfig(ecPublicKeyDer(), "123"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "jwks", `{"crv":"P-384","kid":"123","kty":"EC","x":"QnSqFbKGdEiz_g0CoZFXBAvsMMbMBqqnZ3jJ_zrLxuBShDTmrsRoU7MCOALbt04f","y":"FoRI8H06aS1Wq1XsAbMiMLUqpl5joOWeUChxyPb0v-7B86IE9299UHBpbAfcnthd"}`),
				),
			},
		},
	})
}

func testAccCheckJwksFromKeyDataSourceDestroy(s *terraform.State) error {
	return nil
}

func testAccJwksFromKeyDataSourceConfig(data string) string {
	return fmt.Sprintf(`
data "jwks_from_key" "test" {
  key = <<EOF
%s
EOF
}
	`, data)
}

func testAccJwksFromKeyWithKidDataSourceConfig(data, kid string) string {
	return fmt.Sprintf(`
	data "jwks_from_key" "test" {
		key = <<EOF
%s
EOF
		kid = %s
	}
	`, data, kid)
}
