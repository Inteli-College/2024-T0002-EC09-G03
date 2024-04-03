# Teste de Carga do Banco de Dados

Este documento detalha o teste de carga implementado para verificar a capacidade do sistema de inserir múltiplos documentos em uma coleção do MongoDB. Este teste é crucial para sistemas que necessitam de alta performance e confiabilidade no armazenamento de dados, assegurando que o banco de dados pode lidar com um volume significativo de inserções sem degradação de performance.

### Objetivo

O principal objetivo deste teste é avaliar a eficiência e a robustez do banco de dados MongoDB ao processar múltiplas inserções em uma coleção específica. Isso inclui validar a conexão com o banco de dados, a correta inserção dos documentos, e a capacidade do sistema de gerenciar cargas de trabalho intensas.

### Fluxo do Teste

1. **Configuração Inicial:**
   - Carregamento das variáveis de ambiente necessárias para a conexão com o MongoDB.
   - Estabelecimento da conexão com o banco de dados utilizando a função `NewDBConnection`.

2. **Limpeza da Coleção:**
   - Antes da inserção, garantir que a coleção `testCollection` esteja limpa ou criar uma nova coleção para o teste.

3. **Inserção de Documentos:**
   - Inserir, de forma iterativa, um total de 10 documentos na coleção `testCollection`. Cada documento contém uma chave única e um valor associado, garantindo a diversidade dos dados inseridos.

4. **Verificações:**
   - Após a inserção de todos os documentos, verificar se eles foram corretamente armazenados na coleção.
   - Avaliar a performance da inserção, considerando o tempo total necessário para a operação.

5. **Encerramento e Limpeza:**
   - Desconectar do cliente MongoDB ao final do teste.
   - (Opcional) Remover a coleção `testCollection` ou os documentos inseridos para manter o ambiente de teste limpo.

### Conclusão

A execução bem-sucedida do teste de carga no MongoDB indica que o sistema está preparado para gerenciar eficientemente cargas de trabalho elevadas, mantendo a integridade e a performance do banco de dados. Este teste é fundamental para garantir a confiabilidade do sistema em cenários de uso real, onde a eficiência no processamento de grandes volumes de dados é crucial.

# Teste de Carga da Fila RabbitMQ

Este documento detalha o teste implementado para validar a integração e o desempenho na publicação de mensagens para uma fila do RabbitMQ. Este teste é crucial para sistemas que dependem de mensageria para a comunicação entre diferentes serviços, assegurando que as mensagens enviadas através do RabbitMQ sejam corretamente publicadas na fila especificada sob uma carga de trabalho definida.

### Objetivo

O principal objetivo deste teste é verificar que uma série de mensagens enviadas para uma fila do RabbitMQ sejam corretamente publicadas sem erros, validando a capacidade do sistema de lidar com a carga de publicação especificada e a resiliência da conexão e do canal com o RabbitMQ.

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
