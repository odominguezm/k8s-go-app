# K3s Local Lab: Go App with CI/CD Pipeline 🚀

Este proyecto es un laboratorio de ingeniería de infraestructura y **SRE** diseñado para demostrar el despliegue de microservicios en un clúster de Kubernetes ligero (K3s) utilizando infraestructura como código (IaC), redes avanzadas y automatización profesional.

## 🏗️ Arquitectura del Proyecto

El laboratorio se divide en tres componentes principales:
* **Infraestructura (IaC):** Aprovisionamiento de nodos mediante **Vagrant** y **VirtualBox**.
* **Orquestación y Redes:** Clúster de **K3s** con **Traefik Ingress Controller** configurado en modo `hostNetwork` para exposición directa de puertos.
* **CI/CD:** Pipeline de **GitHub Actions** que construye imágenes de Docker multi-stage y las distribuye vía **GHCR** (GitHub Container Registry).

## 🛠️ Tecnologías Utilizadas

* **Lenguaje:** Go (Golang) con enfoque en SRE.
* **Contenerización:** Docker (Multi-stage builds para optimización).
* **Orquestador:** K3s (Kubernetes).
* **Ingress Controller:** Traefik v2.10.
* **Automatización:** GitHub Actions & Ansible (en desarrollo).
* **Infraestructura:** Vagrant, VirtualBox, Ubuntu Server.

## 📁 Estructura del Repositorio

* **/app**: Código fuente de la aplicación en Go y Dockerfile.
* **/k8s**: Manifiestos de Kubernetes (Deployments, Services, Ingress, RBAC).
* **/.github/workflows**: Definición del pipeline de Integración Continua.
* **Vagrantfile**: Definición de la infraestructura de nodos.

---

## ✅ Estado Actual: Fase 4 Completada

Actualmente, el proyecto ha superado los retos críticos de conectividad de red y exposición de servicios:

* **Infraestructura Base:** Clúster K3s funcional con gestión remota desde el host (Pop!_OS).
* **Ingress Controller:** Despliegue exitoso de **Traefik** vía DaemonSet, superando conflictos de validación de Helm mediante la aplicación de manifiestos nativos.
* **Resolución de Nombres:** Configuración de resolución DNS local en el host para los dominios:
    * `app.orlando.lab` (Punto de entrada de la aplicación).
    * `traefik.orlando.lab` (Acceso al Dashboard).
* **Observabilidad:** Dashboard de Traefik habilitado para el monitoreo de entrypoints y routers.
* **Pipeline CI:** Automatización de construcción y empuje de imágenes a GHCR operativa.

---

## 🔮 Próximos Pasos (Fase 5)

1.  **Despliegue del Microservicio Go:** Migrar de la imagen de prueba al binario propio desarrollado en Go.
2.  **Continuous Deployment (CD):** Implementar la actualización automática del despliegue al detectar nuevas imágenes en el registro.
3.  **Refactorización IaC:** Migrar la configuración manual de Traefik a playbooks de **Ansible** para garantizar la idempotencia total.
4.  **Monitoreo Avanzado:** Integración de Prometheus y Grafana para métricas de rendimiento del clúster.

---

**Desarrollado por:** Orlando - DevOps & Infrastructure Engineer
