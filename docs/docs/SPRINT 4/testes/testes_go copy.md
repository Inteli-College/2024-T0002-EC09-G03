# Teste de Carga da Fila RabbitMQ

Este documento detalha o teste implementado para validar a integração e o desempenho na publicação de mensagens para uma fila do RabbitMQ. Este teste é crucial para sistemas que dependem de mensageria para a comunicação entre diferentes serviços, assegurando que as mensagens enviadas através do RabbitMQ sejam corretamente publicadas na fila especificada sob uma carga de trabalho definida.

### Objetivo

O principal objetivo deste teste é verificar que uma série de mensagens enviadas para uma fila do RabbitMQ sejam corretamente publicadas sem erros, validando a capacidade do sistema de lidar com a carga de publicação especificada e a resiliência da conexão e do canal com o RabbitMQ.

#### Relacionado aos RFs e RNFs:

- **RF3 - Notificações e Alertas Configuráveis**: Validando a publicação eficiente de mensagens no RabbitMQ, este teste assegura que o sistema possa transmitir alertas e notificações sob demanda, mesmo sob alta carga.
- **RNF2 - Desempenho (Aguentar no mínimo 1000 requisições por segundo)**: Ao medir o tempo total para a publicação de mensagens, o teste verifica a capacidade do sistema de manter um alto desempenho sob cargas intensas de trabalho, garantindo que a infraestrutura de mensageria possa suportar uma grande quantidade de requisições simultâneas.
- **RNF6 - Escalabilidade**: Este teste confirma que a arquitetura do sistema é capaz de escalar efetivamente, ajustando-se para lidar com aumentos significativos no volume de mensagens, crucial para sustentar o crescimento do uso do sistema sem degradação do desempenho.
- **RNF11 - Adaptação Dinâmica de Carga**: Avaliando a capacidade do sistema de publicar mensagens sob diferentes cargas, o teste verifica a habilidade do sistema em ajustar dinamicamente seus recursos computacionais, mantendo um desempenho consistente independente das variações na demanda de processamento.

### Fluxo do Teste

1. **Configuração Inicial:**
   - Carrega as variáveis de ambiente necessárias, incluindo a URL de conexão com o RabbitMQ, utilizando a função `LoadEnvVariables` do pacote de inicialização.
   - Estabelece a conexão com o RabbitMQ e abre um canal de comunicação.

2. **Preparação do Teste:**
   - Define o nome da fila (`testQueue`) e o número de mensagens a serem publicadas (`messageCount`).

3. **Execução da Publicação de Mensagens:**
   - Inicia um loop para publicar o número definido de mensagens na fila do RabbitMQ, construindo o corpo da mensagem com um identificador único para cada iteração.
   - Utiliza um contexto com timeout para garantir que a operação de publicação não exceda um limite de tempo pré-estabelecido.

4. **Verificação e Registro do Desempenho:**
   - Registra o tempo total gasto para a publicação de todas as mensagens, permitindo avaliar a eficiência da publicação sob a carga de trabalho definida.

5. **Encerramento:**
   - Fecha o canal e a conexão com o RabbitMQ ao final do teste para liberar recursos.

### Conclusão

A execução bem-sucedida do teste `TestQueueLoad` confirma a eficácia do sistema em lidar com a publicação de uma sequência de mensagens para o RabbitMQ sob a carga de trabalho especificada. Este teste é fundamental para validar a confiabilidade e a eficiência da integração com serviços de mensageria, essenciais para a comunicação entre diferentes componentes do sistema. Garantir a estabilidade e o desempenho dessa integração é crucial para a operação contínua e eficiente de aplicações que dependem de processamento assíncrono e comunicação entre serviços.

