def combinacoes_de(tamanho_comb, elementos):
    if tamanho_comb == 0:
        return [[]]
    if len(elementos) == 0:
        return []
    primeiro_elemento = elementos[0]
    resto_elementos = elementos[1:]
    combinacoes_1 = combinacoes_de(tamanho_comb - 1, resto_elementos)
    for i in range(len(combinacoes_1)):
        combinacoes_1[i] = [primeiro_elemento] + combinacoes_1[i]
    combinacoes_2 = combinacoes_de(tamanho_comb, resto_elementos)
    return combinacoes_1 + combinacoes_2

tamanho_comb = 5
elementos = [1, 2, 3, 4, 5, 6, 7, 8, 9]
resultado = combinacoes_de(tamanho_comb, elementos)

for comb in resultado:
    print(comb)
