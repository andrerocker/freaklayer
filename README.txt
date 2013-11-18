-------------------------------------------------------------------------------------------

0 - Imagem de Build Padrao
    - Debian Squeeze (Heroku build pack ruby use libssl0.9.2)
    - Dependencias do basicas -> build-essential, ruby1.9.1, ca-certificates, curl

2 - Preparar Imagem Builder
    - Docker build -t project-name repositorio-git
    - Montar buildpacks
    - Montar repository

-------------------------------------------------------------------------------------------

1 - Projeto
    - Criar Workspace
    - Clonar Projeto
    - Montar buildpacks
    - Montar Cache

2 - Build Project
    - Buildar Projeto (Heroku Buildpacks Por enquanto)
    - Empacotar Projeto (tar.gz) (poderia ser um debian tbm pah!)
    - Remover git, estrutura velha e talves o cache!

3 - Download
    - Download do Pacotao!
    - Docker

-------------------------------------------------------------------------------------------
