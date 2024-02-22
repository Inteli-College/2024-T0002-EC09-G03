# Diagrama UML de sequência e implantação

Nessa seção são apresentadas as interações detalhadas entre os diferentes componentes do sistema proposto, incluindo os dispositivos de hardware multimodais, a plataforma web de engajamento do cidadão e o dashboard de gerenciamento de dados. Esse diagrama oferece uma visão clara da dinâmica de comunicação entre esses elementos, bem como sua distribuição física nos ambientes urbanos. Ao representar as sequências de eventos entre os sensores, a plataforma web e o sistema de gerenciamento, o diagrama fornece uma base sólida para o planejamento e implementação das soluções propostas, contribuindo significativamente para a compreensão e o desenvolvimento eficaz do projeto.

## Diagrama UML de sequência

![Diagrama UML de sequencia](../../../static/img/uml-sequencia.svg)

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

![Diagrama UML de sequencia](../../../static/img/uml-implantacao.svg)

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
