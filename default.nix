with import <nixpkgs> {};

buildGoModule rec {
  name = "deploymentmanager-tf";
  version = "0.0.1";
  src = ./.;

  modSha256 = "075hjxy2z0zf2gri44h4cq8khlqfigd66il76ayiipnbqvznf5qx"; 

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
