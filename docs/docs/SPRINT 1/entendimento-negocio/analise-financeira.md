# Análise financeira

A análise financeira do projeto determina aspectos importantes para a estratégia econômica que será adotada pela empresa. Essa projeção permite mensurar o desempenho do projeto e elaborar melhores tomadas de decisões. Destaca-se que essa análise inclui entre os fatores um planejamento orçamentário e otimização de recursos. 

## Tabela de despesas e custos

A tabela abaixo ilustra as estimativas de despesas anuais e custos associados ao investimento no projeto. As despesas se relacionam aos compromissos financeiros recorrentes que a empresa tem, como as mensalidades dos serviços AWS. Por outro lado, os custos refletem os investimentos únicos ou fixos, como a aquisição de hardware (exemplo: sensores) e contratação de pessoal.

<table border="1">
  <thead>
    <tr>
      <th>Categorias</th>
      <th>Quantidades</th>
      <th>Descrição</th>
      <th>Gastos mensais</th>
      <th>Valor CLT por funcionário</th>
      <th>Duração (Meses)</th>
      <th>Total Acumulado</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>Estagiário</td>
      <td>1</td>
      <td>
        O estagiário desempenha um papel crucial no desenvolvimento e implementação do projeto, trazendo perspectivas
        novas e habilidades atualizadas.
      </td>
      <td>R$ 2,000.00</td>
      <td>R$ 4,444.44</td>
      <td>12</td>
      <td>R$ 53,328.64</td>
    </tr>
    <tr>
      <td>UX Design/Research</td>
      <td>1</td>
      <td>
        Desenvolve produtos e serviços digitais que satisfazem as necessidades, desejos e expectativas dos usuários, de
        forma eficiente e atrativa.
      </td>
      <td>R$ 6,800.00</td>
      <td>R$ 10,297.00</td>
      <td>4</td>
      <td>R$ 41,188.00</td>
    </tr>
    <tr>
      <td>Engenheiro da Computação</td>
      <td>1</td>
      <td>Responsável pela arquitetura geral do sistema e gerenciamento de hardwares...</td>
      <td>R$ 8,000.00</td>
      <td>R$ 11,688</td>
      <td>12</td>
      <td>R$ 140,016.00</td>
    </tr>
    <tr>
      <td>Técnico em eletrônica</td>
      <td>4</td>
      <td>Responsável para manutenção e instalação dos sensosres</td>
      <td>R$ 2,632.00</td>
      <td>R$ 4,895.00</td>
      <td>12</td>
      <td>31,584.00</td>
    </tr>
    <tr>
      <td>Engenheiro de Software</td>
      <td>2</td>
      <td>Encarregado do desenvolvimento...</td>
      <td>R$ 7,500.00</td>
      <td>11,192.50</td>
      <td>12</td>
      <td>R$ 134,304</td>
    </tr>
      <tr>
      <td>Engenheiro Elétrico</td>
      <td>1</td>
      <td>Encarregado do desenvolvimento de equipamentos elétricos.</td>
      <td>R$ 15,500.00</td>
      <td>19,824.50</td>
      <td>12</td>
      <td>R$ 237,888</td>
    </tr>
    <tr>
      <td>PO</td>
      <td>1</td>
      <td>Líder de equipe de desenvolvimento</td>
      <td>R$ 8,000.00</td>
      <td>R$ 12,688</td>
      <td>12</td>
      <td>R$ 152,256.00</td>
    </tr>
    <tr>
      <td>MacBook Pro para desenvolvimento</td>
      <td>4</td>
      <td>Especificações do MacBook...</td>
      <td>R$ 24,998.00</td>
      <td>-</td>
      <td>-</td>
      <td>R$ 100,000</td>
    </tr>
     <tr>
      <td>Serviços da AWS - Amazon EC2</td>
      <td>1</td>
      <td>-</td>
      <td>USD 1000</td>
      <td>-</td>
      <td>12</td>
      <td>R$ 58,800.60</td>
    </tr>
       <tr>
      <td>Serviços da AWS - Amazon ECS</td>
      <td>1</td>
      <td>-</td>
      <td>USD 500</td>
      <td>-</td>
      <td>12</td>
      <td>R$ 29,400.00</td>
    </tr>
    <tr>
      <td>Serviços da AWS - Amazon RDS</td>
      <td>-</td>
      <td>-</td>
      <td>USD 400</td>
      <td>-</td>
      <td>12</td>
      <td>R$ 23,520.00</td>
    </tr>
     <tr>
      <td>Sensores (temperatura, humidade, gas, pressão)</td>
      <td>10</td>
      <td>-</td>
      <td>USD 1500</td>
      <td>-</td>
      <td>-</td>
      <td>R$ 73,500.00</td>
    </tr>
  </tbody>
  <tr>
    <td colspan="3">Somatório sem imposto e cotação dolar a R$ 4.9</td>
    <td>R$ 65,132.00</td>
    <td>-</td>
    <td>-</td>
    <td>R$ 781,584.00</td>
  </tr>
  <tr>
    <td colspan="3">Somatório incluindo impostos por dentro e cotação dolar a R$ 4.9</td>
    <td>R$ 89,648.66</td>
    <td>-</td>
    <td>-</td>
    <td>R$ 1,075,784.00</td>
  </tr>
</table>

Os custos com os funcionários CLT são baseados em pesquisas de mercado dos sites [Catho](https://paraempresas.catho.com.br/quanto-custa-um-funcionario-para-empresa/) e [iDinheiro](https://www.idinheiro.com.br/). Além dos salários base, foram inclusos benefícios como plano de saúde de R$ 600, Vale Transporte de R$ 400 e VR de R$ 600.
  
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

