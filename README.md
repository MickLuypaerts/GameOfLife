## Docker images
```bash
docker build -t gameoflife .
docker run -d --name gameoflife-container -p 8080:8080 gameoflife
```