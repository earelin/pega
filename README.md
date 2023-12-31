# Pega

[![codecov](https://codecov.io/gh/earelin/pega/graph/badge.svg?token=7CXXF1vn9p)](https://codecov.io/gh/earelin/pega)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=earelin_pega&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=earelin_pega)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=earelin_pega&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=earelin_pega)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=earelin_pega&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=earelin_pega)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=earelin_pega&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=earelin_pega)

O proxecto Pega ten o obxectivo de extraer e amosar datos electorais.

Implementa una API REST para consular los datos de resultados
electorais almacenados no sistema.

## API

### Datos xerais

#### GET `/comunidades-autonomas`

*Valor Retornado*

```
[
  {
    id: long
    nome: string
  }
]
```

#### GET `/comunidade-autonoma/{id}/provincias`

*Valor Retornado*

```
[
  {
    id: long
    nome: string
  }
]
```

#### GET `/provincias`

*Valor Retornado*

```
[
  {
    id: long
    nome: string
  }
]
```

#### GET `/provincia/{id}/concellos`

*Valor Retornado*

```
[
  {
    id: long
    nome: string
  }
]
```

#### GET `/concellos/pescuda/{search}`

*Valor Retornado*

```
[
  {
    id: long
    nome: string
  }
]
```

### Datos de Procesos Electorais

#### GET `/procesos-electorais`

Lista de procesos electorais.

*Parámetros*

`tipo`: Tipo de proceso electoral:
```
 1: Referéndum
 2: Xerais
 3: Congreso
 4: Senado
 5: Municipais
 6: Autonómicas
 7: Cabildos Insulares
10: Parlamento Europeu
15: Partidos Xudiciais e Diputacións Provinciais
```
`ambito`: Ámbito territorial do proceso electoral.
Código INE.

*Valor Retornado*

```
[
  {
    id: long
    tipo: int
    ambito: int
    data: iso string
    primeiroAvanceParticipacion: iso string
    segundoAvanceParticipacion: iso string
  }
]
```

#### GET `/procesos-electorais/{id}/datos-xerais`

Datos xerais do proceso electoral.

*Parametros*

`id`: Id do proceso electoral.

*Valor Retornado*

```
{
  censoIne: number
  censoCera: number
}
```
#### GET `/procesos-electorais/{id}/datos-xerais/{nivel_administrativo}/{id_entidade}`

Datos xerais do proceso electoral nunha entidade administrativa

*Parametros*

`id`: Id do proceso electoral.

`nivel_administrativo`: Nivel administrativo: `mesa`, `seccion`, `distrito`, `concello`, `provincia`,
`comunidade-autonoma`.

`id_entidade`: Id da entidade administrativa.

*Valor Retornado*

```
{
  censoIne: number
  censoCera: number
}
```

#### GET `/procesos-electorais/{id}/resultados`

Resultados xerais do proceso electoral.

*Parametros*

`id`: Id do proceso electoral.

*Valor Retornado*

```
{
  votantesPrimeiroAvanceParticipacion: number
  votantesSegundoAvanceParticipacion: number
  votantesCere: number
  votosEnBranco: number
  votosNulos: number
  votosACandidaturas: number
}
```

#### GET `/procesos-electorais/{id}/resultados/candidaturas`

Resultados xerais por candidaturas.

*Parametros*

`id`: Id do proceso electoral.

*Valor Retornado*

```
[
  {
    candidatura: {
      id: number
      nome: string
    }
    representantesEleitos: number
    votos: number
  }
]
```

#### GET `/procesos-electorais/{id}/resultados/{nivel_administrativo}/{id_entidade}`

Resultados nunha entidate administrativa.

*Parametros*

`id`: Id do proceso electoral.

`nivel_administrativo`: Nivel administrativo: `mesa`, `concello`, `provincia`, `comunidade-autonoma`.

`id_entidade`: Id da entidade administrativa.

*Valor Retornado*

```
{
  votantesPrimeiroAvanceParticipacion: number;
  votantesSegundoAvanceParticipacion: number;
  votantesCere: number;
  votosEnBranco: number;
  votosNulos: number;
  votosACandaturas: number;
  votosPorCandidatura: [
    {
      candidatura: Candidatura;
      representantesEleitos: number;
      votos: number;
    }
  ];
}
```

## Ferramentas

### inebase

Importa datos base do INE á base de datos.

```
inebase [OPCIONS] CONSUNTO_DATOS FICHEIRO
```

`CONXUNTO_DATOS`: Conxunto de datos a importar. Valores soportados: concellos.

`FICHEIRO`: Ruta ao ficheiro cos datos INE.

#### Conxuntos de datos soportados

##### Concellos

Ficheiro xslx co listado de concellos do INE.

### infoelectoral

Importa datos da web infoelectoral do Ministerio de Interior á base de datos.

```
infoelectoral [OPCIONS] FICHEIRO
```

`FICHEIRO`: Ruta ao ficheiro cos datos electorais.
