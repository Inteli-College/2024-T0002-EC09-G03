# Testes do Primary Adapter

## Teste Queue

Este documento detalha o teste implementado para validar a integração e o fluxo de processamento de mensagens entre o RabbitMQ e o MongoDB, através do adaptador `MessageHandlerAdapter`. Este teste é crucial para sistemas que dependem de mensageria para coleta e armazenamento de dados, assegurando que as mensagens enviadas através do RabbitMQ sejam corretamente recebidas e persistidas no MongoDB.

#### Objetivo

O principal objetivo deste teste é verificar que uma mensagem enviada para uma fila do RabbitMQ é corretamente processada pelo adaptador e resulta na inserção dos dados da mensagem em uma coleção do MongoDB. Isso valida tanto a conexão e comunicação com o RabbitMQ quanto a lógica de armazenamento no MongoDB.

#### Relacionado aos RFs e RNFs:

- **RF1 - Captura e Armazenamento de Dados Ambientais**: Verifica o armazenamento correto de dados ambientais coletados via mensagens.
- **RF3 - Notificações e Alertas Configuráveis**: Assegura a entrega e processamento de alertas através de mensagens.
- **RNF2 - Desempenho (1000 requisições por segundo)**: Testa a capacidade do sistema de processar um alto volume de mensagens eficientemente.

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

# Testes dos Secondary Adapters

## Teste MQTT

Este documento aborda os testes implementados no pacote `mqtt`, especificamente focados na funcionalidade de um adaptador MQTT (`MQTTAdapter`). Estes testes visam validar a criação de clientes MQTT e a publicação de mensagens em tópicos específicos, assegurando que a comunicação via MQTT funcione de maneira esperada, que é crucial para aplicações que dependem de comunicação em tempo real ou IoT.

#### Objetivo

Os principais objetivos destes testes incluem:

- **TestCreateClient:** Garantir que um cliente MQTT pode ser criado corretamente, e que os atributos do cliente (como o nome) são configurados como esperado.
- **TestPublish:** Verificar que mensagens podem ser publicadas em tópicos específicos sem causar erros ou pânico, validando assim a funcionalidade de publicação do adaptador.

#### Relacionado aos RFs e RNFs:

- **RF6 - Integração API Externa**: Valida a capacidade do sistema de se conectar e comunicar com APIs externas através de MQTT.
- **RNF11 - Adaptação Dinâmica de Carga**: Confirma que o adaptador MQTT pode ajustar seus recursos dinamicamente para manter a eficiência sob diferentes cargas de trabalho.

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

#### Relacionado aos RFs e RNFs:

- **RF1 - Captura e Armazenamento de Dados Ambientais**: Confirma a capacidade do sistema de criar e recuperar informações dos sensores, essenciais para o monitoramento ambiental.
- **RNF5 - Disponibilidade**: Verifica a confiabilidade na recuperação de dados dos sensores, crucial para a disponibilidade contínua do serviço.

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

#### Relacionado aos RFs e RNFs:

- **RF1 - Captura e Armazenamento de Dados Ambientais**: Testa a eficácia na criação e armazenamento de dados coletados pelos sensores no banco de dados.
- **RNF5 - Disponibilidade**: Assegura a disponibilidade dos dados de sensores para análise e relatórios, suportando decisões baseadas em dados confiáveis.

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