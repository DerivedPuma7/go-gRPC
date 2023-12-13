# go-gRPC

## plugins necessários:

### compilador do protocol buffer:
https://grpc.io/docs/protoc-installation/  
apt install -y protobuf-compiler

## para entender os contratos protoc e transformar no arquivo go:
protoc --go_out=. --go-grpc_out=. proto/course_category.proto

## grpc client
https://github.com/ktr0731/evans

### download evans:
- va para releases no github e baixe a versão mais recente .tar.gz
- extraia o arquivo com:
    - tar -xzf go1.21.5.linux-amd64.tar.gz
- mova o arquivo extraído para:
    - /usr/local/bin/
- adicione a variável de ambiente editando o arquivo ~/.profile
    - export PATH=$PATH:/usr/local/bin/evans
- reinicie a sessão do usuário
- ja deve ser possível rodar o comando evans --version
- para acessar a cli do evans: 
    - evans -r repl
    - com o server grpc rodando na porta padrão
    - em outra porta, é nessessário especificar para o evans

## banco de dados sqlite
- tenha o banco instalado com:
    - sudo apt install sqlite3
- na raiz, acesse a cli do sqlite e crie as tabelas:
    - sqlite3 db.sqlite (acessa cli)
    - create table categories(id string, name string, description string);
