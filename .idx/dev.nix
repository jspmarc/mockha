# To learn more about how to use Nix to configure your environment
# see: https://developers.google.com/idx/guides/customize-idx-env
{ pkgs, ... }: {
  channel = "stable-23.11"; # "stable-23.11" or "unstable"
  # Use https://search.nixos.org/packages to  find packages
  packages = [
    pkgs.go
    pkgs.nodejs_18
    pkgs.nodePackages.nodemon
  ];
  # Sets environment variables in the workspace
  env = {};
  # search for the extension on https://open-vsx.org/ and use "publisher.id"
  idx.extensions = [
    "golang.go"
  ];
  # preview configuration, identical to monospace.json
  idx.previews = {
    enable = false;
    previews = [
      {
        command = [
          "nodemon"
          "--signal" "SIGHUP"
          "-w" "."
          "-e" "go,html"
          "-x" "go run main.go -addr localhost:$PORT"
        ];
        manager = "web";
        id = "web";
      }
    ];
  };
}