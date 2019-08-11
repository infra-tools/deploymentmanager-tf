require 'yaml'
require 'fileutils'

YAML.load_stream File.read("gcp_types.yaml") do |gcp_type| 
  matches = gcp_type["provider"].match "gcp-types/(.*)"
  version = matches[1].split("-")[-1]
  type = matches[1].split("-")[0..-2][0]

  FileUtils.mkdir_p "schemas/#{type}/#{version}"

  gcp_type["types"].each do |schema_type|
    `gcloud beta deployment-manager types describe #{schema_type} --project gcp-types --provider #{matches[1]} > schemas/#{type}/#{version}/#{schema_type}.yaml`
  end
end
