# AZURE GO - API

POC de uma API que retorna um JSON na porta 5000.
**Rotas**
http://localhost:5000/dev


## Docker Instructions  
Cria a imagem do linux alpine com a instalação do go
docker build -t nomedaimagem .

``docker build -t masisiliani/azurego .``

Cria o container 
docker run --publish suaporta:8080 --name nomedocontainer --rm nomedaimagem

``docker run --publish 5000:8080 --name azuregocontainer --rm masisiliani/azurego``

Para o container
docker stop nomedocontainer

``docker stop azuregocontainer``

push para docker.hub
docker login --username nomedoseuusuario -p suasenha

``docker login --username masisiliani -p adm123*``

docker push imagename

``docker push masisiliani/azurego``

## Azure Instructions 
Resource Group - devWomanGo

``az group create --name devWomanGo --location "West Europe"``

Plan Service - azurego

``az appservice plan create --name azurego --resource-group devWomanGo --sku B1 --is-linux``

Web APP - azurego-api

``az webapp create --resource-group devWomanGo --plan azurego --name azurego-api --deployment-container-image-name masisiliani/azurego``

Deletar toda a estrutura criada

``az group delete --name devWomanGo``

## Author

By [Marcela Sisiliani](https://msisiliani.github.io) - 2019
