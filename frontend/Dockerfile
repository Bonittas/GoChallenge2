FROM node:18-alpine
RUN mkdir -p /app

WORKDIR /app

COPY . .
COPY package*.json ./

# Install only production dependencies
RUN npm install 

# Copy the rest of the application code



# Expose port 3000 (the default port for React apps)
EXPOSE 3000

# Serve the built files
CMD [ "npm", "start" ]