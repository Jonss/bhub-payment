# Desafio de Solução Tech - BHub

### 1. Como você pode domar essas regras de negócios que mudam o tempo todo e novas podem surgir a qualquer momento?

Eu usei dois design patterns, o `Observer` e `Chain of responsibility`. Com o primeiro, fiz com que os "observadores" reagissem (tomassem ações) quando forem assinados (subscribed). Assim, quando múltiplas ações forem necessárias, a ação em si está isolada em um escopo menor e só é chamado quando "assinado". Isso pode ser visto no arquivo `observer.go`.
Caso uma nova regra de negócio seja adicionada, um novo observer deve ser criado e executará quando entrar na condição requerida.
Para tratar as condições, usei o design pattern `chain of responsibility` como pode ser visto no arquivo `subscribers_chain`. Esse pattern auxilia para deixar mais limpo uma tomada de decisão muito longa, como no caso do projeto proposto, onde várias regras podem ser adicionadas. Ao invés de usar uma infinidade de if/else, ou mesmo um switch, quando uma nova condição for necessária, basta criar uma nova Chain e adicionar no ciclo.
Um chain chama o próximo chain, até que não hajam mais responsabilidades.

[Chain of responsibility](https://refactoring.guru/design-patterns/chain-of-responsibility)
[Observer](https://refactoring.guru/design-patterns/observer)

### 2. Como você pode construir um sistema que seja flexível o suficiente para lidar com acomplexidade e a necessidade dessas mudanças?

Usando os patterns descrito acima, o código tem uma manutenção razoável porque cada interação com o pagamento é isolada, necessitando apenas que caia na condição específica.
Se uma nova regra de negócio for adicionado, basta criar um novo observer e adicionar a condição em forma de chain of responsibility e adicionar ao fluxo.

### 3. Como você pode fazer isso sem se condenar a vários anos de suporte ao sistema?

Deixando o código com apenas uma responsabilidade e testando os comportamentos esperados dado uma condição. Se algo no futuro mudar, como adicionar ou mesmo remover regras, basta corrigir o chain of responsibility e criar/remover os observers.

## Notas:

- Desculpa ter feito em Go, eu teria feito em Python mas preferi fazer com uma linguagem que está mais a mão.
- Os arquivos `observer.go` contém todos os observers, num projeto real estaria isolado em vários arquivos. Decidi agrupar todos para não obrigar o revisor a abrir vários arquivos para análise. Segui a mesma ideia com o `subscribers.go`.
- Não ficou muito claro o que significava `guia de remessa`, então tratei como uma string qualquer.
- o modelo `product` está muito amplo, podendo ser uma adesão ou um produto. Não pensei numa melhor maneira de deixá-lo mais específico, mas estou ciente que poderia ser melhorado.
