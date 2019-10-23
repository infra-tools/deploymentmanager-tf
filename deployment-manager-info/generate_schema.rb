require 'yaml'

class String
  def underscore
    self.gsub(/::/, '/').
    gsub(/([A-Z]+)([A-Z][a-z])/,'\1_\2').
    gsub(/([a-z\d])([A-Z])/,'\1_\2').
    tr("-", "_").
    downcase
  end
end

test = YAML::load(File.read("schemas/compute/v1/images.yaml"))
type_info = test["type_info"]


name = type_info["name"]
input = type_info["schema"]["input"]

mainSchema = input["mainSchema"]
schemas = input["schemas"]

puts <<EOF
type #{name.capitalize} struct {
EOF

for i in mainSchema["properties"].keys

  l = i.split('')
  l.first.upcase!
  l = l.join

puts <<EOF

  #{l} string `yaml:"#{i},omitempty" hcl:"#{i.underscore},omitempty"`
EOF
end
puts <<EOF
}
EOF
