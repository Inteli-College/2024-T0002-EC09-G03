# Guia de Acesso ao Sistema de Visualização de Dados dos Sensores no Metabase

Este documento é destinado a usuários que precisam acessar o Metabase para visualização de dados dos sensores. O sistema é projetado para permitir o monitoramento e a gestão de informações coletadas por sensores localizados na cidade de São Paulo.

## Pré-requisitos
- Certifique-se de que você possui uma conexão com a Internet.
- Assegure-se de estar autorizado(a) a acessar o sistema com seu e-mail e uma senha temporária, fornecidas pelo administrador após seu cadastro.

## Etapa 1: Configurando Variáveis de Ambiente
Para interagir com o sistema, é necessário configurar seu ambiente com as variáveis necessárias. As instruções abaixo demonstram como configurar estas variáveis para diferentes ambientes de trabalho. Substitua os placeholders pelas credenciais fornecidas de forma segura, que serão ocultadas aqui por questões de segurança.

### Configurações para o serviço de leitura e simulação de sensores:

#### Serviço de Simulação dos Sensores (Publisher)

Abra o arquivo `.env` correspondente ao serviço de simulação dos sensores e atualize-o com as credenciais fornecidas. Existem configurações separadas para execução local e na AWS:

**Local**
```
BROKER_URL=localhost
BROKER_PORT=1883
RABBIT_USER=publisherSensor
RABBIT_PASSWORD= placeholder
MONGODB_URI="mongodb://root:password@localhost:27017/urbanpulsesp?retryWrites=true&connectTimeoutMS=10000&authSource=admin&authMechanism=SCRAM-SHA-1"
```

**AWS**

```
sensors_simulation:
BROKER_URL="34.201.221.151"
BROKER_PORT=1883
RABBIT_USER=publisherSensor
RABBIT_PASSWORD= placeholder
```
Certifique-se de substituir **password** pela senha real fornecida de maneira segura para o seu ambiente específico.

## Etapa 2: Acesso ao Metabase
O Metabase é a plataforma utilizada para a visualização dos dados. Para acessá-la:

1. Digite a URL do Metabase no seu navegador:  http://3.90.150.254:3000/
2. Faça login usando seu e-mail e senha temporária.
3. Após o primeiro login, siga as instruções para criar uma nova senha pessoal e segura.

## Etapa 3: Visualização e Análise de Dados
No Metabase, você terá acesso a várias ferramentas de visualização:

1. Mapa de Sensores: Mostra a localização geográfica dos sensores e pode ser utilizado para avaliar a cobertura e distribuição dos dispositivos.
2. Gráficos de Barras: Fornece uma representação visual da quantidade de dados transmitidos por cada sensor, permitindo detectar padrões ou sensores com desempenho atípico.

## Etapa 4: Executando a Simulação dos Sensores
Para executar a simulação dos sensores, siga estes passos considerando que você já configurou as variáveis de ambiente conforme a Etapa 1. A simulação possui componentes de CONSUMER e PUBLISHER:

**CONSUMER**

```
go run cmd/consumer/main.go <src/simulation/cmd/consumer>
```

**PUBLISHER**

```
go run cmd/publisher/main.go <src/simulation/cmd/publisher> <quantidade_de_sensores>
```

Se a quantidade de sensores não for especificada, o sistema simulará um número padrão de sensores ou todos os disponíveis no banco de dados


## Etapa 5: Compreendendo a Arquitetura do Sistema
Use os diagramas UML dessa documentação para entender as interações entre os componentes do sistema:

- O diagrama UML de sequência detalha o fluxo de dados do usuário final até o banco de dados, passando pelo middleware.
- O diagrama UML de implantação mostra como os componentes do sistema são organizados e como eles interagem em um ambiente de produção.

## Etapa 6: Gerenciamento dos Dados dos Sensores
Responsáveis pelo gerenciamento dos dados dos sensores devem:

- Navegar até a seção de gerenciamento no painel do Metabase.
- Utilizar as ferramentas disponíveis para realizar operações como filtragem, exportação e análise dos dados coletados.

## Video demonstrativo do sistema

O vídeo demonstrativo oferece uma visão geral do processo de acesso e utilização do sistema de visualização de dados dos sensores no Metabase, especificamente projetado para a cidade de São Paulo. 


[![O vídeo apresenta o Metabase do projeto](https://i3.ytimg.com/vi/6M7lcCCiwwQ/maxresdefault.jpg)](https://youtu.be/6M7lcCCiwwQ)


