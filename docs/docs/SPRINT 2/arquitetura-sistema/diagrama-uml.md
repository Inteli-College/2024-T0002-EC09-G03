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
## Diagrama UML de implantação

```
@startuml
package "Sistema IoT" {
    [Painel]
    [Aplicação IoT]
    [Broker MQTT]
    [Sensor 1]
    [Sensor 2]
    [Sensor N]
}

Painel --> AplicacaoIoT: Solicita Dados
AplicacaoIoT --> Broker: Conecta
AplicacaoIoT --> Sensor1: Solicita Dados
AplicacaoIoT --> Sensor2: Solicita Dados
AplicacaoIoT --> SensorN: Solicita Dados
Broker --> Sensor1: Solicita Dados
Broker --> Sensor2: Solicita Dados
Broker --> SensorN: Solicita Dados
@enduml

```