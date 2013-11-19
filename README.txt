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
    - Criar Workspace (estrutura necessaria para build) // curl -X POST http://localhost:3000/workspace/cegonha -d 'repo: git@bacon.com:bacon.git'
    - Clonar Projeto

2 - Build Project
    - Buildar Projeto (Heroku Buildpacks Por enquanto)
    - Empacotar Projeto (tar.gz) (poderia ser um debian tbm pah!)
    - Remover git, estrutura velha e talves o cache!

3 - Download
    - Download do Pacotao!
    - Docker

-------------------------------------------------------------------------------------------

curl -X POST http://localhost:8080/workspace/cegonha/branch/1.9.3 -d "repository=git@code.locaweb.com.br:iaas/cegonha.git"
curl -X POST http://localhost:8080/build/cegonha/image/locaweb-ruby
