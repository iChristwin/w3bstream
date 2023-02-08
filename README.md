# W3bstream

## Overview

W3bStream is a general framework for connecting data generated by devices and machines in the physical world to the blockchain world. In a nutshell, W3bStream uses the IoTeX blockchain to orchestrate a decentralized network of gateways (i.e., W3bStream nodes) that streams encrypted data from IoT devices and machines and generates proofs of real-world facts to different blockchains. An overview graphic of W3bstream is 


![image](https://user-images.githubusercontent.com/448293/196618039-365ab2b7-f50a-49c8-a02d-c28e48acafcb.png)


## Documentation

Please visit [https://docs.w3bstream.com/](https://docs.w3bstream.com/).

Interested in contributing to the doc? Please edit on [Github](https://github.com/machinefi/w3bstream-docs-gitbook) 

## Arch

![w3bstream](__doc__/modules_and_dataflow.png)

## Run W3bstream with prebuilt docker images

### Run W3bstream node with W3bstream Studio
Check it out here [https://github.com/machinefi/w3bstream-studio](https://github.com/machinefi/w3bstream-studio).


### Run W3bstream node without W3bstream Studio

Make a path for w3bstream node. In the path, run the following command

```bash
curl https://raw.githubusercontent.com/machinefi/w3bstream/main/docker-compose.yaml > docker-compose.yaml
```

Edit the config in the `yaml` file if needed. Then run

```bash
docker-compose -p w3bstream -f ./docker-compose.yaml up -d
```

Your node should be up and running. 

Please note: the docker images are hosted at [GitHub Docker Registry](https://github.com/machinefi/w3bstream/pkgs/container/w3bstream)

## Run W3bstream node from code

If you are interested in diving into the code and run the node using a locally built docker, here is the steps of building the docker image from code.

### Build docker image from code

```bash
make build_backend_image
```

### Run server in docker containers

```bash
 make run_docker
 ```

 ### Stop server running in docker containers
 ```bash
 make stop_docker
 ```
 ### Delete docker resources
 ```bash
 make drop_docker
 ```

## Interact with W3bstream using CLI

Please refer to [HOWTO.md](./HOWTO.md) for more details.

## SDKs
### Devices
- Android: https://github.com/machinefi/w3bstream-android-sdk
- iOS: https://github.com/machinefi/w3bstream-ios-sdk
- Embedded: Coming soon!

### WASM
- Golang: https://github.com/machinefi/w3bstream-wasm-golang-sdk
- AssemblyScript: https://github.com/machinefi/w3bstream-wasm-ts-sdk
- Rust: Coming soon!


## Examples

Learning how to get started with W3bstream? Here is a quick get-start example: https://github.com/machinefi/get-started

More code examples: https://github.com/machinefi/w3bstream-examples

Step-by-step tutorials can be found on dev portal: https://developers.iotex.io/

## Community

- Developer portal: https://developers.iotex.io/
- Developer Discord (join #w3bstream channel): https://w3bstream.com/discord
