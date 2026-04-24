# K3s Local Lab: Go App with CI/CD Pipeline 🚀

Este proyecto es un laboratorio de ingeniería de infraestructura y DevOps diseñado para demostrar el despliegue de microservicios en un clúster de Kubernetes ligero (K3s) utilizando infraestructura como código (IaC) y automatización profesional.

## 🏗️ Arquitectura del Proyecto

El laboratorio se divide en tres componentes principales:

1.  **Infraestructura (IaC):** Aprovisionamiento de nodos (Maestro y Worker) mediante **Vagrant** y **VirtualBox**.
2.  **Orquestación:** Clúster de **K3s** configurado para gestión remota desde el host (Pop!_OS).
3.  **CI/CD:** Pipeline de **GitHub Actions** que construye imágenes de Docker multi-stage y las distribuye vía **GHCR** (GitHub Container Registry).

## 🛠️ Tecnologías Utilizadas

* **Lenguaje:** Go (Golang) con Health Checks integrados.
* **Contenerización:** Docker (Multi-stage builds para optimización de tamaño).
* **Orquestador:** K3s (Kubernetes).
* **Automatización:** GitHub Actions.
* **Infraestructura:** Vagrant, VirtualBox, Ubuntu Server.

## 📁 Estructura del Repositorio

* `/app`: Código fuente de la aplicación en Go y Dockerfile.
* `/k8s`: Manifiestos de Kubernetes (Deployments, Services, Secrets).
* `.github/workflows`: Definición del pipeline de Integración Continua.
* `Vagrantfile`: Definición de la infraestructura de nodos.

## 🚀 Estado Actual: Fase 2 Completada
Actualmente, el proyecto cuenta con:
- [x] Clúster K3s funcional con nodos distribuidos.
- [x] Gestión remota de `kubectl` configurada.
- [x] Pipeline de CI exitoso que genera imágenes automáticas.
- [x] Despliegue manual exitoso de 2 réplicas de la aplicación con balanceo de carga vía NodePort.

## 🔮 Próximos Pasos (Fase 3)
* Implementación de **Continuous Deployment (CD)**.
* Configuración de Ingress para gestión de tráfico mediante dominios locales.
* Monitoreo básico de recursos del clúster.

---
**Desarrollado por:** Orlando - DevOps & Infrastructure Engineer
