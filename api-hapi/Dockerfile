# Pull base image from stock node image.
FROM node

# Add the current working folder as a mapped folder at /usr/src/app
ADD . /usr/src/app

# Set the current working directory to the new mapped folder.
WORKDIR /usr/src/app

# Install the express generator which gives you also scaffolding tools.
RUN npm install hapi --save

# Expose the node.js port to the Docker host.
EXPOSE 9292

# This is the stock express binary to start the app.
CMD [ "node", "server.js" ]
