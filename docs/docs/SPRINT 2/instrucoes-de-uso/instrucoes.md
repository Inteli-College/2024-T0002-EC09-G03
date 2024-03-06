# Guia de Acesso ao Sistema de Visualização de Dados dos Sensores no Metabase

Este documento é destinado a usuários que precisam acessar o Metabase para visualização de dados dos sensores. O sistema é projetado para permitir o monitoramento e a gestão de informações coletadas por sensores localizados na cidade de São Paulo.

## Pré-requisitos
- Certifique-se de ter uma conexão com a internet.
- Assegure-se de que você tenha sido autorizado a acessar o sistema com seu e-mail e uma senha temporária fornecida pelo administrador do sistema após o seu cadastro.

## Etapa 1: Configurando Variáveis de Ambiente
Para interagir com o sistema, você precisará configurar seu ambiente com as variáveis necessárias. Como esta documentação é pública, estamos usando credenciais de exemplo abaixo. Você substituirá estas pelos valores reais fornecidos a você de forma segura.

**Para o serviço de leitura de sensores (sensors_reading) e (sensor_simulation):**

1. Abra o arquivo `.env` na pasta de configuração do serviço de leitura.
2. Substitua as variáveis de ambiente pelas credenciais fornecidas. Veja o exemplo:


```
sensors_reading:
RABBITMQ_URL=amqp://consumerSensor:ConsumerSensorPassword@34.201.221.151:5672/
DATABASE_HOST="grupo-03.ctcg2eeii8w3.us-east-1.rds.amazonaws.com"
DATABASE_USER="postgres"
DATABASE_PASSWORD="[senha_bd]"
DATABASE_NAME="postgres"
DATABASE_PORT=5432
```
```
sensors_simulation:
BROKER_URL="34.201.221.151"
BROKER_PORT=1883
RABBIT_USER=publisherSensor
RABBIT_PASSWORD=[senha_rabit]
```
```
DATABASE_HOST="grupo-03.ctcg2eeii8w3.us-east-1.rds.amazonaws.com"
DATABASE_USER="postgres"
DATABASE_NAME="postgres"
DATABASE_PASSWORD="[senha_bd]"
DATABASE_PORT=5432
```
## Etapa 2: Acesso ao Metabase
O Metabase é a plataforma utilizada para a visualização dos dados. Para acessá-la:

1. Digite a URL do Metabase no seu navegador: http://metabase-734478867.us-east-1.elb.amazonaws.com/
2. Faça login usando seu e-mail da Intel e a senha temporária.
3. Após o primeiro login, siga as instruções para criar uma nova senha pessoal e segura.

## Etapa 3: Visualização e Análise de Dados
No Metabase, você terá acesso a várias ferramentas de visualização:

1. Mapa de Sensores: Mostra a localização geográfica dos sensores e pode ser utilizado para avaliar a cobertura e distribuição dos dispositivos.
2. Gráficos de Barras: Fornece uma representação visual da quantidade de dados transmitidos por cada sensor, permitindo detectar padrões ou sensores com desempenho atípico.

## Etapa 4: Compreendendo a Arquitetura do Sistema
Use os diagramas UML dessa documentação para entender as interações entre os componentes do sistema:

- O diagrama UML de sequência detalha o fluxo de dados do usuário final até o banco de dados, passando pelo middleware.
- O diagrama UML de implantação mostra como os componentes do sistema são organizados e como eles interagem em um ambiente de produção.

## Etapa 5: Gerenciamento dos Dados dos Sensores
Responsáveis pelo gerenciamento dos dados dos sensores devem:

- Navegar até a seção de gerenciamento no painel do Metabase.
- Utilizar as ferramentas disponíveis para realizar operações como filtragem, exportação e análise dos dados coletados.

## Video demonstrativo do sistema

O vídeo demonstrativo oferece uma visão geral do processo de acesso e utilização do sistema de visualização de dados dos sensores no Metabase, especificamente projetado para a cidade de São Paulo. 


[![O vídeo apresenta o Metabase do projeto](https://i3.ytimg.com/vi/6M7lcCCiwwQ/maxresdefault.jpg)](https://youtu.be/6M7lcCCiwwQ)


