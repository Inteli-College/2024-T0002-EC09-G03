# Testes de Infraestrutura

## Teste DB

Este documento descreve o script de teste contido no arquivo `db_test.go`em `infra`, destinado a validar a integridade e funcionalidade das conexões de banco de dados dentro de um ambiente de desenvolvimento ou CI/CD. Os testes são essenciais para garantir que a aplicação possa interagir corretamente com o banco de dados MongoDB, realizando operações de leitura e escrita conforme esperado, sem interrupções ou falhas devido a problemas de conexão ou configuração.

### Objetivo

O principal objetivo deste script de teste é assegurar a confiabilidade e robustez da função `NewDBConnection`, que é responsável por estabelecer uma conexão entre a aplicação e o banco de dados MongoDB. Isso inclui a validação da criação do objeto de conexão do banco de dados e do cliente do banco de dados, assegurando que nenhum deles seja nulo, o que indicaria falha na tentativa de conexão.

### Fluxo do Teste

1. **Configuração do Ambiente**: O teste começa configurando o ambiente necessário, definindo a string de conexão ao banco de dados MongoDB através de uma variável de ambiente. Esta string contém informações essenciais, como o URI do banco de dados, credenciais, timeout de conexão, e outras opções relevantes.

2. **Execução do Teste**:
   - A função `NewDBConnection` é chamada para estabelecer a conexão com o banco de dados e retornar o objeto de conexão (`db`) e o cliente do banco de dados (`client`).
   - São realizadas verificações para assegurar que nem o `db` nem o `client` sejam nulos, o que confirmaria a sucesso na conexão com o banco de dados.

3. **Verificações Adicionais** (opcional): O teste sugere a possibilidade de incluir verificações adicionais, como a execução de comandos simples no banco de dados para validar ainda mais a integridade da conexão.

### Conclusão

A execução bem-sucedida deste script de teste garante que a função `NewDBConnection` é capaz de estabelecer conexões válidas e operacionais com o banco de dados MongoDB. Isso é crítico para o desenvolvimento e operação contínua de aplicações que dependem de interações constantes com o banco de dados. Os testes fornecem uma camada essencial de segurança, assegurando que mudanças no código base ou no ambiente de execução não afetem a capacidade da aplicação de se comunicar efetivamente com o banco de dados.

Este documento fornece uma visão geral e orientações sobre a importância, estrutura e execução dos testes contidos no `db_test.go`, fundamentais para manter a qualidade e a estabilidade das funcionalidades de banco de dados da aplicação.

## Teste MQTT 

Este documento descreve a estrutura e o propósito dos testes implementados no arquivo `mqtt_test.go` em `infra`, especificamente projetados para validar a funcionalidade de conexões MQTT dentro da infraestrutura de uma aplicação. A comunicação MQTT é fundamental para aplicações que exigem troca de mensagens eficiente e confiável em ambientes de Internet das Coisas (IoT), telemetria, e mais. Garantir a robustez dessa conexão é vital para a operação suave e confiável desses sistemas.

### Objetivo

O principal objetivo desses testes é assegurar que as conexões MQTT possam ser estabelecidas corretamente, utilizando as configurações fornecidas por variáveis de ambiente, e que tais conexões sejam capazes de suportar as operações de envio e recebimento de mensagens sem falhas. Os testes visam verificar não apenas a capacidade de estabelecer essas conexões, mas também a integridade das sessões de comunicação criadas.

### Fluxo do Teste

1. **Configuração do Ambiente**: Antes da execução dos testes, o ambiente é preparado pela função `setupTestEnvironment`, que configura as variáveis de ambiente necessárias para a conexão MQTT, incluindo URL do broker, porta, usuário e senha.

2. **Teste de Conexão MQTT** (`TestNewMQTTConnection`):
   - O teste inicia definindo um nome para o cliente MQTT e preparando um `sync.WaitGroup` para gerenciar operações concorrentes, se necessário.
   - Invoca a função `NewMQTTConnection`, passando os parâmetros necessários, para tentar estabelecer uma nova conexão MQTT.
   - Verifica se a conexão foi estabelecida com sucesso e se as operações subsequentes, como publicar e subscrever mensagens, podem ser realizadas sem erros.

