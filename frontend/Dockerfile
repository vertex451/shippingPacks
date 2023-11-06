# Use the official Node.js image as the base image
FROM node:14.17.6

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy the package.json and package-lock.json files to the working directory
COPY package*.json ./

# Install the dependencies
RUN npm install

# Copy the remaining source code to the working directory
COPY . .

# Expose the port the React app runs on
EXPOSE 3000

# Start the React application
CMD ["npm", "start"]
