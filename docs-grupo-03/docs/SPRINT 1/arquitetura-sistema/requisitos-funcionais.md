# Requisitos Funcionais

Requisitos funcionais são declarações detalhadas que descrevem as funcionalidades, características e interações específicas que um sistema deve incorporar, delineando de forma precisa o que o sistema deve realizar.

| Requisito | Descrição |
|---|---|
| RF1 | O sistema deve ter, de maneira prévia, o espaço do almoxarifado mapeado ou deve-se realizar o mapeamento manualmente, para garantir que ele possua o caminho da melhor rota para os diferentes pontos do almoxarifado. |
| RF2 | O sistema deve ser capaz de interagir de forma visual e sonora com o usuário, ou seja, através da utilização de servidores STT (Speech-to-text), TTS (Text-to-Speech) e LLM (Large Language Model), o sistema identificará a entrada, realizará a conversão de formato (a depender da entrada) e retornará a resposta.|
| RF3 | O sistema deve ser capaz de interagir com o ambiente proposto, desviando de possíveis obstáculos, através do sensor LIDAR, e agindo como um guia de localização para o usuário. |
| RF4 | O sistema deve identificar qual peça deve ser procurada através da linguagem frequentemente utilizada pelos principais usuários, sendo capaz de pesquisar pelo nome da peça e por numerações específicas de tamanho ou fatores importantes de classificação. |
| RF5 | O sistema deve gerar um arquivo CSV com os logs diários dos serviços prestados pelo próprio. |
| RF6 | Toda e qualquer ordem de serviço deve ser registrada e salva em um arquivo do formato JSON, que ficará na ponta da API do sistema, a fim de facilitar a integração com o ERP da empresa parceira. |