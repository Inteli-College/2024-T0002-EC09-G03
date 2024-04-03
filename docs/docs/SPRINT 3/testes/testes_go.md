# Testes de Infraestrutura

## Teste DB

Este documento descreve o script de teste contido no arquivo `db_test.go`em `infra`, destinado a validar a integridade e funcionalidade das conexões de banco de dados dentro de um ambiente de desenvolvimento ou CI/CD. Os testes são essenciais para garantir que a aplicação possa interagir corretamente com o banco de dados MongoDB, realizando operações de leitura e escrita conforme esperado, sem interrupções ou falhas devido a problemas de conexão ou configuração.

### Objetivo

O principal objetivo deste script de teste é assegurar a confiabilidade e robustez da função `NewDBConnection`, que é responsável por estabelecer uma conexão entre a aplicação e o banco de dados MongoDB. Isso inclui a validação da criação do objeto de conexão do banco de dados e do cliente do banco de dados, assegurando que nenhum deles seja nulo, o que indicaria falha na tentativa de conexão.

#### Relacionado aos RFs e RNFs:

- **RF1 - Captura e Armazenamento de Dados Ambientais**: Este teste assegura que a infraestrutura de banco de dados está pronta para capturar e armazenar uma ampla variedade de dados ambientais de forma segura e redundante.
- **RNF4 - Desempenho (Processamento e geração de arquivos CSV)**: A validação da conexão com o MongoDB e a funcionalidade de leitura/escrita garantem que o processamento e geração de relatórios operacionais sejam eficientes.
- **RNF5 - Disponibilidade**: Assegura que a base de dados estará disponível conforme o exigido, apoiando o uptime de 99,5%.

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

#### Relacionado aos RFs e RNFs:

- **RF6 - Integração API Externa**: Testes de conexão MQTT validam a robustez da integração com APIs externas, crucial para atualizações de dados ambientais.
- **RNF11 - Adaptação Dinâmica de Carga**: A capacidade de estabelecer conexões MQTT robustas sob demanda apoia diretamente a adaptação dinâmica de carga do sistema.

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

#### Relacionado aos RFs e RNFs:

- **RF3 - Notificações e Alertas Configuráveis**: Testando a criação e manipulação de consumidores de fila, este teste suporta a funcionalidade de notificações e alertas do sistema, garantindo que as mensagens cheguem aos usuários em tempo real.
- **RNF7 - Escalabilidade**: A verificação da funcionalidade do adaptador de fila confirma a capacidade do sistema de expandir e suportar um aumento significativo de carga, alinhando-se à exigência de escalabilidade.

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