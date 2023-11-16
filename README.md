# Fundamentos-Golang

Fundamentos de go para quem deseja começar a estudar. Irei colocar uma breve descrição sobre algumas coisas que acho necessário.

Variáveis:

São locais de armazenamento para valores em Go.
Declaradas com var, := ou const dependendo do escopo e da mutabilidade.
Tipos de Dados:

Go possui tipos básicos como inteiros, ponto flutuante, booleanos, strings.
Também suporta tipos compostos como arrays, slices, maps, structs e ponteiros.
Funções:

Bloco de código reutilizável que executa uma tarefa específica.
Declaradas com func.
Operadores:

Go suporta operadores aritméticos, lógicos, relacionais, bit a bit, etc.
Structs:

Estrutura de dados que agrupa campos nomeados sob um nome.
Herança:

Go não tem herança tradicional, mas usa composição para reutilização de código.
Ponteiros:

Variáveis que armazenam endereços de memória.
Declarados usando o operador & e acessados usando o operador *.
Arrays e Slices:

Arrays são coleções fixas de elementos.
Slices são segmentos dinâmicos de arrays.
Map:

Coleção não ordenada de pares chave-valor.
Estruturas de Controle:

Controle de fluxo em programas.
Inclui condicionais (if, else) e loops (for).
Switch:

Estrutura de controle para substituir múltiplos if-else com várias condições.
Loops:

for é o único loop em Go, mas pode ser usado de várias maneiras.
Funções Variáticas:

Funções que aceitam um número variável de argumentos.
Funções Anônimas:

Funções sem nome, também chamadas de closures.
Funções Recursivas:

Funções que chamam a si mesmas.
Defer:

Adia a execução de uma função até que a função envolvente seja concluída.
Panic e Recover:

Mecanismos para lidar com erros graves e recuperar de panics.
Closure:

Função anônima que pode referenciar variáveis de seu escopo circundante.
Funções com Retornos Nomeados:

Permite nomear os resultados de retorno de uma função.
Métodos:

Funções associadas a tipos específicos.
Interfaces:

Conjunto de métodos sem implementação específica, permitindo polimorfismo.
Interfaces Genéricas:

Introduzido em Go 1.18, permite criação de interfaces parametrizadas.
Worker Pools:

Padrão de concorrência onde um grupo de trabalhadores processa tarefas de uma fila.
Generators:

Funções que retornam um canal para produzir sequências de valores.
Multiplexador:

Combina múltiplos canais em um único canal para leitura concorrente.
