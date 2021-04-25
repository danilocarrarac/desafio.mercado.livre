# desafio.mercado.livre

URLs:

Post para todos os Satélites:
http://54.198.86.218:32768/topsecret

Post por Satélite:
http://54.198.86.218:32768/topsecret_split/name/skywalker

Get Por Satélite + nome
http://54.198.86.218:32768/topsecret_split/name/{nome-satelite}/distance/{distance}


Stack do Projeto

GO 1.16
docker
ECR
ECS
EC2
DynamoDB


Doc!

Este projeto possui 3 endpoints Rest

1 - POST /topsecret: recebe um body que é encaminahdo para todos os satélites calcularem a posição da nave;
Os dados requeridos no body da request é um array com objetos do tipo Satélite que por sua vez possuem os parâmetros:

name: Referente ao nome do satélite;
distance: Referente a distância da nave até o satélite;
message: Mensagem secreta transmitida pela nave até os satélites da base aliada.;


2 - POST /topsecret_split/name{nome-satelite}: Recebe no path param o nome do satélite e no body da requsição
as demais características como: Distância e Mensagem. A partir desses dados as funcões de cálculo de rota aplicam
suas regras de negócio para determinar a posiçãso da nave de acordo com o nome de cada satélite.
Os parâmetros para utilização deste endpoint são:

name: Referente ao nome do satélite;
distance: Referente a dis6ancia da nave até o satélite;
message: Mensagem secreta transmitida pela nave até os satélites da base aliada. ;



3 GET /topsecret_split/name{nome-satelite}/distance{distance}: Realiza a consulta na base e retorna o cálculo da rota + a mensagem enviada. Em caso de dados não existentes na base o erro retornado é 404;
Os parâmetros para a consulta do GET são:
name: Referente ao nome do satélite;
distance: Referente a dis6ancia da nave até o satélite;




Roadmap de melhorias na API

 - Implementar o Swagger;

 - Testes dos endpoints;

 - Refatoraçao das Funções em escopos mais atômicos;

 - implementação de uma sessão para rotas da api;

- Retirada de regras de negócio dos controllers;

 - Melhoria na regra de negócio de cálculo de rota para abranger a todo o círculo trigonométrico pois os ângulos são fixos;

- Melhorar o docker file para um processo de build mais automatizado

- Adicão de regras de validação do body da requisição para não aceitar tipos diferentes do



Roadmap infra

- Incluir um lambda proxy para roteamtno do API gateway até o ecs. Para adção de uma nível maior de seguraça;

- Poc com NLB para a repeção de 2 params no path param das requests (testes preliminares o NLB não consegue capturar o request path do API Gateway)


- Integração do Code Pipeline com o github para automação do processo de deploy;


