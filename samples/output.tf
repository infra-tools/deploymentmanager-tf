resource "test-vm" "google_compute_instance" {
  name                = "test"
  can_ip_forward      = false
  deletion_protection = true

  shielded_instance_config {
  }
  metadata = { bitnami-base-password = "redacted", google-logging-enable = "0", google-monitoring-enable = "0" }

  network_interface {
    network    = "https://www.googleapis.com/compute/v1/projects/test/global/networks/test-vpc"
    subnetwork = "https://www.googleapis.com/compute/v1/projects/test/regions/northamerica-northeast1/subnetworks/shared-services-subnet"
  }
}
resource "test-vm-no-2" "google_compute_instance" {
  name                = "test-props-2"
  can_ip_forward      = false
  deletion_protection = true

  shielded_instance_config {
  }
  metadata = { bitnami-base-password = "redacted", google-logging-enable = "0", google-monitoring-enable = "0" }

  network_interface {
    network    = "https://www.googleapis.com/compute/v1/projects/test/global/networks/test-vpc"
    subnetwork = "https://www.googleapis.com/compute/v1/projects/test/regions/northamerica-northeast1/subnetworks/shared-services-subnet"
  }
}
