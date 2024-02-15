# Simulador MQTT


## Sumário

1. [Descrição Geral](#descrição-geral)
2. [Requisitos](#requisitos)
3. [Configuração](#configuração)
4. [Uso](#uso)
5. [Arquivos](#arquivos)


## Descrição Geral

Este projeto é uma prova de conceito que simula a leitura de dados de sensores e o envio desses dados para um broker MQTT. O código lê dados contínuos de um arquivo CSV que simula a saída de sensores físicos, como sensores de radiação solar e temperatura ambiente, e publica esses dados em tópicos MQTT específicos.

## Requisitos

- Python 3.x
- Bibliotecas Python:
  - `paho-mqtt`: Para interação com o broker MQTT.
  - `pytest`: Para executar os testes (opcional, apenas se você deseja rodar os testes).
- Um broker MQTT, como `broker.hivemq.com`, ou uma instalação local do Mosquitto.

## Configuração

1. Instale as dependências necessárias usando pip:

   ```sh
   pip install paho-mqtt pytest
   ```

2. Certifique-se de que o broker MQTT esteja acessível. Se estiver usando um broker local, inicie o serviço do broker.

3. Prepare um arquivo CSV `sensors.csv` com dados simulados para os sensores. O arquivo deve ter um cabeçalho e pelo menos algumas linhas de dados, por exemplo:

   ```csv
   radiacao_solar,temperatura
   300,25
   320,24
   ```

## Uso

Para executar a simulação, use o seguinte comando:

```sh
python sensor_reading.py
```

Isso iniciará o script que se conecta ao broker MQTT especificado, lê os dados do sensor do arquivo CSV e publica esses dados nos tópicos MQTT configurados.

## Arquivos

- `sensor_reading.py`: O script principal que executa a simulação dos sensores e a publicação MQTT.
- `sensors.csv`: Arquivo CSV contendo os dados simulados dos sensores.
- `tests/test_sensors_pytest.py`: Arquivo contendo os testes para validar as funções do script principal.
Para uma documentação mais aprofundada dos testes realizados no projeto de simulação de sensores com MQTT, é importante detalhar não apenas como executar os testes, mas também explicar o propósito de cada teste, o que cada um verifica, e como eles contribuem para a validade e a confiabilidade do sistema como um todo. Abaixo, você encontrará um modelo aprimorado de documentação com foco nos testes.
