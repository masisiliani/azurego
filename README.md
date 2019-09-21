# AZURE GO - API

POC de uma API que retorna um JSON na porta 5000.
**Rotas**
http://localhost:5000/devas


## Docker Instructions  
Cria a imagem do linux alpine com a instalação do go
docker build -t nomedaimagem .

``docker build -t masisiliani/azurego .``

Cria o container 
docker run --publish suaporta:5000 --name nomedocontainer --rm nomedaimagem
Estamos informando que a porta 5000 no Host é aberta e deve ser mapeada na porta 5000 do container.
``docker run --publish 5000:5000 --name azuregocontainer --rm masisiliani/azurego``

Parar o container
docker stop nomedocontainer

``docker stop azuregocontainer``

push para docker.hub
docker login --username nomedoseuusuario -p suasenha

``docker login --username [nomedoseuusuario] -p [suasenha]``

Para não expor a senha usar.

``docker login`` 


docker push imagename

``docker push masisiliani/azurego``

## Azure Instructions 

Acessar o [Portal Azure](https://portal.azure.com/#home) (Nesse momento é necessário já ter o cadastro da conta)
E utilizar o cli para executar os comandos abaixo: 

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

## Referências

[Comandos Docker](https://woliveiras.com.br/posts/comandos-mais-utilizados-no-docker/)
[Cadastro no Portal Azure](https://azure.microsoft.com/pt-br/free/)
[Portal Azure](https://portal.azure.com/#home)
[Docker Hub](https://hub.docker.com)
[Dicionário de comandos da Azure](https://docs.microsoft.com/en-us/cli/azure/group?view=azure-cli-latest)