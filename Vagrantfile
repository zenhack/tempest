# -*- mode: ruby -*-
# vi: set ft=ruby :

gopath = "/home/vagrant/go"

Vagrant.configure("2") do |config|
  config.vm.box = "fedora/25-cloud-base"
  config.vm.network :private_network, ip: "169.254.254.2"
  config.vm.synced_folder ".", "/#{gopath}/src/zenhack.net/go/sandstorm-next", type: "nfs"
  config.vm.provision "shell", inline: <<-EOF
    set -exuo pipefail
    dnf install -y make gcc go
    echo 'export GOPATH="${HOME}/go"' >> ~vagrant/.bashrc

    # This is owned by root by default, which prevents go build from working.
    chown vagrant:vagrant ~vagrant/go
  EOF
end
