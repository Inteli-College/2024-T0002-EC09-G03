import csv
import time
import json
from datetime import datetime
from paho.mqtt import client as mqtt_client

broker = 'broker.hivemq.com'
port = 8000
client_id = "clientId-Qro3K1kRhV"
csv_file_path = "Ponderada 1\data\sensors.csv"

def read_sensor_data_from_csv(csv_path):
    while True:  # Loop infinito para simular dados contínuos
        with open(csv_path, mode='r') as csvfile:
            csvreader = csv.DictReader(csvfile)
            for row in csvreader:
                yield row
        # Opcional: Pode-se adicionar uma pausa ou condição de saída aqui

# Criar um gerador
sensor_data_generator = read_sensor_data_from_csv(csv_file_path)

def generate_data(sensor_name):
    row = next(sensor_data_generator)  # Pega a próxima linha de dados do gerador
    return row[sensor_name]

# Configurações dos sensores utilizando o gerador para dados
sensores = {
    "radiacao_solar": {
        "topic": "meuTesteIoT/sensor/radiacao_solar",
        "generate_data": lambda: generate_data("radiacao_solar"),
        "unidade": "W/m²",
        "nome": "Sensor de Radiação Solar"
    },
    "temperatura_ambiente": {
        "topic": "meuTesteIoT/sensor/temperatura",
        "generate_data": lambda: generate_data("temperatura"),
        "unidade": "°C",
        "nome": "Sensor de Temperatura Ambiente"
    },
    # Adicione mais sensores conforme necessário
}

def connect_mqtt():
    def on_connect(client, userdata, flags, rc):
        if rc == 0:
            print("Connected to MQTT Broker!")
        else:
            print("Failed to connect, return code %d\n", rc)

    client = mqtt_client.Client(client_id)
    client.on_connect = on_connect
    client.connect(broker, port)
    return client

def publish(client):
    while True:
        for sensor_name, sensor_info in sensores.items():
            valor = sensor_info["generate_data"]()
            mensagem = json.dumps({
                "sensor": sensor_info["nome"],
                "valor": valor,
                "unidade": sensor_info["unidade"],
                "timestamp": datetime.now().isoformat()
            })
            print(f"Sent `{mensagem}` to topic `{sensor_info['topic']}`")
            result = client.publish(sensor_info["topic"], mensagem)
            # Tratar o resultado da publicação como antes
            time.sleep(1)  # Ajuste o tempo de espera conforme necessário

def run():
    client = connect_mqtt()
    client.loop_start()
    publish(client)

if __name__ == "__main__":
   run()
