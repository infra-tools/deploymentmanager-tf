with import <nixpkgs> {};

buildGoModule rec {
  name = "deploymentmanager-tf";
  version = "0.0.1";
  src = ./.;

  modSha256 = "09x57vq395kqaj6w942g19xnggh5s89c8ml67vzphigpzclc35cf"; 

  goPackagePath = "github.com/juliosueiras/deploymentmanager-tf";
  subPackages = [ "." ];

  meta = with stdenv.lib; {
    description = "Deployment Manager to Terraform";
    homepage = https://github.com/juliosueiras/deploymentmanager-tf;
    license = licenses.mit;
    maintainers = with maintainers; [ juliosueiras ];
    platforms = platforms.all;
  };
}
