0 - Imagem de Build Padrao
    - Debian Squeeze (Heroku build pack ruby use libssl0.9.2)
    - Dependencias do basicas -> build-essential, ruby1.9.1, ca-certificates, curl

1 - Projeto
    - Criar Workspace
    - Clonar Projeto
    - Existe Dockerfile
    - Copiar buildpacks

2 - Preparar Imagem Builder
    - Docker build -t project-name repositorio-git
    - Montar buildpacks
    - Montar repository

3 - Builder
    - Docker