### Conclusão

A conclusão bem-sucedida dos testes no `mqtt_test.go` indica que as conexões MQTT dentro da infraestrutura da aplicação são robustas, confiáveis e estão configuradas corretamente. Isso garante que a aplicação possa efetivamente comunicar-se usando o protocolo MQTT, fundamental para sistemas que dependem de comunicação em tempo real ou IoT. A documentação detalha a importância e o processo destes testes, sublinhando seu papel na manutenção da qualidade e estabilidade das funcionalidades relacionadas à comunicação MQTT na aplicação.

## Teste Queue

#### Introdução

Este documento descreve os testes implementados no arquivo `queue_test.go`, pertencente ao pacote `infra`. Os testes são projetados para validar a funcionalidade de um adaptador de fila, especificamente a criação e manipulação de consumidores de fila dentro de um sistema de mensagens. Esses testes asseguram a robustez e a confiabilidade das operações de fila, que são cruciais para o processamento assíncrono de tarefas e comunicação entre diferentes partes de uma aplicação.

#### Objetivo

O principal objetivo destes testes é verificar a capacidade do `QueueAdapter` de criar corretamente um consumidor para uma fila especificada e limpar os recursos após o teste. Isso inclui:

- A correta declaração de uma fila de teste.
- A verificação de que um consumidor pode ser gerado para a fila especificada.
- A limpeza adequada da fila de teste e o fechamento da conexão após a execução do teste.

#### Fluxo do Teste

1. **Configuração Inicial (`setUp`)**:
   - Criação de uma instância do `QueueAdapter`.
   - Declaração de uma fila de teste com configurações específicas (não durável, auto-deletável, não exclusiva).
   - Verificação e tratamento de possíveis erros durante a declaração da fila.

2. **Execução do Teste (`TestGenerateConsumer`)**:
   - Inicialização do ambiente de teste usando `setUp` para preparar os recursos necessários.
   - Chamada do método `GenerateConsumer` no adaptador de fila para gerar um consumidor para a fila de teste.
   - Verificação de que o consumidor foi adicionado ao mapa de canais de mensagens do adaptador de fila.
   - Utilização de asserções para garantir que o consumidor foi criado corretamente.

3. **Limpeza (`tearDown`)**:
   - Exclusão da fila de teste e verificação de erros.
   - Fechamento da conexão com o sistema de mensagens.
   - Garantia de que todos os recursos são liberados adequadamente após o teste.

#### Conclusão

A execução bem-sucedida do teste `TestGenerateConsumer` assegura que o `QueueAdapter` funciona conforme esperado em termos de gerenciamento de consumidores de fila. Isso valida a implementação do adaptador de fila, garantindo que ele possa criar e gerenciar consumidores de forma eficiente, o que é fundamental para o processamento de mensagens e a comunicação assíncrona dentro da aplicação. Este teste contribui para a confiança na estabilidade e confiabilidade do sistema de filas implementado, permitindo que a aplicação processe tarefas de forma assíncrona e eficiente.

# Testes da Repository

## Teste Sensor

Este documento fornece uma visão geral dos testes implementados no arquivo `sensor_test.go`, que faz parte do pacote `repository`. Os testes visam validar a funcionalidade relacionada à geração e manipulação de sensores dentro de um domínio específico, utilizando uma abordagem baseada em mock para simular interações com a porta de sensores (`SensorPort`). A utilização de mocks permite verificar o comportamento esperado das funções em teste sem a necessidade de interações reais com as dependências externas.

#### Objetivo

O objetivo principal desses testes é assegurar que o repositório de sensores funciona corretamente sob diferentes condições, incluindo a geração de sensores com base em critérios específicos e o tratamento de entradas inválidas ou estados inesperados das dependências. Isso é alcançado por meio de várias funções de teste que verificam:

