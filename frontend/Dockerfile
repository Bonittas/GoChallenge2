# Use the official Node.js image as the base image
FROM node:alpine

# Set the working directory in the container
WORKDIR /app

# Copy the package.json and package-lock.json files to the container
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the remaining application code to the container
COPY . .

# Build the React app for production
RUN npm run build

# Expose port 80 to the outside world
EXPOSE 8080

# Start the React app
CMD ["npm", "start"]
