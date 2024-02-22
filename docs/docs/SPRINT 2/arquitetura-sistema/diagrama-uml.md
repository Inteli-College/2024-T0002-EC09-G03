# Diagrama UML de sequência e implantação


## Diagrama UML de sequência

```
@startuml
actor Usuário
participant "Painel" as Painel
participant "Aplicação IoT" as AplicacaoIoT
participant "Broker MQTT" as Broker
participant "Sensor 1" as Sensor1
participant "Sensor 2" as Sensor2
participant "Sensor N" as SensorN

Usuário -> Painel: Acessa o Painel
activate Painel
Painel -> AplicacaoIoT: Solicita Dados
activate AplicacaoIoT
AplicacaoIoT -> AplicacaoIoT: Autentica Usuário
AplicacaoIoT --> Painel: Autenticação Bem-Sucedida
AplicacaoIoT -> AplicacaoIoT: Autoriza Usuário
AplicacaoIoT --> Painel: Autorização Concedida
AplicacaoIoT -> Broker: Conecta
activate Broker
Broker -> Sensor1: Solicita Dados
activate Sensor1
Sensor1 --> Broker: Envia Dados
deactivate Sensor1
Broker -> Sensor2: Solicita Dados
activate Sensor2
Sensor2 --> Broker: Envia Dados
deactivate Sensor2
Broker -> SensorN: Solicita Dados
activate SensorN
SensorN --> Broker: Envia Dados
deactivate SensorN
Broker --> AplicacaoIoT: Dados Recebidos
deactivate Broker
AplicacaoIoT --> Painel: Envia Dados
deactivate AplicacaoIoT
deactivate Painel
@enduml

```