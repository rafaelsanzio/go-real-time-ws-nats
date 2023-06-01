# Stream live ䷧

O propósito desse repositório é fazer o streaming de dados que inclui a implementação de um WebSocket e do sistema de streaming NATS.
O WebSocket é responsável por receber as mensagens provenientes de diversas fontes e enviá-las para processamento nos respectivos tópicos do NATS.
Uma vez processadas, as mensagens ficam disponíveis para serem consumidas por meio de subscrição ao tópico. Essa arquitetura permite uma comunicação assíncrona e distribuída entre os componentes do sistema, garantindo a disponibilidade e escalabilidade necessárias para lidar com grandes volumes de mensagens em tempo real.

## Iniciando ▶️

### Pré-requesitos

É necessário efetuar a instalação do gerenciador de container:

- Docker: https://www.docker.com/

### Instalação

1. Clone the repositório: https://github.com/rafaelsanzio/go-stream-live
2. Crie um arquivo .env como o .env.exemple mostra. Se necessário mude os valores.
3. É necessário criar uma network no docker, para conexão com outros serviços:
   ```sh
   docker network create app_network
   ```
4. Por fim, para iniciar a aplicação basta rodar os comandos de build e up.
   ```sh
   docker-compose build && docker-compose up
   ```
