# Teste de Carga do Banco de Dados

Este documento detalha o teste de carga implementado para verificar a capacidade do sistema de inserir múltiplos documentos em uma coleção do MongoDB. Este teste é crucial para sistemas que necessitam de alta performance e confiabilidade no armazenamento de dados, assegurando que o banco de dados pode lidar com um volume significativo de inserções sem degradação de performance.

### Objetivo

O principal objetivo deste teste é avaliar a eficiência e a robustez do banco de dados MongoDB ao processar múltiplas inserções em uma coleção específica. Isso inclui validar a conexão com o banco de dados, a correta inserção dos documentos, e a capacidade do sistema de gerenciar cargas de trabalho intensas.

#### Relacionado aos RFs e RNFs:

- **RF1 - Captura e Armazenamento de Dados Ambientais**: Este teste verifica a capacidade do MongoDB de gerenciar com eficácia a inserção de grandes volumes de dados ambientais, essencial para o armazenamento confiável e acessível dessas informações.
- **RNF2 - Desempenho (Aguentar no mínimo 1000 requisições por segundo)**: Avaliando a eficiência nas operações de inserção no banco de dados, este teste contribui para a garantia de que o sistema pode manter um desempenho elevado, mesmo sob cargas intensas de inserção de dados.
- **RNF4 - Desempenho (Processamento e geração de arquivos CSV para relatórios de operações deve ser concluído em menos de 30 segundos)**: Indiretamente, ao assegurar a rápida inserção de dados no banco, este teste suporta a eficiência na geração subsequente de relatórios e análises, facilitando a conversão de dados brutos em informações úteis em tempo hábil.
- **RNF5 - Disponibilidade (99,5% de tempo de operação)**: Confirmar que o banco de dados pode lidar com inserções pesadas sem interrupções assegura a alta disponibilidade dos serviços, fundamental para a confiabilidade contínua do sistema.
- **RNF6 - Escalabilidade**: Este teste confirma que o banco de dados e, por extensão, o sistema como um todo, são capazes de escalar para suportar aumentos no volume de dados gerados ou coletados, mantendo a performance e a eficiência.

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