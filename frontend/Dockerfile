# Use the official Nginx image as a base image
FROM nginx:alpine

# Copy the frontend files to the default Nginx public html directory
COPY ../templates /usr/share/nginx/html

# Expose port 80 to the outside world
EXPOSE 80

# Command to run Nginx
CMD ["nginx", "-g", "daemon off;"]