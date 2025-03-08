1. Iniciamos el manejador
   `go mod init rps`
   rock-paper-scisor

2. Crear las routas y handlers

3. Ingresar a https://github.com/air-verse/air e instalar con
   `go install github.com/air-verse/air@latest`
   (para evitar tener que bajar y re-levantar el servidor con cada cambio)

Comprobar
`air -v`

4. Ejecutar la config inicial
   `air init`

5. Ejecutar el servidor con
   `air`

(Si no anda, ir a .profile o en .bashrc)

```bash
cd..
cd..
nano .profile
```

```bash
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

O sino en Windows revisar las variables de entorno
