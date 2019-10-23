type Images struct {

  ArchiveSizeBytes string `yaml:"archiveSizeBytes,omitempty" hcl:"archive_size_bytes,omitempty"`

  Deprecated string `yaml:"deprecated,omitempty" hcl:"deprecated,omitempty"`

  Description string `yaml:"description,omitempty" hcl:"description,omitempty"`

  DiskSizeGb string `yaml:"diskSizeGb,omitempty" hcl:"disk_size_gb,omitempty"`

  Family string `yaml:"family,omitempty" hcl:"family,omitempty"`

  ForceCreate string `yaml:"forceCreate,omitempty" hcl:"force_create,omitempty"`

  GuestOsFeatures string `yaml:"guestOsFeatures,omitempty" hcl:"guest_os_features,omitempty"`

  ImageEncryptionKey string `yaml:"imageEncryptionKey,omitempty" hcl:"image_encryption_key,omitempty"`

  LabelFingerprint string `yaml:"labelFingerprint,omitempty" hcl:"label_fingerprint,omitempty"`

  Labels string `yaml:"labels,omitempty" hcl:"labels,omitempty"`

  LicenseCodes string `yaml:"licenseCodes,omitempty" hcl:"license_codes,omitempty"`

  Licenses string `yaml:"licenses,omitempty" hcl:"licenses,omitempty"`

  Name string `yaml:"name,omitempty" hcl:"name,omitempty"`

  RawDisk string `yaml:"rawDisk,omitempty" hcl:"raw_disk,omitempty"`

  RequestId string `yaml:"requestId,omitempty" hcl:"request_id,omitempty"`

  SourceDisk string `yaml:"sourceDisk,omitempty" hcl:"source_disk,omitempty"`

  SourceDiskEncryptionKey string `yaml:"sourceDiskEncryptionKey,omitempty" hcl:"source_disk_encryption_key,omitempty"`

  SourceImage string `yaml:"sourceImage,omitempty" hcl:"source_image,omitempty"`

  SourceImageEncryptionKey string `yaml:"sourceImageEncryptionKey,omitempty" hcl:"source_image_encryption_key,omitempty"`

  SourceSnapshot string `yaml:"sourceSnapshot,omitempty" hcl:"source_snapshot,omitempty"`

  SourceSnapshotEncryptionKey string `yaml:"sourceSnapshotEncryptionKey,omitempty" hcl:"source_snapshot_encryption_key,omitempty"`

  SourceType string `yaml:"sourceType,omitempty" hcl:"source_type,omitempty"`
}
