# Requisitos Funcionais, não Funcionais e Restrições
**Introdução:**

O projeto em parceria com a PRODAM representa um marco inovador no monitoramento ambiental urbano, com o objetivo de utilizar tecnologia sensorial avançada para coletar, analisar e disponibilizar dados sobre diversos aspectos ambientais. Este projeto, liderado pela Empresa de Tecnologia da Informação e Comunicação do Município de São Paulo - Prodam-SP S.A., visa aprimorar a qualidade de vida nas cidades, promovendo um ambiente mais sustentável e saudável. EcoVigilância busca integrar cidadãos, gestores públicos e tecnologias emergentes em um esforço coletivo para enfrentar desafios ambientais, como a poluição do ar, ruído urbano, e gestão de resíduos, através de uma abordagem holística e inovadora.

**Requisitos Funcionais:**

1. **Captura e Armazenamento de Dados Ambientais**: O sistema coletará dados em intervalos de 5 minutos, armazenando-os de forma segura e redundante. A variedade de dados coletados inclui indicadores de qualidade da água, ar, solo e biodiversidade urbana.
2. **Dashboard Dinâmico**: Atualização de dados a cada 10 minutos, com capacidade para visualizações históricas e análise comparativa.
3. **Notificações e Alertas Configuráveis**: Sistema de alerta com notificação de usuários em menos de 1 minuto após a detecção de condições anormais.
4. **Feedback e Participação Cidadã**: Portal para suportar até 10.000 usuários simultâneos, promovendo a interação e o engajamento.
5. **Relatórios e Exportação de Dados**: Facilidades para a geração de relatórios personalizados e exportação em diversos formatos.
6. **Integração API Externa**: Atualizações de dados de APIs externas a cada 30 minutos, enriquecendo as análises.
7. **Análise Avançada de Dados**: Facilitar por meio dos dados e organização dos mesmos a implementação futura de algoritmos de Machine Learning para otimizar a análise dos dados ambientais, permitindo previsões mais precisas sobre eventos ambientais adversos.
8. **Interação Melhorada no Portal de Participação Cidadã**: Desenvolver funcionalidades interativas, como fóruns e enquetes, para aumentar o engajamento e a participação dos cidadãos nas decisões ambientais.
9. **Suporte a Dispositivos Móveis**: Garantir que o dashboard e o portal de participação cidadã sejam totalmente responsivos e otimizados para dispositivos móveis, permitindo acesso fácil e em qualquer lugar.

**Requisitos Não Funcionais:**

1. **Desempenho**: Garantia de resposta em menos de 2 segundos para 95% das requisições.
2. **Desempenho**: Aguentar no mínimo 1000 requisições por segundos. 
3. **Desempenho**: Comunicar com ao menos 10 clientes ao mesmo tempo. 
4. **Desempenho**: Processamento e geração de arquivos CSV para relatórios de operações deve ser concluído em menos de 30 segundos.
4. **Disponibilidade**: 99,5% de tempo de operação, com manutenções programadas fora de picos.
5. **Segurança**: Adesão a padrões de segurança e conformidade com legislações de proteção de dados (LGPD e GDPR). 
6. **Escalabilidade**: Capacidade de expansão para suportar um aumento de carga de até 200%.
7. **Usabilidade**: Testes de usabilidade para assegurar uma experiência positiva ao usuário.
8. **Monitoramento e Logs**: Sistema de monitoramento contínuo e armazenamento de logs por 6 meses (30 GB aproximadamente).
9. **Compatibilidade Cross-Platform**: Assegurar que o sistema seja compatível com diferentes sistemas operacionais e navegadores, maximizando a acessibilidade.
10. **Sustentabilidade do Sistema**: Implementar práticas de desenvolvimento sustentável, incluindo a otimização do uso de recursos e energia pelo sistema.
11. **Adaptação Dinâmica de Carga**: O sistema deve ser capaz de ajustar dinamicamente seus recursos computacionais em resposta a variações na demanda, garantindo desempenho consistente.


**Restrições:**

1. **Conformidade Legal**: Adesão estrita a regulamentações de proteção de dados.
2. **Prazo de Entrega**: Meta de operacionalidade de teste em 12 meses.
3. **Orçamento**: Limitação orçamentária de R$ 5 milhões.
4. **Tecnologia e Infraestrutura**: Utilização da infraestrutura de nuvem AWS.
5. **Capacitação**: Programa de treinamento para equipe interna e usuários externos.
6. **Impacto Ambiental da Infraestrutura Tecnológica**: Minimizar o impacto ambiental associado à infraestrutura tecnológica utilizada pelo projeto, incluindo servidores e dispositivos de rede.
7. **Acessibilidade**: Assegurar que o sistema seja acessível para pessoas com deficiência, seguindo as diretrizes de acessibilidade da web (WCAG).
8. **Extensibilidade**: O design do sistema deve permitir fácil expansão e inclusão de novos tipos de sensores ambientais no futuro.

**Conclusão:**

Em suma, tratamos de um projeto ambicioso que coloca a tecnologia a serviço do meio ambiente e da sociedade, promovendo uma gestão urbana mais inteligente e sustentável. Através da implementação dos requisitos funcionais e não funcionais detalhados, além de aderir às restrições estabelecidas, este projeto tem o potencial de transformar a maneira como as cidades monitoram e respondem a questões ambientais. Com foco na inovação, segurança, e engajamento cidadão, EcoVigilância está destinado a se tornar um modelo para futuras iniciativas de cidades inteligentes e sustentáveis, reforçando o compromisso da Prodam-SP S.A. com a melhoria contínua da qualidade de vida urbana.