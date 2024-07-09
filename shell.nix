{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  name = "benkyou";

  buildInputs = with pkgs; [
    go
    air
    nodejs
    goose
    sqlc
  ];

  shellHook = ''
    exec zsh
  '';
}
