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
      <td>Encarregado do desenvolvimento...</td>
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

### Despesas com serviços da AWS
Nessa tabela, as despesas relacionadas ao ambiente de desenvolvimento na AWS são distribuídas da seguinte forma:

- 1 EC2 a $ 2,000 mensais para operacionalizar os serviços do robô.
- RDS com custo mensal de R$ 200 para o banco MySQL.

### Cálculo do imposto por dentro

No cálculo do custo total de R$ 381,266.64, já foram incluídos impostos, que representam 12% do custo antes dos impostos. Para entender o valor dos impostos nesse contexto, usamos a seguinte fórmula:

- Custo Total = Custo Sem Imposto + (Custo Sem Imposto × Taxa de Imposto).

Reorganizando a fórmula para encontrar o custo sem imposto, temos:

- Custo Sem Imposto = Custo Total / (1 + Taxa de Imposto).

Substituindo com nossos valores:

- Custo Sem Imposto = R$ 381,266.64 / (1 + 0.12).

Calculando, temos:

- Custo Sem Imposto ≈ R$ 299,766.00.

Portanto, o valor dos impostos é:

- R$ 381,266.64 - R$ 632,400.46 = R$ 75,888.06.

### Hipótese de Estimativa de Receitas

Para calcular a estimativa de receitas deste projeto de implementação de um robô de autoatendimento, será considerado o cenário hipotético de desempenho financeiro descrito a seguir.

#### Premissas iniciais

- **Investimento Inicial**: R$ 1.000.000.
- **Gastos com Manutenção**: Consideremos que estes sejam uma fração anual do investimento inicial, por exemplo, 5% ao ano, que equivale a R$ 50.000.
- **Melhoria em Eficiência**: Suponhamos que a implementação do robô melhore a eficiência operacional, levando a uma economia de custos de 10% sobre os custos operacionais anuais da empresa.
- **Aumento na Capacidade Produtiva**: Presumamos um aumento de 15% na receita devido ao aumento da produção e satisfação do cliente.
- **Redução de Erros**: Estimemos uma economia de 5% nos custos associados a erros operacionais.

#### Cálculo da receita

- **Receita Anual Antes do Robô**: Para este exemplo, temos R$ 10.000.000.
- **Custos Operacionais Anuais Antes do Robô**: Suponhamos R$ 7.000.000.
- **Economia de Custos Operacionais com Robô**: 10% de R$ 7.000.000 = R$ 700.000.
- **Redução de Custos por Erros**: 5% de R$ 7.000.000 = R$ 350.000.
- **Aumento na Receita devido ao Robô**: 15% de R$ 10.000.000 = R$ 1.500.000.


Esta hipótese de lucros leva em conta o investimento inicial como um gasto único.
