# Initial setup

```bash
curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.25.4/install.sh | bash
nvm install `cat .nvmrc`
nvm alias default `cat .nvmrc`
npm install -g node-dev
npm install
```
*Note, if you use zsh, you'll need to add the nvm alias stuff to your zshrc somewhere*
 
## Operations

### To run tests
```make test```

### To run tests in watch mode
```make test-w```

### To run the server
```make server```

### To run the server in watch mode
```make server-w```
