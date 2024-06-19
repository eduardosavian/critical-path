# critic-path

Made by Eduardo Savian and Marcos Fehlauer

## Description

Fazer um programa que receba uma tabela de atividades, duração e precedentes e monte o grafo de PERT/CPM, identificando o caminho crítico e as folgas nas atividades.

Utilizar a notação de aula, com tempo máximo e mínimo para inicio e fim das atividades, com a atividade sendo representado pelos vértices do grafo, conforme exemplo a seguir.

O trabalho deverá ser apresentado no dia da entrega, explicando como ele foi feito e demonstrando seu funcionamento.

## How to use

### Build

```bash
go build -o critical src/main.go src/utils.go src/critical.go
```

### Put the path table in data/data.csv file

ex:

```csv
A;10;-
B;4;A
C;7;A
D;5;C
E;5;B,D
F;2;C
```

### Run

```bash
./critical
```

or

```ps1
.\critical.exe
```
