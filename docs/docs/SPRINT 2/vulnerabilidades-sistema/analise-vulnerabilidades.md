# Análise de vulnerabilidades do sistema

A arquitetura proposta para um sistema de dashboard, banco de dados, broker MQTT e sensores IoT representa um ambiente complexo e interconectado, fundamental para a coleta, processamento e visualização de dados em tempo real para a tomada de decisões. Com isso, a robustez e segurança desse sistema se fazem necessárias, especialmente considerando o quão estratégicos são os dados ambientais urbanos envolvidos, bem como as regras de negócio em nome de segurança informadas pelo parceiro de mercado durante o workshop e também presentes no TAPI.

Esta análise visa examinar as principais áreas de segurança e vulnerabilidade presentes na arquitetura proposta, identificando potenciais pontos de falha e oferecendo recomendações para mitigar riscos e fortalecer a segurança do sistema como um todo. Desde a autenticação e autorização dos usuários até a proteção dos dados em trânsito e em repouso, levando em consideração os componentes específicos do presente projeto, como o uso de um arquivo definitions.json para o controle de acesso, a integração com a Metabase e a configuração de um Brojer RabbitMQ.

### 1. Senhas e Autenticação:

**1.1. Senhas Fracas**

Há possibilidade de os usuários escolherem senhas fracas (por falta de educação em segurança digitial) que podem ser facilmente comprometidas.

**1.2. Políticas de Senha**

Não há políticas de senha robustas que exijam senhas fortes, incluindo combinações de letras, números e caracteres especiais.

**1.3 Ataques de Força Bruta**

Não há mecanismos de proteção adequados para ataques de força bruta, onde os invasores podem tentar realizar ataques de força bruta para descobrir as senhas dos usuários. É necessário implementar medidas de segurança, como bloqueio de contas após um número específico de tentativas de login falhas.

**1.4. Autenticação Fraca**

Falta de um método de autenticação de dois fatores (2FA).

### 2. Autorização e Controle de Acesso

**2.1. Uso de um arquivo definitions.json**

O arquivo de definições deve ser adequadamente protegido, pois pode ser alvo de adulteração (injeção), resultando em concessão indevida de privilégios.

**2.2. Metabase**

A escolha pela Metabase faz com que seja necessária a constante checagem de suas atualizações para que esta possa contar sempre com as últimas correções de segurança. A falta de atualizações pode deixar o sistema vulnerável a explorações conhecidas.

### 3. Comunicação e Transferência de Dados

**3.1. Segurança do Protocolo MQTT**

As mensagens MQTT não estão protegidas por criptografia ou método similar, podendo lidas facilmente se interceptadas por terceiros.

**3.1. Integridade dos Dados**

Não há um mecanismo para garantir a integridade dos dados transmitidos (assinatura/checksum), o que pode levar fazer com que ataques de manipulação de dados sutis passem despercebidos.

### 4. Segurança do Broker MQTT e RabbitMQ

**4.1. Configuração do Broker**
Uma configuração inadequada do RabbitMQ pode resultar em vulnerabilidades de segurança, como a exposição de portas não essenciais ou permissões excessivas concedidas a usuários não autorizados. Recomenda-se a configuração disponibilizada pela própria documentação do RabbitMQ, que oferece um guia de boas práticas para a configuração segura do broker.

**4.2. Monitoramento e Auditoria**
É essencial implementar sistemas de monitoramento que possam detectar atividades suspeitas, como tentativas de acesso não autorizado ao RabbitMQ ou comportamento anormal nos dados transmitidos.

### 5. Gestão de Sessões e Controle de Acesso

**5.1. Tempo de Sessão**
Não há gerencimento adequado de tempo de sessão paraos usuários do dashboard, isso pode resultar em sessões inativas que permanecem abertas por tempo indeterminado, aumentando o risco de algum acesso não autorizado.

### 6. Backup e Recuperação de Dados

**6.1. Planos de Backup**
Não estão previstos planos de backup robustos para garantir a disponibilidade e a integridade dos dados em caso de falhas de hardware, corrupção de dados ou ataques de ransomware.

### 7. Auditoria e Registro de Atividades

**7.1. Logs de Auditoria**
Logs de auditoria precisam ser implantados e adequadamente configurados e monitorados, essa medida vai permitir a detecção de atividades maliciosas, bem como a identificação da origem de incidentes de segurança.

### 8. Educação e Conscientização do Usuário:

**8.1. Treinamento em Segurança**
Não há previsão da promoção de programas de treinamento em segurança da informação para educar os usuários sobre práticas seguras de uso do sistema e reconhecimento de ameaças.

### 9. Conformidade com Padrões de Segurança

**9.1. GDPR, HIPAA, etc.**
Ainda está em aberto a avaliaçõa de conformidade do sistema sendo desenvolvido quanto as regulamentações relevantes, como o GDPR (Regulamento Geral de Proteção de Dados) ou LGPD (Lei Geral de Proteção de Dados Pessoais), no caso da plataforma de contribuição pelo cidadão.
