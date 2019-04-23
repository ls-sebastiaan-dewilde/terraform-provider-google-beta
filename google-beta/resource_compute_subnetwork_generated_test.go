// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccComputeSubnetwork_subnetworkBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeSubnetworkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSubnetwork_subnetworkBasicExample(context),
			},
			{
				ResourceName:      "google_compute_subnetwork.network-with-private-secondary-ip-ranges",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeSubnetwork_subnetworkBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_subnetwork" "network-with-private-secondary-ip-ranges" {
  name          = "test-subnetwork-%{random_suffix}"
  ip_cidr_range = "10.2.0.0/16"
  region        = "us-central1"
  network       = "${google_compute_network.custom-test.self_link}"
  secondary_ip_range {
    range_name    = "tf-test-secondary-range-update1"
    ip_cidr_range = "192.168.10.0/24"
  }
}

resource "google_compute_network" "custom-test" {
  name                    = "test-network-%{random_suffix}"
  auto_create_subnetworks = false
}
`, context)
}

func TestAccComputeSubnetwork_subnetworkLoggingConfigBetaExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeSubnetworkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeSubnetwork_subnetworkLoggingConfigBetaExample(context),
			},
		},
	})
}

func testAccComputeSubnetwork_subnetworkLoggingConfigBetaExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_subnetwork" "subnet-with-logging" {
  provider      = "google-beta" 
  name          = "log-test-subnetwork-%{random_suffix}"
  ip_cidr_range = "10.2.0.0/16"
  region        = "us-central1"
  network       = "${google_compute_network.custom-test.self_link}"

  enable_flow_logs = true
  log_config {
    aggregation_interval = "INTERVAL_10_MIN"
    flow_sampling = 0.5
    metadata = "INCLUDE_ALL_METADATA"
  }
}

resource "google_compute_network" "custom-test" {
  provider                = "google-beta"
  name                    = "log-test-network-%{random_suffix}"
  auto_create_subnetworks = false
}

provider "google-beta"{
  region = "us-central1"
  zone   = "us-central1-a"
}
`, context)
}

func testAccCheckComputeSubnetworkDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_compute_subnetwork" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(rs, "https://www.googleapis.com/compute/beta/projects/{{project}}/regions/{{region}}/subnetworks/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", url, nil)
		if err == nil {
			return fmt.Errorf("ComputeSubnetwork still exists at %s", url)
		}
	}

	return nil
}