- A capacidade de gerar uma quantidade específica de sensores.
- O tratamento correto de entradas inválidas.
- A robustez do sistema em face de estados inesperados nas dependências.
- A integração e cooperação adequadas entre diferentes componentes do sistema.

#### Fluxo do Teste

1. **Configuração e Mocking**:
   - Criação de um mock para `SensorPort`, permitindo simular o comportamento de dependências externas.
   - Definição de comportamentos esperados para os métodos `GetAllSensors` e `CreateSensor` utilizando o mock.

2. **Teste de Geração de Sensores**:
   - Verificação da capacidade de gerar uma quantidade específica de sensores e da presença dos sensores gerados na estrutura de dados esperada.

3. **Teste com Entradas Inválidas**:
   - Avaliação da resposta do sistema a várias entradas inválidas, incluindo números negativos, zero e números excessivamente altos.

4. **Teste de Estados Inesperados das Dependências**:
   - Simulação de um retorno nulo por parte das dependências para testar a resiliência do sistema em face de falhas ou comportamentos inesperados.

5. **Teste de Integração Entre Componentes**:
   - Teste da cooperação entre o repositório de sensores e suas dependências, verificando se a lógica de negócios funciona conforme esperado quando as condições são ideais.

#### Conclusão

A conclusão bem-sucedida desses testes garante que o repositório de sensores e sua integração com as portas de sensor funcionem de maneira eficaz, tratando corretamente entradas inválidas e sendo resilientes a estados inesperados das dependências. Isso contribui para a estabilidade e confiabilidade do sistema, permitindo que ele gerencie e manipule dados de sensores de maneira eficiente. Este conjunto de testes desempenha um papel crucial na manutenção da qualidade e na prevenção de regressões no sistema de gerenciamento de sensores.

## Teste SensorData

Este documento descreve os testes implementados para verificar a funcionalidade de criação de dados de sensor a partir de JSON no arquivo pertencente ao pacote `repository`. Os testes cobrem cenários de sucesso e falha na deserialização de dados JSON para uma estrutura de dados de sensor, validando assim a robustez e a confiabilidade do processo de parseamento de dados de entrada.

#### Objetivo

O principal objetivo destes testes é assegurar que a função `NewSensorData` possa corretamente converter um JSON de entrada em uma estrutura de dados de sensor (`SensorData`) quando o JSON está bem formado e identificar e reportar um erro de forma apropriada quando o JSON está malformado.

#### Fluxo dos Testes

**Teste de Sucesso (`TestNewSensorDataSuccess`)**:

1. **Preparação dos Dados de Teste**:
   - Cria um payload JSON bem formado representando os dados de um sensor.

2. **Execução da Função de Teste**:
   - Chama a função `NewSensorData` com o payload JSON de teste.

3. **Verificações**:
   - Confirma que nenhum erro foi retornado pela função.
   - Verifica que o objeto resultante não é nulo.
   - Assegura que o ID do sensor no objeto resultante corresponde ao esperado.
   - Realiza verificações adicionais conforme necessário para outros campos do objeto resultante.

**Teste de Falha (`TestNewSensorDataFailure`)**:

1. **Preparação dos Dados de Teste**:
   - Cria um payload JSON malformado (por exemplo, com erro de sintaxe).

2. **Execução da Função de Teste**:
   - Chama a função `NewSensorData` com o payload JSON inválido.

3. **Verificações**:
   - Confirma que um erro foi retornado pela função devido ao JSON malformado.
   - Verifica que o objeto resultante é nulo, indicando que a deserialização falhou.

#### Conclusão

A execução bem-sucedida destes testes confirma a capacidade do sistema de corretamente processar e converter dados de entrada JSON em estruturas de dados de sensor, bem como de maneira adequada manejar dados de entrada inválidos. Isso garante a fiabilidade do processo de deserialização, um componente crítico na recepção e processamento de dados de sensor. Garantir a acurácia e a robustez desse processo é essencial para a integridade dos dados e a operação eficiente da aplicação.

## Teste Generator

#### Introdução

