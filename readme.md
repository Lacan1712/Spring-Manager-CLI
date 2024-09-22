# Spring Boot Manager CLI (smc)

Bem-vindo ao **Spring Boot CLI (smc)**! Esta ferramenta foi desenvolvida para melhorar e agilizar a manutenção e a criação de aplicações Spring Boot.

## Objetivo

O **smc** tem como objetivo simplificar o processo de configuração e inicialização de projetos Spring Boot, permitindo que desenvolvedores criem e gerenciem aplicações de maneira mais eficiente.

## Instalação

1. **Baixe o binário**: Acesse a seção de [Releases](https://github.com/Lacan1712/Spring-Manager-CLI/releases) do repositório e baixe o binário correspondente à sua versão do sistema operacional.

2. **Adicione ao PATH**: Coloque o binário e pasta 'src' em um diretório que esteja no seu `PATH`, para que você possa executá-lo de qualquer lugar no terminal.

## Uso

Após a instalação, você pode utilizar a CLI com o comando padrão:

```bash
smc --command
```

```bash

### Comandos Disponíveis

#### 1. `init`

O comando `init` é utilizado para inicializar um novo projeto Spring Boot com as configurações padrão.

Uso:

smc init


Descrição:

Este comando cria uma nova estrutura de projeto Spring Boot em um diretório específico.
O projeto será configurado com as dependências e arquivos básicos necessários para iniciar o desenvolvimento.

```
```bash
#### Comandos disponíveis

#### 2. `init --custom`

O comando `init -custom` é utilizado para inicializar um novo projeto mas com parâmetros personalizados.

*Uso:*

smc init --custom


Descrição:

Este comando cria uma nova estrutura de projeto Spring Boot em um diretório específico.
O projeto será configurado com as dependências e arquivos básicos necessários para iniciar o desenvolvimento mas com parâmetros personalizados.
```
```bash
### Comandos Disponíveis

#### 3. `create`

O comando `create` é utilizado para criar novos componentes dentro de um projeto Spring Boot, e em cada de componente pode ser passado um diretório path no padrão 'meu/caminho/component'.

*Uso:*

smc create --entity <NomeDaEntidade>

*Descrição:*

Este comando cria uma nova entidade no projeto Spring Boot, gerando a classe correspondente na estrutura de pacotes definida.
```
---
```bash
#### 4. `create --repository`

O comando `create --repository` é utilizado para criar um repositório associado a uma entidade.

*Uso:*

smc create --repository <NomeDaEntidade>

*Descrição:*

Este comando cria um novo repositório para a entidade especificada, permitindo a realização de operações de persistência no banco de dados.
```
---
```bash
#### 5. `create --controller`

O comando `create --controller` é utilizado para criar um controlador associado a uma entidade.

*Uso:*

smc create --controller <NomeDaEntidade>

*Descrição:*

Este comando cria um novo controlador para a entidade especificada, permitindo a manipulação de requisições e respostas HTTP no projeto.
```
