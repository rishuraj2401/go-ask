# Use an official Node.js runtime as a base image
FROM node:18-alpine

# Set the working directory in the container
WORKDIR /usr/src/app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the local code into the container
COPY . .

# Build the Vite + React app
RUN npm run build

# Expose the port that the app will run on
EXPOSE 8080

# Define the command to run the application
CMD ["npm", "run", "dev"]
