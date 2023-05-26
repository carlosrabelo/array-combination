# Algoritmo de Combinações em Go

Este algoritmo foi criado como um exemplo para auxiliar uma pessoa interessada em combinar uma lista de elementos em matrizes menores. A função combinacoesDe implementa a lógica necessária para gerar todas as combinações possíveis, de acordo com os parâmetros fornecidos. A seguir, detalhamos o funcionamento do algoritmo:

1. A função `combinacoesDe` recebe dois parâmetros: `tamanhoComb` e `elementos`. `tamanhoComb` representa o número de elementos desejados em cada combinação, e `elementos` é a lista de entrada.

2. Primeiro, a função verifica duas condições de parada:
   - Se `tamanhoComb` for igual a 0, isso significa que não são necessários mais elementos na combinação. Nesse caso, a função retorna uma lista contendo uma combinação vazia (representada por `[][]int{{}}`).
   - Se o comprimento de `elementos` for igual a 0, isso significa que não há mais elementos disponíveis para combinar. Nesse caso, a função retorna uma lista vazia (representada por `[][]int{}`).

3. Se nenhuma das condições de parada for atendida, o algoritmo continua.

4. O algoritmo pega o primeiro elemento da lista `elementos` e o armazena na variável `primeiroElemento`.

5. Em seguida, cria uma nova lista `restoElementos` contendo todos os elementos de `elementos` a partir do segundo elemento.

6. Chama recursivamente a função `combinacoesDe` passando `tamanhoComb-1` (pois estamos buscando combinações com um elemento a menos) e `restoElementos`.

7. A variável `combinacoes1` recebe o resultado da chamada recursiva.

8. Em um loop `for`, itera sobre cada combinação em `combinacoes1` e insere o `primeiroElemento` no início de cada combinação usando `append`.

9. Em seguida, chama recursivamente a função `combinacoesDe` novamente com `tamanhoComb` e `restoElementos` (sem o `primeiroElemento`).

10. A variável `combinacoes2` recebe o resultado da segunda chamada recursiva.

11. Por fim, retorna a combinação dos resultados `combinacoes1` e `combinacoes2` usando `append`.

No programa principal (`main`), é definido o `tamanhoComb` desejado e a lista de `elementos`. A função `combinacoesDe` é chamada com esses parâmetros e o resultado é percorrido em um loop `for range`.
