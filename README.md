## How to setup a Go environment

1. Install go using gvm (https://larry-price.com/blog/2015/01/18/managing-a-go-environment-in-ubuntu)
2. Create the initial directory structure as show here
3. Create a new environment : gvm pkgset create env_name
4. Switch to the environment : gvm pkgset use env_name
5. Then run the export.sh script (. ./export.sh) to set the necessary paths, modify PROJECT_DIR as needed
6. go build (cd to the location of server.go)
7. go install
8. In the command window, run server
