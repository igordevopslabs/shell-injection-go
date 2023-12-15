# Shell Injection Go Application

## Descrição
Este é um aplicativo web simples em Go, criado para fins educacionais, que demonstra uma vulnerabilidade de injeção de shell. O aplicativo possui um endpoint `/inject` que permite a execução de comandos shell arbitrários enviados via parâmetro de URL.

## Aviso de Segurança
⚠️ Este aplicativo contém vulnerabilidades de segurança intencionais e nunca deve ser usado em ambientes de produção ou exposto à Internet. É estritamente para uso em um ambiente controlado para fins de aprendizado sobre segurança em aplicações web.

## Como Usar
Para testar a vulnerabilidade de injeção de shell:

1. Inicie o servidor.
2. Envie comandos via `curl` ou navegador para `http://localhost:9000/inject?cmd=[seu_comando]`.

Exemplo: ```curl "http://localhost:9000/inject?cmd=ls"```


## Endpoints
- `/`: Retorna uma mensagem indicando que o aplicativo está pronto.
- `/ready`: Endpoint para checagem de prontidão.
- `/liveness`: Endpoint para checagem de vitalidade.
- `/inject`: Endpoint vulnerável para injeção de shell.

## Execução com Docker

### Construindo a Imagem

1.  `docker build -t shell-injection-go-app .`
2. `docker run -d -p 9000:9000 shell-injection-go-app`

## Licença
Este projeto é distribuído sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.