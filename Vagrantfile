Vagrant.configure("2") do |config|
  config.vm.box = "bento/ubuntu-22.04" # Imagen optimizada y ligera

  # Configuración del Nodo Maestro
  config.vm.define "k3s-master" do |master|
    master.vm.hostname = "k3s-master"
    master.vm.network "private_network", ip: "192.168.56.10"
    master.vm.provider "virtualbox" do |vb|
      vb.memory = "2048"
      vb.cpus = 2
    end
  end

  # Configuración del Nodo Trabajador
  config.vm.define "k3s-worker" do |worker|
    worker.vm.hostname = "k3s-worker"
    worker.vm.network "private_network", ip: "192.168.56.11"
    worker.vm.provider "virtualbox" do |vb|
      vb.memory = "2048"
      vb.cpus = 2
    end
  end
end