Este documento fornece detalhes sobre uma série de testes implementados para validar a funcionalidade de geração de sensores no contexto de um `repository`, utilizando mocks para simular interações com interfaces externas. Esses testes são fundamentais para garantir a correta geração e manipulação de entidades de sensores, uma operação crítica em sistemas que dependem de dados de sensores para funcionamento e análise.

#### Objetivo

O principal objetivo desses testes é assegurar que o repositório de sensores possa gerar sensores de forma correta e robusta, lidando apropriadamente com diferentes cenários, incluindo entradas válidas e inválidas, e estados inesperados das dependências. Isso é crucial para manter a integridade dos dados e a estabilidade do sistema.

#### Fluxo dos Testes

**Teste Básico de Geração (`TestGenerateSensors`):**

1. **Mock Setup:** Configuração do mock para a interface `SensorPort`, simulando a disponibilidade de sensores.
2. **Geração de Sensores:** Chamada da função de geração com um número específico de sensores a serem gerados.
3. **Verificações:** Assegura que o número correto de sensores foi gerado e que os objetos resultantes são válidos.

**Teste com Entradas Inválidas (`TestGenerateSensorsInvalidInput`):**

1. **Cenários de Teste:** Definição de casos de teste para diferentes valores de entrada inválidos (números negativos, zero, números excessivamente altos).
2. **Execução e Verificação:** Para cada caso, verifica-se que a função de geração retorna `nil`, indicando o tratamento adequado de entradas inválidas.

**Teste com Estados Inesperados das Dependências (`TestGenerateSensorsUnexpectedDependencyState`):**

1. **Mock Setup:** Configuração do mock para retornar `nil`, simulando uma falha ou estado inesperado na dependência.
2. **Execução e Verificação:** Verifica-se que a função de geração retorna `nil` nesse cenário, indicando a resiliência do sistema a falhas das dependências.

**Teste de Integração entre Componentes (`TestIntegrationBetweenComponents`):**

1. **Mock Setup:** Configuração do mock para simular a disponibilidade de sensores.
2. **Geração de Sensores:** Chamada da função de geração com um número de sensores que a lógica de negócios permite gerar sem problemas.
3. **Verificações:** Confirma que os sensores são gerados conforme esperado quando as condições são ideais.

#### Conclusão

A execução bem-sucedida desses testes confirma a capacidade do repositório de gerar entidades de sensores de forma eficaz e robusta, tratando corretamente entradas inválidas e sendo resiliente a estados inesperados das dependências. Isso é essencial para garantir a confiabilidade e estabilidade do sistema de gerenciamento de sensores, permitindo que a aplicação opere de maneira eficiente e precisa.

# Testes dos Secondary Adapters

## Teste MQTT

Este documento aborda os testes implementados no pacote `mqtt`, especificamente focados na funcionalidade de um adaptador MQTT (`MQTTAdapter`). Estes testes visam validar a criação de clientes MQTT e a publicação de mensagens em tópicos específicos, assegurando que a comunicação via MQTT funcione de maneira esperada, que é crucial para aplicações que dependem de comunicação em tempo real ou IoT.

#### Objetivo

Os principais objetivos destes testes incluem:

- **TestCreateClient:** Garantir que um cliente MQTT pode ser criado corretamente, e que os atributos do cliente (como o nome) são configurados como esperado.
- **TestPublish:** Verificar que mensagens podem ser publicadas em tópicos específicos sem causar erros ou pânico, validando assim a funcionalidade de publicação do adaptador.

#### Fluxo dos Testes

**TestCreateClient:**

1. **Criação de Cliente:** Inicializa um `MQTTAdapter` e chama `CreateClient` com um nome de cliente específico.
2. **Verificações:** 
   - Confirma que o objeto cliente dentro do adaptador não é nulo após a criação.
   - Assegura que o nome do cliente configurado no adaptador corresponde ao nome fornecido na criação.

**TestPublish:**

