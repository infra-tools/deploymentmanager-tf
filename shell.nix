with import <nixpkgs> {};

mkShell {
  buildInputs = [
    go_1_12
  ];

  shellHook = ''
    export GO111MODULE=on
    export GOROOT=${pkgs.go_1_12}/share/go
  '';
}
