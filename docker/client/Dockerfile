# Start your image with a node base image
FROM node:18-alpine as deps

# The /app directory should act as the main application directory
WORKDIR /app

# Copy local directories to the current local directory of our docker image (/app)
COPY ./client .

# Install node packages, install serve, build the app, and remove dependencies at the end
RUN yarn install
# for production
# RUN npm run build
RUN yarn build 

FROM nginx:alpine

# Copy the built React app to Nginx's web server directory
COPY --from=deps /app/dist /usr/share/nginx/html
COPY --from=deps /app/nginx.conf /etc/nginx/conf.d/default.conf

# Expose port 80 for the Nginx server
EXPOSE 80

# Start Nginx when the container runs
CMD ["nginx", "-g", "daemon off;"]