1. **Configuração Inicial:** Cria um `MQTTAdapter` e simula a conexão com um broker MQTT, utilizando um nome de cliente específico.
2. **Publicação de Mensagem:** Tenta publicar uma mensagem em um tópico específico.
3. **Verificações:** 
   - Verifica que a tentativa de publicação não causa pânico.
   - (Opcional) Adicionalmente, pode-se verificar se a mensagem foi efetivamente recebida por um assinante, para testar a entrega da mensagem.

#### Conclusão

A conclusão bem-sucedida desses testes assegura que o `MQTTAdapter` está apto a criar clientes MQTT e publicar mensagens em tópicos sem incidentes. Isso é essencial para a integração eficaz de componentes em sistemas que utilizam comunicação MQTT, como aplicações IoT, garantindo que a comunicação entre dispositivos e serviços ocorra de forma confiável e eficiente. Estes testes são um passo crucial na validação da funcionalidade de comunicação do adaptador MQTT, contribuindo para a robustez e a confiabilidade da aplicação como um todo.

## Teste Sensor

Este documento detalha os testes implementados para as funcionalidades de criação e recuperação de sensores dentro do pacote `sensor`, utilizando MongoDB como banco de dados. Os testes validam o correto funcionamento do adaptador de dados do sensor (`SensorDataAdapter`), garantindo que os sensores possam ser criados e recuperados de maneira eficaz, o que é fundamental para o gerenciamento de dados de sensores em aplicações que requerem monitoramento ou coleta de dados ambientais.

#### Objetivo

O objetivo desses testes é assegurar a integridade das operações de criação e leitura de entidades de sensores no banco de dados, verificando que:

- Sensores podem ser criados com os atributos corretos e armazenados no MongoDB (`TestCreateSensor`).
- Todos os sensores armazenados podem ser recuperados corretamente, validando a funcionalidade de listagem de sensores (`TestGetAllSensors`).

#### Fluxo dos Testes

**TestCreateSensor:**

1. **Configuração Inicial:** Conexão com o MongoDB e seleção da coleção de teste.
2. **Criação de Sensor:** Utiliza o `SensorDataAdapter` para criar um novo sensor no banco de dados.
3. **Verificações:**
   - Confirma que nenhum erro ocorreu durante a criação do sensor.
   - Assegura que o sensor criado possui um identificador MongoDB válido (`Id`), diferente de `NilObjectID`.

**TestGetAllSensors:**

1. **Configuração Inicial:** Mesmo processo de conexão com o MongoDB e preparação da coleção de teste.
2. **Preparação de Dados de Teste:** Inserção direta de um sensor na coleção de teste para simular dados pré-existentes.
3. **Recuperação de Todos os Sensores:** Utilização do `SensorDataAdapter` para recuperar todos os sensores armazenados.
4. **Verificações:**
   - Confirma que a lista de sensores recuperada não é nula e contém dados.
   - Verifica que os atributos dos sensores recuperados, como o nome, correspondem aos dados inseridos para teste.

#### Conclusão

A execução bem-sucedida desses testes confirma a funcionalidade adequada do adaptador de dados do sensor para operações de criação e leitura no MongoDB. Isso inclui a capacidade de persistir sensores com atributos corretos e a habilidade de recuperar todos os sensores armazenados de forma confiável. Essas operações são essenciais para aplicações que dependem de dados de sensores para monitoramento ambiental, análise de dados, ou outras funcionalidades relacionadas, garantindo assim a confiabilidade e a precisão dos dados manipulados pelo sistema.

## Teste SensorData

Este documento detalha os testes implementados para validar as funcionalidades de criação e recuperação de dados de sensores no contexto do pacote `sensor`. Utilizando o MongoDB como banco de dados, estes testes são essenciais para assegurar que as operações de persistência e consulta de dados de sensores funcionem como esperado, um componente chave para aplicações que dependem de monitoramento ou coleta de dados ambientais.

#### Objetivo

O objetivo dos testes é verificar a correta implementação das funcionalidades de:

- **Criação de Sensor (`TestCreateSensor`):** Garantir que um sensor possa ser criado e armazenado corretamente no MongoDB, com a geração de um identificador único.
- **Recuperação de Todos os Sensores (`TestGetAllSensors`):** Assegurar que a funcionalidade de recuperação de todos os sensores cadastrados no banco de dados funcione corretamente, retornando os dados esperados.

