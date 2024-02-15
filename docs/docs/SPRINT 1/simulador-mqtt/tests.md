
# Testes

## Visão Geral dos Testes

Os testes deste projeto foram projetados para garantir a confiabilidade e a precisão da simulação dos dados dos sensores e da publicação dos dados via MQTT. Eles verificam a integridade da leitura de dados do arquivo CSV, a correta geração de dados com base nos valores do sensor, e a capacidade do sistema de se conectar e publicar para um broker MQTT.

## Ambiente de Teste

Antes de executar os testes, certifique-se de que o ambiente de teste está devidamente configurado com todas as dependências necessárias instaladas. Isso inclui a biblioteca `pytest` para execução dos testes e a `paho-mqtt` para a simulação de publicações MQTT.

## Execução dos Testes

Para executar os testes, utilize o comando:

```sh
pytest test_sensors_pytest.py
```

Este comando irá procurar por todas as funções no arquivo `test_sensors_pytest.py` que começam com `test_` e as executará.

## Testes Implementados

### Teste de Leitura de Dados do CSV

**Nome do Teste**: `test_read_sensor_data_from_csv`

**Objetivo**: Verificar se a função `read_sensor_data_from_csv` consegue ler dados de um arquivo CSV e gerar um dicionário para cada linha, contendo as chaves e valores corretos correspondentes aos sensores.

**Metodologia**: Este teste invoca a função `read_sensor_data_from_csv` passando o caminho para um arquivo CSV de teste. Então, ele verifica se o primeiro dicionário gerado possui as chaves esperadas (`radiacao_solar`, `temperatura`), assegurando que os dados estão sendo corretamente interpretados.

### Teste de Geração de Dados

**Nome do Teste**: `test_generate_data`

**Objetivo**: Assegurar que a função `generate_data` consiga extrair e retornar o valor correto para um determinado nome de sensor a partir dos dados fornecidos pelo gerador de dados do sensor.

**Metodologia**: O teste invoca a função `generate_data` com um nome de sensor específico (por exemplo, `radiacao_solar`). Verifica-se então se o valor retornado não é `None`, indicando que a função conseguiu acessar e extrair o valor do sensor a partir dos dados do gerador.

## Considerações Importantes

- **Validade dos Dados**: Os testes assumem que os dados fornecidos no arquivo CSV de teste são válidos e representativos dos dados reais que seriam gerados por sensores físicos.

- **Conexão MQTT**: Devido à natureza do projeto, que envolve a publicação de dados para um broker MQTT, seria ideal implementar testes adicionais que simulam essa interação. Tais testes podem incluir a verificação de se a conexão com o broker é estabelecida com sucesso e se as mensagens são publicadas corretamente nos tópicos esperados. No entanto, isso pode requerer a configuração de um broker MQTT de teste ou o uso de mocks para simular as funcionalidades do broker.

- **Cobertura de Testes**: Para garantir a robustez do sistema, é recomendável expandir a cobertura de testes para incluir cenários de erro, como arquivos CSV malformados, falhas de conexão MQTT e dados de sensores inválidos.

## Conclusão

Os testes são fundamentais para assegurar a confiabilidade e a precisão do sistema de simulação de sensores e publicação MQTT. Uma documentação detalhada dos testes não só facilita a manutenção e a expansão do sistema, mas também fornece uma base sólida para a validação do comportamento esperado do sistema sob várias condições.