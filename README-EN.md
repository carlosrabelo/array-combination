## Combination Algorithm in Go

This algorithm was created as an example to help someone interested in combining a list of elements into smaller arrays. The `combinacoesDe` function implements the necessary logic to generate all possible combinations based on the provided parameters. Below, we provide a detailed explanation of how the algorithm works:

1. The `combinacoesDe` function takes two parameters: `tamanhoCombinacao` (combination size) and `elementos` (elements). `tamanhoCombinacao` represents the number of elements desired in each combination, and `elementos` is the input list.

2. The function follows a recursive approach to generate the combinations. It checks for two base cases:
   - If `tamanhoCombinacao` is equal to 0, it means no more elements are needed in the combination. In this case, the function returns an array containing an empty combination (`[][]int{{}}`).
   - If the length of `elementos` is equal to 0, it means there are no more elements available to combine. In this case, the function returns an empty array (`[][]int{}`).

3. If none of the base cases is met, the algorithm continues.

4. It takes the first element from the `elementos` list and stores it in a variable called `primeiroElemento` (first element).

5. It creates a new list called `elementosRestantes` (remaining elements) that contains all the elements from `elementos` except the first one.

6. The function recursively calls itself with `tamanhoCombinacao - 1` (since we are looking for combinations with one less element) and `elementosRestantes`.

7. The variable `combinacoes1` (combinations1) receives the result of the recursive call.

8. It iterates over each combination in `combinacoes1` and appends `primeiroElemento` to the beginning of each combination.

9. The function then recursively calls itself again with `tamanhoCombinacao` and `elementosRestantes` (without the `primeiroElemento`).

10. The variable `combinacoes2` (combinations2) receives the result of the second recursive call.

11. Finally, the function returns the combination of `combinacoes1` and `combinacoes2`.

This algorithm provides a way to generate all possible combinations of a given size from a list of elements.