#### Fluxo dos Testes

**Teste de Criação de Sensor (`TestCreateSensor`):**

1. **Configuração Inicial:** Estabelece conexão com o MongoDB e seleciona a coleção de testes.
2. **Execução do Teste:** Cria um sensor usando a função `CreateSensor` e verifica se o sensor foi criado com um identificador único no banco de dados.
3. **Verificações:** 
   - Confirma que não ocorreram erros durante a criação.
   - Verifica se o sensor retornado não é nulo e possui um identificador único.

**Teste de Recuperação de Todos os Sensores (`TestGetAllSensors`):**

1. **Configuração Inicial:** Similar ao teste anterior, inclui a inserção direta de dados de teste na coleção de sensores para simular um ambiente pré-populado.
2. **Execução do Teste:** Recupera todos os sensores cadastrados utilizando a função `GetAllSensors`.
3. **Verificações:** 
   - Confirma que a lista de sensores recuperada não é nula e contém os dados esperados.
   - Verifica se os dados dos sensores correspondem aos dados de teste inseridos anteriormente.

#### Conclusão

A execução bem-sucedida destes testes assegura a funcionalidade correta das operações de criação e recuperação de sensores no sistema, validando a integração com o MongoDB e a lógica de negócio associada. Isso é fundamental para garantir a confiabilidade e precisão dos dados de sensores manipulados pela aplicação, suportando assim funcionalidades críticas de monitoramento e análise de dados. Estes testes contribuem significativamente para a manutenção da qualidade e estabilidade do sistema, facilitando a detecção precoce de regressões ou falhas.

# Testes do Primary Adapter

## Teste Queue

Este documento detalha o teste implementado para validar a integração e o fluxo de processamento de mensagens entre o RabbitMQ e o MongoDB, através do adaptador `MessageHandlerAdapter`. Este teste é crucial para sistemas que dependem de mensageria para coleta e armazenamento de dados, assegurando que as mensagens enviadas através do RabbitMQ sejam corretamente recebidas e persistidas no MongoDB.

#### Objetivo

O principal objetivo deste teste é verificar que uma mensagem enviada para uma fila do RabbitMQ é corretamente processada pelo adaptador e resulta na inserção dos dados da mensagem em uma coleção do MongoDB. Isso valida tanto a conexão e comunicação com o RabbitMQ quanto a lógica de armazenamento no MongoDB.

#### Fluxo do Teste

1. **Configuração Inicial:**
   - Estabelece uma conexão com o MongoDB e assegura que a coleção de destino para testes está limpa.
   - Configura a conexão com o RabbitMQ, incluindo a declaração da fila de teste.

2. **Publicação de Mensagem no RabbitMQ:**
   - Publica uma mensagem de teste na fila do RabbitMQ configurada anteriormente.

3. **Processamento da Mensagem:**
   - Aguarda um intervalo de tempo para garantir que a mensagem seja consumida e processada pelo sistema (neste passo, assume-se que há uma lógica fora deste teste que consome a mensagem da fila e insere os dados no MongoDB).

4. **Verificação no MongoDB:**
   - Realiza uma consulta ao MongoDB para confirmar se a mensagem foi corretamente armazenada na coleção especificada.

5. **Assertivas:**
   - Verifica se não houve erros durante a consulta ao MongoDB e se o resultado não está vazio, indicando que a mensagem foi processada e armazenada corretamente.

#### Conclusão

A execução bem-sucedida do `TestMessageHandlerAdapter` confirma a eficácia do sistema em processar mensagens de um broker RabbitMQ e persisti-las no MongoDB. Este teste é fundamental para validar a integridade do fluxo de dados entre componentes de mensageria e armazenamento, essencial para aplicações que dependem de processamento assíncrono de dados. Garantir a confiabilidade desta integração é crucial para a estabilidade e eficácia da aplicação, suportando operações críticas de coleta e análise de dados.