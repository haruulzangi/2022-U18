docker build -t docker_image_name -f Dockerfile .
docker run -d -p 5000:5000 docker_image_name

browse - http://127.0.0.1:5000/
