# Simulador de Comunicação de IoT-Servdior
Estudo de simular uma comunicação de um dispositivo IoT para um servidor
## Pastas
public - pasta que contém arquivo html para visualizar as mensagens enviadas com as informações do dispositivo IoT (ID e temperatura)

video - pasta que contém um vídeo do funcionamento do simulador numa maquina rodando windows 11

\ - código fonte do simulador em Go e os arquivos do módulo do projeto

## Dependências
### Instalação do compilador e bibliotecas necessárias para execução do código em Go
Para executar o código é necessário instalar o compilador do Go e algumas bibliotecas como **paho.mqtt.golang** para comunicação MQTT e **net/http** para comunicação HTTP.

Baixe o instalador do Go para Windows no site oficial [Golang](https://go.dev/dl/).

Então crie uma pasta para o projeto e execute o comando abaixo no terminal do windows para criar o módulo do Go para executar o código.

```
go mod init nome-do-seu-modulo
```

Por fim, execute a linha de comando abaxio no terminal do windows para instalar a biblioteca necessária.

```
go get -u github.com/eclipse/paho.mqtt.golang
```

### Instalação do cliente MQTT no sistema
Baixe o instalador instalador do [MQTT Explore](https://mqtt-explorer.com/) e coloque com essa configuração da imagem abaixo.

![Captura de tela 2024-01-03 143204](https://github.com/valderlan/SimuladorMQTT/assets/71195621/a936d92a-34a9-400d-bf54-f4435f71e5bf)

Click em CONNECT e tela seguinte pesquise pelo o tópico criado pelo o programa, que no caso é **iot/data**. Dessa forma terá acesso a tela abaixo, onde será mostrado os valores da temperatura monitorada pelo o dispositivo **device123**.

![Captura de tela 2024-01-03 144046](https://github.com/valderlan/SimuladorMQTT/assets/71195621/e00c5bbd-f584-473c-a829-2830b6cea828)





