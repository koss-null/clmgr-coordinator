#Define the list of machines
slurm_cluster = {
    :node1 => {
        :hostname => "node1",
        :ipaddress => "10.10.10.3"
    },
    :node2 => {
        :hostname => "node2",
        :ipaddress => "10.10.10.4"
    },
    :node3 => {
        :hostname => "node3",
        :ipaddress => "10.10.10.5"
    }
}

$provision_script = <<SCRIPT
echo "10.10.10.3    node1" >> /etc/hosts
echo "10.10.10.4    node2" >> /etc/hosts
echo "10.10.10.5    node3" >> /etc/hosts

#provision dependencies
sudo zypper addrepo http://download.opensuse.org/distribution/11.4/repo/oss/ oss
sudo zypper addrepo http://download.opensuse.org/distribution/openSUSE-stable/repo/oss/ oss1
sudo zypper update
sudo zypper in -y go go-doc

sudo chmod -R 777 /tmp
sudo mkdir -p /go/src/myproj.com/clmgr-lrm/
sudo mkdir -p /opt/clmgr/config/
sudo chmod 777 -R /go
export GOPATH=/go/
export PATH=$PATH+:/usr/bin/go

SCRIPT

Vagrant.configure("2") do |global_config|
    slurm_cluster.each_pair do |name, options|
        global_config.vm.define name do |config|
            # to accure notmally running vms
            # u should manually create /go dir and change
            # permissions got /opt and /go to 777
            # then perform vagrant up again


            # workaround of ssh problem
            config.ssh.private_key_path = "~/.ssh/id_rsa"
            #config.ssh.forward_agent = true
            config.ssh.username = "root"
            config.ssh.password = "vagrant"
            #config.ssh.insert_key = true

            config.vm.box = "suse/sles12sp1"
            config.vm.hostname = "#{name}"
            config.vm.network :private_network, ip: options[:ipaddress]

            config.vm.synced_folder '.', '/vagrant'

            config.vm.provider :virtualbox do |v|
                v.customize ["modifyvm", :id, "--memory", "1536"]
            end

            config.vm.provision :shell, :inline => $provision_script
        end
    end
end