# -*- mode: ruby -*-
# vi: set ft=ruby :
#

LINUX_BASE_BOX = "ubuntu/focal64"

Vagrant.configure(2) do |config|
	# Compilation and development boxes
	config.vm.define "linux", autostart: true, primary: true do |vmCfg|
		vmCfg.vm.box = LINUX_BASE_BOX
		vmCfg.vm.hostname = "linux"
		vmCfg = configureProviders vmCfg,
			cpus: suggestedCPUCores()

		vmCfg = configureLinuxProvisioners(vmCfg)

		vmCfg.vm.synced_folder '.',
			'/opt/gopath/src/github.com/hashicorp/nomad-driver-lxc'

		vmCfg.vm.provision "shell",
			privileged: false,
			path: './scripts/vagrant-linux-unpriv-bootstrap.sh'
	end
end

def configureLinuxProvisioners(vmCfg)
	vmCfg.vm.provision "shell",
		privileged: true,
		inline: 'rm -f /home/vagrant/linux.iso'

	vmCfg.vm.provision "shell",
		privileged: true,
		path: './scripts/vagrant-linux-priv-go.sh'

    vmCfg.vm.provision "shell",
        privileged: true,
        path: './scripts/vagrant-linux-priv-config.sh'

	return vmCfg
end

def configureProviders(vmCfg, cpus: "2", memory: "4096")
	vmCfg.vm.provider "virtualbox" do |v|
		v.customize ["modifyvm", :id, "--cableconnected1", "on"]
		v.memory = memory
		v.cpus = cpus
	end

	["vmware_fusion", "vmware_workstation"].each do |p|
		vmCfg.vm.provider p do |v|
			v.enable_vmrun_ip_lookup = false
			v.vmx["memsize"] = memory
			v.vmx["numvcpus"] = cpus
		end
	end

	vmCfg.vm.provider "virtualbox" do |v|
		v.customize ["modifyvm", :id, "--cableconnected1", "on"]
		v.memory = memory
		v.cpus = cpus
	end

	return vmCfg
end

def suggestedCPUCores()
	case RbConfig::CONFIG['host_os']
	when /darwin/
		Integer(`sysctl -n hw.ncpu`) / 2
	when /linux/
		Integer(`grep -c ^processor /proc/cpuinfo`) / 2
	else
		2
	end
end
