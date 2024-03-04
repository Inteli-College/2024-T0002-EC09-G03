# Testes do sistema

Para assegurar a robustez e confiabilidade do sistema, uma série de testes automatizados foram implementados. Esses testes são fundamentais para verificar o correto funcionamento das interações entre os componentes do sistema, bem como a integração com serviços externos, como o RabbitMQ. A seguir, detalhamos os procedimentos e testes específicos que compõem nossa suíte de testes, garantindo a qualidade e a eficácia do nosso sistema em operação contínua.

# Workflow de Testes em Go

O arquivo YAML fornecido define um workflow no GitHub Actions que é acionado por eventos de push e pull_request na branch main. A seguir, descreve-se cada parte e passo definidos no workflow:

### Eventos de Gatilho

```
on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]
```

Essa seção estipula que o workflow é executado quando ocorre um push ou um pull_request para a branch main.

### Jobs

```
jobs:
  build:
    name: Build
    runs-on: ubuntu-latest

```

Nesta parte, é definido um job chamado build, que é executado em uma máquina virtual com a versão mais recente do Ubuntu.

## Passos do Job

Cada passo no job build é definido com - name: e explicado em detalhes a seguir:

1. **Configuração do Go:**
- O ambiente virtual é configurado com a versão especificada do Go, utilizando a ação setup-go@v2.

2. **Obtenção do Código Fonte:**
- O código fonte é clonado para o diretório de módulos Go na máquina virtual, através da ação checkout@v2.

3. **Download de Dependências:**
- O comando go mod download é executado para baixar as dependências do módulo Go, assegurando que todas as bibliotecas necessárias estejam disponíveis.

4. **Adição de Dependências Ausentes:**

- Dependências adicionais são instaladas utilizando go get, como no caso do pacote github.com/gin-gonic/gin.
Compilação:
- O projeto é compilado com o comando go build -v ./..., verificando a ausência de erros de compilação.

5. **Testes:**

- Testes são executados com go test -v ./..., garantindo que as mudanças no código não introduzam regressões.

## Passos do Workflow para Testes

### Configuração do Pipeline de CI com GitHub Actions
Para assegurar a qualidade e estabilidade do código, um pipeline de CI foi configurado utilizando GitHub Actions. Este pipeline é ativado automaticamente por qualquer push ou pull_request na branch main. Abaixo, os passos do workflow de CI são descritos em detalhe:

- Instalação do Go: O ambiente de execução é preparado com a versão desejada do Go.
- Clonagem do Código: O código é clonado para a máquina virtual do GitHub Actions.
- Download de Dependências: Todas as dependências necessárias são baixadas para garantir uma compilação bem-sucedida.
- Inclusão de Dependências Ausentes: Bibliotecas adicionais são obtidas conforme necessário para o projeto.
- Compilação do Projeto: O código é compilado para verificar a ausência de erros.
- Execução de Testes: Os testes são realizados para confirmar que não existem regressões ou novos erros introduzidos.

Este pipeline é crucial para manter a integridade do código e deve ser atualizado conforme necessário à medida que o projeto evolui.

# Testes do Pacote `rabbitmq`

Dentro do nosso sistema, utilizamos o RabbitMQ para intermediar as mensagens entre os serviços. Para garantir a correta implementação da nossa integração com o RabbitMQ, incluímos testes específicos no pacote rabbitmq.

## Teste de Geração de Consumidor
Localizado em src/sensors_reading/connections/rabbitmq, o arquivo rabbitmq_test.go contém o teste para a função GenerateConsumer.

## Propósito do Teste
O teste TestGenerateConsumer foi projetado para validar se a função GenerateConsumer está criando de forma eficiente uma instância não nula de um consumidor do RabbitMQ. Isso é essencial para garantir que o sistema possa processar as mensagens recebidas adequadamente.

## Execução do Teste
Para executar este teste, é necessário que o ambiente de teste tenha acesso a uma instância operacional do RabbitMQ. Esta instância pode ser local, para desenvolvimento, ou pode estar situada em um ambiente de integração contínua que imite as condições de produção.

## Resultado Esperado
A expectativa para o teste é que a função GenerateConsumer retorne uma instância não nula do consumidor. Se o consumidor retornado for nulo, o teste falhará, indicando um possível problema na lógica de criação que necessitará ser investigado e corrigido.





