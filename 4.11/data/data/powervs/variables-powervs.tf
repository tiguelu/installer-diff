################################################################
# Configure the IBM Cloud provider
################################################################
variable "powervs_api_key" {
  type        = string
  description = "IBM Cloud API key associated with user's identity"
  default     = "<key>"
}

variable "powervs_vpc_region" {
  type        = string
  description = "The IBM Cloud region where you want to create the resources"
  default     = "eu-gb"
}

variable "powervs_vpc_zone" {
  type        = string
  description = "The IBM Cloud zone associated with the VPC region you're using"
}

variable "powervs_region" {
  type        = string
  description = "The IBM Cloud region where you want to create the resources"
  default     = "lon"
}

variable "powervs_zone" {
  type        = string
  description = "The IBM Cloud zone associated with the region you're using"
}

variable "powervs_resource_group" {
  type        = string
  description = "The cloud instance resource group"
}

variable "powervs_cloud_instance_id" {
  type        = string
  description = "The cloud instance ID of your account"
}

################################################################
# Configure storage
################################################################
variable "powervs_cos_instance_location" {
  type        = string
  description = "The location of your COS instance"
  default     = "global"
}

variable "powervs_cos_bucket_location" {
  type        = string
  description = "The location to create your COS bucket"
  default     = "us-east"
}

variable "powervs_cos_storage_class" {
  type        = string
  description = "The plan used for your COS instance"
  default     = "smart"
}

################################################################
# Configure instances
################################################################
variable "powervs_bootstrap_memory" {
  type        = string
  description = "Amount of memory, in  GiB, used by the bootstrap node."
  default     = "32"
}

variable "powervs_bootstrap_processors" {
  type        = string
  description = "Number of processors used by the bootstrap node."
  default     = "0.5"
}

variable "powervs_master_memory" {
  type        = string
  description = "Amount of memory, in  GiB, used by each master node."
  default     = "32"
}

variable "powervs_master_processors" {
  type        = string
  description = "Number of processors used by each master node."
  default     = "0.5"
}

variable "powervs_proc_type" {
  type        = string
  description = "The type of processor mode for all nodes (shared/dedicated)"
  default     = "shared"
}

variable "powervs_sys_type" {
  type        = string
  description = "The type of system (s922/e980)"
  default     = "s922"
}

variable "powervs_key_name" {
  type        = string
  description = "The name for the SSH key created in the Service Instance"
  default     = ""
}

variable "powervs_ssh_key" {
  type        = string
  description = "Public key for keypair used to access cluster. Required when creating 'ibm_pi_instance' resources."
  default     = ""
}

variable "powervs_image_bucket_name" {
  type        = string
  description = "Name of the COS bucket containing the image to be imported."
}

variable "powervs_image_bucket_file_name" {
  type        = string
  description = "File name of the image in the COS bucket."
}

variable "powervs_image_storage_type" {
  type        = string
  description = "Storage type used when storing image in Power VS."
  default     = "tier3"
}

variable "powervs_expose_bootstrap" {
  type        = bool
  description = "Setting this to false allows the bootstrap resources to be removed from the cluster load balancers."
  default     = true
}

################################################################
# Configure DNS
################################################################
## TODO: Pass the CIS CRN from the installer program, refer the IBM Cloud code to see the implementation.
variable "powervs_cis_crn" {
  type        = string
  description = "The CRN of CIS instance to use."
}

################################################################
# Output Variables
################################################################
variable "bootstrap_ip" { default = "" }
variable "control_plane_ips" { default = [] }
