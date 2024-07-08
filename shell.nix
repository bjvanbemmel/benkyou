{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  name = "benkyou";

  buildInputs = with pkgs; [
    go
    air
    nodejs
  ];

  shellHook = ''
    exec zsh
  '';
}
