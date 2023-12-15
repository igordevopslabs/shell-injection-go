## Cenário de exploração

1. Faça o deploy de um pod com o utilitário do curl:

```kubectl run mycurlpod --image=curlimages/curl -i --tty -- sh```

2. Faça uma requisição para o service da aplicação vulnerável:

```curl http://shell-injection-go-app-service/inject?cmd=ls%20/```

### Protegendo o ambiente com kubearmor policy

1. aplique a seguinte regra:
```yaml
apiVersion: security.kubearmor.com/v1
kind: KubeArmorPolicy
metadata:
  name: block-shell
  namespace: default
spec:
  selector:
    matchLabels:
      app: shell-injection-go-app
  process:
    matchPaths:
      - path: /bin/sh
      - path: /bin/bash
      - path: /bin/busybox
  action:
    Block
```

Isso irá impedir que a aplicação seja explorada via shell injection