# Docker Quick Guide - Lyvora Server

## Building the Image

### Linux
```bash
sudo docker build -t lyvora-server .
```

### Windows
```powershell
docker build -t lyvora-server .
```

## Running the Service

### Using Docker Compose

#### Linux
```bash
# Start the service
sudo docker-compose up -d

# Or rebuild and start
sudo docker-compose up --build -d
```

#### Windows (PowerShell/CMD)
```powershell
# Start the service
docker-compose up -d

# Or rebuild and start
docker-compose up --build -d
```

### Using Docker Directly

#### Linux
```bash
sudo docker run -d -p 8080:8080 -v $(pwd)/data:/root/data --name lyvora lyvora-server
```

#### Windows (PowerShell)
```powershell
docker run -d -p 8080:8080 -v ${PWD}/data:/root/data --name lyvora lyvora-server
```

#### Windows (CMD)
```cmd
docker run -d -p 8080:8080 -v %cd%/data:/root/data --name lyvora lyvora-server
```

## Stopping the Service

### Using Docker Compose

#### Linux
```bash
sudo docker-compose down
```

#### Windows
```powershell
docker-compose down
```

### Using Docker Directly

#### Linux
```bash
# Stop the container
sudo docker stop lyvora

# Stop and remove the container
sudo docker rm -f lyvora
```

#### Windows
```powershell
# Stop the container
docker stop lyvora

# Stop and remove the container
docker rm -f lyvora
```

## Checking Status

### Linux
```bash
# View running containers
sudo docker ps

# View logs (Docker Compose)
sudo docker-compose logs -f

# View logs (Docker)
sudo docker logs -f lyvora
```

### Windows
```powershell
# View running containers
docker ps

# View logs (Docker Compose)
docker-compose logs -f

# View logs (Docker)
docker logs -f lyvora
```

## Accessing the Service

Once running, access your API at: **http://localhost:8080**

## Testing the Service

### Linux
```bash
curl http://localhost:8080/api/tracks
```

### Windows (PowerShell)
```powershell
curl http://localhost:8080/api/tracks
```

### Windows (CMD)
```cmd
curl http://localhost:8080/api/tracks
```
