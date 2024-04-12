# Análise financeira

A análise financeira do projeto determina aspectos importantes para a estratégia econômica que será adotada pela empresa. Essa projeção permite mensurar o desempenho do projeto e elaborar melhores tomadas de decisões. Destaca-se que essa análise inclui entre os fatores um planejamento orçamentário e otimização de recursos. 

## Tabela de despesas e custos

A tabela abaixo ilustra as estimativas de despesas anuais e custos associados ao investimento no projeto. As despesas se relacionam aos compromissos financeiros recorrentes que a empresa tem, como as mensalidades dos serviços AWS. Por outro lado, os custos refletem os investimentos únicos ou fixos, como a aquisição de hardware (exemplo: sensores) e contratação de pessoal.

### Custos com Folha de Pagamento

| Categorias              | Quantidades | Descrição                                                                                           | Gastos mensais | Valor CLT por funcionário | Duração (Meses) | Total Acumulado |
|-------------------------|-------------|----------------------------------------------------------------------------------------------------|----------------|--------------------------|------------------|-----------------|
| Estagiário              | 1           | O estagiário desempenha um papel crucial no desenvolvimento e implementação do projeto, trazendo perspectivas novas e habilidades atualizadas.                  | R$ 2,000.00    | R$ 4,444.44              | 12               | R$ 53,328.64    |
| UX Design/Research      | 1           | Desenvolve produtos e serviços digitais que satisfazem as necessidades, desejos e expectativas dos usuários, de forma eficiente e atrativa.                      | R$ 6,800.00    | R$ 10,297.00             | 4                | R$ 41,188.00    |
| Engenheiro da Computação| 1           | Responsável pela arquitetura geral do sistema e gerenciamento de hardwares...                                                                                   | R$ 8,000.00    | R$ 11,688                | 12               | R$ 140,016.00   |
| Técnico em eletrônica  | 4           | Responsável para manutenção e instalação dos sensores                                                                                                          | R$ 2,632.00    | R$ 4,895.00              | 12               | R$ 31,584.00    |
| Engenheiro de Software  | 2           | Encarregado do desenvolvimento...                                                                                                                                 | R$ 7,500.00    | R$ 11,192.50             | 12               | R$ 134,304.00   |
| Engenheiro Elétrico     | 1           | Encarregado do desenvolvimento de equipamentos elétricos.                                                                                                        | R$ 15,500.00   | R$ 19,824.50             | 12               | R$ 237,888.00   |
| PO                      | 1           | Líder de equipe de desenvolvimento                                                                                                                               | R$ 8,000.00    | R$ 12,688                | 12               | R$ 152,256.00   |

### Custos de Infraestrutura e Serviços

| Descrição                        | Quantidades | Detalhes                              | Gastos mensais | -     | Duração (Meses) | Total Acumulado |
|----------------------------------|-------------|---------------------------------------|----------------|-------|------------------|-----------------|
| MacBook Pro para desenvolvimento | 4           | Especificações do MacBook...          | R$ 24,998.00   | -     | -                | R$ 100,000.00   |
| Serviços da AWS - Amazon EC2     | 1           | -                                     | USD 1000       | -     | 12               | R$ 58,800.60    |
| Serviços da AWS - Amazon ECS     | 1           | -                                     | USD 500        | -     | 12               | R$ 29,400.00    |
| Serviços da AWS - Amazon RDS     | -           | -                                     | USD 400        | -     | 12               | R$ 23,520.00    |
| Sensores (temperatura, humidade, gas, pressão) | 10  | -                                     | USD 1500       | -     | -                | R$ 73,500.00    |

### Totais Gerais

|                                 |                 |                                         |                  |       |                  |                  |
|---------------------------------|-----------------|-----------------------------------------|------------------|-------|------------------|------------------|
| Somatório sem imposto e cotação dolar a R$ 4.9 | -               | -                                       | R$ 65,132.00     | -     | -                | R$ 781,584.00    |
| Somatório incluindo impostos por dentro e cotação dolar a R$ 4.9 | -               | -                                       | R$ 89,648.66     | -     | -                | R$ 1,075,784.00  |

-------

Os custos com os funcionários CLT são baseados em pesquisas de mercado dos sites [Catho](https://paraempresas.catho.com.br/quanto-custa-um-funcionario-para-empresa/) e [iDinheiro](https://www.idinheiro.com.br/). Além dos salários base, foram inclusos benefícios como plano de saúde de R$ 600, Vale Transporte de R$ 400 e VR de R$ 600.
  
-------
### Cálculo do imposto funcionário

O cálculo do imposto para um funcionário CLT (Consolidação das Leis do Trabalho) no Brasil envolveu a aplicação de diversos descontos sobre o salário bruto. Os principais descontos são o Imposto de Renda Retido na Fonte (IRRF), que varia conforme a faixa salarial, e as contribuições previdenciárias (INSS), que têm percentuais específicos. Além disso, podem existir outros descontos, como vale-transporte e benefícios. O salário líquido é obtido subtraindo-se esses descontos do salário bruto. O cálculo detalhado é feito mensalmente, considerando as particularidades de cada empregado e as regras fiscais vigentes.

#### Sendo atribuida a seguinte expressão: Custo Total = Salario Bruto + INSS + IRRF + FGTS + (VR + VT + PLANO DE SAUDE)

### Custo operacional
O custo operacional em um projeto de software e IoT inclui despesas contínuas relacionadas à manutenção, suporte e operação do sistema após sua implementação. Isso abrange gastos como hospedagem de servidores, manutenção de software, manutenção de dispositivos IoT, monitoramento de desempenho, entre outros. Considerar esses custos é crucial para garantir a sustentabilidade e eficiência contínua do projeto ao longo do tempo. E os valores foram categorizados como:

#### Hopedagem

As despesas relacionadas de hospedagem são direcionadas ao ambiente de desenvolvimento na AWS que estão distribuídas da seguinte forma:

- 1 EC2 a $ 1,000 mensais para operacionalizar os serviços do dashboard.
- RDS com custo mensal de $ 500 para o banco Postgres.
- ECS com custo mensal de $ 400 para orquestração dos serviços.
- 1 Engenheiro de software para configuração dos serviços

Custo mensal: R$ 20,502

#### Manutençao de software 

Recuros destinados a gerenciamento do software, e disponibilidade de recursos graficos e sistemáticos:

- 1 Engenheiros de software
- 1 Estagiário
- UX/UI Design
- PO (Projeto como um todo)

  Custo mensal: R$ 38,591

#### Dispositivos IoT
Manutenção e criação do ambiente IoT ao qual o sistema vai ser mantido, para isso requer equipe tanto para manutençao quanto para criação das interfaces e integrações com os demais sistemas:

- 1 Engenheiro da computação
- 1 Engenheiro Elétrico
- 4 Técnico em eletrônica

  Custo mensal: R$ 51,092

