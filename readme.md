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
### Comandos Disponíveis
#### 1. `init`
```bash

```bash


O comando `init` é utilizado para inicializar um novo projeto Spring Boot com as configurações padrão.

Uso:

smc init


Descrição:

Este comando cria uma nova estrutura de projeto Spring Boot em um diretório específico.
O projeto será configurado com as dependências e arquivos básicos necessários para iniciar o desenvolvimento.

```
#### 2. `init --custom`

O comando `init -custom` é utilizado para inicializar um novo projeto mas com parâmetros personalizados.

*Uso:*

smc init --custom


Descrição:

Este comando cria uma nova estrutura de projeto Spring Boot em um diretório específico.
O projeto será configurado com as dependências e arquivos básicos necessários para iniciar o desenvolvimento mas com parâmetros personalizados.
```
#### 3.  `database --listables -n myConnection`
```bash
O comando `init -listables` é utilizado para listar todas as tabelas de um database passando o argumento -n ou --name como nome da connection.

*Uso:*

smc database --listables -n myConnection


Descrição:
Para utilizar os comandos que precisam de banco de dados você precisa definir o json para a conexão da CLI com o banco, veja como definir em Como definir uma nova conection
```
#### 4.  `create --entity`
```bash
O comando `create --entity` é utilizado para criar uma nova entity spring vazia, o comando espera que seja passado um caminho para a nova entity
que pode ser no formato "/caminho/entity/minhaEntity" ou em formato de pacote "caminho.entity.minhaEntity"

*Uso:*

smc create --entity entity.perfil.usuarios  #cria uma nova entity chamada "usuarios" no caminho especificado
```
* ##### 4.1  `create --entity -d`
```bash
O comando `create --entity` quano utilizado com a flag -d ou --database pode gerar uma nova entity com base em uma tabela existente, a flag espera o nome definido na configuração
de uma nova connection (ùltima seção) e espera que o último nome informado no caminho da criação da entity corresponda ao nome da tabela
*Uso:*

smc create --entity entity.perfil.usuarios -d myConnection #A CLI espera que o último nome informado no caminho da criação da entity seja o nome da tabela
                                                           #O argumento `-d` ou `--database` especifica qual "connectionName" deve ser usado (Veja a seção que explica a criação de uma connection)
                                                           #Cria uma nova entity da tabela usuarios

Descrição:
Agora é possível criar uma Entity com base em uma tabela existente, com a flag -d ou --database que espera o nome de uma connection definida no arquivo database.json
```

#### 5.  `create --controller`
```bash
O comando create --controller é utilizado para criar um novo controller Spring Boot vazio. Ele segue a mesma lógica do comando de entity, onde você pode passar um caminho no formato /caminho/controller/meuController ou no formato de pacote caminho.controller.meuController.

*Uso:*

smc create --conteoller entity.controller.meuController  #cria um novo controller chamado "meuController" no caminho especificado
```
#### 6.  `create --repository`
```bash
O comando create --repository é utilizado para criar um novo repository Spring Boot vazio. Assim como os outros comandos, você pode passar o caminho no formato /caminho/repository/meuRepository ou no formato de pacote caminho.repository.meuRepository.

*Uso:*

smc create --repository repository.perfil.usuarios #cria um novo repository chamado usuario

Observações da versão:
As atuais versões da CLI ainda não oferecem suporte para integração com database para criação de repositorys, estamos trabalhando para que seja possível em futuras versões.
```


## Criando uma nova Connection

## Estrutura de Configuração de Conexão com Banco de Dados

A configuração das conexões com bancos de dados será feita utilizando o seguinte formato JSON:
* O arquivo se encontra em json/database.json após download e extração do zip da release

```json
{
    "connections": [
        {   
            "connectionName":"myConnection",
            "driveDatabase": "driveDatabase",
            "databaseName": "databaseName",
            "host": "localhost",
            "port":"8888",
            "username": "usernames",
            "password": "123456789"
        }
    ]
}
```
* ### Atributos de Conexão

  - #### connections
    Um array que contém objetos, onde cada objeto representa uma configuração de conexão com o banco de dados.

    - **connectionName (string):**  
      Identificador único para a conexão.  
      Este valor deve ser passado como argumento nos comandos que precisam se conectar ao banco de dados.  
      **Exemplo:** `" create --entity entity.perfil.usuarios -d myConnection"`.

    - **driveDatabase (string):**  
      Especifica o tipo de banco de dados utilizado
      Os drivers suportados devem ser adicionados como "postgres", "mysql", ou "sqlserve".  
      Define qual driver será usado para conectar ao banco.

    - **databaseName (string):**  
      O nome do banco de dados ao qual se deseja conectar.  
      **Exemplo:** `"databaseName"`.

    - **host (string):**  
      O endereço do servidor onde o banco de dados está hospedado.  
      Pode ser `"localhost"` ou um endereço IP, como `"127.0.0.1"`.

    - **port (string):**  
      O número da porta utilizada para a conexão com o banco de dados.  
      **Exemplos:** `"5432"` para PostgreSQL, `"3306"` para MySQL, ou `"8888"` (personalizado).

    - **username (string):**  
      Nome de usuário para autenticação no banco de dados.

    - **password (string):**  
      Senha usada para autenticação no banco de dados.